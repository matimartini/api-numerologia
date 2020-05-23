package service

import (
	"strings"
)

type SoulAmbitionService struct{}

func (SoulAmbitionService) CalculateNumberSoulAmbition(name string) int {
	vowels := mapVowels()
	nameLetters := strings.Split(name, "")

	numberSoulAmbition := CalculateNumberLetters(vowels, nameLetters)
	return numberSoulAmbition
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
