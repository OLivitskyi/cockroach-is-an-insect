package antfarm

import (
	"fmt"
	"os"
	"strconv"
)

type Ant struct {
	Name    string
	Current int
	Queue   int
}

type ants struct {
	Number int
	All    map[string]*Ant
	Queues [][]*Ant
}

var AllMoves [][]string

func (ants *ants) Distribute() {
	ants.Queues = make([][]*Ant, len(Graph.Paths))
	for i := 0; i < ants.Number; i++ {
		name := strconv.Itoa(i + 1)
		ant := ants.All[name]
		for j := range ants.Queues {
			pathLen := len(Graph.Paths[j]) - 1
			queueLen := len(ants.Queues[j])
			if pathLen+queueLen <= Graph.Turns {
				ants.Queues[j] = append(ants.Queues[j], ant)
				ant.Queue = j
				ant.Current = 0
				break
			}
		}
	}
}

func (ants *ants) Step(webVisualisation bool) {
	var movesOnStep []string
NEXT_QUEUE:
	for _, queue := range ants.Queues {
		if len(queue) > 0 {
			for _, ant := range queue {
				if ant.Current < len(Graph.Paths[ant.Queue])-1 {
					ant.Current++

					ants.Print(ant, webVisualisation)
					concat := fmt.Sprintf("%s-%s", ant.Name, Graph.Paths[ant.Queue][ant.Current].Name)
					movesOnStep = append(movesOnStep, concat)

					if ant.Current == 1 {
						continue NEXT_QUEUE
					}
				}
			}
		}
	}
	AllMoves = append(AllMoves, movesOnStep)
}

func (ants *ants) Move(webVisualisation bool) {
	if !webVisualisation && Graph.Turns > 600 {
		for i := 0; i < Graph.Turns; i++ {
			if i%100 == 0 {
				fmt.Printf("%d turns have been shown, there are %d more. To continue, press ENTER, or N+ENTER to stop.\n", i, Graph.Turns-i)
				var i rune
				n, _ := fmt.Scanf("%c", &i)
				if n == 1 && (i == 'n' || i == 'N') {
					fmt.Println("...")
					fmt.Printf("Moved %d ants along %v disjoint paths in %v turns.\n", Ants.Number, len(Graph.Paths), Graph.Turns)
					fmt.Printf("Found altogether %v paths, %v best paths in %v.\n", len(Paths.All), len(Graph.Paths), Graph.Time)
					os.Exit(0)
				}
			}
			ants.Step(webVisualisation)
			if !webVisualisation {
				fmt.Println()
			}
		}
	} else {
		for i := 0; i < Graph.Turns; i++ {
			ants.Step(webVisualisation)
			if !webVisualisation {
				fmt.Println()
			}
		}
	}
}

func (ants *ants) Print(ant *Ant, webVisualization bool) {
	if !webVisualization {
		fmt.Printf("L%s-%s ", ant.Name, Graph.Paths[ant.Queue][ant.Current].Name)
	}

}

var Ants = &ants{}
