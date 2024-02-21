package main

import (
	"bufio"
	"cockroach/src/main/antfarm"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var (
		ants             int
		webVisualisation bool
	)
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fi, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}
		if fi.Mode()&os.ModeNamedPipe == 0 {
			fmt.Println("Error: missing filename")
			fmt.Println("USAGE: go run . [-ants <int>] <filename|path>")
			os.Exit(1)
		}
	}

	antfarm.Ants.Number = ants
	start := time.Now()
	fileName, _, _, _ := parseInput(args)

	antfarm.Paths.Find()
	antfarm.Ants.Distribute()
	antfarm.Graph.Time = time.Since(start)
	printInput(fileName, webVisualisation)
	antfarm.Ants.Move(webVisualisation)
	fmt.Println()
}

func parseInput(args []string) (string, string, string, string) {
	var file *os.File
	fileName := ""
	var scanner *bufio.Scanner
	if len(args) > 0 {
		file, fileName = openFile(args)
		defer file.Close()
		scanner = bufio.NewScanner(file)
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}
	scanner.Scan()
	line := scanner.Text()
	getAnts(line)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#") {
			if line == "##start" || line == "##end" {
				scanner.Scan()
				nextLine := scanner.Text()
				vertex := setVertex(nextLine)
				vertex.Capacity = antfarm.Ants.Number
				if line == "##start" {
					antfarm.Graph.Start = vertex
					antfarm.StartJSON = vertex.Name
				} else {
					antfarm.Graph.End = vertex
				}
			}
		} else if strings.Contains(line, " ") {
			setVertex(line)
		} else if strings.Contains(line, "-") {
			setEdge(line)
		}
	}
	antfarm.Graph.Check()
	return fileName, "", "", ""
}

func openFile(args []string) (*os.File, string) {
	if len(args) < 1 {
		fmt.Println("No file specified.")
		os.Exit(1)
	}
	fileName := args[0]

	if !strings.Contains(fileName, "/") {
		fileName = "resources/" + fileName
	}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file.")
		os.Exit(1)
	}
	return file, fileName
}

func getAnts(line string) {
	number, err := strconv.Atoi(line)
	if err != nil {
		antfarm.FaultyData(fmt.Sprintf("invalid number of ants: %s", line))
	}

	if !FlagIsPassed("ants") {
		antfarm.Ants.Number = number
	}
	number = antfarm.Ants.Number
	if number < 1 || number > math.MaxInt32 {
		antfarm.FaultyData(fmt.Sprintf("invalid number of ants: %d", number))
	}
	antfarm.Ants.All = make(map[string]*antfarm.Ant)
	for i := 0; i < number; i++ {
		ant := &antfarm.Ant{}
		ant.Name = strconv.Itoa(i + 1)
		antfarm.Ants.All[ant.Name] = ant
	}
}

func setVertex(line string) *antfarm.GraphVertex {
	vertex := &antfarm.GraphVertex{
		Edges:  make(map[*antfarm.GraphVertex]*antfarm.PathProcessing),
		Sorted: make([]*antfarm.GraphVertex, 0),
	}
	fields := strings.Split(line, " ")
	vertex.Name = fields[0]
	vertex.Position.X, _ = strconv.Atoi(fields[1])
	vertex.Position.Y, _ = strconv.Atoi(fields[2])
	vertex.Capacity = 1
	antfarm.Graph.Vertices = append(antfarm.Graph.Vertices, vertex)
	antfarm.Graph.Edges[vertex.Name] = []string{}
	return vertex
}

func setEdge(line string) {
	fields := strings.Split(line, "-")
	v1Name := fields[0]
	v2Name := fields[1]
	if v1Name == v2Name {
		antfarm.FaultyData(fmt.Sprintf("Invalid edge: %s", line))
	}
	vertex1 := antfarm.Graph.FindVertex(v1Name)
	vertex2 := antfarm.Graph.FindVertex(v2Name)
	if _, exists := vertex1.Edges[vertex2]; !exists {
		vertex1.Edges[vertex2] = &antfarm.PathProcessing{}
		vertex2.Edges[vertex1] = &antfarm.PathProcessing{}
		vertex1.Sorted = append(vertex1.Sorted, vertex2)
		vertex2.Sorted = append(vertex2.Sorted, vertex1)
		antfarm.Graph.Edges[v1Name] = append(antfarm.Graph.Edges[v1Name], v2Name)
		antfarm.Graph.Edges[v2Name] = append(antfarm.Graph.Edges[v2Name], v1Name)
		pair := antfarm.EdgePair{
			V1Name: v1Name,
			V2Name: v2Name,
		}
		antfarm.EdgePairs = append(antfarm.EdgePairs, pair)
	}
}

func printInput(fileName string, webVisualisation bool) {
	if !webVisualisation {
		input, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading file")
			os.Exit(0)
		}
		spec := string(input)
		i := strings.Index(spec, "\n")
		fmt.Printf("%d%s\n\n", antfarm.Ants.Number, spec[i:])
	}
}

func FlagIsPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
