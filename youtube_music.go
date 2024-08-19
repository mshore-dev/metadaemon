package metadaemon

import (
	"context"
	"encoding/json"
	"os/exec"
	"time"
)

type youtubeMusicTrackMetadata struct {
	Track     string   `json:"track"`
	Thumbnail string   `json:"thumbnail"`
	Album     string   `json:"album"`
	Artists   []string `json:"artists"`
	URL       string   `json:"original_url"`
}

func (m *MetaDaemon) GetYouTubeMusicMetadata(id string) (*ParsedMetadata, error) {

	// https://music.youtube.com/watch?v=:id

	// TODO: don't hardcode timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "yt-dlp", "-j", "https://music.youtube.com/watch?v="+id)

	out, err := cmd.Output()
	if err != nil {
		return &ParsedMetadata{}, err
	}

	var meta youtubeMusicTrackMetadata

	err = json.Unmarshal(out, &meta)
	if err != nil {
		return &ParsedMetadata{}, err
	}

	return &ParsedMetadata{
		URL:       meta.URL,
		Artists:   meta.Artists,
		Name:      meta.Track,
		Album:     meta.Album,
		Thumbnail: meta.Thumbnail,
	}, nil

}
