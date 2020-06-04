package service

import (
	"fmt"

	"github.com/matimartini/api-numerologia/database"
)

type GoalLife struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (goalLife *GoalLife) GetDescriptionGoalLife(number int) {
	document := database.GetDesciption(number, "goal-life")

	goalLife.Number = number
	document.DataTo(&goalLife)

	if goalLife.Description == "" {
		fmt.Println("Error empty description goalLife. Number: ", number)
	}
}
