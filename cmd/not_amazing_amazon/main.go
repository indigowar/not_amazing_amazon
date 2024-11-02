package main

import (
	"log"

	"github.com/indigowar/not_amazing_amazon/internal/common/app"
	"github.com/indigowar/not_amazing_amazon/internal/common/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
