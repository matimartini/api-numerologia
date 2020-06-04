package service

import (
	"fmt"
	"strings"

	"github.com/matimartini/api-numerologia/database"
)

type SoulAmbition struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (soulAmbition *SoulAmbition) CalculateNumberSoulAmbition(name string) {
	vowels := mapVowels()
	nameLetters := strings.Split(name, "")

	numberSoulAmbition := CalculateNumberLetters(vowels, nameLetters)
	document := database.GetDesciption(numberSoulAmbition, "soul-ambition")

	soulAmbition.Number = numberSoulAmbition
	document.DataTo(&soulAmbition)

	if soulAmbition.Description == "" {
		fmt.Println("Error empty description soulAmbition. number: ", numberSoulAmbition)
	}
}

func mapVowels() map[string]int {
	var vowels = make(map[string]int)
	vowels["A"] = 1
	vowels["E"] = 5
	vowels["I"] = 9
	vowels["O"] = 6
	vowels["U"] = 3
	return vowels
}
