package loader

import (
	"context"
	"log"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/go-resty/resty/v2"
	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/ent/character"
	"github.com/samuelagm/moviex/ent/predicate"
	"github.com/samuelagm/moviex/internal/loader/types"
)

var httpC *resty.Client

func init() {
	httpC = resty.New()
	httpC.SetBaseURL("https://swapi.dev/api")
}

func Load(ctx context.Context, entC *ent.Client) {
	downloadCharacters(ctx, entC, "1")
	downloadFilms(ctx, entC)
}

func downloadCharacters(
	ctx context.Context,
	entC *ent.Client,
	pageNumber string,
) error {
	resp, err := httpC.R().
		SetHeader("Accept", "application/json").
		SetQueryParam("page", pageNumber).
		SetResult(types.CharacterQueryResult{}).
		Get("/people")
	if err != nil {
		log.Fatal(err)
	}

	result := resp.Result().(*types.CharacterQueryResult)

	for _, c := range result.Results {
		if _, err := entC.Character.
			Create().SetName(c.Name).
			SetHeight(c.Height).
			SetMass(c.Mass).
			SetHairColor(c.HairColor).
			SetSkinColor(c.SkinColor).
			SetEyeColor(c.EyeColor).
			SetGender(c.Gender).
			SetEdited(c.Edited).
			SetBirthYear(c.BirthYear).
			SetFilms(c.Films).
			SetURL(c.URL).
			SetCreated(c.Created).
			Save(ctx); err != nil {
			log.Fatal(err)
		}
	}

	if len(result.Next) != 0 {
		if parsedURL, err := url.Parse(result.Next); err == nil {
			nextPage := parsedURL.Query()["page"][0]
			downloadCharacters(ctx, entC, nextPage)
		} else {
			log.Fatal(err)
		}
	}

	return nil
}

func downloadFilms(ctx context.Context, entC *ent.Client) error {
	resp, err := httpC.R().
		SetHeader("Accept", "application/json").
		SetResult(types.FilmsQueryResult{}).
		Get("/films")

	if err != nil {
		log.Fatal(err)
	}

	for _, m := range resp.Result().(*types.FilmsQueryResult).Results {
		if characters, err := entC.Character.Query().
			Where(
				predicate.Character(func(s *sql.Selector) {
					s.Where(sqljson.ValueContains(character.FieldFilms, m.URL))
				}),
			).
			All(ctx); err == nil {
			if _, err = entC.Movie.
				Create().
				SetTitle(m.Title).
				SetEpisodeID(m.EpisodeID).
				SetOpeningCrawl(m.OpeningCrawl).
				SetDirector(m.Director).
				SetProducer(m.Producer).
				SetReleaseDate(time.Time(m.ReleaseDate)).
				SetCharacters(m.Characters).
				SetCreated(m.Created).
				SetEdited(m.Edited).
				SetURL(m.URL).
				AddPeople(characters...).
				Save(ctx); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	return nil
}
