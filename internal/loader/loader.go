package loader

import (
	"context"
	"crypto/tls"
	"log"
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
	// Skip TLS verification for development
	httpC.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	httpC.SetBaseURL("https://swapi.info/api")
}

func Load(ctx context.Context, entC *ent.Client) {
	downloadCharacters(ctx, entC)
	downloadFilms(ctx, entC)
}

func downloadCharacters(
	ctx context.Context,
	entC *ent.Client,
) error {
	resp, err := httpC.R().
		SetHeader("Accept", "application/json").
		SetResult([]types.Character{}).
		Get("/people")
	if err != nil {
		log.Fatal(err)
	}

	characters := resp.Result().(*[]types.Character)

	for _, c := range *characters {
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

	return nil
}

func downloadFilms(ctx context.Context, entC *ent.Client) error {
	resp, err := httpC.R().
		SetHeader("Accept", "application/json").
		SetResult([]types.Film{}).
		Get("/films")

	if err != nil {
		log.Fatal(err)
	}

	films := resp.Result().(*[]types.Film)
	for _, m := range *films {
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
				SetURL(m.URL).
				SetCreated(m.Created).
				SetEdited(m.Edited).
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
