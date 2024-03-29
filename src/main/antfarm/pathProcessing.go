package antfarm

import (
	"math"
)

type PathProcessing struct {
	Positions []Position
	Marks     []rune
}

func (path *PathProcessing) Add(pos Position, mark rune) {
	path.Positions = append(path.Positions, pos)
	path.Marks = append(path.Marks, mark)
}

type paths struct {
	All [][]*GraphVertex
}

func (paths *paths) String() string {
	str := ""
	for _, path := range paths.All {
		for _, v := range path {
			str += v.Name + " "
		}
		str += "\n"
	}
	return str
}

func (paths *paths) Find() {
	Explorer.Explore()
	paths.Sort()
	paths.Disjoin()
}

func (paths *paths) Sort() {
	for i := 0; i < len(paths.All)-1; i++ {
		for j := i; j < len(paths.All); j++ {
			if len(paths.All[i]) > len(paths.All[j]) {
				paths.All[i], paths.All[j] = paths.All[j], paths.All[i]
			}
		}
	}
}

func (paths *paths) Disjoin() {
	var candidates = make(map[int][]Combination)
	for i, path := range paths.All {
		var disjoint [][]*GraphVertex
		disjoint = append(disjoint, path)
		if i == 0 {
			candidates[1] = []Combination{{paths: disjoint}}
		}
		paths.AddCandidates(i, disjoint, &candidates)
	}
	paths.SelectBest(candidates)
}

func (paths *paths) AddCandidates(i int, disjoint [][]*GraphVertex, candidates *map[int][]Combination) {
	for j := i + 1; j < len(paths.All); j++ {
		if paths.areDisjoint(disjoint, paths.All[j]) {
			disjoint := append(disjoint, paths.All[j])
			if _, exists := (*candidates)[len(disjoint)]; !exists {
				(*candidates)[len(disjoint)] = []Combination{{paths: disjoint}}
			} else {
				(*candidates)[len(disjoint)] = append((*candidates)[len(disjoint)], Combination{paths: disjoint})
			}
			paths.AddCandidates(j, disjoint, candidates)
		}
	}
}

func (paths *paths) SelectBest(candidates map[int][]Combination) {
	Graph.Turns = math.MaxInt32
	for combLen, combinations := range candidates {
		for _, combination := range combinations {
			elements := Ants.Number + paths.length(combination.paths) - combLen
			steps := elements / combLen
			if elements%combLen > 0 {
				steps++
			}
			if steps < Graph.Turns {
				Graph.Paths = combination.paths
				Graph.Turns = steps
			}
		}
	}
}

func (paths *paths) length(disjoint [][]*GraphVertex) int {
	length := 0
	for _, path := range disjoint {
		length += len(path) - 1
	}
	return length
}

func (paths *paths) areDisjoint(disjointPaths [][]*GraphVertex, path2 []*GraphVertex) bool {
	for _, path1 := range disjointPaths {
		for _, v1 := range path1[1 : len(path1)-1] {
			for _, v2 := range path2[1 : len(path2)-1] {
				if v1 == v2 {
					return false
				}
			}
		}
	}
	return true
}

var Paths = &paths{}
