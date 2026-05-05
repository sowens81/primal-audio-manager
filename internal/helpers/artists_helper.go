package helpers

import (
	"strings"

	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

func ConvertDiscogsArtistToString(artists []models.Artist) string {
	if len(artists) == 0 {
		return ""
	}

	var result strings.Builder

	for i, artist := range artists {
		result.WriteString(artist.Name)

		if artist.Join != "" {
			result.WriteString(" ")
			result.WriteString(artist.Join)
			result.WriteString(" ")
		} else if i < len(artists)-1 {
			result.WriteString(" ")
		}
	}

	return strings.TrimSpace(result.String())
}
