package main

import (
	"dsa/graphs"
	"fmt"
)

func main() {
	wg := graphs.NewWeightedGraph()
	wg.AddEdge(0, 1, 4)
	wg.AddEdge(0, 2, 1)
	wg.AddEdge(2, 1, 2)
	wg.AddEdge(1, 3, 1)
	wg.AddEdge(2, 3, 5)

	wg.PrettyPrint()

	distances := wg.Dijkstra(0)
	fmt.Println("Shortest distances from node 0:", distances)
}
