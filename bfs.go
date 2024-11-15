// File:		bfs.go
// Created by:	Hoven
// Created on:	2024-11-15
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package pathplaning

import (
	"errors"

	"github.com/go-puzzles/puzzles/pqueue"
)

func BFSSearch(graph Graph, start, goal Point) ([]Point, error) {
	if !graph.IsPointReachable(goal) {
		return nil, errors.New("goal point not reachable")
	}

	if !graph.IsPointReachable(start) {
		return nil, errors.New("start point not reachable")
	}

	if start.Equals(goal) {
		return []Point{start}, nil
	}

	cameFrom := make(map[Point]Point)

	queue := pqueue.NewMemoryQueue[Point]()
	queue.Enqueue(start)
	graph.SetVisited(start)

	for {
		isEmpty, _ := queue.IsEmpty()
		if isEmpty {
			break
		}

		current, _ := queue.Dequeue()

		if current.Equals(goal) {
			path := []Point{}
			way := current
			for way != nil {
				path = append([]Point{way}, path...)
				way = cameFrom[way]
			}
			return path, nil
		}

		for _, next := range graph.Neighbors(current) {
			if graph.IsVisited(next) {
				continue
			}

			cameFrom[next] = current
			graph.SetVisited(next)
			queue.Enqueue(next)
		}
	}

	return nil, nil
}
