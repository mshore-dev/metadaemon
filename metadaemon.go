package metadaemon

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	YTDLPBinary         string
	SpotifyClientID     string
	SpotifyClientSecret string
}

type ParsedMetadata struct {
	URL string `json:"url"`
	// ReleaseDate time.Time
	Artists   []string `json:"artists"`
	Name      string   `json:"name"`
	Album     string   `json:"album"`
	Thumbnail string   `json:"thumbnail"`
}

type spotifyAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

type MetaDaemon struct {
	AccessToken string
	TokenExpiry time.Time

	Config *Config

	httpClient *http.Client
}

func New(config Config) *MetaDaemon {
	return &MetaDaemon{
		Config:     &config,
		httpClient: &http.Client{},
	}
}

// authenticates with any metadata provider that requires it
func (m *MetaDaemon) Authorize() error {

	// currently, this is only Spotify
	resp, err := http.PostForm("https://accounts.spotify.com/api/token", url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {m.Config.SpotifyClientID},
		"client_secret": {m.Config.SpotifyClientSecret},
	})

	if err != nil {
		return err
	}

	var authorzieResponse spotifyAccessToken

	err = json.NewDecoder(resp.Body).Decode(&authorzieResponse)
	if err != nil {
		return err
	}

	m.AccessToken = authorzieResponse.AccessToken
	m.TokenExpiry = time.Now().Add(time.Duration(authorzieResponse.ExpiresIn * int64(time.Second)))

	return nil
}
