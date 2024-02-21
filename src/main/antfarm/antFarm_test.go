package antfarm_test

import (
	"cockroach/src/main/antfarm"
	"testing"
)

func TestGraphVertexInitialization(t *testing.T) {
	vertex := &antfarm.GraphVertex{Name: "Start", Capacity: 1}
	if vertex.Name != "Start" {
		t.Errorf("Expected vertex name to be 'Start', got %s", vertex.Name)
	}

	if vertex.Capacity != 1 {
		t.Errorf("Expected vertex capacity to be 1, got %d", vertex.Capacity)
	}
}

// TestSortEdgesByDegrees - тест правильності сортування ребер за кутами
func TestSortEdgesByDegreesAnts(t *testing.T) {
	vertex := &antfarm.GraphVertex{Position: antfarm.Position{X: 0, Y: 0}}
	v1 := &antfarm.GraphVertex{Position: antfarm.Position{X: -1, Y: 1}} // 135 degrees
	v2 := &antfarm.GraphVertex{Position: antfarm.Position{X: 1, Y: 0}}  // 0 degrees
	vertex.Sorted = []*antfarm.GraphVertex{v1, v2}
	vertex.SortEdgesByDegrees()

	if !(vertex.Sorted[0] == v2 && vertex.Sorted[1] == v1) {
		t.Errorf("Expected vertices to be sorted by degrees, but they were not")
	}
}
