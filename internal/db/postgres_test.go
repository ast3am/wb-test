package db

import (
	"context"
	"github.com/ast3am/wb-test/internal/db/mocks"
	"github.com/ast3am/wb-test/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

var testCfg = models.Config{
	ListenPort: "",
	SqlConfig: models.SqlConfig{
		UsernameDB: "postgres",
		PasswordDB: "password",
		HostDB:     "localhost",
		PortDB:     "5430",
		DBName:     "WB_db_test",
	},
	NatsConfig: models.NatsConfig{ChannelName: "", ClusterID: "", ClientID: "", Host: "", Port: ""},
	LogLevel:   "",
}

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)

	testTable := []struct {
		name string
		cfg  models.Config
	}{
		{"positive",
			testCfg,
		},
		{"negative",
			models.Config{
				ListenPort: "",
				SqlConfig: models.SqlConfig{
					UsernameDB: "postgres",
					PasswordDB: "password",
					HostDB:     "localhost",
					PortDB:     "54320",
					DBName:     "WB_db_test",
				},
				NatsConfig: models.NatsConfig{ChannelName: "", ClusterID: "", ClientID: "", Host: "", Port: ""},
				LogLevel:   "",
			},
		},
	}
	for _, test := range testTable {
		if test.name == "positive" {
			log.On("DebugMsg", "connection to DB is OK").Return()
			db, _ := NewClient(ctx, &test.cfg, log)
			require.NotNil(t, db, "")
			defer db.dbConnect.Close(ctx)
		} else {
			db, _ := NewClient(ctx, &test.cfg, log)
			assert.Nil(t, db, "")
		}
	}
}

func TestDB_InsertOrders(t *testing.T) {

	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	testTable := []struct {
		name      string
		testOrder models.Orders
	}{
		{"positive",
			models.Orders{
				Order:    models.Order{"1", "WBILMTESTTRACK", "WBILl", "en", "sad", "test", "meest", "9", 99, time.Date(2021, time.November, 26, 6, 22, 10, 0, time.UTC), "1"},
				Delivery: models.Delivery{"Test Testov", "+79211552101", "2639809", "Kiryat Mozkin", "Ploshad Mira 15", "Kraiot", "test@gmail.com"},
				Payment:  models.Payment{"b563feb7b2b84b6test", "12345", "USD", "wbpay", 1817, 1637907627, "alpha", 1500, 317, 0},
				Items: []models.Item{{1, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202},
					{2, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202}},
			},
		},
		{
			"negative",
			models.Orders{
				Order:    models.Order{"1", "WBILMTESTTRACK", "WBILl", "en", "sad", "test", "meest", "9", 99, time.Date(2021, time.November, 26, 6, 22, 10, 0, time.UTC), "1"},
				Delivery: models.Delivery{"Test Testov", "+79211552101", "2639809", "Kiryat Mozkin", "Ploshad Mira 15", "Kraiot", "test@gmail.com"},
				Payment:  models.Payment{"b563feb7b2b84b6test", "12345", "USD", "wbpay", 1817, 1637907627, "alpha", 1500, 317, 0},
				Items: []models.Item{{1, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202},
					{2, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202}},
			},
		},
	}

	for _, test := range testTable {
		if test.name == "positive" {
			err := db.InsertOrders(ctx, test.testOrder)
			assert.Nil(t, err, "")
		} else {
			err := db.InsertOrders(ctx, test.testOrder)
			assert.NotNil(t, err, "")
		}
	}
}

func TestDB_GetOrders(t *testing.T) {
	ctx := context.Background()
	log := mocks.NewLogger(t)
	log.On("DebugMsg", "connection to DB is OK").Return()
	db, _ := NewClient(ctx, &testCfg, log)
	defer db.dbConnect.Close(ctx)

	expected := make([]*models.Orders, 1, 1)
	expected[0] = &models.Orders{
		Order:    models.Order{"1", "WBILMTESTTRACK", "WBILl", "en", "sad", "test", "meest", "9", 99, time.Date(2021, time.November, 26, 6, 22, 10, 0, time.UTC), "1"},
		Delivery: models.Delivery{"Test Testov", "+79211552101", "2639809", "Kiryat Mozkin", "Ploshad Mira 15", "Kraiot", "test@gmail.com"},
		Payment:  models.Payment{"b563feb7b2b84b6test", "12345", "USD", "wbpay", 1817, 1637907627, "alpha", 1500, 317, 0},
		Items: []models.Item{{1, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202},
			{2, "WBILMTESTTRACK", 453, "ab4319087a764ae0btest", "Mascaras", 30, "0", 317, 2389232, "Vivienne Sabo", 202}},
	}

	result, err := db.GetOrders(ctx)
	assert.NotNil(t, result, "")
	require.Nil(t, err, "")
	for i, k := range expected {
		assert.Equal(t, k, result[i])
	}
}
