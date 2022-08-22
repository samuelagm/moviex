package types

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/ent/character"
	"github.com/samuelagm/moviex/ent/movie"
	"github.com/samuelagm/moviex/ent/predicate"
	commontypes "github.com/samuelagm/moviex/internal/common/types"
)

var sortFn map[string]func(...string) ent.OrderFunc
var sortField map[string]string
var filterField map[string]string

func init() {
	sortFn = map[string]func(...string) ent.OrderFunc{
		"asc":  ent.Asc,
		"desc": ent.Desc,
	}
	sortField = map[string]string{
		"name":   character.FieldName,
		"height": character.FieldHeight,
		"gender": character.FieldGender,
	}
	filterField = map[string]string{
		"male":   "male",
		"female": "female",
	}
}

func NewApiHelper(ctx context.Context, dbClient *ent.Client) *ApiHelper {
	return &ApiHelper{
		EntClient: dbClient,
		Context:   ctx,
	}
}

// @BasePath /api/v1

// Movies 			godoc
// @Summary 		list all movies
// @Schemes
// @Description 	Returns a list of all star wars movies
// @Accept       	json
// @Produce 		json
// @Success 		200 {array} FilmResponse
// @Failure      	500  {object}  ErrorResponse
// @Router 			/movies [get]
func (h *ApiHelper) Movies(gctx *gin.Context) {
	if movies, err := h.EntClient.Movie.Query().
		Order((ent.Asc(movie.FieldReleaseDate))).
		All(h.Context); err == nil {
		result := []FilmResponse{}
		for _, m := range movies {
			result = append(result, FilmResponse{
				Title:        m.Title,
				EpisodeID:    m.EpisodeID,
				OpeningCrawl: m.OpeningCrawl,
				Director:     m.Director,
				Producer:     m.Producer,
				ReleaseDate:  commontypes.ReleaseDate(m.ReleaseDate),
				Characters:   m.Characters,
				Created:      m.Created,
				Edited:       m.Edited,
				URL:          m.URL,
			})
		}
		gctx.JSON(http.StatusOK, result)
	} else {
		log.Println(err)
		gctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "something went wrong",
		})
	}
}

// @BasePath /api/v1

// Movies 			godoc
// @Summary 		list all characters
// @Schemes
// @Description 	Returns a list of all star wars characters
// @Accept       	json
// @Produce 		json
// @Success 		200 {array} CharacterResponse
// @Failure      	500  {object}  ErrorResponse
// @Router 			/people [get]
func (h *ApiHelper) Characters(gctx *gin.Context) {

	sortOp := ent.Asc(character.FieldName)
	filterOp := predicate.Character(func(s *sql.Selector) {})

	if sort, sortOk := gctx.GetQuery("sort"); sortOk {
		sComponents := strings.Split(sort, " ")
		if fn, fnOk := sortFn[sComponents[1]]; fnOk {
			if field, Ok := sortField[sComponents[0]]; Ok {
				sortOp = fn(field)
			}
		}
	}

	if field, filterOk := gctx.GetQuery("filter"); filterOk {
		if field, Ok := filterField[field]; Ok {
			filterOp = character.GenderEQ(field)
		} else {
			filterOp = character.GenderEQ("*$@!0p") //produce an empty result
		}
	}

	if movies, err := h.EntClient.Character.Query().
		Where(filterOp).
		Order(sortOp).
		All(h.Context); err == nil {
		result := []CharacterResponse{}
		for _, c := range movies {
			result = append(result, CharacterResponse{
				Name:      c.Name,
				Height:    commontypes.HtoFeetInces(c.Height),
				Mass:      c.Mass,
				HairColor: c.HairColor,
				SkinColor: c.SkinColor,
				EyeColor:  c.EyeColor,
				BirthYear: c.BirthYear,
				Gender:    c.Gender,
				Films:     c.Films,
				Created:   c.Created,
				Edited:    c.Edited,
				URL:       c.URL,
			})
		}
		gctx.JSON(http.StatusOK, result)
	} else {
		log.Println(err)
		gctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "something went wrong",
		})
	}
}

// @BasePath /api/v1

// Movies 			godoc
// @Summary 		list all movies
// @Schemes
// @Description 	Returns a list of all star wars movies
// @Accept       	json
// @Produce 		json
// @Success 		200 {array} FilmResponse
// @Failure      	500  {object}  ErrorResponse
// @Router 			/movies [get]
func (h *ApiHelper) Comments(gctx *gin.Context) {
	if comments, err := h.EntClient.Comment.Query().
		//Order((ent.Asc(movie.FieldReleaseDate))).
		All(h.Context); err == nil {
		result := []CommentResponse{}
		for _, m := range comments {
			result = append(result, CommentResponse{
				Name:    m.Name,
				Text:    m.Text,
				IP:      m.IP,
				Created: m.Created,
			})
		}
		gctx.JSON(http.StatusOK, result)
	} else {
		log.Println(err)
		gctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: "something went wrong",
		})
	}
}

// @BasePath /api/v1

// Movies 			godoc
// @Summary 		creates a new comment
// @Schemes
// @Description 	adds a comment to a movie
// @Accept       	json
// @Produce 		json
// @Success 		200 {array}   CommentResponse
// @Failure      	500 {object}  ErrorResponse
// @Router 			/movies [get]
func (h *ApiHelper) NewComment(gctx *gin.Context) {
	var comment Comment
	if gctx.ShouldBindJSON(&comment) == nil {
		return
	}
	if _, err := h.EntClient.Comment.Create().
		SetName(comment.Name).
		SetIP(gctx.ClientIP()).
		SetCreated(time.Now()).
		SetText(comment.Text).
		Save(h.Context); err != nil {
		log.Println(err)
		gctx.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	gctx.JSON(http.StatusCreated, comment)
}
