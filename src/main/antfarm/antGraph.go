package antfarm

import (
	"fmt"
	"os"
	"time"
)

type antGraph struct {
	Start    *GraphVertex
	End      *GraphVertex
	Vertices []*GraphVertex
	Edges    map[string][]string
	Paths    [][]*GraphVertex
	Turns    int
	Time     time.Duration
}

func (graph *antGraph) FindVertex(name string) *GraphVertex {
	var found *GraphVertex
	for _, v := range graph.Vertices {
		if v.Name == name {
			found = v
			break
		}
	}
	if found == nil {
		FaultyData("vertex not found: " + name)
	}
	return found
}

func (graph *antGraph) Check() {
	if len(graph.Vertices) == 0 {
		FaultyData("no rooms specified")
	}
	if numberOfEdges() == 0 {
		FaultyData("no edges specified")
	}
	if graph.Start == nil {
		FaultyData("no start room specified")
	}
	if graph.End == nil {
		FaultyData("no end room specified")
	}
}

func numberOfEdges() int {
	var count int
	for _, edges := range Graph.Edges {
		count += len(edges)
	}
	return count
}

var Graph = &antGraph{
	Edges:    make(map[string][]string),
	Paths:    [][]*GraphVertex{},
	Vertices: []*GraphVertex{},
}

func FaultyData(msg string) {
	fmt.Printf("ERROR: invalid data format, %s\n", msg)
	os.Exit(0)
}
