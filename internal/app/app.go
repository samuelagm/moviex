package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "modernc.org/sqlite"
	"github.com/samuelagm/moviex/ent"
	"github.com/samuelagm/moviex/internal/loader"
	"github.com/samuelagm/moviex/internal/server"
)

func Run() {
	ch := make(chan os.Signal, 1)
	ctx, cancelCTX := context.WithCancel(context.Background())
	defer cancelCTX()

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	loader.Load(ctx, client)
	go server.Listen(ctx, client)

	signal := <-ch

	fmt.Printf("\nGot signal: %s, exiting...\n", signal)
	os.Exit(0)
}
