package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	var endpoint Point
	endpoint.x = s.start.x + int(s.a)
	endpoint.y = s.start.y + int(s.a)
	return endpoint
}

func (s Square) Area() uint {
	return a * a
}

func (s Square) Perimeter() uint {
	return 2 * (s.a + s.a)
}
