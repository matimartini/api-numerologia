package service

import (
	"strings"

	"github.com/matimartini/api-numerologia/utils"
)

func CalculateNumberLetters(mapa map[string]int, letters []string) int {
	number := 0
	for _, letter := range letters {
		letterTopUpper := strings.ToUpper(letter)
		num := mapa[letterTopUpper]
		number += num
	}
	result := utils.SumNumberInteger(number)
	return result
}