package main

import "math"

type Edge struct {
	src, dest, weight int
}

func Find_longest_path(edges []Edge, vertices, source int) []int {
	dist := make([]int, vertices)
	for i := range dist {
		dist[i] = math.MinInt32
	}
	dist[source] = 0

	for i := 0; i < vertices-1; i++ {
		for _, edge := range edges {
			u := edge.src
			v := edge.dest
			w := edge.weight

			if dist[u] != math.MinInt32 && dist[u]+w > dist[v] {
				dist[v] = dist[u] + w
			}
		}
	}
	return dist
}

func NewEdgeList(matrix [][]int) []Edge {
	edgeList := make([]Edge, 0)

	for i, line := range matrix {
		for j, dist := range line {
			if dist == 0 {
				continue
			}
			edgeList = append(edgeList, Edge{src: i, dest: j, weight: dist})
		}
	}
	return edgeList
}
