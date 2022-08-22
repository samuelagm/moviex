package app

import (
	"context"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/internal/loader"
	"github.com/samuelagm/moviex/internal/server"
)

func Run() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	loader.Load(context.Background(), client)
	go server.Listen(context.Background(), client)
	select {}
}
