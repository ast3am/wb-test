package nats_streaming

import (
	"github.com/ast3am/wb-test/internal/models"
	"github.com/ast3am/wb-test/internal/nats_streaming/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

var testCfg = models.Config{
	ListenPort: "",
	SqlConfig:  models.SqlConfig{UsernameDB: "", PasswordDB: "", HostDB: "", PortDB: "", DBName: ""},
	NatsConfig: models.NatsConfig{
		ChannelName: "my-channel",
		ClusterID:   "test-cluster",
		ClientID:    "my-client2",
		Host:        "localhost",
		Port:        "4220",
	},
	LogLevel: "",
}

func TestInitNats(t *testing.T) {
	w := mocks.NewWriter(t)
	log := mocks.NewLogger(t)
	nats := InitNats(w, log)
	require.NotNil(t, nats, "")
}

func TestNats_Connect(t *testing.T) {
	w := mocks.NewWriter(t)
	log := mocks.NewLogger(t)
	nats := InitNats(w, log)

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
				SqlConfig:  models.SqlConfig{UsernameDB: "", PasswordDB: "", HostDB: "", PortDB: "", DBName: ""},
				NatsConfig: models.NatsConfig{
					ChannelName: "my-channel",
					ClusterID:   "test-cluster",
					ClientID:    "my-client2",
					Host:        "localhost",
					Port:        "42250",
				},
				LogLevel: "",
			},
		},
	}

	for _, test := range testTable {
		if test.name == "positive" {
			log.On("DebugMsg", "connect to nuts is OK").Return()
			err := nats.Connect(&testCfg)
			defer nats.NatsConnection.Close()
			require.Nil(t, err)
			require.NotNil(t, nats.NatsConnection)
		} else {
			err := nats.Connect(&testCfg)
			require.NotNil(t, err)
		}
	}
}

func TestNats_Subscribe(t *testing.T) {
	w := mocks.NewWriter(t)
	log := mocks.NewLogger(t)

	log.On("DebugMsg", "connect to nuts is OK").Return()
	nats := InitNats(w, log)
	nats.Connect(&testCfg)

	log.On("DebugMsg", "subscribe to channel "+testCfg.NatsConfig.ChannelName+" is OK").Return()
	err := nats.Subscribe("my-channel")
	require.Nil(t, err)
}
