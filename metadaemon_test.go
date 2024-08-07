package metadaemon

import (
	"os"
	"testing"
)

var md *MetaDaemon

func TestAuthorize(t *testing.T) {
	md = New(Config{
		SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
	})

	err := md.Authorize()
	if err != nil {
		t.Fatalf("failed to authorize: %v\n", err)
	}
}
