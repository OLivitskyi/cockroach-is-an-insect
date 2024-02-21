package antfarm_test

import (
	"testing"

	"cockroach/src/main/antfarm"
	"github.com/stretchr/testify/assert"
)

func TestSortEdgesCross(t *testing.T) {
	v := &antfarm.GraphVertex{
		Name:     "A",
		Position: antfarm.Position{X: 0, Y: 0},
		Sorted: []*antfarm.GraphVertex{
			{Name: "B", Position: antfarm.Position{X: 5, Y: 0}},
			{Name: "C", Position: antfarm.Position{X: 5, Y: 5}},
		},
	}
	expectedSorted := []*antfarm.GraphVertex{
		{Name: "B", Position: antfarm.Position{X: 5, Y: 0}},
		{Name: "C", Position: antfarm.Position{X: 5, Y: 5}},
	}
	v.SortEdgesCross()
	assert.Equal(t, expectedSorted, v.Sorted)
}

func TestSortEdgesByDegrees(t *testing.T) {
	v := &antfarm.GraphVertex{
		Name:     "A",
		Position: antfarm.Position{X: 0, Y: 0},
		Sorted: []*antfarm.GraphVertex{
			{Name: "C", Position: antfarm.Position{X: 2, Y: 1}},
			{Name: "B", Position: antfarm.Position{X: 3, Y: 4}},
		},
	}

	expectedSorted := []*antfarm.GraphVertex{
		{Name: "C", Position: antfarm.Position{X: 2, Y: 1}},
		{Name: "B", Position: antfarm.Position{X: 3, Y: 4}},
	}

	v.SortEdgesByDegrees()

	assert.Equal(t, expectedSorted, v.Sorted)
}

func TestSinCos(t *testing.T) {
	pos1 := antfarm.Position{X: 0, Y: 0}
	pos2 := antfarm.Position{X: 3, Y: 4}

	sin, cos := antfarm.SinCos(pos1, pos2)

	assert.Equal(t, 0.8, sin)
	assert.Equal(t, 0.6, cos)
}

func TestLinedUp(t *testing.T) {
	pos1 := antfarm.Position{X: 0, Y: 0}
	pos2 := antfarm.Position{X: 1, Y: 1}
	pos3 := antfarm.Position{X: 2, Y: 0}
	pos4 := antfarm.Position{X: 0, Y: 2}
	assert.True(t, antfarm.LinedUp(pos1, pos2))
	assert.True(t, antfarm.LinedUp(pos1, pos3))
	assert.True(t, antfarm.LinedUp(pos1, pos4))
}

func TestDiagonal(t *testing.T) {
	pos1 := antfarm.Position{X: 0, Y: 0}
	pos2 := antfarm.Position{X: 2, Y: 2}
	pos3 := antfarm.Position{X: 3, Y: 2}
	assert.True(t, antfarm.Diagonal(pos1, pos2))
	assert.False(t, antfarm.Diagonal(pos1, pos3))
}

func TestAbs(t *testing.T) {
	assert.Equal(t, 5, antfarm.Abs(-5))
	assert.Equal(t, 5, antfarm.Abs(5))
	assert.Equal(t, 0, antfarm.Abs(0))
}
