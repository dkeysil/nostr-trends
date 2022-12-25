package domain

import "time"

// Message represents a internal nostr message struct.
type Message struct {
	ID        string    `json:"id"`
	Content   string    `json:"content"`
	Relay     string    `json:"relay"`
	PubKey    string    `json:"pubkey"`
	CreatedAt time.Time `json:"created_at"`
	Kind      int       `json:"kind"`
}
