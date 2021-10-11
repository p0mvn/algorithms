package main

type Color int

const(
    WHITE Color = iota // Initial
    GRAY               // Processing
    BLACK              // Processed
)

func DFS (edges [][]int, colors []Color, vertexIdx int) bool {
    if colors[vertexIdx] == GRAY {
        return true
    }
    
    colors[vertexIdx] = GRAY
    
    for _, nextVertexIdx := range edges[vertexIdx] {
        if DFS(edges, colors, nextVertexIdx) {
            return true
        }
    }
    
    colors[vertexIdx] = BLACK
    
    return false
}

func CycleInGraph(edges [][]int) bool {
    
    colors := make([]Color, len(edges))
    
    for idx, _ := range edges {
        if DFS(edges, colors, idx) {
            return true
        }
    }
    
    return false
}
