package service

import (
	"strings"

	"github.com/matimartini/api-numerologia/processes"
)

type Lack struct {
	Number int
	Exists bool
	Cant   int
}

type ListLack struct {
	ListLack []*Lack
}

func (list *ListLack) appendLack(lack *Lack) {
	list.ListLack = append(list.ListLack, lack)
}

func (listLack *ListLack) newListLack() *ListLack {
	lackBase := make([]Lack, 9)
	for i := range lackBase {
		lackBase[i].Number = i + 1
		listLack.appendLack(&lackBase[i])
	}
	return listLack
}

func (listLack *ListLack) CalculateLack(name string) {
	listLack.newListLack()

	nameLetters := strings.Split(name, "")

	var mapLetters = processes.MapLetters()

	for _, letter := range nameLetters {
		letterTopUpper := strings.ToUpper(letter)
		num := mapLetters[letterTopUpper]
		listLack.setValueLack(num)
	}
}

func (ListLack *ListLack) setValueLack(number int) {

	for _, lack := range ListLack.ListLack {
		if lack.Number == number {
			lack.markLack()
			break
		}
	}
}

func (lack *Lack) markLack() {
	lack.Exists = true
	lack.Cant++
}
