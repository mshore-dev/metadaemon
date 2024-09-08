package metadaemon

import "testing"

// TODO: test Spotify matcher, just to be sure it works here.
func TestMatchYouTube(t *testing.T) {
	meta, err := md.Match("https://www.youtube.com/watch?v=3hEwM6o9k4c")
	if err != nil {
		t.Fatalf("failed to get track metadata: %v\n", err)
	}

	if meta.Name != "少女レイ/歌ってみた【ゆーり】" {
		t.Fatalf("fetched track name did not match expected")
	}

	if meta.URL != "https://www.youtube.com/watch?v=3hEwM6o9k4c" {
		t.Fatalf("fetched track url did not match expected")
	}

	if meta.Artists[0] != "ゆーり🍎🐥〔23〕・yuri" {
		t.Fatal("fetched artist 0 did not match expected")
	}
}
