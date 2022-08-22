package types

import (
	"time"

	"github.com/samuelagm/moviex/internal/common/types"
)

type FilmsQueryResult struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Film      `json:"results"`
}

type CharacterQueryResult struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Character `json:"results"`
}

type Film struct {
	Title        string            `json:"title"`
	EpisodeID    int               `json:"episode_id"`
	OpeningCrawl string            `json:"opening_crawl"`
	Director     string            `json:"director"`
	Producer     string            `json:"producer"`
	ReleaseDate  types.ReleaseDate `json:"release_date"`
	Characters   []string          `json:"characters"`
	Planets      []string          `json:"planets"`
	Starships    []string          `json:"starships"`
	Vehicles     []string          `json:"vehicles"`
	Species      []string          `json:"species"`
	Created      time.Time         `json:"created"`
	Edited       time.Time         `json:"edited"`
	URL          string            `json:"url"`
}

type Character struct {
	Name      string        `json:"name"`
	Height    string        `json:"height"`
	Mass      string        `json:"mass"`
	HairColor string        `json:"hair_color"`
	SkinColor string        `json:"skin_color"`
	EyeColor  string        `json:"eye_color"`
	BirthYear string        `json:"birth_year"`
	Gender    string        `json:"gender"`
	Homeworld string        `json:"homeworld"`
	Films     []string      `json:"films"`
	Species   []interface{} `json:"species"`
	Vehicles  []string      `json:"vehicles"`
	Starships []string      `json:"starships"`
	Created   time.Time     `json:"created"`
	Edited    time.Time     `json:"edited"`
	URL       string        `json:"url"`
}
