package writer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/example/internal/models"
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
	fmt.Printf("%s\n", text)
	o := models.Order{}
	d := models.Delivery{}
	p := models.Payment{}
	i := models.Items{}
	res := models.Orders{
		o,
		d, p, i,
	}
	err := json.Unmarshal(text, &res)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
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
