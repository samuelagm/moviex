package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/samuelagm/moviex/docs"
	"github.com/samuelagm/moviex/ent"
	apitypes "github.com/samuelagm/moviex/internal/api/types"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	docs.SwaggerInfo.BasePath = "/api/v1"
}

func Listen(ctx context.Context, dbClient *ent.Client) {

	api := apitypes.NewApiHelper(ctx, dbClient)
	r := gin.Default()

	env1 := os.Getenv("ENV_1")
	sec1 := os.Getenv("SEC_1")
	log.Printf("ENV_1=%s SEC_1=%s", env1, sec1)
	healthMessage := fmt.Sprintf("Alive and Well 28 | ENV_1=%s SEC_1=%s", env1, sec1)

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, fmt.Sprintf("Welcome, see: /api/v1/docs/index.html. %s", healthMessage))
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, healthMessage)
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/movies", api.Movies)
		v1.GET("/characters/:episodeId", api.Characters)
		v1.GET("/comments/:episodeId", api.Comments)
		v1.POST("/comments/:episodeId", api.NewComment)
		v1.POST("/movies", api.NewMovie)
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
