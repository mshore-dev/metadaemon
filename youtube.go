package metadaemon

import (
	"context"
	"encoding/json"
	"os/exec"
	"time"
)

type youtubeTrackMetadata struct {
	Title     string `json:"title"`
	Thumbnail string `json:"thumbnail"`
	Uploader  string `json:"uploader"`
	URL       string `json:"original_url"`
}

func (m *MetaDaemon) GetYouTubeMetadata(id string) (*ParsedMetadata, error) {

	// https://www.youtube.com/watch?v=:id

	// TODO: don't hardcode timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "yt-dlp", "-j", "https://www.youtube.com/watch?v="+id)

	out, err := cmd.Output()
	if err != nil {
		return &ParsedMetadata{}, err
	}

	var meta youtubeTrackMetadata

	err = json.Unmarshal(out, &meta)
	if err != nil {
		return &ParsedMetadata{}, err
	}

	return &ParsedMetadata{
		URL: meta.URL,
		Artists: []string{
			meta.Uploader,
		},
		Name:      meta.Title,
		Thumbnail: meta.Thumbnail,
	}, nil

}
