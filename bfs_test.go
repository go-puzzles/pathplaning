// File:		bfs_test.go
// Created by:	Hoven
// Created on:	2024-11-15
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package pathplaning

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkPass(t *testing.T, start, goal Point, path []Point) {
	assert.NotEmpty(t, path, "Expected a non-empty path")

	assert.Equal(t, start, path[0], "The first point of the path should be the start point")
	assert.Equal(t, goal, path[len(path)-1], "The last point of the path should be the goal point")

	for i := 0; i < len(path)-1; i++ {
		assert.True(t, areAdjacent(path[i], path[i+1]), "Points in the path should be adjacent")
	}
}

func TestBfs(t *testing.T) {
	graph := NewSimpleGraph(6, 6)

	graph.SetBlock(&SimplePoint{X: 1, Y: 0})
	graph.SetBlock(&SimplePoint{X: 2, Y: 1})
	graph.SetBlock(&SimplePoint{X: 3, Y: 2})

	start := &SimplePoint{X: 0, Y: 0}
	goal := &SimplePoint{X: 4, Y: 3}

	path, err := BFSSearch(graph, start, goal)
	if !assert.Nil(t, err) {
		return
	}
	t.Log(path)
}

func TestBfsWithValidPath(t *testing.T) {
	graph := NewSimpleGraph(6, 6)

	start := &SimplePoint{X: 0, Y: 0}
	goal := &SimplePoint{X: 4, Y: 3}

	path, err := BFSSearch(graph, start, goal)
	if !assert.Nil(t, err) {
		return
	}
	checkPass(t, start, goal, path)
}

func TestBfsWithOnePass(t *testing.T) {
	graph := NewSimpleGraph(6, 6)
	graph.SetBlock(&SimplePoint{X: 3, Y: 2})
	graph.SetBlock(&SimplePoint{X: 3, Y: 4})
	graph.SetBlock(&SimplePoint{X: 2, Y: 3})

	start := &SimplePoint{X: 0, Y: 0}
	goal := &SimplePoint{X: 3, Y: 3}

	path, err := BFSSearch(graph, start, goal)
	if !assert.Nil(t, err) {
		return

	}
	t.Log(path)
	checkPass(t, start, goal, path)
}

func TestBfsWithNoPath(t *testing.T) {
	graph := NewSimpleGraph(6, 6)
	graph.SetBlock(&SimplePoint{X: 3, Y: 2})
	graph.SetBlock(&SimplePoint{X: 3, Y: 4})
	graph.SetBlock(&SimplePoint{X: 2, Y: 3})
	graph.SetBlock(&SimplePoint{X: 4, Y: 3})

	start := &SimplePoint{X: 0, Y: 0}
	goal := &SimplePoint{X: 3, Y: 3}

	path, err := BFSSearch(graph, start, goal)
	if !assert.Nil(t, err) {
		return
	}
	t.Log(path)
	assert.Empty(t, path, "Expected an empty path when there is no valid path")
}

func areAdjacent(p1, p2 Point) bool {
	return (p1.GetX() == p2.GetX() && abs(p1.GetY()-p2.GetY()) == 1) ||
		(p1.GetY() == p2.GetY() && abs(p1.GetX()-p2.GetX()) == 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
