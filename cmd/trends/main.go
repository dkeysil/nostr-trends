package main

import (
	"github.com/dkeysil/nostr-trends/internal/config"
	"github.com/dkeysil/nostr-trends/internal/service"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	var cfg config.Config
	envconfig.MustProcess("trends", &cfg)

	service.RunApplication(cfg)
}
