package api

import (
	"encoding/json"
	"github.com/example/internal/models"
	"github.com/gorilla/mux"
	"net/http"
)

type Cache interface {
	GetOrdersById(Uid string) models.Orders
}

type handler struct {
	cache Cache
}

func NewHandler(cache Cache) *handler {
	return &handler{
		cache: cache,
	}
}

func (h *handler) Register(router *mux.Router) {
	router.HandleFunc("/{id}", h.getOrdersByID).Methods("GET")
}

func (h *handler) getOrdersByID(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	itemID := param["id"]

	orders := h.cache.GetOrdersById(itemID)
	res := transformToDTO(orders)
	if orders.Order.OrderUid == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User not found"))
		return
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Marshal error \n" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
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
