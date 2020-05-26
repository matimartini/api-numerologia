package service

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/matimartini/api-numerologia/utils"
)

type PathOfLifeService struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (p PathOfLifeService) CalculateNumberPathOfLife(date string) int {
	dateFormat, _ := time.Parse("2006-01-02", date)
	day := sumAllTheDigits(dateFormat.Day())
	month := sumAllTheDigits(int(dateFormat.Month()))
	year := sumAllTheDigits(dateFormat.Year())

	num := day + month + year

	log.Println(num)
	numberPathOfLife := utils.SumNumberInteger(num)
	return numberPathOfLife
}

func sumAllTheDigits(number int) int {
	sum := 0
	digits := strings.Split(strconv.Itoa(number), "")

	for _, n := range digits {
		numberConvert, _ := strconv.Atoi(n)
		sum += numberConvert
	}
	return sum
}
