package metadaemon

import "testing"

func TestSpotifyGetMetadata(t *testing.T) {

	meta, err := md.GetSpotifyMetadata("6kNHVx9UxZRoWhRD3947W5")
	if err != nil {
		t.Fatalf("failed to get track metadata: %v\n", err)
	}

	if meta.Name != "Machine Heart" {
		t.Fatalf("fetched track name did not match expected")
	}

	if meta.URL != "https://open.spotify.com/track/6kNHVx9UxZRoWhRD3947W5" {
		t.Fatalf("fetched track url did not match expected")
	}

	if meta.Artists[0] != "Icarus" {
		t.Fatal("fetched artist 0 did not match expected")
	}

	t.Logf("%v\n", meta)

}
