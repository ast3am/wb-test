package writer

import (
	"context"
	"github.com/ast3am/wb-test/internal/models"
	"github.com/ast3am/wb-test/internal/writer/mocks"
	"github.com/ast3am/wb-test/internal/writer/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewWriter(t *testing.T) {
	db := mocks.NewDb(t)
	cache := mocks.NewCache(t)
	log := mocks.NewLogger(t)
	writer := NewWriter(db, cache, log)
	require.NotNil(t, writer)
}

func TestWriter_GetCacheOnStart(t *testing.T) {
	ctx := context.Background()
	db := mocks.NewDb(t)
	cache := mocks.NewCache(t)
	log := mocks.NewLogger(t)

	DBexpected := make([]*models.Orders, 1, 1)
	DBexpected[0] = &models.Orders{
		Order:    models.Order{OrderUid: "1"},
		Delivery: models.Delivery{Name: "test 1 name"},
		Payment:  models.Payment{Transaction: "test 1 transaction"},
		Items:    []models.Item{{ChrtID: 15}, {ChrtID: 12}},
	}
	db.On("GetOrders", ctx).Return(DBexpected, nil)
	cache.On("InsertOrdersToCache", *DBexpected[0]).Return()

	w := NewWriter(db, cache, log)
	err := w.GetCacheOnStart(ctx)
	require.Nil(t, err)
}

func TestWriter_Write(t *testing.T) {
	ctx := context.Background()
	db := mocks.NewDb(t)
	cache := mocks.NewCache(t)
	log := mocks.NewLogger(t)
	w := NewWriter(db, cache, log)
	testData := []byte(`{"OrderUid": "123", ...}`)
	err := w.Write(testData)

	testTable := []struct {
		name      string
		testOrder []byte
	}{{
		"positive",
		test_data.TestPositiveData,
	}, {
		"negative",
		[]byte(`{"OrderUid": "123", ...}`),
	}, {
		"negative",
		test_data.TestNegativeData1,
	}, {
		"negative",
		test_data.TestNegativeData2,
	}, {
		"negative",
		test_data.TestNegativeData3,
	}, {
		"negative",
		test_data.TestNegativeData4,
	},
	}

	orders := models.Orders{
		Order:    models.Order{"1", "WBILMTESTTRACK", "WBILl", "en", "sad", "test", "meest", "9", 99, time.Date(2021, time.November, 26, 6, 22, 19, 0, time.UTC), "1"},
		Delivery: models.Delivery{"Test Testov", "+79211552101", "2639809", "Kiryat Mozkin", "Ploshad Mira 15", "Kraiot", "test@gmail.com"},
		Payment:  models.Payment{"b563feb7b2b84b6test", "12345", "USD", "wbpay", 1817, 1637907627, "alpha", 1500, 317, 0},
		Items:    []models.Item{{1, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202}},
	}

	for _, test := range testTable {
		if test.name == "positive" {
			db.On("InsertOrders", ctx, orders).Return(nil)
			cache.On("InsertOrdersToCache", orders).Return(0)
			log.On("InfoMsgf", "added new order with UID %s", orders.Order.OrderUid).Return(0)
			err := w.Write(test.testOrder)
			assert.Nil(t, err, "")
		} else {
			err := w.Write(test.testOrder)
			assert.NotNil(t, err, "")
		}
	}

	require.NotNil(t, err)
}
