import "math"

// Dijkstra finds the shortest path from start to all other nodes in the graph.
func (wg *WeightedGraph) Dijkstra(start int) map[int]int {
    distances := make(map[int]int)
    for vertex := range wg.vertices {
        distances[vertex] = math.MaxInt32
    }
    distances[start] = 0

    visited := make(map[int]bool)
    for len(visited) < len(wg.vertices) {
        // Find the vertex with the smallest distance
        minVertex := -1
        minDistance := math.MaxInt32
        for vertex, dist := range distances {
            if !visited[vertex] && dist < minDistance {
                minVertex = vertex
                minDistance = dist
            }
        }

        if minVertex == -1 {
            break // All remaining vertices are inaccessible
        }

        visited[minVertex] = true
        for _, edge := range wg.vertices[minVertex] {
            alt := distances[minVertex] + edge.Weight
            if alt < distances[edge.Node] {
                distances[edge.Node] = alt
            }
        }
    }

    return distances
}
