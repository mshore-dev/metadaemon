package metadaemon

import (
	"errors"
	"regexp"
)

var (
	reSpotifyLink      = regexp.MustCompile(`https://open.spotify.com/track/(\w{22})`)
	reYouTubeMusicLink = regexp.MustCompile(`https://music.youtube.com/watch\?v=([a-zA-Z0-9_-]{11})`)
	reYouTubeLink      = regexp.MustCompile(`https://(?:(?:www.youtube.com/watch\?v=)|(?:youtu.be/))([a-zA-Z0-9_-]{11})`)

	errNoMatcher = errors.New("provided url had no matching handler")
)

func (m *MetaDaemon) Match(url string) (*ParsedMetadata, error) {

	var match []string

	match = reYouTubeMusicLink.FindStringSubmatch(url)
	if len(match) != 0 {
		// log.Printf("got ytm url with id %s\n", match[1])
		return m.GetYouTubeMusicMetadata(match[1])
	}

	match = reYouTubeLink.FindStringSubmatch(url)
	if len(match) != 0 {
		// log.Printf("got yt url with id %s\n", match[1])
		return m.GetYouTubeMetadata(match[1])
	}

	match = reSpotifyLink.FindStringSubmatch(url)
	if len(match) != 0 {
		// log.Printf("got spotify url with id %s\n", match[1])
		return m.GetSpotifyMetadata(match[1])
	}

	return &ParsedMetadata{}, errNoMatcher

}
