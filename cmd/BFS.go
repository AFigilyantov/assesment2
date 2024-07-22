package main

import (
	"math"
	"slices"
)

type GraphLists struct {
	// first key - vertex
	// second key - adjacent vertex to prev vertex
	// value - weight of edge
	adjList map[int]map[int]int
}
type Dist struct {
	Dist     int
	EdgeFrom int
}

func (g *GraphLists) BFSShort(start, target int) ([]int, int) {
	d := make([]Dist, len(g.adjList))
	for i := range d {
		d[i].Dist = math.MaxInt32
		d[i].EdgeFrom = -1
	}
	d[start].Dist = 0
	var queue []int
	queue = append(queue, start)
	for len(queue) != 0 {
		currentVert := queue[0]
		for adjVert, weight := range g.adjList[currentVert] {
			if d[adjVert].Dist > d[currentVert].Dist+weight {
				d[adjVert].EdgeFrom = currentVert
				d[adjVert].Dist = d[currentVert].Dist + weight
				queue = append(queue, adjVert)
			}
		}
		queue = queue[1:]
	}
	if d[target].EdgeFrom == -1 {
		// there is no path to target vertex
		return []int{}, 0
	}
	minDist := d[target].Dist
	var path []int
	for target != -1 {
		path = append(path, target)
		target = d[target].EdgeFrom
	}
	slices.Reverse(path)
	return path, minDist
}

func (g *GraphLists) FillAdjList(matrix [][]int) {

	g.adjList = make(map[int]map[int]int)

	l := len(matrix)

	for i := 0; i < l; i++ {
		g.adjList[i] = make(map[int]int)
	}

	for i, line := range matrix {
		for j, weight := range line {
			if weight != 0 {
				m := g.adjList[i]
				m[j] = weight
			}

		}
	}
}
