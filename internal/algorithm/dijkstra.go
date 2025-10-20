package algorithm

import (
	"dijkstra-visualizer/internal/graph"
	"math"
)

type Step struct {
	Distances      map[string]int
	Visited        map[string]bool
	CurrentVertex  string
	PreviousVertex map[string]string
	ExploringEdges []EdgeExploration
	Description    string
}

type EdgeExploration struct {
	From   string
	To     string
	Weight int
}

func Dijkstra(g *graph.Graph, start string) []Step {
	distances := make(map[string]int)
	visited := make(map[string]bool)
	previous := make(map[string]string)
	steps := []Step{}

	for _, v := range g.GetAllVertices() {
		distances[v] = math.MaxInt32
		visited[v] = false
	}
	distances[start] = 0

	steps = append(steps, Step{
		Distances:      copyMap(distances),
		Visited:        copyBoolMap(visited),
		CurrentVertex:  start,
		PreviousVertex: copyStringMap(previous),
		Description:    "Initial state",
	})

	pq := graph.NewPriorityQueue()
	pq.Insert(start, 0)

	for !pq.IsEmpty() {
		current := pq.ExtractMin()
		currentVertex := current.Vertex

		if visited[currentVertex] {
			continue
		}

		visited[currentVertex] = true

		steps = append(steps, Step{
			Distances:      copyMap(distances),
			Visited:        copyBoolMap(visited),
			CurrentVertex:  currentVertex,
			PreviousVertex: copyStringMap(previous),
			Description:    "Visiting vertex " + currentVertex,
		})

		exploringEdges := []EdgeExploration{}
		for _, edge := range g.GetNeighbors(currentVertex) {
			if visited[edge.To] {
				continue
			}

			newDist := distances[currentVertex] + edge.Weight
			exploringEdges = append(exploringEdges, EdgeExploration{
				From:   currentVertex,
				To:     edge.To,
				Weight: edge.Weight,
			})

			if newDist < distances[edge.To] {
				distances[edge.To] = newDist
				previous[edge.To] = currentVertex
				pq.Insert(edge.To, newDist)

				steps = append(steps, Step{
					Distances:      copyMap(distances),
					Visited:        copyBoolMap(visited),
					CurrentVertex:  currentVertex,
					PreviousVertex: copyStringMap(previous),
					ExploringEdges: exploringEdges,
					Description:    "Updated distance to " + edge.To,
				})
			}
		}
	}

	steps = append(steps, Step{
		Distances:      copyMap(distances),
		Visited:        copyBoolMap(visited),
		CurrentVertex:  "",
		PreviousVertex: copyStringMap(previous),
		Description:    "Algorithm complete",
	})

	return steps
}

func copyMap(m map[string]int) map[string]int {
	result := make(map[string]int)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func copyBoolMap(m map[string]bool) map[string]bool {
	result := make(map[string]bool)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func copyStringMap(m map[string]string) map[string]string {
	result := make(map[string]string)
	for k, v := range m {
		result[k] = v
	}
	return result
}
