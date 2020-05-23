package service

import (
	"strings"
)

type Personality struct{}

func (Personality) CalculateNumberPersonality(name string) int {
	consonants := mapConsonants()
	nameLetters := strings.Split(name, "")

	numberPersonality := CalculateNumberLetters(consonants, nameLetters)
	return numberPersonality
}

func mapConsonants() map[string]int {
	var consonants = make(map[string]int)
	consonants["J"] = 1
	consonants["S"] = 1
	consonants["B"] = 2
	consonants["K"] = 2
	consonants["T"] = 2
	consonants["C"] = 3
	consonants["L"] = 3
	consonants["D"] = 4
	consonants["M"] = 4
	consonants["V"] = 4
	consonants["N"] = 5
	consonants["Ñ"] = 5
	consonants["W"] = 5
	consonants["F"] = 6
	consonants["X"] = 6
	consonants["G"] = 7
	consonants["P"] = 7
	consonants["Y"] = 7
	consonants["H"] = 8
	consonants["Q"] = 8
	consonants["Z"] = 8
	consonants["R"] = 9
	return consonants
}
