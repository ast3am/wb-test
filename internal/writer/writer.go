package writer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/example/internal/models"
	"github.com/go-playground/validator/v10"
)

type DB interface {
	InsertOrders(ctx context.Context, orders models.Orders) error
	GetOrders(ctx context.Context) ([]*models.Orders, error)
}

type Cache interface {
	InsertOrdersToCache(orders models.Orders)
}

type writer struct {
	db    DB
	Cache Cache
}

func NewWriter(db DB, cache Cache) *writer {
	return &writer{
		db:    db,
		Cache: cache,
	}
}

func (w *writer) Write(text []byte) error {
	jsonModel := models.JsonReadDTO{}

	err := json.Unmarshal(text, &jsonModel)
	if err != nil {
		fmt.Println("Error unmarshal:", err)
		return err
	}
	orders, err := transformData(jsonModel)
	if err != nil {
		fmt.Println("Error unmarshal items:", err)
		return err
	}
	err = validate(*orders)
	if err != nil {
		fmt.Println("Non valid date:", err)
		return err
	}
	ctx := context.Background()
	err = w.db.InsertOrders(ctx, *orders)
	if err != nil {
		fmt.Print("Error:", err)
		return err
	}
	w.Cache.InsertOrdersToCache(*orders)
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
	i := []models.Items{}

	for _, itemJson := range dto.Items {
		var item models.Items
		err := json.Unmarshal(itemJson, &item)
		if err != nil {
			fmt.Println("Error unmarshal items:", err)
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
		o,
		d, p, i,
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
