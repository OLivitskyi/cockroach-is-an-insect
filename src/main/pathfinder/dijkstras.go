package pathfinder

import (
	"cockroach/src/main/antfarm"
	"math"
)

func Dijkstra(antFarm *antfarm.AntFarm, startRoom *antfarm.Room, endRoom *antfarm.Room) []*antfarm.Room {
	distance := make(map[*antfarm.Room]int)
	previous := make(map[*antfarm.Room]*antfarm.Room)

	for _, room := range antFarm.Rooms {
		distance[room] = math.MaxInt32
	}

	distance[startRoom] = 0
	queue := []*antfarm.Room{startRoom}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighbour := range current.NeighbourRooms {
			alt := distance[current] + 1 // assuming all rooms are of equal distance
			if alt < distance[neighbour] {
				distance[neighbour] = alt
				previous[neighbour] = current
				queue = append(queue, neighbour)
			}
		}
	}

	path, ok := backtracePath(previous, startRoom, endRoom)
	if !ok {
		return nil
	}

	return path
}

func backtracePath(previous map[*antfarm.Room]*antfarm.Room, startRoom *antfarm.Room, endRoom *antfarm.Room) ([]*antfarm.Room, bool) {
	var path []*antfarm.Room
	current := endRoom

	for current != startRoom {
		path = append([]*antfarm.Room{current}, path...)
		p, ok := previous[current]
		if !ok {
			return nil, false
		}
		current = p
	}

	// Finally, include the start room.
	path = append([]*antfarm.Room{startRoom}, path...)

	return path, true
}
