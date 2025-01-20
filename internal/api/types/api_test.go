package types

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
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
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	loader.Load(context.Background(), dbClient)

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
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiHelper_Characters(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/characters/4", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApiHelper_Comments(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/v1/comments/4", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestApiHelper_NewComments(t *testing.T) {
	var jsonData = []byte(`{
		"name": "job",
		"text": "moviex"
	}`)
	req, _ := http.NewRequest("POST", "/api/v1/comments/4",
		bytes.NewBuffer(jsonData))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
