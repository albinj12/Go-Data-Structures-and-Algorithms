package main

import (
	"fmt"
)

type vertex struct {
	data   int
	vertex []*vertex
}

type graph struct {
	vertices []*vertex
	isDirected bool
}

func main() {
	g := graph{isDirected: true,}
	for i := 0; i < 4; i++ {
		g.addVertex(i)
	}
	g.addEdge(1, 2)
	g.addEdge(2, 0)
	g.addEdge(0, 2)
	g.addEdge(0, 1)
	g.addEdge(2, 3)
	g.addEdge(3, 3)
	g.adjacencyList()

	depthFirstTraversal(g.vertices[1], map[*vertex]bool{})
	fmt.Println()
	breadthFirstTraversal(g.vertices[2])

}

func (g *graph) addVertex(value int) {
	if vertexExists(g.vertices, value) {
		fmt.Println("Vertex already exists")
		return
	}
	g.vertices = append(g.vertices, &vertex{data: value})
}

func (g *graph) addEdge(from, to int) {
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	if fromVertex == nil || toVertex == nil {
		fmt.Println("Cannot add edge")
	} else if vertexExists(fromVertex.vertex, toVertex.data) {
		fmt.Println("Vertex already exists")
	} else {
		fromVertex.vertex = append(fromVertex.vertex, toVertex)
		if !g.isDirected{
			toVertex.vertex = append(toVertex.vertex, fromVertex)
		}
	}
}

func vertexExists(vertices []*vertex, value int) bool {
	for _, v := range vertices {
		if v.data == value {
			return true
		}
	}
	return false
}

func (g *graph) getVertex(value int) *vertex {
	for i, v := range g.vertices {
		if v.data == value {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *graph) adjacencyList() {
	for _, v := range g.vertices {
		fmt.Printf("%v :", v.data)
		for _, v := range v.vertex {
			fmt.Printf("%v ", v.data)
		}
		fmt.Println()
	}
}

func depthFirstTraversal(startingVertex *vertex, visitedVerticesDFS map[*vertex]bool) {
	visitedVerticesDFS[startingVertex] = true
	fmt.Printf("%v ", startingVertex.data)
	for _, v := range startingVertex.vertex {
		if !visitedVerticesDFS[v] {
			depthFirstTraversal(v, visitedVerticesDFS)
		}
	}
}

func breadthFirstTraversal(startingVertex *vertex) {
	visitedVerticesBFS := make(map[*vertex]bool)
	visitedVerticesBFS[startingVertex] = true
	adjacentVertices := []*vertex{startingVertex}
	for len(adjacentVertices) > 0 {
		currentVertex := adjacentVertices[0]
		fmt.Printf("%v ", currentVertex.data)
		adjacentVertices = append(adjacentVertices[:0], adjacentVertices[1:]...)
		for _, v := range currentVertex.vertex {
			if !visitedVerticesBFS[v] {
				visitedVerticesBFS[v] = true
				adjacentVertices = append(adjacentVertices, v)
			}
		}
	}
}
