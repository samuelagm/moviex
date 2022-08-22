package types

import (
	"fmt"
	"math"
)

func HtoFeetInces(h int) string {
	hFeet := 0
	hInch := 0
	hComposite := (float64(h) * 0.3937) / 12
	hFeet = int(math.Floor(hComposite))
	hInch = int(math.Ceil(hComposite))

	return fmt.Sprintf(`%dcm, %d'%d"`, h, hFeet, hInch)
}
