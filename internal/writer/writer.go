package writer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ast3am/wb-test/internal/models"
	"github.com/go-playground/validator/v10"
)

//go:generate mockery --name db
type db interface {
	InsertOrders(ctx context.Context, orders models.Orders) error
	GetOrders(ctx context.Context) ([]*models.Orders, error)
}

//go:generate mockery --name cache
type cache interface {
	InsertOrdersToCache(orders models.Orders)
}

//go:generate mockery --name logger
type logger interface {
	InfoMsgf(format string, v ...interface{})
}

type writer struct {
	db    db
	Cache cache
	log   logger
}

func NewWriter(db db, cache cache, log logger) *writer {
	return &writer{
		db:    db,
		Cache: cache,
		log:   log,
	}
}

func (w *writer) Write(text []byte) error {
	jsonModel := models.JsonReadDTO{}

	err := json.Unmarshal(text, &jsonModel)
	if err != nil {
		err = fmt.Errorf("error unmarshal: %s", err)
		return err
	}
	orders, err := transformData(jsonModel)
	if err != nil {
		err = fmt.Errorf("error unmarshal items: %s", err)
		return err
	}
	err = validate(*orders)
	if err != nil {
		err = fmt.Errorf("non valid date: %s", err)
		return err
	}
	ctx := context.Background()
	err = w.db.InsertOrders(ctx, *orders)
	if err != nil {
		err = fmt.Errorf("insert To Sql with UID:  %s %s", orders.Order.OrderUid, err)
		return err
	}
	w.Cache.InsertOrdersToCache(*orders)
	w.log.InfoMsgf("added new order with UID %s", orders.Order.OrderUid)
	return nil

}

func (w *writer) GetCacheOnStart(ctx context.Context) error {
	temp := make([]*models.Orders, 0, 10)
	temp, err := w.db.GetOrders(ctx)
	if err != nil {
		return err
	}
	if len(temp) != 0 {
		for _, k := range temp {
			w.Cache.InsertOrdersToCache(*k)
		}
	}
	return nil
}

func transformData(dto models.JsonReadDTO) (*models.Orders, error) {

	o := models.Order{}
	d := models.Delivery{}
	p := models.Payment{}
	i := make([]models.Item, 0, 5)

	for _, itemJson := range dto.Items {
		var item models.Item
		err := json.Unmarshal(itemJson, &item)
		if err != nil {
			return nil, err
		}
		i = append(i, item)
	}

	o.OrderUid = dto.OrderUid
	o.TrackNumber = dto.TrackNumber
	o.Entry = dto.Entry
	o.Locale = dto.Locale
	o.InternalSignature = dto.InternalSignature
	o.CustomerID = dto.CustomerID
	o.DeliveryService = dto.DeliveryService
	o.Shardkey = dto.Shardkey
	o.SmID = dto.SmID
	o.DateCreated = dto.DateCreated
	o.OofShard = dto.OofShard
	d = dto.Delivery
	p = dto.Payment

	res := models.Orders{
		Order:    o,
		Delivery: d,
		Payment:  p,
		Items:    i,
	}

	return &res, nil
}

func validate(orders models.Orders) error {
	validate := validator.New()
	err := validate.Struct(orders.Order)
	if err != nil {
		return err
	}
	err = validate.Struct(orders.Delivery)
	if err != nil {
		return err
	}
	err = validate.Struct(orders.Payment)
	if err != nil {
		return err
	}
	for _, item := range orders.Items {
		err = validate.Struct(item)
		if err != nil {
			return err
		}
	}
	return err
}
