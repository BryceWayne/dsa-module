package graphs

import "fmt"

// BFS performs a Breadth-First Search on the graph starting from vertex `start`.
func (g *Graph) BFS(start int) {
    visited := make(map[int]bool)
    queue := []int{start}

    for len(queue) > 0 {
        vertex := queue[0]
        queue = queue[1:]

        if !visited[vertex] {
            visited[vertex] = true
            fmt.Print(vertex, " ")

            for _, neighbor := range g.vertices[vertex] {
                if !visited[neighbor] {
                    queue = append(queue, neighbor)
                }
            }
        }
    }
    fmt.Println()
}
