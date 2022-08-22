package types

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/internal/loader"
	"gotest.tools/assert"
)

var dbClient *ent.Client
var router *gin.Engine

func TestMain(m *testing.M) {
	initVars()
	os.Exit(m.Run())
}

func initVars() {
	var err error
	dbClient, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	loader.Load(context.Background(), dbClient)
	//defer client.Close()
	api := NewApiHelper(context.Background(), dbClient)
	router = gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/movies", api.Movies)
		v1.GET("/characters/:episodeId", api.Characters)
		v1.GET("/comment/:episodeId", api.Comments)
		v1.POST("/comment/:episodeId", api.NewComment)
	}
}

func TestApiHelper_Movies(t *testing.T) {

	req, _ := http.NewRequest("GET", "/api/v1/movies", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestApiHelper_Characters(t *testing.T) {
// 	type fields struct {
// 		EntClient *ent.Client
// 		Context   context.Context
// 	}
// 	type args struct {
// 		gctx *gin.Context
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &ApiHelper{
// 				EntClient: tt.fields.EntClient,
// 				Context:   tt.fields.Context,
// 			}
// 			h.Characters(tt.args.gctx)
// 		})
// 	}
// }

// func TestApiHelper_Comments(t *testing.T) {
// 	type fields struct {
// 		EntClient *ent.Client
// 		Context   context.Context
// 	}
// 	type args struct {
// 		gctx *gin.Context
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &ApiHelper{
// 				EntClient: tt.fields.EntClient,
// 				Context:   tt.fields.Context,
// 			}
// 			h.Comments(tt.args.gctx)
// 		})
// 	}
// }

// func TestApiHelper_NewComment(t *testing.T) {
// 	type fields struct {
// 		EntClient *ent.Client
// 		Context   context.Context
// 	}
// 	type args struct {
// 		gctx *gin.Context
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			h := &ApiHelper{
// 				EntClient: tt.fields.EntClient,
// 				Context:   tt.fields.Context,
// 			}
// 			h.NewComment(tt.args.gctx)
// 		})
// 	}
// }
