package app

import (
	"github.com/dkeysil/nostr-trends/internal/app/command"
	"github.com/dkeysil/nostr-trends/internal/app/query"
)

type Queries struct {
	query.WordTrends
}

type Commands struct {
	command.IndexMessage
}
