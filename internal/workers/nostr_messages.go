package workers

import (
	"context"

	"github.com/dkeysil/nostr-trends/internal/app/command"
	"github.com/dkeysil/nostr-trends/internal/domain"
	"github.com/nbd-wtf/go-nostr"
	"go.uber.org/zap"
)

type NostrMessagesWorker struct {
	logger *zap.Logger

	pool                *nostr.RelayPool
	indexMessageCommand command.IndexMessage
}

func NewNostrMessagesWorker(
	logger *zap.Logger,
	pool *nostr.RelayPool,
	indexMessageCommand command.IndexMessage,
) *NostrMessagesWorker {
	return &NostrMessagesWorker{
		logger:              logger,
		pool:                pool,
		indexMessageCommand: indexMessageCommand,
	}
}

func (w *NostrMessagesWorker) Run() {
	subID, events := w.pool.Sub(nostr.Filters{{
		Kinds: []int{
			nostr.KindTextNote,
		}},
	})

	w.logger.Debug("starting listenting for messages", zap.String("subID", subID))

	for event := range events {
		w.logger.Debug(
			"got event",
			zap.String("subID", subID),
			zap.String("content", event.Event.Content),
			zap.Time("createdAt", event.Event.CreatedAt),
		)

		message := domain.Message{
			ID:        event.Event.ID,
			Content:   event.Event.Content,
			Relay:     event.Relay,
			PubKey:    event.Event.PubKey,
			Kind:      event.Event.Kind,
			CreatedAt: event.Event.CreatedAt,
		}
		w.indexMessageCommand.Handle(context.Background(), message)
	}
}
