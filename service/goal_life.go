package service

import "github.com/matimartini/api-numerologia/database"

type GoalLife struct {
	Number      int    `json:"number"`
	Description string `json:"description"`
}

func (goalLife *GoalLife) GetDescriptionGoalLife(number int) {
	document := database.GetDesciption(number, "goal-life")
	var goal GoalLife
	document.DataTo(&goal)
	goalLife.Number = number
	goalLife.Description = goal.Description
}