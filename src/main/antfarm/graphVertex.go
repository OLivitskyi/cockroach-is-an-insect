package antfarm

import (
	"math"
)

type GraphVertex struct {
	Name     string
	Edges    map[*GraphVertex]*PathProcessing
	Sorted   []*GraphVertex
	Capacity int
	Position Position
	Paths    map[*GraphVertex][]*PathProcessing
}

func (v *GraphVertex) SortEdgesCross() {
	for i := 0; i < len(v.Sorted)-1; i++ {
		for j := i + 1; j < len(v.Sorted); j++ {
			pos1 := v.Sorted[i].Position
			pos2 := v.Sorted[j].Position
			if !LinedUp(v.Position, pos1) && LinedUp(v.Position, pos2) {
				v.Sorted[i], v.Sorted[j] = v.Sorted[j], v.Sorted[i]
			}
		}
	}
}

func (v *GraphVertex) SortEdgesByDegrees() {
	pos0 := v.Position
	for i := 0; i < len(v.Sorted)-1; i++ {
		for j := i + 1; j < len(v.Sorted); j++ {
			v1 := v.Sorted[i]
			pos1 := v1.Position
			sin1, cos1 := SinCos(pos0, pos1)
			v2 := v.Sorted[j]
			pos2 := v2.Position
			sin2, cos2 := SinCos(pos0, pos2)
			if cos1 < 0 && cos2 > 0 {
				v.Sorted[i], v.Sorted[j] = v.Sorted[j], v.Sorted[i]
			} else if (cos1 < 0 && cos2 < 0) || (cos1 > 0 && cos2 > 0) {
				if sin1 > sin2 {
					v.Sorted[i], v.Sorted[j] = v.Sorted[j], v.Sorted[i]
				}
			}
		}
	}
}

func SinCos(pos1, pos2 Position) (float64, float64) {
	diff := struct{ X, Y float64 }{float64(pos2.X - pos1.X), float64(pos2.Y - pos1.Y)}
	hypotenuse := math.Hypot(float64(diff.X), float64(diff.Y))
	sin := diff.Y / hypotenuse
	cos := diff.X / hypotenuse
	return sin, cos
}

func LinedUp(pos1, pos2 Position) bool {
	return pos1.Y == pos2.Y || pos1.X == pos2.X || Diagonal(pos1, pos2)
}

func Diagonal(pos1, pos2 Position) bool {
	return Abs(pos1.Y-pos2.Y) == Abs(pos1.X-pos2.X)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
