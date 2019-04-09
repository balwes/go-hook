package math

import (
    "math"
)

func PointInsideRect(p *Point, r *Rect) bool {
    return p.X >= r.X && p.X <= r.X + r.Width && p.Y >= r.Y && p.Y <= r.Y + r.Height
}

func PointPointDist(from *Point, to *Point) float32 {
    dx := from.X - to.X
    dy := from.Y - to.Y
    return float32(math.Sqrt(float64(dx*dx + dy*dy)))
}

func Round(f float32) int32 {
	return int32(math.Round(float64(f)))
}

func FloatIsWhole(f float32) bool {
	return f == float32(int32(f))
}

func Clamp(value, min, max float32) float32 {
    if value < min {
        return min
    } else if value > max {
        return max
    } else {
        return value
    }
}
