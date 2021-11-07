package main

type OrgChart struct {
	Name          string
	DirectReports []*OrgChart
}

type Result struct {
	lowestCommon *OrgChart
	reportCount int
}

func getLowestCommonManager(cur, reportOne, reportTwo *OrgChart) *Result {
	reportCountSoFar := 0
	var lowestCommon *OrgChart = nil
	for _, dr := range cur.DirectReports {
		result := getLowestCommonManager(dr, reportOne, reportTwo)
		
		if result.lowestCommon != nil {
			return result
		}
		reportCountSoFar += result.reportCount
		
		if reportCountSoFar == 2 {
			lowestCommon = cur
			break
		}
	}
	
	if cur == reportOne || cur == reportTwo {
		reportCountSoFar++
	}

	if reportCountSoFar == 2 {
		lowestCommon = cur
	}
	
	return &Result{
		lowestCommon,
		reportCountSoFar,
	}
}

// Problem: find lowst common manager of the two reports
//
// Solution: note that the lowest common manager is the root of the smallest subtree
// that contains both reports. Therfore, we can search for a subtree with the 
// root satisfying the criteria above
func GetLowestCommonManager(org, reportOne, reportTwo *OrgChart) *OrgChart {
	return getLowestCommonManager(org, reportOne, reportTwo).lowestCommon
}
