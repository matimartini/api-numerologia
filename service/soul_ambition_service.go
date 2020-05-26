package service

import (
	"strings"

	"github.com/matimartini/api-numerologia/database"
)

type SoulAmbition struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (s *SoulAmbition) CalculateNumberSoulAmbition(name string) {
	vowels := mapVowels()
	nameLetters := strings.Split(name, "")

	numberSoulAmbition := CalculateNumberLetters(vowels, nameLetters)
	document := database.GetDesciption(numberSoulAmbition, "soul-ambition")

	var c SoulAmbition
	document.DataTo(&c)

	s.Description = c.Description
	s.Number = numberSoulAmbition
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
