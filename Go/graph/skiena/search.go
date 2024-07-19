package main

type SearchData struct {
	Discovered map[int]struct{}

	ProcessVertexEarly func(v int)

	ProcessVertexLate func(v int)

	ProcessEdge func(x, y int)
}
