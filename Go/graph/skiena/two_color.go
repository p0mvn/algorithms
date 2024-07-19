package main

var color = make(map[int]colorType)

type colorType int

const (
	colorUnknown colorType = -1
	colorWhite   colorType = 0
	colorBlack   colorType = 1
)

func twocolor(graph AdjacencyListGraph) {
	searchData := &SearchData{
		Discovered:         make(map[int]struct{}, len(graph.uniqueVertices)),
		ProcessVertexEarly: defaultProcessVertexEarly,
		ProcessVertexLate:  defaultProcessVertexLate,
		ProcessEdge:        twoColorProcessEdge,
	}

	for i := range graph.uniqueVertices {
		color[i] = colorUnknown
	}

	for i := range graph.uniqueVertices {
		_, ok := searchData.Discovered[i]
		if !ok {
			color[i] = colorWhite
			BFS(graph, searchData, i)
		}
	}

	printColors()
}

func twoColorProcessEdge(x, y int) {
	if color[x] == color[y] {
		println("Not two colorable")
	} else {
		color[y] = complement(color[x])
	}
}

func complement(c colorType) colorType {
	if c == colorWhite {
		return colorBlack
	}
	return colorWhite
}

func printColors() {
	println("Colors:")
	for i, c := range color {
		println(i, ":", c)
	}
}
