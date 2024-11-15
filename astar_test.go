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
	graph.PrintGraph(start, goal, path)
	t.Log(path)
	checkPass(t, start, goal, path)
}

func TestAstarSearch_NoPath(t *testing.T) {
	graph := NewSimpleGraph(3, 3)
	graph.SetBlock(&SimplePoint{X: 1, Y: 0})
	graph.SetBlock(&SimplePoint{X: 1, Y: 1})
	graph.SetBlock(&SimplePoint{X: 1, Y: 2})

	start := &SimplePoint{X: 0, Y: 0}
	goal := &SimplePoint{X: 2, Y: 2}

	path, err := AstarSearch(&SimpleAstarGraph{SimpleGraph: graph}, start, goal)
	if !assert.Nil(t, err) {
		return
	}
	t.Log(path)
	assert.Empty(t, path)
}

func TestAstarSearchWithBlock(t *testing.T) {
	graph := NewSimpleGraph(10, 10)
	for i := 3; i < 10; i++ {
		graph.SetBlock(&SimplePoint{X: i, Y: 4})
	}

	for j := 5; j < 9; j++ {
		graph.SetBlock(&SimplePoint{X: 6, Y: j})
	}
	graph.SetBlock(&SimplePoint{X: 7, Y: 8})
	graph.SetBlock(&SimplePoint{X: 8, Y: 8})
	graph.SetBlock(&SimplePoint{X: 8, Y: 7})

	/*
		o o o o o o o o o o
		S o o o o o o o o o
		o o o o o o o o o o
		o o o o o o o o o o
		o o o x x x x x x x
		o o o o o o x o o o
		o o o o o o x o o o
		o o o o o o x G o o
		o o o o o o x x x o
		o o o o o o o o o o
	*/

	start := &SimplePoint{X: 0, Y: 1}
	goal := &SimplePoint{X: 7, Y: 7}

	path, err := AstarSearch(&SimpleAstarGraph{SimpleGraph: graph}, start, goal)
	if !assert.Nil(t, err) {
		return
	}
	graph.PrintGraph(start, goal, path)

	t.Log(path)
	checkPass(t, start, goal, path)
}
