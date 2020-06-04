package service

import (
	"fmt"

	"github.com/matimartini/api-numerologia/database"
)

type Expression struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (expression *Expression) GetDescriptionExpression(number int) {
	document := database.GetDesciption(number, "expression")
	expression.Number = number
	document.DataTo(&expression)

	if expression.Description == "" {
		fmt.Println("Error empty description expression. number: ", number)
	}
}
