package sprint

import "fmt"

/* type Point struct {
	X    float32
	Y    float32
	Text string
} */

func PointText(p Point) Point {

	p.Text = fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return p
}
