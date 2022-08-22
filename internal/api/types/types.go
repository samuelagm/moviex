package types

import (
	"context"
	"time"

	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/internal/common/types"
)

type ApiHelper struct {
	EntClient *ent.Client
	Context   context.Context
}

type ErrorResponse struct {
	Message string
}

type FilmResponse struct {
	Title        string            `json:"title"`
	EpisodeID    int               `json:"episode_id"`
	OpeningCrawl string            `json:"opening_crawl"`
	Director     string            `json:"director"`
	Producer     string            `json:"producer"`
	ReleaseDate  types.ReleaseDate `json:"release_date"`
	Characters   []string          `json:"characters"`
	Created      time.Time         `json:"created"`
	Edited       time.Time         `json:"edited"`
	URL          string            `json:"url"`
}

type Comment struct {
	Name string `json:"title"`
	Text string `json:"text"`
}

type CommentResponse struct {
	Name    string    `json:"title"`
	Text    string    `json:"text"`
	IP      string    `json:"ip"`
	Created time.Time `json:"created"`
}

type CharacterResponse struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Films     []string  `json:"films"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
}
