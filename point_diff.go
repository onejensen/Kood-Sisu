package sprint

type Point struct {
	X    float32
	Y    float32
	Text string
}

func PointDiff(p1, p2 Point) Point {

	pos1 := p1.X*p1.X + p1.Y*p1.Y
	pos2 := p2.X*p2.X + p2.Y*p2.Y

	if pos1 >= pos2 {
		return p1
	} else {
		return p2
	}
}
