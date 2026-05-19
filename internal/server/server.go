package server

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/samuelagm/moviex/docs"
	"github.com/samuelagm/moviex/ent"
	apitypes "github.com/samuelagm/moviex/internal/api/types"
	"github.com/samuelagm/moviex/internal/auth"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed ui
var uiFS embed.FS

func init() {
	docs.SwaggerInfo.BasePath = "/api/v1"
}

func Listen(ctx context.Context, dbClient *ent.Client) {

	api := apitypes.NewApiHelper(ctx, dbClient)
	authStore := auth.NewStore()
	r := gin.Default()

	env1 := os.Getenv("ENV_1")
	sec1 := os.Getenv("SEC_1")
	log.Printf("ENV_1=%s SEC_1=%s", env1, sec1)
	healthMessage := fmt.Sprintf("Alive and Well 35 | ENV_1=%s SEC_1=%s", env1, sec1)

	r.GET("/", func(c *gin.Context) {
		serveEmbedded(c, "ui/index.html")
	})
	r.GET("/movies/new", func(c *gin.Context) {
		serveEmbedded(c, "ui/movies-new.html")
	})
	r.GET("/movies/:id", func(c *gin.Context) {
		serveEmbedded(c, "ui/movie.html")
	})
	r.GET("/login", func(c *gin.Context) {
		serveEmbedded(c, "ui/login.html")
	})
	r.GET("/register", func(c *gin.Context) {
		serveEmbedded(c, "ui/register.html")
	})

	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, healthMessage)
	})

	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/register", authStore.Register)
		v1.POST("/auth/login", authStore.Login)

		v1.GET("/movies", api.Movies)
		v1.GET("/movies/:episodeId", api.Movie)
		v1.GET("/characters", api.AllCharacters)
		v1.GET("/characters/:episodeId", api.Characters)
		v1.GET("/comments/:episodeId", api.Comments)
		v1.GET("/stats", api.Stats)

		protected := v1.Group("/", authStore.Middleware())
		{
			protected.POST("/comments/:episodeId", api.NewComment)
			protected.DELETE("/comments/:episodeId/:commentId", api.DeleteComment)
			protected.POST("/movies", api.NewMovie)
			protected.PUT("/movies/:episodeId", api.UpdateMovie)
			protected.POST("/characters", api.NewCharacter)
		}

		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func serveEmbedded(c *gin.Context, path string) {
	data, err := uiFS.ReadFile(path)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", data)
}
