package antfarm_test

import (
	"testing"

	"cockroach/src/main/antfarm"
	"github.com/stretchr/testify/assert"
)

func TestDistribute(t *testing.T) {
	antfarm.Ants.Number = 5
	antfarm.Ants.All = map[string]*antfarm.Ant{
		"1": {Name: "1"},
		"2": {Name: "2"},
		"3": {Name: "3"},
		"4": {Name: "4"},
		"5": {Name: "5"},
	}
	antfarm.Graph.Paths = [][]*antfarm.GraphVertex{
		{{}, {}, {}},
		{{}, {}},
		{{}},
	}
	antfarm.Graph.Turns = 5

	antfarm.Ants.Distribute()

	assert.Equal(t, 4, len(antfarm.Ants.Queues[0]))
	assert.Equal(t, 1, len(antfarm.Ants.Queues[1]))
	assert.Equal(t, 0, len(antfarm.Ants.Queues[2]))
}
