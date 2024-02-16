package antfarm

type GraphExplorer struct {
	visited map[*GraphVertex]bool
	current *GraphVertex
	path    []*GraphVertex
}

type PathData struct {
	Paths []string `json:"paths"`
}

func (explorer *GraphExplorer) Explore() {
	var prev *GraphVertex
	if explorer.current == nil {
		explorer.current = Graph.Start
	}
	for v, _ := range explorer.current.Edges {
		if explorer.current == Graph.Start {
			explorer.path = []*GraphVertex{Graph.Start}
			explorer.visited[Graph.Start] = true
		}
		if !explorer.visited[v] {
			explorer.path = append(explorer.path, v)
			explorer.visited[v] = true
			prev = explorer.current
			explorer.current = v
			if v == Graph.End {
				tmp := make([]*GraphVertex, len(explorer.path))
				copy(tmp, explorer.path)
				Paths.All = append(Paths.All, tmp)
				AllPaths = append(AllPaths, tmp)
			} else {
				explorer.Explore()
			}
			explorer.visited[v] = false
			explorer.current = prev
			if len(explorer.path) > 1 {
				explorer.path = explorer.path[:len(explorer.path)-1]
			}
		}
	}
}

var Explorer = &GraphExplorer{current: Graph.Start, visited: make(map[*GraphVertex]bool)}
