package service

import (
	"github.com/matimartini/api-numerologia/utils"
)

type Person struct {
	FullName         string       `json:"nombre"`
	BirthDate        string       `json:"fecha_nacimiento"`
	NumberPathOfLife int          `json:"camino_de_vida"`
	Personality      Personality  `json:"personalidad"`
	Expression       Expression   `json:"expresion"`
	SoulAmbition     SoulAmbition `json:"deseo_del_alma"`
	GoalLife         GoalLife     `json:"meta_de_vida"`
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
