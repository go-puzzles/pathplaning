// File:		graph.go
// Created by:	Hoven
// Created on:	2024-11-15
//
// This file is part of the Example Project.
//
// (c) 2024 Example Corp. All rights reserved.

package pathplaning

type Graph interface {
	Neighbors(p Point) []Point
	IsPointReachable(p Point) bool
	IsInGraph(p Point) bool
	IsBlocked(p Point) bool
	SetBlock(p Point)
	IsVisited(p Point) bool
	SetVisited(p Point)
	Cost(from, to Point) int
}

var (
	directions = []SimplePoint{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
)

type SimpleGraph struct {
	graph [][]int
}

func NewSimpleGraph(width, height int) *SimpleGraph {
	graph := make([][]int, height)
	for i := range graph {
		graph[i] = make([]int, width)
	}
	return &SimpleGraph{graph: graph}
}

func (g *SimpleGraph) Neighbors(p Point) []Point {
	var neighbors []Point
	for _, dir := range directions {
		newPoint := p.Shift(dir.GetX(), dir.GetY())
		if !g.IsInGraph(newPoint) || g.IsBlocked(newPoint) {
			continue
		}
		neighbors = append(neighbors, newPoint)
	}
	return neighbors
}

func (g *SimpleGraph) IsPointReachable(p Point) bool {
	return g.IsInGraph(p) && !g.IsBlocked(p)
}

func (g *SimpleGraph) IsInGraph(p Point) bool {
	return p.GetX() >= 0 && p.GetX() < len(g.graph) && p.GetY() >= 0 && p.GetY() < len(g.graph[0])
}

func (g *SimpleGraph) IsBlocked(p Point) bool {
	return g.graph[p.GetX()][p.GetY()] == -1
}

func (g *SimpleGraph) SetBlock(p Point) {
	g.graph[p.GetX()][p.GetY()] = -1
}

func (g *SimpleGraph) SetVisited(p Point) {
	g.graph[p.GetX()][p.GetY()] = 1
}

func (g *SimpleGraph) IsVisited(p Point) bool {
	return g.graph[p.GetX()][p.GetY()] == 1
}

func (g *SimpleGraph) Cost(from, to Point) int {
	return abs(from.GetX()-to.GetX()) + abs(from.GetY()-to.GetY())
}

func (g *SimpleGraph) PrintGraph(start, end Point, path []Point) {
	for _, p := range path {
		g.graph[p.GetX()][p.GetY()] = -100
	}

	for y := 0; y < len(g.graph); y++ {
		for x := 0; x < len(g.graph[0]); x++ {
			p := &SimplePoint{X: x, Y: y}
			if p.Equals(start) {
				print("S ")
			} else if p.Equals(end) {
				print("G ")
			} else if g.IsBlocked(p) {
				print("X ")
			} else if g.graph[p.GetX()][p.GetY()] == -100 {
				print("o ")
			} else {
				print(". ")
			}
		}
		println()
	}
	println()
}
