package graphs

import "fmt"

// DFS performs a Depth-First Search on the graph starting from vertex `start`.
func (g *Graph) DFS(start int) {
    visited := make(map[int]bool)
    g.dfsRecursive(start, visited)
    fmt.Println()
}

func (g *Graph) dfsRecursive(v int, visited map[int]bool) {
    visited[v] = true
    fmt.Print(v, " ")

    for _, neighbor := range g.vertices[v] {
        if !visited[neighbor] {
            g.dfsRecursive(neighbor, visited)
        }
    }
}
