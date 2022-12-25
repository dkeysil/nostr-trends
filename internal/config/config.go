package config

type Config struct {
	// IndexFolder path to the folder where the index will be stored.
	IndexFolder string `envconfig:"INDEX_FOLDER"`
	// NostrRelayURL URL of the Nostr relay to connect to.
	NostrRelayURL string `envconfig:"NOSTR_RELAY_URL" default:"wss://relay.damus.io"`
	// HTTPPort port on which the HTTP server will listen.
	HTTPPort int `envconfig:"HTTP_PORT"       default:"8000"`
}
