package main

type Color int

const (
	WHITE Color = iota // Initial
	GRAY               // Processing
	BLACK              // Processed
)

func DFS2(edges [][]int, colors []Color, vertexIdx int) bool {
	if colors[vertexIdx] == GRAY {
		return true
	}

	colors[vertexIdx] = GRAY

	for _, nextVertexIdx := range edges[vertexIdx] {
		if DFS2(edges, colors, nextVertexIdx) {
			return true
		}
	}

	colors[vertexIdx] = BLACK

	return false
}

func CycleInGraph2(edges [][]int) bool {

	colors := make([]Color, len(edges))

	for idx, _ := range edges {
		if DFS2(edges, colors, idx) {
			return true
		}
	}

	return false
}
