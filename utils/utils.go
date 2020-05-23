package utils

import (
	"strconv"
	"strings"
)

func SumNumberInteger(number int) int {
	sum := 0
	numbers := splitString(strconv.Itoa(number))
	for _, n := range numbers {
		numberConvert, _ := strconv.Atoi(n)
		sum += numberConvert
	}

	if isMoreDigits(sum) {
		sum = SumNumberInteger(sum)
	}
	return sum
}

func isMoreDigits(number int) bool {
	return len(splitString(strconv.Itoa(number))) > 1
}

func splitString(text string) []string {
	return strings.Split(text, "")
}
