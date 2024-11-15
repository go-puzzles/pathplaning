// File:		astar.go
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

type AstarGraph interface {
	Graph
	Heuristic(from, to Point) int
}

type SimpleAstarGraph struct {
	*SimpleGraph
}

func (g *SimpleAstarGraph) Heuristic(from, to Point) int {
	return abs(from.GetX()-to.GetX()) + abs(from.GetY()-to.GetY())
}

type priorityPoint struct {
	Point
	priority int
}

func (p *priorityPoint) Priority() int {
	return p.priority
}

func AstarSearch(graph AstarGraph, start, goal Point) ([]Point, error) {
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
	pointCost := make(map[Point]int)
	pointCost[start] = 0

	queue := pqueue.NewPriorityQueue[*priorityPoint](pqueue.WithPriorityMode(1))
	queue.Enqueue(&priorityPoint{start, 0})

	const maxIterations = 10000
	iterCount := 0

	for {
		iterCount++
		if iterCount > maxIterations {
			return nil, errors.New("exceeded maximum iterations")
		}

		isEmpty, _ := queue.IsEmpty()
		if isEmpty {
			break
		}

		popItem, _ := queue.Dequeue()
		current := popItem.Point

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
			newCost := pointCost[current] + graph.Cost(current, next)
			if _, exists := pointCost[next]; !exists || newCost < pointCost[next] {
				cameFrom[next] = current
				pointCost[next] = newCost
				priority := newCost + graph.Heuristic(next, goal)

				queue.Enqueue(&priorityPoint{next, priority})
			}
		}
	}

	return nil, nil
}