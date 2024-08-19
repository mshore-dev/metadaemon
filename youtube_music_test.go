package metadaemon

import "testing"

func TestYouTubeMusicGetMetadata(t *testing.T) {

	meta, err := md.GetYouTubeMusicMetadata("uQTBzmBDSv0")
	if err != nil {
		t.Fatalf("failed to get track metadata: %v\n", err)
	}

	if meta.Name != "The Town Inside Me" {
		t.Fatalf("fetched track name did not match expected")
	}

	if meta.URL != "https://music.youtube.com/watch?v=uQTBzmBDSv0" {
		t.Fatalf("fetched track url did not match expected")
	}

	if meta.Artists[0] != "AISHA" {
		t.Fatal("fetched artist 0 did not match expected")
	}

	t.Logf("%v\n", meta)

}
