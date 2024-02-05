package sprint

import "math"

func Casting(num float64) int {
	result := (math.Round(num))
	return int(result)
}
