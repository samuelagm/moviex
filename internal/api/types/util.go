package types

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/ent/movie"
)

func getConnectedMovie(gctx *gin.Context, h *ApiHelper) *ent.Movie {
	var m *ent.Movie
	if id, err := strconv.Atoi(gctx.Param("episodeId")); err == nil {
		m, err = h.EntClient.Movie.Query().
			Where(movie.EpisodeIDEQ(id)).
			Unique(true).First(h.Context)
		if err != nil {
			gctx.JSON(http.StatusNotFound, ErrorResponse{
				Message: fmt.Sprintf("movie with episode_id %d not found", id),
			})
		}
	}
	return m
}
