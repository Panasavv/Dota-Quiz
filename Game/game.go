package game

import (
	"fmt"
	"interfaces"
	"questions"
)

func StartGame(qPlayed interfaces.QCategory) (interface{}, error) {
	qsToBePlayed := QuestionHandler(qPlayed)
	user1 := AskForUser()
	user2 := AskForUser()
	user1.Points += 1
	fmt.Println(user1.Points)
	fmt.Println(user2.Name)

	fmt.Println(qsToBePlayed[0])
	return nil, nil
}

func AskForUser() interfaces.User {
	fmt.Println("Give a name for this User")
	var resp string
	fmt.Scan(&resp)
	user := interfaces.User{
		Name:   resp,
		Points: 0,
		Lifelines: interfaces.Helpers{
			Fifty:  true,
			Phone:  true,
			Double: true,
		},
	}
	return user
}

func QuestionHandler(qPlayed interfaces.QCategory) [8]interface{} {
	var allQuestions [8]interface{}
	if qPlayed.Category1.IsPicked {
		allQuestions[0], _ = questions.GetQuestion("QFolder/Category1.json", qPlayed.Category1.OnePointers, qPlayed.Category1.TwoPointers, qPlayed.Category1.ThreePointers)
	}
	if qPlayed.Category2.IsPicked {
		allQuestions[1], _ = questions.GetQuestion("QFolder/Category2.json", qPlayed.Category2.OnePointers, qPlayed.Category2.TwoPointers, qPlayed.Category2.ThreePointers)
	}
	if qPlayed.Category3.IsPicked {
		allQuestions[2], _ = questions.GetQuestion("QFolder/Category3.json", qPlayed.Category3.OnePointers, qPlayed.Category3.TwoPointers, qPlayed.Category3.ThreePointers)
	}
	if qPlayed.Category4.IsPicked {
		allQuestions[3], _ = questions.GetQuestion("QFolder/Category4.json", qPlayed.Category4.OnePointers, qPlayed.Category4.TwoPointers, qPlayed.Category4.ThreePointers)
	}
	if qPlayed.Category5.IsPicked {
		allQuestions[4], _ = questions.GetQuestion("QFolder/Category5.json", qPlayed.Category5.OnePointers, qPlayed.Category5.TwoPointers, qPlayed.Category5.ThreePointers)
	}
	if qPlayed.Category6.IsPicked {
		allQuestions[5], _ = questions.GetQuestion("QFolder/Category6.json", qPlayed.Category6.OnePointers, qPlayed.Category6.TwoPointers, qPlayed.Category6.ThreePointers)
	}
	if qPlayed.Category7.IsPicked {
		allQuestions[6], _ = questions.GetQuestion("QFolder/Category7.json", qPlayed.Category7.OnePointers, qPlayed.Category7.TwoPointers, qPlayed.Category7.ThreePointers)
	}
	if qPlayed.Category8.IsPicked {
		allQuestions[7], _ = questions.GetQuestion("QFolder/Category8.json", qPlayed.Category8.OnePointers, qPlayed.Category8.TwoPointers, qPlayed.Category8.ThreePointers)
	}
	return allQuestions
}
