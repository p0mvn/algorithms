package main

import (
	"strings"
)

const asciiLetterBase = 97
const asciiNumBase = 48

func digitToStartAscii(d int) int  {
	switch d {
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 5:
			fallthrough
		case 6:
			fallthrough
		case 7:
			return asciiLetterBase + (d - 2) * 3
		case 8:
			fallthrough
		case 9:
			return asciiLetterBase + (d - 3) * 3 + 4
		case 10:
			return digitToStartAscii(9) + 4
		default:
	}
	return asciiNumBase + d
}

func constructAllLettersForDigit(d int) []string {
	start := digitToStartAscii(d)
	result := []string{string(rune(start))}
	if d == 0 || d == 1 {
		return result
	}
	
	for i := start + 1; i < digitToStartAscii(d + 1); i++ {
		result = append(result, string(rune(i)))
	}
	return result
}

func constructMnemonics(phoneNumber string, curIdx int)[]string {
	if curIdx == len(phoneNumber) - 1 {
		return constructAllLettersForDigit(int(phoneNumber[curIdx]) - asciiNumBase)
	}
	
	curDigitMnemonics := constructAllLettersForDigit(int(phoneNumber[curIdx]) - asciiNumBase)
	nextResult := constructMnemonics(phoneNumber, curIdx + 1)
	result := make([]string, 0)
	
	for _, cur := range curDigitMnemonics {
		var sb strings.Builder
		
		for _, next := range nextResult {
			sb.Reset()
			sb.WriteString(cur)
			sb.WriteString(next)
			result = append(result, sb.String())
		}
	}
	
	return result
}

func PhoneNumberMnemonics(phoneNumber string) []string {
	return constructMnemonics(phoneNumber, 0)
}
