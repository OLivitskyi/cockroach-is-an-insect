package antfarm_test

import (
	"testing"

	"cockroach/src/main/antfarm"
	"github.com/stretchr/testify/assert"
)

func TestExplore(t *testing.T) {
	antfarm.Graph.Start = &antfarm.GraphVertex{Name: "A"}
	antfarm.Graph.End = &antfarm.GraphVertex{Name: "D"}

	// Building graph structure
	vertexB := &antfarm.GraphVertex{Name: "B"}
	vertexC := &antfarm.GraphVertex{Name: "C"}

	antfarm.Graph.Start.Edges = make(map[*antfarm.GraphVertex]*antfarm.PathProcessing)
	antfarm.Graph.Start.Edges[vertexB] = &antfarm.PathProcessing{}
	antfarm.Graph.Start.Edges[vertexC] = &antfarm.PathProcessing{}

	vertexB.Edges = make(map[*antfarm.GraphVertex]*antfarm.PathProcessing)
	vertexB.Edges[antfarm.Graph.End] = &antfarm.PathProcessing{}

	vertexC.Edges = make(map[*antfarm.GraphVertex]*antfarm.PathProcessing)
	vertexC.Edges[antfarm.Graph.End] = &antfarm.PathProcessing{}

	antfarm.Paths.All = make([][]*antfarm.GraphVertex, 0)
	antfarm.Explorer.Explore()

	path1 := []*antfarm.GraphVertex{
		antfarm.Graph.Start,
		vertexB,
		antfarm.Graph.End,
	}

	path2 := []*antfarm.GraphVertex{
		antfarm.Graph.Start,
		vertexC,
		antfarm.Graph.End,
	}

	expectedPaths := [][]*antfarm.GraphVertex{path1, path2}

	assert.NotNil(t, antfarm.Paths.All)
	assert.Equal(t, 2, len(antfarm.Paths.All))
	assert.ElementsMatch(t, expectedPaths, antfarm.Paths.All)
}
