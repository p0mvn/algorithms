package main

type SearchData struct {
	Discovered map[int]struct{}

	Processed map[int]struct{}

	ProcessVertexEarly func(v int)

	ProcessVertexLate func(v int)

	ProcessEdge func(x, y int)

	DFSData
}

type DFSData struct {
	EntryTime map[int]int
	ExitTime  map[int]int

	Time int
}
