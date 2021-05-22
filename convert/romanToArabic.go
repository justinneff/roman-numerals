package convert

import (
	"regexp"
	"strings"
)

var conversionMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func CanConvertToArabic(roman string) bool {
	re := regexp.MustCompile("[^CDILMVX]")
	return !re.MatchString(strings.ToUpper(roman))
}

func ToArabic(roman string) int {
	sum := 0
	roman = strings.ToUpper(roman)
	lastVal := 0

	for i := len(roman) - 1; i >= 0; i-- {
		char := string(roman[i])
		val := conversionMap[char]
		if val < lastVal {
			sum -= val
		} else {
			sum += val
		}
		lastVal = val
	}

	return sum
}
