package main

import(
	"strings"
)

func instertChars(str string) string {
	var b strings.Builder
	for _, c := range str {
		b.WriteRune('#')
		b.WriteRune(c)
	}
	
	b.WriteRune('#')
	
	return b.String()
}

// Manacher's algorithm
// https://en.wikipedia.org/wiki/Longest_palindromic_substring#Manacher's_algorithm
func LongestPalindromicSubstring(str string) string {
	if len(str) == 0 {
		return ""
	}
	
	newStr := instertChars(str)
	
	P := make([]int, len(newStr))
	
	center := 0
	radius := 0
	maxRadiusCenter := 0
	maxRadius := 0
	
	for center < len(newStr) {
		// Expand from C
		for center - radius - 1 >= 0 && center + radius + 1 < len(newStr) && newStr[center - radius - 1] == newStr[center + radius + 1] {
			radius = radius + 1
		}
		
		P[center] = radius
		
		// keep track of maximum
		if radius > maxRadius {
			maxRadius = P[center]
			maxRadiusCenter = center
		}
		
		oldCenter := center
		oldRadius := radius
		radius = 0
		center = center + 1
		
		// use the property of palindromes to expand on the mirror within the allowed radius
		for center <= oldCenter + oldRadius {
			mirrorCenter := oldCenter - (center - oldCenter) 
			mirrorMaxRadius := oldCenter + oldRadius - center
			
			// Case 1: mirror is within the max interval 
			if P[mirrorCenter] < mirrorMaxRadius {
				P[center] = P[mirrorCenter]
				center = center + 1
				
				// Case 2: mirror is not within the max interval
			} else if P[mirrorCenter] > mirrorMaxRadius {
				P[center] = mirrorMaxRadius
				center = center + 1
				
				// Case 3: extends exactly up to the border, continue expanding from here
				//         with the external loop
				// P[mirrorC] == mirrorMaxRadius
			} else {
				radius = mirrorMaxRadius
				break
			}
		}
	}
	
	return strings.ReplaceAll(newStr[maxRadiusCenter - P[maxRadiusCenter]:maxRadiusCenter + P[maxRadiusCenter]], "#", "")
}