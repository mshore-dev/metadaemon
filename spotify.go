package metadaemon

import (
	"encoding/json"
	"net/http"
)

type spotifyTrackMetadata struct {
	Album struct {
		Name   string `json:"name"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	} `json:"album"`
	Artists []struct {
		ID           string `json:"id"`
		Name         string `json:"name"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
	} `json:"artists"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Name string `json:"name"`
}

func (m *MetaDaemon) doSpotifyGet(url string) (*http.Response, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+m.AccessToken)

	return m.httpClient.Do(req)

}

func (m *MetaDaemon) GetSpotifyMetadata(id string) (*ParsedMetadata, error) {

	// https://api.spotify.com/v1/tracks/:id

	resp, err := m.doSpotifyGet("https://api.spotify.com/v1/tracks/" + id)
	if err != nil {
		return &ParsedMetadata{}, err
	}

	var trackMeta spotifyTrackMetadata

	err = json.NewDecoder(resp.Body).Decode(&trackMeta)
	if err != nil {
		return &ParsedMetadata{}, err
	}

	var pm ParsedMetadata

	pm.Name = trackMeta.Name
	pm.URL = trackMeta.ExternalUrls.Spotify
	pm.Album = trackMeta.Album.Name
	pm.Thumbnail = trackMeta.Album.Images[0].URL

	for i := 0; i < len(trackMeta.Artists); i++ {
		pm.Artists = append(pm.Artists, trackMeta.Artists[i].Name)
	}

	return &pm, nil
}
