package pathplaning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAstarSearch(t *testing.T) {
	graph := NewSimpleGraph(7, 7)

	start := &SimplePoint{X: 0, Y: 0}
	goal := &SimplePoint{X: 6, Y: 5}

	path, err := AstarSearch(&SimpleAstarGraph{SimpleGraph: graph}, start, goal)
	if !assert.Nil(t, err) {
		return
	}

	t.Log(path)
	checkPass(t, start, goal, path)
}

func TestAstarSearch_NoPath(t *testing.T) {
	graph := NewSimpleGraph(3, 3)
	graph.SetBlock(&SimplePoint{X: 1, Y: 1})

	start := &SimplePoint{X: 0, Y: 0}
	goal := &SimplePoint{X: 2, Y: 2}

	path, err := AstarSearch(&SimpleAstarGraph{SimpleGraph: graph}, start, goal)
	if !assert.Nil(t, err) {
		return
	}
	checkPass(t, start, goal, path)
}
