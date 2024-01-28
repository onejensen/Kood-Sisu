package sprint

/* type Point struct {
	X    float32
	Y    float32
	Text string
} */

func MakePoint(x, y float32, text string) Point {

	var p1 Point = Point{x, y, text}
	return p1

}
