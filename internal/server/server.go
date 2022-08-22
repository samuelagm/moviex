package server

import (
	"context"

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

	v1 := r.Group("/api/v1")
	{
		v1.GET("/movies", api.Movies)
		v1.GET("/characters/:episodeId", api.Characters)
		v1.GET("/comment/:episodeId", api.Comments)
		v1.POST("/comment/:episodeId", api.NewComment)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
