package service

import "github.com/matimartini/api-numerologia/database"

type Expression struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (expression *Expression) GetDescriptionExpression(number int) {
	document := database.GetDesciption(number, "expression")
	var e Expression
	document.DataTo(&e)
	expression.Number = number
	expression.Description = e.Description
}
