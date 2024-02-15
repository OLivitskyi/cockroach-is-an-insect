package antfarm

import (
	"errors"
	"fmt"
)

type AntFarm struct {
	Rooms []*Room
	Ants  []*Ant
}

func (a *AntFarm) GetStartRoom() (*Room, error) {
	for _, room := range a.Rooms {
		if room.IsStart {
			return room, nil
		}
	}
	return nil, errors.New("start room not found")
}

func (a *AntFarm) GetEndRoom() (*Room, error) {
	for _, room := range a.Rooms {
		if room.IsEnd {
			return room, nil
		}
	}
	return nil, errors.New("end room not found")
}

func (a *AntFarm) BreadthFirstSearch() error {
	startRoom, err := a.GetStartRoom()
	if err != nil {
		return fmt.Errorf("could not find start room: %v", err)
	}

	endRoom, err := a.GetEndRoom()
	if err != nil {
		return fmt.Errorf("could not find end room: %v", err)
	}

	queue := []*Room{startRoom}

	for len(queue) > 0 {
		room := queue[0]
		queue = queue[1:]

		room.Visited = true

		for _, neighbour := range room.NeighbourRooms {
			if !neighbour.Visited {
				neighbour.Distance = room.Distance + 1
				neighbour.Parent = room

				if neighbour == endRoom {
					return nil
				}

				queue = append(queue, neighbour)
			}
		}
	}

	return errors.New("no path found from start to end")
}

func (a *AntFarm) MoveAnts() {
	steps := 0

	for {
		movedAnts := 0

		for _, ant := range a.Ants {
			if ant.CurrentRoom == nil {
				ant.Move(a.Rooms[0])
				movedAnts++
			} else if ant.CurrentRoom.IsEnd {
				continue
			} else {
				nextRoom := ant.CurrentRoute[0]
				ant.CurrentRoute = ant.CurrentRoute[1:]
				if len(nextRoom.Ants) == 0 {
					ant.Move(nextRoom)
					movedAnts++
				}
			}
		}

		if movedAnts == 0 {
			break
		}

		fmt.Println(a)

		steps++
	}

	fmt.Printf("Completed in %d steps\n", steps)
}
