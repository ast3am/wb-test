package api

import (
	"encoding/json"
	"github.com/example/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

//go:generate mockery --name cache
type cache interface {
	GetOrdersById(Uid string) models.Orders
}

//go:generate mockery --name logger
type logger interface {
	HandlerLog(r *http.Request, status int, msg string)
	HandlerErrorLog(r *http.Request, status int, msg string, err error)
}

type handler struct {
	cache cache
	log   logger
}

func NewHandler(cache cache, log logger) *handler {
	return &handler{
		cache: cache,
		log:   log,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/{id}", h.GetOrdersByID).Methods("GET")
}

func (h *handler) GetOrdersByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	itemID := param["id"]

	orders := h.cache.GetOrdersById(itemID)
	res := transformToDTO(orders)
	if orders.Order.OrderUid == "" {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Order not found"))
		h.log.HandlerLog(r, http.StatusOK, "order with ID "+itemID+" not found")
		return
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Marshal error \n" + err.Error()))
		h.log.HandlerErrorLog(r, http.StatusOK, "", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
	h.log.HandlerLog(r, http.StatusOK, "return order data with UID "+itemID)
}

func transformToDTO(orders models.Orders) *models.JsonWriteDTO {
	return &models.JsonWriteDTO{
		orders.Order.OrderUid,
		orders.Order.TrackNumber,
		orders.Order.Entry,
		orders.Delivery,
		orders.Payment,
		orders.Items,
		orders.Order.Locale,
		orders.Order.InternalSignature,
		orders.Order.CustomerID,
		orders.Order.DeliveryService,
		orders.Order.Shardkey,
		orders.Order.SmID,
		orders.Order.DateCreated,
		orders.Order.OofShard,
	}
}
