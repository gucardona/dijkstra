package visualizer

import (
	"dijkstra-visualizer/internal/algorithm"
	"dijkstra-visualizer/internal/graph"
	"fmt"
	"math"
	"strings"
)

func GenerateDOT(g *graph.Graph, step algorithm.Step) string {
	var sb strings.Builder

	sb.WriteString("digraph G {\n")
	sb.WriteString("  layout=neato;\n")
	sb.WriteString("  overlap=false;\n")
	sb.WriteString("  splines=true;\n")
	sb.WriteString("  rankdir=LR;\n")
	sb.WriteString("  node [shape=circle, style=filled, fontsize=16, width=0.8, height=0.8];\n")
	sb.WriteString("  edge [fontsize=14, fontcolor=blue, dir=none];\n")
	sb.WriteString("  labelloc=\"t\";\n")
	sb.WriteString(fmt.Sprintf("  label=\"%s\";\n", step.Description))
	sb.WriteString("  fontsize=20;\n")
	sb.WriteString("  fontname=\"Arial Bold\";\n\n")

	positions := map[string]string{
		"A": "0,2!",
		"B": "2,2.8!",
		"C": "3,2!",
		"D": "5,2.5!",
		"E": "7,2!",
		"F": "4,1.2!",
	}

	for vertex := range g.Vertices {
		color := getVertexColor(vertex, step)
		distLabel := getDistanceLabel(step.Distances[vertex])
		pos := positions[vertex]
		if pos == "" {
			pos = "0,2!"
		}

		sb.WriteString(fmt.Sprintf("  %s [fillcolor=\"%s\", label=\"%s\\n%s\", pos=\"%s\"];\n",
			vertex, color, vertex, distLabel, pos))
	}

	sb.WriteString("\n")

	processedEdges := make(map[string]bool)
	for from, edges := range g.Vertices {
		for _, edge := range edges {
			edgeKey := getEdgeKey(from, edge.To)
			if processedEdges[edgeKey] {
				continue
			}
			processedEdges[edgeKey] = true

			edgeColor := getEdgeColor(from, edge.To, step)
			edgeWidth := getEdgeWidth(from, edge.To, step)
			sb.WriteString(fmt.Sprintf("  %s -> %s [label=\"%d\", color=\"%s\", penwidth=%d];\n",
				from, edge.To, edge.Weight, edgeColor, edgeWidth))
		}
	}

	sb.WriteString("\n")
	sb.WriteString(generateTable(step))

	sb.WriteString("}\n")
	return sb.String()
}

func generateTable(step algorithm.Step) string {
	var sb strings.Builder

	vertices := make([]string, 0, len(step.Distances))
	for v := range step.Distances {
		vertices = append(vertices, v)
	}

	sortVertices := append([]string(nil), vertices...)
	for i := 0; i < len(sortVertices); i++ {
		for j := i + 1; j < len(sortVertices); j++ {
			if sortVertices[i] > sortVertices[j] {
				sortVertices[i], sortVertices[j] = sortVertices[j], sortVertices[i]
			}
		}
	}

	tableLabel := "<<TABLE BORDER=\"2\" CELLBORDER=\"1\" CELLSPACING=\"0\" CELLPADDING=\"12\" BGCOLOR=\"white\">"
	tableLabel += "<TR><TD BGCOLOR=\"#4A90E2\"><FONT COLOR=\"white\" POINT-SIZE=\"18\"><B>Vertex</B></FONT></TD>"
	tableLabel += "<TD BGCOLOR=\"#4A90E2\"><FONT COLOR=\"white\" POINT-SIZE=\"18\"><B>Distance</B></FONT></TD>"
	tableLabel += "<TD BGCOLOR=\"#4A90E2\"><FONT COLOR=\"white\" POINT-SIZE=\"18\"><B>Previous</B></FONT></TD>"
	tableLabel += "<TD BGCOLOR=\"#4A90E2\"><FONT COLOR=\"white\" POINT-SIZE=\"18\"><B>Visited</B></FONT></TD></TR>"

	for _, v := range sortVertices {
		dist := step.Distances[v]
		prev := step.PreviousVertex[v]
		visited := step.Visited[v]

		bgColor := "#FFFFFF"
		if v == step.CurrentVertex {
			bgColor = "#FFD700"
		} else if visited {
			bgColor = "#90EE90"
		}

		distStr := getDistanceLabel(dist)
		prevStr := "-"
		if prev != "" {
			prevStr = prev
		}
		visitedStr := "No"
		if visited {
			visitedStr = "Yes"
		}

		tableLabel += fmt.Sprintf("<TR><TD BGCOLOR=\"%s\"><FONT POINT-SIZE=\"16\"><B>%s</B></FONT></TD>", bgColor, v)
		tableLabel += fmt.Sprintf("<TD BGCOLOR=\"%s\"><FONT POINT-SIZE=\"16\">%s</FONT></TD>", bgColor, distStr)
		tableLabel += fmt.Sprintf("<TD BGCOLOR=\"%s\"><FONT POINT-SIZE=\"16\">%s</FONT></TD>", bgColor, prevStr)
		tableLabel += fmt.Sprintf("<TD BGCOLOR=\"%s\"><FONT POINT-SIZE=\"16\">%s</FONT></TD></TR>", bgColor, visitedStr)
	}

	tableLabel += "</TABLE>>"

	sb.WriteString(fmt.Sprintf("  table [shape=plaintext, label=%s, pos=\"3.5,0!\"];\n", tableLabel))

	return sb.String()
}

func getVertexColor(vertex string, step algorithm.Step) string {
	if vertex == step.CurrentVertex {
		return "#FFD700"
	}
	if step.Visited[vertex] {
		return "#90EE90"
	}
	return "#ADD8E6"
}

func getDistanceLabel(distance int) string {
	if distance == math.MaxInt32 {
		return "∞"
	}
	return fmt.Sprintf("%d", distance)
}

func getEdgeColor(from, to string, step algorithm.Step) string {
	for _, exploring := range step.ExploringEdges {
		if (exploring.From == from && exploring.To == to) ||
			(exploring.From == to && exploring.To == from) {
			return "#FF4500"
		}
	}

	if step.PreviousVertex[to] == from || step.PreviousVertex[from] == to {
		if step.Visited[from] && step.Visited[to] {
			return "#32CD32"
		}
	}

	return "#000000"
}

func getEdgeWidth(from, to string, step algorithm.Step) int {
	for _, exploring := range step.ExploringEdges {
		if (exploring.From == from && exploring.To == to) ||
			(exploring.From == to && exploring.To == from) {
			return 4
		}
	}

	if step.PreviousVertex[to] == from || step.PreviousVertex[from] == to {
		if step.Visited[from] && step.Visited[to] {
			return 3
		}
	}

	return 2
}

func getEdgeKey(v1, v2 string) string {
	if v1 < v2 {
		return v1 + "-" + v2
	}
	return v2 + "-" + v1
}
