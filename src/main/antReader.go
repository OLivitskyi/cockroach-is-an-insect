package main

import (
	"bufio"
	"cockroach/src/main/antfarm"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type antReader struct {
	file *os.File
}

func (ar *antReader) NewReader(filename string) (*antReader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return &antReader{file: file}, nil
}

func (r *antReader) ParseStartOrEndRoom(line string, isStart bool) *antfarm.Room {
	fields := strings.Fields(line)
	room := &antfarm.Room{ID: fields[0], Name: fields[0], IsStart: isStart, IsEnd: !isStart}
	return room
}

func (r *antReader) ParseTunnel(line string, rooms []*antfarm.Room) []*antfarm.Room {
	ids := strings.Split(line, "-")
	var room1, room2 *antfarm.Room

	for _, room := range rooms {
		if room.ID == ids[0] {
			room1 = room
		} else if room.ID == ids[1] {
			room2 = room
		}
	}

	if room1 == nil || room2 == nil {
		return rooms
	}

	room1.NeighbourRooms = append(room1.NeighbourRooms, room2)
	room2.NeighbourRooms = append(room2.NeighbourRooms, room1)

	return rooms
}

func (r *antReader) ParseNormalRoom(line string) *antfarm.Room {
	fields := strings.Fields(line)
	if len(fields) != 3 {
		// ignore this line
		return nil
	}

	room := &antfarm.Room{ID: fields[0], Name: fields[0]}
	return room
}

func (r *antReader) Parse() (*antfarm.AntFarm, error) {
	scanner := bufio.NewScanner(r.file)
	antFarm := &antfarm.AntFarm{}
	var startRoom, endRoom *antfarm.Room

	// Assumption: first line is number of ants
	if scanner.Scan() {
		count, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("can not convert ant count to integer: %v", err)
		}

		for i := 0; i < count; i++ {
			ant := &antfarm.Ant{
				Name: fmt.Sprintf("Ant #%d", i+1),
			}
			antFarm.Ants = append(antFarm.Ants, ant)
		}
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "#") {
			// ignore comments
			continue
		}

		if strings.HasPrefix(line, "##") {
			if strings.HasSuffix(line, "start") && scanner.Scan() {
				startRoom = r.ParseStartOrEndRoom(scanner.Text(), true)
				antFarm.Rooms = append(antFarm.Rooms, startRoom)
			} else if strings.HasSuffix(line, "end") && scanner.Scan() {
				endRoom = r.ParseStartOrEndRoom(scanner.Text(), false)
				antFarm.Rooms = append(antFarm.Rooms, endRoom)
			}
		} else if strings.Contains(line, "-") {
			antFarm.Rooms = r.ParseTunnel(line, antFarm.Rooms)
		} else {
			room := r.ParseNormalRoom(line)
			if room != nil {
				antFarm.Rooms = append(antFarm.Rooms, room)
			}
		}
	}

	// start and end room search based on ID now
	startRoom, endRoom = getStartAndEndRoom(antFarm)
	if startRoom == nil || endRoom == nil {
		return nil, fmt.Errorf("start or end room not found")
	}

	return antFarm, nil
}

func (r *antReader) Close() error {
	return r.file.Close()
}
