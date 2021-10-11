package main

func DFS(edges [][]int, parents map[int]interface{}, curVIdx int) bool {
    if _, ok := parents[curVIdx]; ok {
        return true
    }
    parents[curVIdx] = true
    
    nextVs := edges[curVIdx]
    
    for _, nextV := range nextVs {
        detectedCycle := DFS(edges, parents, nextV)
        if detectedCycle {
            return true
        }
    }
    
    delete(parents, curVIdx)
    
    return false
}

func CycleInGraph(edges [][]int) bool {
    
    parents := make(map[int]interface{})
    
    for idx, _ := range edges {
        detectedCycle := DFS(edges, parents, idx)
        if detectedCycle {
            return true
        }
    }
    return false
}
