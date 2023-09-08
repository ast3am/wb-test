package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	OrderUid          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Items struct {
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

type Orders struct {
	Order    Order    `json:"order"`
	Delivery Delivery `json:"delivery"`
	Payment  Payment  `json:"payment"`
	Items    Items    `json:"items"`
}

func main() {
	jsonStr := `{
        "order_uid": "b563feb7b2b84b6test",
        "track_number": "WBILMTESTTRACK",
        "entry": "WBILl",
        "delivery": {
            "name": "Test Testov",
            "phone": "+9721111111",
            "zip": "2639809",
            "city": "Kiryat Mozkin",
            "address": "Ploshad Mira 15",
            "region": "Kraiot",
            "email": "test@gmail.com"
        },
        "payment": {
            "transaction": "b563feb7b2b84b6test",
            "request_id": "",
            "currency": "USD",
            "provider": "wbpay",
            "amount": 1817,
            "payment_dt": 1637907627,
            "bank": "alpha",
            "delivery_cost": 1500,
            "goods_total": 317,
            "custom_fee": 0
        },
        "items": [
            {
                "chrt_id": 9934922,
                "track_number": "WBILMTESTTRACK",
                "price": 453,
                "rid": "ab4319087a764ae0btest",
                "name": "Mascaras",
                "sale": 30,
                "size": "0",
                "total_price": 317,
                "nm_id": 2389232,
                "brand": "Vivienne Sabo",
                "status": 202
            }
        ],
        "locale": "en",
        "internal_signature": "",
        "customer_id": "test",
        "delivery_service": "meest",
        "shardkey": "9",
        "sm_id": 99,
        "date_created": "2021-11-26T06:22:19Z",
        "oof_shard": "1"
    }`

	var ordersData struct {
		Order    Order             `json:"order"`
		Delivery Delivery          `json:"delivery"`
		Payment  Payment           `json:"payment"`
		Items    []json.RawMessage `json:"items"`
	}

	err := json.Unmarshal([]byte(jsonStr), &ordersData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var items []Items
	for _, rawItem := range ordersData.Items {
		var item Items
		err := json.Unmarshal(rawItem, &item)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		items = append(items, item)
	}

	orders := Orders{
		Order:    ordersData.Order,
		Delivery: ordersData.Delivery,
		Payment:  ordersData.Payment,
	}
	if len(items) > 0 {
		orders.Items = items[0]
	}

	fmt.Printf("Orders: %+v\n", orders)
}
