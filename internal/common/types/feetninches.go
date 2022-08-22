package types

import (
	"fmt"
	"math"
	"strconv"
)

func HtoFeetInces(hstring string) string {
	hFeet := 0
	hInch := 0
	if hInt, err := strconv.Atoi(hstring); err == nil {
		hComposite := (float64(hInt) * 0.3937) / 12
		hFeet = int(math.Floor(hComposite))
		hInch = int(math.Ceil(hComposite))
	} else {
		return hstring
	}
	return fmt.Sprintf(`%scm, %d'%d"`, hstring, hFeet, hInch)
}
