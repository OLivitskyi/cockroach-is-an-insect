package antfarm

type Room struct {
	ID             string
	Name           string
	IsStart        bool
	IsEnd          bool
	Ants           []*Ant
	NeighbourRooms []*Room
	Distance       int
	Parent         *Room
	Visited        bool
}
