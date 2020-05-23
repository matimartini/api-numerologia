package service

import (
	"log"

	"github.com/matimartini/api-numerologia/utils"
)

type Person struct {
	FullName          string `json:"full_name"`
	BirthDate         string `json:"birth_date"`
	NumberPathOfLife  int    `json:"number_path_of_life"`
	SoulAmbition      int    `json:"soul_ambition"`
	NumberPersonality int    `json:"number_personality"`
	NumberExpression  int    `json:"number_expression"`
	GoalLife          int    `json:"goal_life"`
}

func (p *Person) NewPerson(name string, date string) {
	log.Println("Ingresa creacion persona!!!   ", name, date)
	p.FullName = name
	p.BirthDate = date
	p.calculateNumberPathOfLife()
	p.calculateSoulAmbition()
	p.calculatePersonality()
	p.calculateNumberExpression()
	p.calculateGoalLife()
}

func (p *Person) calculateNumberPathOfLife() {
	servicePathOfLife := PathOfLifeService{}
	p.NumberPathOfLife = servicePathOfLife.CalculateNumberPathOfLife(p.BirthDate)
}

func (p *Person) calculateSoulAmbition() {
	serviceSoulAmbition := SoulAmbitionService{}
	p.SoulAmbition = serviceSoulAmbition.CalculateNumberSoulAmbition(p.FullName)
}

func (p *Person) calculatePersonality() {
	servicePersonality := Personality{}
	p.NumberPersonality = servicePersonality.CalculateNumberPersonality(p.FullName)
}

func (p *Person) calculateNumberExpression() {
	numberExpression := utils.SumNumberInteger(p.SoulAmbition + p.NumberPersonality)
	p.NumberExpression = numberExpression
}

func (p *Person) calculateGoalLife() {
	p.GoalLife = utils.SumNumberInteger(p.NumberExpression + p.NumberPathOfLife)
}
