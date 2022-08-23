package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/internal/loader"
	"github.com/samuelagm/moviex/internal/server"
)

func Run() {
	ch := make(chan os.Signal, 1)
	ctx, cancelCTX := context.WithCancel(context.Background())

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	loader.Load(ctx, client)
	go server.Listen(ctx, client)

	signal := <-ch

	fmt.Printf("\nGot signal: %s, exiting...\n", signal)
	client.Close()
	cancelCTX()
	os.Exit(0)
}
