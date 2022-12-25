package service

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/dkeysil/nostr-trends/internal/adapter/indexer"
	"github.com/dkeysil/nostr-trends/internal/app"
	"github.com/dkeysil/nostr-trends/internal/app/command"
	"github.com/dkeysil/nostr-trends/internal/app/query"
	"github.com/dkeysil/nostr-trends/internal/config"
	"github.com/dkeysil/nostr-trends/internal/ports"
	"github.com/dkeysil/nostr-trends/internal/workers"
	"github.com/nbd-wtf/go-nostr"
	"go.uber.org/zap"
)

func RunApplication(cfg config.Config) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	var index bleve.Index

	_, err = os.Stat(cfg.IndexFolder)
	if os.IsNotExist(err) {
		index, err = bleve.New(cfg.IndexFolder, bleve.NewIndexMapping())
		if err != nil {
			logger.Fatal("failed to open index", zap.Error(err))
		}
	} else {
		index, err = bleve.Open(cfg.IndexFolder)
		if err != nil {
			logger.Fatal("failed to open index", zap.Error(err))
		}
	}

	indexer := indexer.New(index)

	queries := app.Queries{
		WordTrends: query.NewWordTrendsHandler(logger, indexer),
	}

	commands := app.Commands{
		IndexMessage: command.NewIndexMessageHandler(logger, indexer),
	}

	pool := nostr.NewRelayPool()
	errCh := pool.Add(cfg.NostrRelayURL, nostr.SimplePolicy{Read: true})
	if err = <-errCh; err != nil {
		logger.Fatal("failed to connect to relay", zap.Error(err))
	}

	nostrMessages := workers.NewNostrMessagesWorker(logger, pool, commands.IndexMessage)
	go nostrMessages.Run()

	logger.Info("starting server")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.HTTPPort), ports.Handler(queries)); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
