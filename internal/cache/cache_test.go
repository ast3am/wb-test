package cache

import (
	"github.com/ast3am/wb-test/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockLogger struct {
	debugMsg []string
}

func (m *mockLogger) DebugMsg(msg string) {
	m.debugMsg = append(m.debugMsg, msg)
}

// т.к. данный пакет не валидирует данные, нам будет достаточно неполных данных для теста
var testsData = []models.Orders{
	{
		Order:    models.Order{OrderUid: "1"},
		Delivery: models.Delivery{Name: "test 1 name"},
		Payment:  models.Payment{Transaction: "test 1 transaction"},
		Items:    []models.Item{{ChrtID: 15}, {ChrtID: 12}},
	}, {
		Order:    models.Order{OrderUid: "2"},
		Delivery: models.Delivery{Name: "test 2 name"},
		Payment:  models.Payment{Transaction: "test 2 transaction"},
		Items:    []models.Item{{ChrtID: 15}, {ChrtID: 12}},
	},
}

func TestCacheInit(t *testing.T) {
	mockLog := &mockLogger{}
	cache := CacheInit(mockLog)

	assert.NotNil(t, cache, "assert not nil cache")
	assert.Equal(t, "init cache is OK", mockLog.debugMsg[0], "assert 'init cache is OK' msg")
}

func TestInsertOrdersToCache(t *testing.T) {
	mockLog := &mockLogger{}
	cache := CacheInit(mockLog)

	for _, test := range testsData {
		cache.InsertOrdersToCache(test)
		require.Equal(t, test, cache.cache[test.Order.OrderUid])
	}
}

func TestGetOrdersById(t *testing.T) {
	mockLog := &mockLogger{}
	cache := CacheInit(mockLog)

	for _, test := range testsData {
		cache.InsertOrdersToCache(test)
		require.Equal(t, test, cache.GetOrdersById(test.Order.OrderUid))
	}
}
