# Project
***
This project does the following:
1. Serves endpoint for the following routes:
   - `GET /api/v1/movies` list all movies.
   - `GET /api/v1/characters/:episodeId` list all characters in a movie with episode_id.
   - `GET /api/v1/comments/:episodeId` list all characters in a movie with episode_id.
   - `POST /api/v1/comment/:episodeId` add a comment to a movie with episode_id.

2. Downloads movie data into an in memory sqlite3 database

### App Url
[https://morning-chamber-96103.herokuapp.com](https://morning-chamber-96103.herokuapp.com/)
### API Docs
   [https://morning-chamber-96103.herokuapp.com/api/v1/docs/index.html](https://morning-chamber-96103.herokuapp.com/api/v1/docs/index.html)
### Prerequisites

1. Install go dependencies  
   `go install`


### Unit Tests
   - `go test ./internal/api/types/...`

### Execution

`go run main.go`


