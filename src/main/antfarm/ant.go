package antfarm

type Ant struct {
	Name         string
	CurrentRoom  *Room
	CurrentRoute []*Room
}

func (a *Ant) Move(room *Room) {
	// Update the current room
	a.CurrentRoom = room
	// Remove the room from the route
	if len(a.CurrentRoute) > 0 {
		a.CurrentRoute = a.CurrentRoute[1:]
	}
}
func (a *Ant) GetName() string {
	return a.Name
}
