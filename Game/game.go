package game

import (
	"errors"
	"fmt"
	"interfaces"
	"math/rand"
	"net/http"
	"questions"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var games = make([]interfaces.Game, 0)

func InitGame(c *gin.Context) {

	newGame := interfaces.Game{
		Date:               time.Now().String(),
		GameID:             uuid.New().String(),
		Participants:       make([]interfaces.User, 2),
		ActivePlayer:       interfaces.User{},
		StartingQuestions:  interfaces.AllQuestions{},
		QuestionPicked:     interfaces.Question{},
		QuestionsRemaining: interfaces.AllQuestions{},
		TotalQuestions:     0,
		Winner:             interfaces.User{},
	}
	games = append(games, newGame)
	fmt.Println(games[0])
	//response,err
	unires := interfaces.UniResponse[interfaces.Game]{
		Message: "Hello",
		Data:    newGame,
		Status:  "200",
	}
	c.IndentedJSON(http.StatusOK, unires)
}

func getgameID(id string) (*interfaces.Game, error) {
	for i, b := range games {
		if b.GameID == id {
			return &games[i], nil
		}
	}
	return nil, errors.New("game not found")
}

func GetGameByID(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "id is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}

	game, err := getgameID(id)
	if err != nil {
		unires := interfaces.UniResponse[string]{
			Message: "Game not found",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}
	unires := interfaces.UniResponse[interfaces.Game]{
		Message: "Hello",
		Data:    *game,
		Status:  "200",
	}
	c.IndentedJSON(http.StatusOK, unires)
}

func SetUsers(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "id is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}

	game, err := getgameID(id)
	if err != nil {
		unires := interfaces.UniResponse[string]{
			Message: "Game not found",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}

	user1, ok := c.GetQuery("user1")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "user1 is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}
	user2, ok := c.GetQuery("user2")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "user2 is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}

	game.Participants[0].Name = user1
	game.Participants[1].Name = user2
	coin := rand.Intn(2)
	game.ActivePlayer = game.Participants[coin]

	unires := interfaces.UniResponse[string]{
		Message: string("Success in setting users. Active player is " + game.ActivePlayer.Name),
		Data:    "",
		Status:  "200",
	}
	c.IndentedJSON(http.StatusOK, unires)
}

func SetQuestions(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "id is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}

	game, err := getgameID(id)
	if err != nil {
		unires := interfaces.UniResponse[string]{
			Message: "Game not found",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}

	qsToBePlayed := QuestionHandler(interfaces.QuestionsPlayed)
	totalNumberOfQs := FindTotalQuestions(interfaces.QuestionsPlayed)

	game.QuestionsRemaining = qsToBePlayed
	game.StartingQuestions = qsToBePlayed
	game.TotalQuestions = totalNumberOfQs

	unires := interfaces.UniResponse[string]{
		Message: "Success in setting questions in this game",
		Data:    "",
		Status:  "200",
	}
	c.IndentedJSON(http.StatusOK, unires)
}

func StartGame(qPlayed interfaces.QCategory) (interface{}, error) {
	qsToBePlayed := QuestionHandler(qPlayed)
	var qPicked interfaces.Question
	//startingQuestions := qsToBePlayed
	totalNumberOfQs := FindTotalQuestions(qPlayed)
	fmt.Println("Give a name for the first player")
	user1 := AskForUser() //First User
	fmt.Println("Give a name for the second player")
	user2 := AskForUser() //Second User
	isActive := interfaces.User{}

	//Decide who plays first
	fmt.Println("Heads or tails:")
	fmt.Println(user1.Name + " Chooses! Heads? [Y/n]")
	var isHeads string = YNScanner()

	coin := rand.Intn(2)
	if isHeads == "Y" {
		if coin == 1 {
			fmt.Println("Heads!")
			isActive = user1
		} else {
			fmt.Println("Tails!")
			isActive = user2
		}
	} else {
		if coin == 1 {
			fmt.Println("Heads!")
			isActive = user2
		} else {
			fmt.Println("Tails!")
			isActive = user1
		}
	}

	//Play all the questions picked
	for i := 0; i < totalNumberOfQs; i++ {
		var correctAnswer string = ""
		var needHelp, doubleHelp string = "", ""
		//Active player chooses the question they wish
		qPicked = chooseQuestion(isActive, qPlayed, qsToBePlayed)
		qPoints, err := strconv.Atoi(qPicked.Points)

		if err != nil {
			fmt.Println(err)
		}

		//Helps
		if isActive.Lifelines.Double {
			fmt.Println("Do you want to double the points of this question? [Y/n]")
			doubleHelp = YNScanner()
			if doubleHelp == "Y" {
				qPoints = 2 * qPoints
				if isActive.Name == user1.Name {
					user1.Lifelines.Double = false
				} else {
					user2.Lifelines.Double = false
				}
			}
		}
		fmt.Println(qPicked.Text)

		if isActive.Lifelines.Fifty {
			fmt.Println("Do you need to use the 50/50 Lifeline? [Y/n]")
			needHelp = YNScanner()
			if needHelp == "Y" {
				fmt.Println(qPicked.Fifty)
				qPoints = 1
				if isActive.Name == user1.Name {
					user1.Lifelines.Fifty = false
				} else {
					user2.Lifelines.Fifty = false
				}
			}
		}
		if isActive.Lifelines.Phone {
			fmt.Println("Do you need to use the Phone Lifeline? [Y/n]")
			needHelp = YNScanner()
			if needHelp == "Y" {
				if isActive.Name == user1.Name {
					user1.Lifelines.Phone = false
				} else {
					user2.Lifelines.Phone = false
				}
			}
		}

		fmt.Println("Was the question answered correctly? [Y/n]")
		correctAnswer = YNScanner()
		if correctAnswer == "Y" {
			if isActive.Name == user1.Name {
				user1.Points += qPoints
				isActive = user2
			} else {
				user2.Points += qPoints
				isActive = user1
			}
		} else {
			if isActive.Name == user1.Name {
				isActive = user2
			} else {
				isActive = user1
			}
		}
		fmt.Printf(user1.Name+" has: %d points\n", user1.Points)
		fmt.Printf(user2.Name+" has: %d points\n", user2.Points)
	}
	return nil, nil
}

func chooseQuestion(activePlayer interfaces.User, qPlayed interfaces.QCategory, qsRemaining interfaces.AllQuestions) interfaces.Question {
	var questionIndex int
	var questionPicked interfaces.Question
	fmt.Println(activePlayer.Name + " chooses next question! Available questions are: ")
	//Print all questions still available
	cat1String, cat1Counter := PrintQuestions(0, qPlayed.Category1, qsRemaining.Category1)
	fmt.Println("Category 1: ", cat1String)
	cat2String, cat2Counter := PrintQuestions(cat1Counter, qPlayed.Category2, qsRemaining.Category2)
	fmt.Println("Category 2: ", cat2String)
	cat3String, cat3Counter := PrintQuestions(cat2Counter, qPlayed.Category3, qsRemaining.Category3)
	fmt.Println("Category 3: ", cat3String)
	cat4String, cat4Counter := PrintQuestions(cat3Counter, qPlayed.Category4, qsRemaining.Category4)
	fmt.Println("Category 4: ", cat4String)
	cat5String, cat5Counter := PrintQuestions(cat4Counter, qPlayed.Category5, qsRemaining.Category5)
	fmt.Println("Category 5: ", cat5String)
	cat6String, cat6Counter := PrintQuestions(cat5Counter, qPlayed.Category6, qsRemaining.Category6)
	fmt.Println("Category 6: ", cat6String)
	cat7String, cat7Counter := PrintQuestions(cat6Counter, qPlayed.Category7, qsRemaining.Category7)
	fmt.Println("Category 7: ", cat7String, cat7Counter)
	cat8String, cat8Counter := PrintQuestions(cat7Counter, qPlayed.Category8, qsRemaining.Category8)
	fmt.Println("Category 8: ", cat8String, cat8Counter)

	//sumCounter := cat1Counter + cat2Counter + cat3Counter + cat4Counter + cat5Counter + cat6Counter + cat7Counter + cat8Counter
	for {
		//User inputs the question they wish
		fmt.Scan(&questionIndex)
		questionIndex--
		if questionIndex < 0 {
			fmt.Println("This number is negative! Questions can't be negative, can they?")
		} else if questionIndex >= 0 && questionIndex < cat1Counter {
			questionPicked = chosenQuestion(questionIndex, qPlayed.Category1, qsRemaining.Category1)
			break
		} else if questionIndex < cat2Counter {
			questionPicked = chosenQuestion(questionIndex-cat1Counter, qPlayed.Category2, qsRemaining.Category2)
			break
		} else if questionIndex < cat3Counter {
			questionPicked = chosenQuestion(questionIndex-cat2Counter, qPlayed.Category3, qsRemaining.Category3)
			break
		} else if questionIndex < cat4Counter {
			questionPicked = chosenQuestion(questionIndex-cat3Counter, qPlayed.Category4, qsRemaining.Category4)
			break
		} else if questionIndex < cat5Counter {
			questionPicked = chosenQuestion(questionIndex-cat4Counter, qPlayed.Category5, qsRemaining.Category5)
			break
		} else if questionIndex < cat6Counter {
			questionPicked = chosenQuestion(questionIndex-cat5Counter, qPlayed.Category6, qsRemaining.Category6)
			break
		} else if questionIndex < cat7Counter {
			questionPicked = chosenQuestion(questionIndex-cat6Counter, qPlayed.Category7, qsRemaining.Category7)
			break
		} else if questionIndex < cat8Counter {
			questionPicked = chosenQuestion(questionIndex-cat7Counter, qPlayed.Category8, qsRemaining.Category8)
			break
		} else {
			fmt.Println("This number is bigger than the number of questions left!")
		}
	}
	return questionPicked
}

func chosenQuestion(questionIndex int, questionCategory interfaces.Category, qCategoryTBP interfaces.AllPoints) interfaces.Question {
	if questionCategory.OnePointers == 0 {
		if questionCategory.TwoPointers == 0 {
			return qCategoryTBP.ThreePointers[questionIndex]
		} else if questionIndex < questionCategory.TwoPointers {
			return qCategoryTBP.TwoPointers[questionIndex]
		} else {
			return qCategoryTBP.ThreePointers[questionIndex-questionCategory.TwoPointers]
		}
	} else if questionCategory.TwoPointers == 0 {
		if questionIndex < questionCategory.OnePointers {
			return qCategoryTBP.OnePointers[questionIndex]
		} else {
			return qCategoryTBP.ThreePointers[questionIndex-questionCategory.OnePointers]
		}
	} else {
		if questionIndex < questionCategory.OnePointers {
			return qCategoryTBP.OnePointers[questionIndex]
		} else if questionIndex < (questionCategory.TwoPointers + questionCategory.OnePointers) {
			return qCategoryTBP.TwoPointers[questionIndex-questionCategory.OnePointers]
		} else {
			return qCategoryTBP.ThreePointers[questionIndex-questionCategory.OnePointers-questionCategory.TwoPointers]
		}
	}
}

func PrintQuestions(counter int, questionCategory interfaces.Category, qCategoryTBP interfaces.AllPoints) (string, int) {
	var catString string = ""
	if questionCategory.IsPicked {
		for i := 0; i < questionCategory.OnePointers; i++ {
			if !qCategoryTBP.OnePointers[i].IsPlayed {
				counter++
				catString += "#" + strconv.Itoa(counter) + " " + qCategoryTBP.OnePointers[i].Points + " "
			}
		}
		for j := 0; j < questionCategory.TwoPointers; j++ {
			if !qCategoryTBP.TwoPointers[j].IsPlayed {
				counter++
				catString += "#" + strconv.Itoa(counter) + " " + qCategoryTBP.TwoPointers[j].Points + " "
			}
		}
		for k := 0; k < questionCategory.ThreePointers; k++ {
			if !qCategoryTBP.ThreePointers[k].IsPlayed {
				counter++
				catString += "#" + strconv.Itoa(counter) + " " + qCategoryTBP.ThreePointers[k].Points + " "
			}
		}
	}
	return catString, counter
}

func AskForUser() interfaces.User {
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

func FindTotalQuestions(qPlayed interfaces.QCategory) int {
	var totalSum int = 0
	if qPlayed.Category1.IsPicked {
		totalSum += qPlayed.Category1.OnePointers + qPlayed.Category1.TwoPointers + qPlayed.Category1.ThreePointers
	}
	if qPlayed.Category2.IsPicked {
		totalSum += qPlayed.Category2.OnePointers + qPlayed.Category2.TwoPointers + qPlayed.Category2.ThreePointers
	}
	if qPlayed.Category3.IsPicked {
		totalSum += qPlayed.Category3.OnePointers + qPlayed.Category3.TwoPointers + qPlayed.Category3.ThreePointers
	}
	if qPlayed.Category4.IsPicked {
		totalSum += qPlayed.Category4.OnePointers + qPlayed.Category4.TwoPointers + qPlayed.Category4.ThreePointers
	}
	if qPlayed.Category5.IsPicked {
		totalSum += qPlayed.Category5.OnePointers + qPlayed.Category5.TwoPointers + qPlayed.Category5.ThreePointers
	}
	if qPlayed.Category6.IsPicked {
		totalSum += qPlayed.Category6.OnePointers + qPlayed.Category6.TwoPointers + qPlayed.Category6.ThreePointers
	}
	if qPlayed.Category7.IsPicked {
		totalSum += qPlayed.Category7.OnePointers + qPlayed.Category7.TwoPointers + qPlayed.Category7.ThreePointers
	}
	if qPlayed.Category8.IsPicked {
		totalSum += qPlayed.Category8.OnePointers + qPlayed.Category8.TwoPointers + qPlayed.Category8.ThreePointers
	}
	return totalSum
}

func QuestionHandler(qPlayed interfaces.QCategory) interfaces.AllQuestions {
	allQuestions := interfaces.AllQuestions{}
	if qPlayed.Category1.IsPicked {
		allQuestions.Category1, _ = questions.GetQuestion("QFolder/Category1.json", qPlayed.Category1.OnePointers, qPlayed.Category1.TwoPointers, qPlayed.Category1.ThreePointers)
	}
	if qPlayed.Category2.IsPicked {
		allQuestions.Category2, _ = questions.GetQuestion("QFolder/Category2.json", qPlayed.Category2.OnePointers, qPlayed.Category2.TwoPointers, qPlayed.Category2.ThreePointers)
	}
	if qPlayed.Category3.IsPicked {
		allQuestions.Category3, _ = questions.GetQuestion("QFolder/Category3.json", qPlayed.Category3.OnePointers, qPlayed.Category3.TwoPointers, qPlayed.Category3.ThreePointers)
	}
	if qPlayed.Category4.IsPicked {
		allQuestions.Category4, _ = questions.GetQuestion("QFolder/Category4.json", qPlayed.Category4.OnePointers, qPlayed.Category4.TwoPointers, qPlayed.Category4.ThreePointers)
	}
	if qPlayed.Category5.IsPicked {
		allQuestions.Category5, _ = questions.GetQuestion("QFolder/Category5.json", qPlayed.Category5.OnePointers, qPlayed.Category5.TwoPointers, qPlayed.Category5.ThreePointers)
	}
	if qPlayed.Category6.IsPicked {
		allQuestions.Category6, _ = questions.GetQuestion("QFolder/Category6.json", qPlayed.Category6.OnePointers, qPlayed.Category6.TwoPointers, qPlayed.Category6.ThreePointers)
	}
	if qPlayed.Category7.IsPicked {
		allQuestions.Category7, _ = questions.GetQuestion("QFolder/Category7.json", qPlayed.Category7.OnePointers, qPlayed.Category7.TwoPointers, qPlayed.Category7.ThreePointers)
	}
	if qPlayed.Category8.IsPicked {
		allQuestions.Category8, _ = questions.GetQuestion("QFolder/Category8.json", qPlayed.Category8.OnePointers, qPlayed.Category8.TwoPointers, qPlayed.Category8.ThreePointers)
	}
	return allQuestions
}

func YNScanner() string {
	var tester string = ""
	for {
		fmt.Scan(&tester)
		if tester == "Y" || tester == "n" {
			break
		}
	}
	return tester
}
