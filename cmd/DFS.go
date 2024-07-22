package main

type Graph struct {
	vertices int
	adjList  map[int][]int
}

func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adjList:  make(map[int][]int),
	}
}

func (g *Graph) AddEdge(source, dest int) {
	g.adjList[source] = append(g.adjList[source], dest)
	g.adjList[dest] = append(g.adjList[dest], source)
}

func (g *Graph) DFSUtil(vertex int, visited map[int]bool) {
	visited[vertex] = true

	for _, v := range g.adjList[vertex] {
		if !visited[v] {
			g.DFSUtil(v, visited)
		}
	}
}

func (g *Graph) DFS(startVertex int) {
	visited := make(map[int]bool)
	g.DFSUtil(startVertex, visited)
}
