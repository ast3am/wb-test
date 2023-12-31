package db

import (
	"context"
	"fmt"
	"github.com/ast3am/wb-test/internal/models"
	"github.com/jackc/pgx/v4"
)

//go:generate mockery --name logger
type logger interface {
	DebugMsg(msg string)
	ErrorMsg(msg string, err error)
}

type DB struct {
	dbConnect *pgx.Conn
	log       logger
}

func NewClient(ctx context.Context, cfg *models.Config, log logger) (*DB, error) {
	DB := DB{dbConnect: nil, log: log}
	var err error
	posgresURL := "postgresql://" + cfg.SqlConfig.UsernameDB + ":" + cfg.SqlConfig.PasswordDB + "@" + cfg.SqlConfig.HostDB + ":" + cfg.SqlConfig.PortDB + "/" + cfg.SqlConfig.DBName
	DB.dbConnect, err = pgx.Connect(ctx, posgresURL)
	if err != nil {
		return nil, err
	}

	err = DB.dbConnect.Ping(ctx)
	if err != nil {
		return nil, err
	}
	log.DebugMsg("connection to DB is OK")
	return &DB, nil
}

func (db *DB) Close(ctx context.Context) {
	err := db.dbConnect.Close(ctx)
	if err != nil {
		db.log.ErrorMsg("error closing connection to BD", err)
	} else {
		db.log.DebugMsg("closing connection to BD is OK")
	}
}

func (db *DB) InsertOrders(ctx context.Context, orders models.Orders) error {

	// order
	order := orders.Order
	queryOrder := `
	INSERT INTO orders(order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.dbConnect.Exec(ctx, queryOrder, order.OrderUid, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.Shardkey, order.SmID, order.DateCreated, order.OofShard)
	if err != nil {
		return err
	}

	// delivery
	delivery := orders.Delivery
	queryDelivery := `
	INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.dbConnect.Exec(ctx, queryDelivery, orders.Order.OrderUid, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email)
	if err != nil {
		return err
	}

	// payment
	payment := orders.Payment
	queryPayment := `
	INSERT INTO payment (order_uid, transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err = db.dbConnect.Exec(ctx, queryPayment, orders.Order.OrderUid, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee)
	if err != nil {
		return err
	}

	//items
	items := orders.Items
	for _, item := range items {
		queryItems := `
	INSERT INTO items (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`
		_, err = db.dbConnect.Exec(ctx, queryItems, orders.Order.OrderUid, item.ChrtID, item.TrackNumber, item.Price, item.Rid, item.Name, item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *DB) GetOrders(ctx context.Context) ([]*models.Orders, error) {
	count := 0
	queryCount := `SELECT count(*) FROM orders`
	err := db.dbConnect.QueryRow(ctx, queryCount).Scan(&count)
	if err != nil {
		return nil, err
	}

	query := `
	SELECT * FROM orders o
	JOIN delivery d on d.order_uid = o.order_uid
	JOIN payment p on p.order_uid = o.order_uid
	`

	rows, err := db.dbConnect.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ordersArr := make([]*models.Orders, 0, count)

	for rows.Next() {
		o := models.Order{}
		d := models.Delivery{}
		p := models.Payment{}
		err = rows.Scan(&o.OrderUid, &o.TrackNumber, &o.Entry, &o.Locale, &o.InternalSignature, &o.CustomerID, &o.DeliveryService, &o.Shardkey, &o.SmID, &o.DateCreated, &o.OofShard,
			&o.OrderUid, &d.Name, &d.Phone, &d.Zip, &d.City, &d.Address, &d.Region, &d.Email,
			&o.OrderUid, &p.Transaction, &p.RequestID, &p.Currency, &p.Provider, &p.Amount, &p.PaymentDt, &p.Bank, &p.DeliveryCost, &p.GoodsTotal, &p.CustomFee)
		if err != nil {
			return nil, err
		}

		orders := models.Orders{
			Order:    o,
			Delivery: d,
			Payment:  p,
			Items:    nil,
		}
		ordersArr = append(ordersArr, &orders)
	}

	queryItems := `
	SELECT * FROM items WHERE order_uid = $1
	`
	for _, o := range ordersArr {
		rowsItems, err := db.dbConnect.Query(ctx, queryItems, o.Order.OrderUid)
		if err != nil {
			fmt.Printf("Ошибка чтения g: %s\n", err)
			return nil, err
		}
		defer rowsItems.Close()
		iArr := make([]models.Item, 0, 5)
		for rowsItems.Next() {
			i := models.Item{}
			err = rowsItems.Scan(&o.Order.OrderUid, &i.ChrtID, &i.TrackNumber, &i.Price, &i.Rid, &i.Name, &i.Sale, &i.Size, &i.TotalPrice, &i.NmID, &i.Brand, &i.Status)
			if err != nil {
				fmt.Printf("Ошибка записи: %s\n", err)
				return nil, err
			}
			iArr = append(iArr, i)
		}
		o.Items = iArr
	}

	return ordersArr, nil
}
