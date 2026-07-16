package practice

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed catalog.json
var catalogJSON []byte

type CatalogExercise struct {
	Day          int    `json:"day"`
	Chapter      int    `json:"chapter"`
	Exercise     int    `json:"exercise"`
	Folder       string `json:"folder"`
	ChapterTitle string `json:"chapter_title"`
	Prompt       string `json:"prompt"`
}

func LoadCatalog() ([]CatalogExercise, error) {
	var out []CatalogExercise
	if err := json.Unmarshal(catalogJSON, &out); err != nil {
		return nil, fmt.Errorf("parse embedded catalog: %w", err)
	}
	return out, nil
}
