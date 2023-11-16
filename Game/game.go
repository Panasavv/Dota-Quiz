package game

import (
	"fmt"
	"interfaces"
	"questions"
)

func StartGame() (interface{}, error) {
	qPlayed := &interfaces.QCategory{
		Category1: true,
		Category2: false,
	}
	if qPlayed.Category1 {

		q1, _ := questions.GetQuestion("QFolder/Category1.json", 0, 2, 3)
		fmt.Println(q1)
	}
	return nil, nil
}
