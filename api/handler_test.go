package api

import (
	"github.com/example/api/mocks"
	"github.com/example/internal/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// ответ обязательно в одну строку!
var TestPositiveResponse = []byte(`{"order_uid":"1","track_number":"WBILMTESTTRACK","entry":"WBILl","delivery":{"name":"Test Testov","phone":"+79211552101","zip":"2639809","city":"Kiryat Mozkin","address":"Ploshad Mira 15","region":"Kraiot","email":"test@gmail.com"},"payment":{"transaction":"b563feb7b2b84b6test","request_id":"12345","currency":"USD","provider":"wbpay","amount":1817,"payment_dt":1637907627,"bank":"alpha","delivery_cost":1500,"goods_total":317,"custom_fee":0},"items":[{"chrt_id":1,"track_number":"WBILMTESTTRACK","price":453,"rid":"ab4319087a764ae0btest","name":"Mascaras","sale":30,"size":"0","total_price":317,"nm_id":2389232,"brand":"Vivienne Sabo","status":202}],"locale":"en","internal_signature":"sad","customer_id":"test","delivery_service":"meest","shardkey":"9","sm_id":99,"date_created":"2021-11-26T06:22:19Z","oof_shard":"1"}`)

func TestNewHandler(t *testing.T) {
	cache := mocks.NewCache(t)
	log := mocks.NewLogger(t)
	h := NewHandler(cache, log)
	require.NotNil(t, h)
}

func TestHandler_GetOrdersByID(t *testing.T) {
	cache := mocks.NewCache(t)
	log := mocks.NewLogger(t)
	h := NewHandler(cache, log)
	router := mux.NewRouter()
	h.Register(router)

	testTable := []struct {
		name                 string
		request              string
		expectedStatusCode   int
		expectedResponseBody []byte
	}{
		{
			"positive",
			"/1",
			http.StatusOK,
			TestPositiveResponse,
		},
		{
			"negative",
			"/2",
			http.StatusOK,
			[]byte("Order not found"),
		},
	}

	orders := models.Orders{
		Order:    models.Order{"1", "WBILMTESTTRACK", "WBILl", "en", "sad", "test", "meest", "9", 99, time.Date(2021, time.November, 26, 6, 22, 19, 0, time.UTC), "1"},
		Delivery: models.Delivery{"Test Testov", "+79211552101", "2639809", "Kiryat Mozkin", "Ploshad Mira 15", "Kraiot", "test@gmail.com"},
		Payment:  models.Payment{"b563feb7b2b84b6test", "12345", "USD", "wbpay", 1817, 1637907627, "alpha", 1500, 317, 0},
		Items:    []models.Item{{1, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202}},
	}

	emptyorders := models.Orders{}

	for _, test := range testTable {
		if test.name == "positive" {
			cache.On("GetOrdersById", "1").Return(orders)
			log.On("HandlerLog", mock.AnythingOfType("*http.Request"), test.expectedStatusCode, "return order data with UID 1").Return(0)
		} else {
			cache.On("GetOrdersById", "2").Return(emptyorders)
			log.On("HandlerLog", mock.AnythingOfType("*http.Request"), test.expectedStatusCode, "order with ID 2 not found").Return(0)
		}

		req, err := http.NewRequest("GET", test.request, nil)
		r := httptest.NewRecorder()
		router.ServeHTTP(r, req)
		responseBody := r.Body.Bytes()

		assert.Nil(t, err)
		assert.Equal(t, test.expectedStatusCode, r.Code)
		assert.Equal(t, test.expectedResponseBody, responseBody)
	}
}
