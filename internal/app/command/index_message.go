package command

import (
	"context"

	"github.com/dkeysil/nostr-trends/internal/adapter/indexer"
	"github.com/dkeysil/nostr-trends/internal/common"
	"github.com/dkeysil/nostr-trends/internal/domain"
	"go.uber.org/zap"
)

type (
	IndexMessage common.CommandHandler[domain.Message]
)

type indexMessageHandler struct {
	logger  *zap.Logger
	indexer *indexer.Indexer
}

func NewIndexMessageHandler(logger *zap.Logger, indexer *indexer.Indexer) IndexMessage {
	return &indexMessageHandler{
		logger:  logger,
		indexer: indexer,
	}
}

func (h *indexMessageHandler) Handle(_ context.Context, message domain.Message) error {
	h.logger.Debug("indexing message", zap.String("id", message.ID))

	return h.indexer.IndexMessage(message.ID, domain.Message(message))
}
