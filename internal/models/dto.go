package models

import (
	"encoding/json"
	"time"
)

type JsonReadDTO struct {
	OrderUid          string            `json:"order_uid" validate:"required"`
	TrackNumber       string            `json:"track_number" validate:"required"`
	Entry             string            `json:"entry" validate:"required"`
	Locale            string            `json:"locale" validate:"required"`
	InternalSignature string            `json:"internal_signature" validate:"required"`
	CustomerID        string            `json:"customer_id" validate:"required"`
	DeliveryService   string            `json:"delivery_service" validate:"required"`
	Shardkey          string            `json:"shardkey" validate:"required"`
	SmID              int               `json:"sm_id" validate:"required"`
	DateCreated       time.Time         `json:"date_created" validate:"required"`
	OofShard          string            `json:"oof_shard" validate:"required"`
	Delivery          Delivery          `json:"delivery"`
	Payment           Payment           `json:"payment"`
	Items             []json.RawMessage `json:"items"`
}

type JsonWriteDTO struct {
	OrderUid          string    `json:"order_uid" validate:"required"`
	TrackNumber       string    `json:"track_number" validate:"required"`
	Entry             string    `json:"entry" validate:"required"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Items   `json:"items"`
	Locale            string    `json:"locale" validate:"required"`
	InternalSignature string    `json:"internal_signature" validate:"required"`
	CustomerID        string    `json:"customer_id" validate:"required"`
	DeliveryService   string    `json:"delivery_service" validate:"required"`
	Shardkey          string    `json:"shardkey" validate:"required"`
	SmID              int       `json:"sm_id" validate:"required"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`
}
