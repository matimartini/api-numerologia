package service

import (
	"time"

	"github.com/matimartini/api-numerologia/utils"
)

type PathOfLifeService struct{}

func (p PathOfLifeService) CalculateNumberPathOfLife(date string) int {
	dateFormat, _ := time.Parse("2006-01-02", date)

	num := int(dateFormat.Day()) + int(dateFormat.Month()) + int(dateFormat.Year())

	numberPathOfLife := utils.SumNumberInteger(num)
	return numberPathOfLife
}

/*
func sumNumberInteger(number int) int {
	sum := 0
	numbers := splitString(strconv.Itoa(number))
	for _, n := range numbers {
		numberConvert, _ := strconv.Atoi(n)
		sum += numberConvert
	}

	if isMoreDigits(sum) {
		sum = sumNumberInteger(sum)
	}
	return sum
}

func isMoreDigits(number int) bool {
	return len(splitString(strconv.Itoa(number))) > 1
}

func splitString(text string) []string {
	return strings.Split(text, "")
}
*/
