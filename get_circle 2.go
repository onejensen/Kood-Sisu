package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {

	newc := Circle{
		Radius:    r,
		Diameter:  r * 2,
		Area:      r * r * 3.14,
		Perimeter: 2 * r * 3.14,
	}

	return newc

}
