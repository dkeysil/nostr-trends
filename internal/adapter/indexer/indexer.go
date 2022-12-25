package indexer

import (
	"errors"
	"time"

	"github.com/blevesearch/bleve"
	"github.com/dkeysil/nostr-trends/internal/domain"
)

type Indexer struct {
	index bleve.Index
}

func New(index bleve.Index) *Indexer {
	return &Indexer{
		index: index,
	}
}

// IndexMessage indexes a message, where id is nostr message id
func (i *Indexer) IndexMessage(id string, message domain.Message) error {
	return i.index.Index(id, message)
}

type Words struct {
	Word  string
	Count int
}

// GetWordsCount
func (i *Indexer) GetWordsCount(count int, from time.Time, to time.Time) ([]Words, error) {
	query := bleve.NewDateRangeQuery(from, to)

	searchRequest := bleve.NewSearchRequest(query)

	facetRequest := bleve.NewFacetRequest("top_words", count)
	facetRequest.Field = "content"

	searchRequest.AddFacet("top_words", facetRequest)

	searchResult, err := i.index.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	facet, ok := searchResult.Facets["top_words"]
	if !ok {
		return nil, errors.New("top_words aggregation not found")
	}

	words := make([]Words, 0, len(facet.Terms))

	for _, term := range facet.Terms {
		words = append(words, struct {
			Word  string
			Count int
		}{
			Word:  term.Term,
			Count: term.Count,
		})
	}

	return words, nil
}
