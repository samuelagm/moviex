package main

import (
	"fmt"
	"os"

	"github.com/samuelagm/moviex/internal/app"
	_ "github.com/samuelagm/moviex/internal/database"
)

// @title Movie API
// @version 0.0.5
// @description StarWars Movie Server.
// @termsOfService http://swagger.io/terms/
// @contact.name Samuel Agbonkpolo
// @contact.email samuelagm@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if os.Getenv("MOVIEX_ENV_PROBE") == "" {
		fmt.Fprintln(os.Stderr, "MOVIEX_ENV_PROBE is required")
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "MOVIEX_ENV_PROBE is set")
	app.Run()
}

// monk-test-1777467649
