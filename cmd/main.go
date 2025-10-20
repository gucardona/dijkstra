package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"dijkstra-visualizer/internal/algorithm"
	"dijkstra-visualizer/internal/graph"
	"dijkstra-visualizer/internal/visualizer"
)

func main() {
	if err := os.MkdirAll("output", 0755); err != nil {
		log.Fatal(err)
	}

	g := createSampleGraph()
	start := "A"

	fmt.Println("========================================")
	fmt.Println("   DIJKSTRA'S ALGORITHM VISUALIZATION")
	fmt.Println("========================================")
	fmt.Printf("Start vertex: %s\n", start)
	fmt.Println()

	steps := algorithm.Dijkstra(g, start)

	printStepByStep(steps)
	printResults(steps[len(steps)-1])

	fmt.Printf("\nGenerated %d visualization steps\n", len(steps))

	tempDir := "temp_frames"
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n----------------------------------------")
	fmt.Println("Generating visualization frames...")
	if err := visualizer.GenerateFrames(g, steps, tempDir); err != nil {
		log.Fatal(err)
	}

	outputGIF := filepath.Join("output", "dijkstra_animation.gif")
	fmt.Println("Creating animated GIF...")
	if err := visualizer.CreateGIF(tempDir, outputGIF, 2000); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cleaning up temporary DOT files...")
	if err := visualizer.CleanupDotFiles(tempDir); err != nil {
		log.Printf("Warning: failed to clean up DOT files: %v", err)
	}

	fmt.Println("\n========================================")
	fmt.Printf("✓ Animation created: %s\n", outputGIF)
	fmt.Printf("✓ Step frames saved: %s\n", tempDir)
	fmt.Println("========================================")
}

func createSampleGraph() *graph.Graph {
	g := graph.NewGraph()

	g.AddEdge("A", "B", 7)
	g.AddEdge("A", "C", 9)
	g.AddEdge("A", "F", 14)
	g.AddEdge("B", "C", 10)
	g.AddEdge("B", "D", 15)
	g.AddEdge("C", "D", 11)
	g.AddEdge("C", "F", 2)
	g.AddEdge("D", "E", 6)
	g.AddEdge("E", "F", 9)

	return g
}

func printStepByStep(steps []algorithm.Step) {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║        ALGORITHM STEP-BY-STEP          ║")
	fmt.Println("╚════════════════════════════════════════╝")

	for i, step := range steps {
		fmt.Printf("\n[Step %d] %s\n", i, step.Description)
		fmt.Println(strings.Repeat("-", 50))

		if step.CurrentVertex != "" {
			fmt.Printf("→ Current vertex: %s\n", step.CurrentVertex)

			if dist, ok := step.Distances[step.CurrentVertex]; ok {
				if dist == math.MaxInt32 {
					fmt.Printf("  Distance from start: ∞\n")
				} else {
					fmt.Printf("  Distance from start: %d\n", dist)
				}
			}
		}

		if len(step.ExploringEdges) > 0 {
			fmt.Println("\n  Exploring edges:")
			for _, edge := range step.ExploringEdges {
				oldDist := "∞"
				if d := step.Distances[edge.To]; d != math.MaxInt32 {
					oldDist = fmt.Sprintf("%d", d)
				}

				newDist := step.Distances[step.CurrentVertex] + edge.Weight
				fmt.Printf("    %s → %s (weight: %d)\n", edge.From, edge.To, edge.Weight)
				fmt.Printf("      Old distance to %s: %s\n", edge.To, oldDist)
				fmt.Printf("      New distance to %s: %d\n", edge.To, newDist)

				if step.Distances[edge.To] == newDist && newDist < math.MaxInt32 {
					fmt.Printf("      ✓ Updated! New path through %s\n", edge.From)
				}
			}
		}

		visited := []string{}
		unvisited := []string{}
		for v, vis := range step.Visited {
			if vis {
				visited = append(visited, v)
			} else {
				unvisited = append(unvisited, v)
			}
		}
		sort.Strings(visited)
		sort.Strings(unvisited)

		if len(visited) > 0 {
			fmt.Printf("\n  Visited: %s\n", strings.Join(visited, ", "))
		}
		if len(unvisited) > 0 {
			fmt.Printf("  Unvisited: %s\n", strings.Join(unvisited, ", "))
		}

		fmt.Println("\n  Current distances:")
		vertices := make([]string, 0, len(step.Distances))
		for v := range step.Distances {
			vertices = append(vertices, v)
		}
		sort.Strings(vertices)

		for _, v := range vertices {
			dist := step.Distances[v]
			distStr := "∞"
			if dist != math.MaxInt32 {
				distStr = fmt.Sprintf("%d", dist)
			}

			prev := "-"
			if p, ok := step.PreviousVertex[v]; ok && p != "" {
				prev = p
			}

			marker := " "
			if v == step.CurrentVertex {
				marker = "►"
			} else if step.Visited[v] {
				marker = "✓"
			}

			fmt.Printf("    %s %s: distance=%s, previous=%s\n", marker, v, distStr, prev)
		}
	}
}

func printResults(finalStep algorithm.Step) {
	fmt.Println("\n╔════════════════════════════════════════╗")
	fmt.Println("║           FINAL RESULTS                ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println("\nShortest paths from start vertex:")
	fmt.Println(strings.Repeat("=", 50))

	vertices := make([]string, 0, len(finalStep.Distances))
	for v := range finalStep.Distances {
		vertices = append(vertices, v)
	}
	sort.Strings(vertices)

	for _, v := range vertices {
		dist := finalStep.Distances[v]
		if dist == math.MaxInt32 {
			fmt.Printf("\n  %s: UNREACHABLE\n", v)
		} else {
			path := reconstructPath(v, finalStep.PreviousVertex)
			fmt.Printf("\n  %s:\n", v)
			fmt.Printf("    Distance: %d\n", dist)
			fmt.Printf("    Path: %s\n", path)
		}
	}
	fmt.Println()
}

func reconstructPath(vertex string, previous map[string]string) string {
	path := []string{vertex}
	current := vertex

	for {
		prev, exists := previous[current]
		if !exists || prev == "" {
			break
		}
		path = append([]string{prev}, path...)
		current = prev
	}

	return strings.Join(path, " → ")
}
