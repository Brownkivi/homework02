package logic

type shape interface {
	Area(a int, b int) float64
	Perimeter(c int, d int) float64
}

type Circle struct {
	Radius int
}

type Rectangle struct {
	Rength int
	Width  int
}

func (c Circle) Area(a int, b int) float64 {
	return 3.14 * float64(a) * float64(a)
}

func (c Circle) Perimeter(c1 int, d int) float64 {
	return 2 * 3.14 * float64(c1)
}

func (r Rectangle) Area(a int, b int) float64 {
	return float64(a) * float64(b)
}

func (r Rectangle) Perimeter(c int, d int) float64 {
	return 2 * (float64(c) + float64(d))
}
