package helpers

import (
	"github.com/sowens81/primal-audio-manager/pkg/discogs/models"
)

func ConvertDiscogsLabelsToListString(labels []models.Label) []string {
	result := make([]string, 0, len(labels))

	for _, l := range labels {
		if l.Name != "" {
			result = append(result, l.Name)
		}
	}

	return result
}

func ConvertDiscogsCatNoToListString(labels []models.Label) []string {
	result := make([]string, 0, len(labels))

	for _, l := range labels {
		if l.CatNo != "" {
			result = append(result, l.CatNo)
		}
	}

	return result
}
