package query

import (
	"context"
	"time"

	"github.com/dkeysil/nostr-trends/internal/adapter/indexer"
	"github.com/dkeysil/nostr-trends/internal/common"
	"go.uber.org/zap"
)

type (
	Params struct {
		Count int
	}

	Words []struct {
		Word  string
		Count int
	}

	WordTrends common.QueryHandler[Params, *Words]
)

type wordTrendsHandler struct {
	indexer *indexer.Indexer
	logger  *zap.Logger
}

func NewWordTrendsHandler(logger *zap.Logger, indexer *indexer.Indexer) WordTrends {
	return &wordTrendsHandler{
		indexer: indexer,
		logger:  logger,
	}
}

func (h *wordTrendsHandler) Handle(ctx context.Context, params Params) (*Words, error) {
	h.logger.Debug("searching for top words", zap.Int("count", params.Count))
	words, err := h.indexer.GetWordsCount(10, time.Now().Add(-24*time.Hour), time.Now())

	if err != nil {
		return nil, err
	}

	responseWords := make(Words, 0, len(words))
	for _, word := range words {
		responseWords = append(responseWords, struct {
			Word  string
			Count int
		}{
			Word:  word.Word,
			Count: word.Count,
		})
	}

	return &responseWords, nil
}
