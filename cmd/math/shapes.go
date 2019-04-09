package math

type Point struct {
	X float32
	Y float32
}

type Circle struct {
	X float32
	Y float32
	Radius float32
}

func (c *Circle) CenterPoint() *Point {
	return &Point{c.X, c.Y}
}

type Line struct {
	X1 float32
	Y1 float32
	X2 float32
	Y2 float32
}

func (l *Line) StartPoint() *Point {
	return &Point{l.X1, l.Y1}
}

func (l *Line) EndPoint() *Point {
	return &Point{l.X2, l.Y2}
}

type Triangle struct {
	X1 float32
	Y1 float32
	X2 float32
	Y2 float32
	X3 float32
	Y3 float32
}

func (t Triangle) V1() *Point {
	return &Point{t.X1, t.Y1}
}

func (t Triangle) V2() *Point {
	return &Point{t.X2, t.Y2}
}

func (t Triangle) V3() *Point {
	return &Point{t.X3, t.Y3}
}

type Rect struct {
	X float32
	Y float32
	Width float32
	Height float32
}

func (r *Rect) Center() (float32, float32) {
	return r.X + r.Width/2, r.Y + r.Height/2
}

func (r *Rect) MinPoint() *Point {
	return &Point{r.X, r.Y}
}

func (r *Rect) Max() (float32, float32) {
	return r.X + r.Width, r.Y + r.Height
}

func (r *Rect) MaxPoint() *Point {
	return &Point{r.X + r.Width, r.Y + r.Height}
}

type Polygon struct {
	Vertices []*Point
}
