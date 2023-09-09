package main

import (
	"fmt"
	"github.com/example/internal/config"
)

func main() {
	cfg := config.GetConfig("config.yml")
	fmt.Printf("%+v\n", cfg)
}
