package service

import (
	"github.com/matimartini/api-numerologia/utils"
)

type Person struct {
	FullName         string       `json:"full_name"`
	BirthDate        string       `json:"birth_date"`
	NumberPathOfLife int          `json:"number_path_of_life"`
	Personality      Personality  `json:"personality"`
	Expression       Expression   `json:"expression"`
	GoalLife         GoalLife     `json:"goal_life"`
	SoulAmbition     SoulAmbition `json:"soul_ambition"`
}

func (p *Person) NewPerson(name string, date string) {
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
	soulAmbition := &SoulAmbition{}
	soulAmbition.CalculateNumberSoulAmbition(p.FullName)
	p.SoulAmbition = *soulAmbition
}

func (p *Person) calculatePersonality() {
	personality := &Personality{}
	personality.CalculateNumberPersonality(p.FullName)
	p.Personality = *personality
}

func (p *Person) calculateNumberExpression() {
	numberExpression := utils.SumNumberInteger(p.SoulAmbition.Number + p.Personality.Number)
	expression := &Expression{}
	expression.GetDescriptionExpression(numberExpression)
	p.Expression = *expression
}

func (p *Person) calculateGoalLife() {
	number := utils.SumNumberInteger(p.Expression.Number + p.NumberPathOfLife)
	goalLife := &GoalLife{}
	goalLife.GetDescriptionGoalLife(number)
	p.GoalLife = *goalLife
}
