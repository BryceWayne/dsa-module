package graphs

import (
    "fmt"
    "sort"
)

// Graph structure with a map
type Graph struct {
    vertices map[int][]int
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(u, v int) {
    g.vertices[u] = append(g.vertices[u], v)
}

func (g *Graph) PrettyPrint() {
    keys := make([]int, 0, len(g.vertices))
    for k := range g.vertices {
        keys = append(keys, k)
    }
    sort.Ints(keys) // Sort the keys for consistent order

    fmt.Println("Graph:")
    for _, k := range keys {
        fmt.Printf("%d -> ", k)
        for i, v := range g.vertices[k] {
            if i > 0 {
                fmt.Print(", ")
            }
            fmt.Print(v)
        }
        fmt.Println()
    }
}
