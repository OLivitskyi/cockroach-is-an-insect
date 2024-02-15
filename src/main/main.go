package main

import (
	"cockroach/src/main/antfarm"
	"cockroach/src/main/pathfinder"
	"fmt"
	"log"
	os "os"
	"path/filepath"
)

func main() {
	workDir, _ := os.Getwd()
	filePath := filepath.Join(workDir, "resources", os.Args[1])

	reader := &antReader{}
	var err error
	reader, err = reader.NewReader(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	antFarm, err := reader.Parse()
	if err != nil {
		log.Fatal(err)
	}

	startRoom, endRoom := getStartAndEndRoom(antFarm)
	antPaths := createAntPaths(antFarm, startRoom, endRoom)

	fmt.Println("Simulation starts:")
	turn := 0
	for {
		fmt.Printf("Turn %d:\n", turn)
		if !moveAnts(antFarm.Ants, antPaths) {
			break
		}
		turn++
	}
}

func getStartAndEndRoom(antFarm *antfarm.AntFarm) (startRoom, endRoom *antfarm.Room) {
	for _, room := range antFarm.Rooms {
		if room.ID == "1" {
			startRoom = room
		}
		if room.ID == "0" {
			endRoom = room
		}
	}
	return startRoom, endRoom
}

func createAntPaths(antFarm *antfarm.AntFarm, startRoom, endRoom *antfarm.Room) [][]*antfarm.Room {
	antPaths := make([][]*antfarm.Room, len(antFarm.Ants))
	for i, ant := range antFarm.Ants {
		shortestPath := pathfinder.Dijkstra(antFarm, startRoom, endRoom)
		ant.CurrentRoute = shortestPath[1:] // exclude start room
		for _, room := range shortestPath {
			if room != startRoom {
				antPaths[i] = append(antPaths[i], room)
			}
		}
	}
	return antPaths
}

func moveAnts(ants []*antfarm.Ant, paths [][]*antfarm.Room) bool {
	moveMade := false
	for i, ant := range ants {
		if len(paths[i]) > 0 {
			next := paths[i][0]
			ant.Move(next)
			paths[i] = paths[i][1:]
			moveMade = true
		}
	}
	return moveMade
}

func nextStep(room *antfarm.Room, end *antfarm.Room) *antfarm.Room {
	// This is a placeholder function. You might want to replace this code with an actual room picking logic.
	for _, neighbour := range room.NeighbourRooms {
		return neighbour
	}
	return nil
}
