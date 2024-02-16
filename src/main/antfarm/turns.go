package antfarm

import (
	"fmt"
	"slices"
	"strings"
)

type Turn struct {
	InStart int
	InEnd   int
	EnRoute map[*Ant]*Vertex
}

func (t *Turn) String() string {
	return fmt.Sprintf("in start: %d, in route: %d, in end: %d", t.InStart, len(t.EnRoute), t.InEnd)
}

type turns struct {
	InStart int
	InEnd   int
	Data    []*Turn
}

func (t *turns) String() string {
	for i, turn := range t.Data {
		fmt.Printf("Turn %d: %s\n", i, turn)
	}
	return ""
}

func (t *turns) Parse(line string) {
	turn := &Turn{
		InStart: t.InStart,
		InEnd:   t.InEnd,
		EnRoute: make(map[*Ant]*Vertex),
	}
	fields := strings.Split(strings.Trim(line, " "), " ")
	for _, field := range fields {
		parts := strings.Split(field, "-")
		antName := parts[0][1:]
		ant := Ants.All[antName]
		vertexName := parts[1]
		vertex := Graph.FindVertex(vertexName)

		if ant.Current == 0 {
			turn.InStart--
		}
		if vertex == Graph.End {
			turn.InEnd++
			turn.EnRoute[ant] = vertex
		} else {
			turn.EnRoute[ant] = vertex
		}
		ant.Current++
	}
	t.Data = append(t.Data, turn)
	t.InStart = turn.InStart
	t.InEnd = turn.InEnd
}

func (t *turns) ExtractPaths() {
	paths := make(map[*Ant][]*Vertex)
	for _, ant := range Ants.All {
		paths[ant] = []*Vertex{}
	}
	for _, turn := range t.Data {
		for ant, vertex := range turn.EnRoute {
			paths[ant] = append(paths[ant], vertex)
		}
	}
	for ant, path := range paths {
		path = append([]*Vertex{Graph.Start}, path...)
		path = append(path, Graph.End)
		var found bool
		for i, p := range Graph.Paths {
			if slices.Equal(p, path) {
				found = true
				Ants.Queues[i] = append(Ants.Queues[i], ant)
				ant.Queue = i
				break
			}
		}
		if !found {
			Graph.Paths = append(Graph.Paths, path)
			Ants.Queues = append(Ants.Queues, []*Ant{ant})
			ant.Queue = len(Graph.Paths) - 1
		}
	}
	for _, ant := range Ants.All {
		ant.Current = 0
	}
}
