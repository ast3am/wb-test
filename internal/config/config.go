package config

import (
	"github.com/ast3am/wb-test/internal/models"
	"github.com/ast3am/wb-test/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg *models.Config

func GetConfig(path string) *models.Config {
	logger := logging.GetLogger("")
	logger.InfoMsg("read config")
	cfg = &models.Config{}
	err := cleanenv.ReadConfig(path, cfg)
	if err != nil {
		logger.FatalMsg("", err)
	}
	return cfg
}
