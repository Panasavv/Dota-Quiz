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

func GetQuestions(c *gin.Context) {
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
	questionResponse := printQs(*game)

	unires := interfaces.UniResponse[interfaces.QuestionResponse]{
		Message: "Questions remaining are: ",
		Data:    questionResponse,
		Status:  "200",
	}
	c.IndentedJSON(http.StatusOK, unires)
}

func printQs(game interfaces.Game) interfaces.QuestionResponse {
	var questionResponse interfaces.QuestionResponse
	cat1String, cat1Counter := PrintQuestions(0, interfaces.QuestionsPlayed.Category1, game.QuestionsRemaining.Category1)
	questionResponse.Category1.CatString = string("Category 1: " + cat1String)
	questionResponse.Category1.CatCount = cat1Counter
	cat2String, cat2Counter := PrintQuestions(cat1Counter, interfaces.QuestionsPlayed.Category2, game.QuestionsRemaining.Category2)
	questionResponse.Category2.CatString = string("Category 2: " + cat2String)
	questionResponse.Category2.CatCount = cat2Counter
	cat3String, cat3Counter := PrintQuestions(cat2Counter, interfaces.QuestionsPlayed.Category3, game.QuestionsRemaining.Category3)
	questionResponse.Category3.CatString = string("Category 3: " + cat3String)
	questionResponse.Category3.CatCount = cat3Counter
	cat4String, cat4Counter := PrintQuestions(cat3Counter, interfaces.QuestionsPlayed.Category4, game.QuestionsRemaining.Category4)
	questionResponse.Category4.CatString = string("Category 4: " + cat4String)
	questionResponse.Category4.CatCount = cat4Counter
	cat5String, cat5Counter := PrintQuestions(cat4Counter, interfaces.QuestionsPlayed.Category5, game.QuestionsRemaining.Category5)
	questionResponse.Category5.CatString = string("Category 5: " + cat5String)
	questionResponse.Category5.CatCount = cat5Counter
	cat6String, cat6Counter := PrintQuestions(cat5Counter, interfaces.QuestionsPlayed.Category6, game.QuestionsRemaining.Category6)
	questionResponse.Category6.CatString = string("Category 6: " + cat6String)
	questionResponse.Category6.CatCount = cat6Counter
	cat7String, cat7Counter := PrintQuestions(cat6Counter, interfaces.QuestionsPlayed.Category7, game.QuestionsRemaining.Category7)
	questionResponse.Category7.CatString = string("Category 7: " + cat7String)
	questionResponse.Category7.CatCount = cat7Counter
	cat8String, cat8Counter := PrintQuestions(cat7Counter, interfaces.QuestionsPlayed.Category8, game.QuestionsRemaining.Category8)
	questionResponse.Category8.CatString = string("Category 8: " + cat8String)
	questionResponse.Category8.CatCount = cat8Counter
	return questionResponse
}

func GetChosenQuestion(c *gin.Context) {
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
	questionResponse := printQs(*game)

	qNumberPre, ok := c.GetQuery("qNumber")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "qNumber is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}
	qNumber, _ := strconv.Atoi(qNumberPre)
	qNumber--
	var questionPicked interfaces.Question

	if qNumber < 0 {
		unires := interfaces.UniResponse[string]{
			Message: "qNumber is negative or zero, please try again",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	} else if qNumber >= 0 && qNumber < questionResponse.Category1.CatCount {
		questionPicked, game.QuestionsRemaining.Category1 = chosenQuestion(qNumber, interfaces.QuestionsPlayed.Category1, game.QuestionsRemaining.Category1)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category1 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category1)
		//edw vazw
	} else if qNumber < questionResponse.Category2.CatCount {
		questionPicked, game.QuestionsRemaining.Category2 = chosenQuestion(qNumber-questionResponse.Category1.CatCount, interfaces.QuestionsPlayed.Category2, game.QuestionsRemaining.Category2)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category2 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category2)
	} else if qNumber < questionResponse.Category3.CatCount {
		questionPicked, game.QuestionsRemaining.Category3 = chosenQuestion(qNumber-questionResponse.Category2.CatCount, interfaces.QuestionsPlayed.Category3, game.QuestionsRemaining.Category3)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category3 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category3)
	} else if qNumber < questionResponse.Category4.CatCount {
		questionPicked, game.QuestionsRemaining.Category4 = chosenQuestion(qNumber-questionResponse.Category3.CatCount, interfaces.QuestionsPlayed.Category4, game.QuestionsRemaining.Category4)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category4 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category4)
	} else if qNumber < questionResponse.Category5.CatCount {
		questionPicked, game.QuestionsRemaining.Category5 = chosenQuestion(qNumber-questionResponse.Category4.CatCount, interfaces.QuestionsPlayed.Category5, game.QuestionsRemaining.Category5)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category5 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category5)
	} else if qNumber < questionResponse.Category6.CatCount {
		questionPicked, game.QuestionsRemaining.Category6 = chosenQuestion(qNumber-questionResponse.Category5.CatCount, interfaces.QuestionsPlayed.Category6, game.QuestionsRemaining.Category6)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category6 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category6)
	} else if qNumber < questionResponse.Category7.CatCount {
		questionPicked, game.QuestionsRemaining.Category7 = chosenQuestion(qNumber-questionResponse.Category6.CatCount, interfaces.QuestionsPlayed.Category7, game.QuestionsRemaining.Category7)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category7 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category7)
	} else if qNumber < questionResponse.Category8.CatCount {
		questionPicked, game.QuestionsRemaining.Category8 = chosenQuestion(qNumber-questionResponse.Category7.CatCount, interfaces.QuestionsPlayed.Category8, game.QuestionsRemaining.Category8)
		game.TotalQuestions--
		interfaces.QuestionsPlayed.Category8 = removeQuestions(questionPicked, interfaces.QuestionsPlayed.Category8)
	} else {
		unires := interfaces.UniResponse[string]{
			Message: "The number you typed is bigger than the number of questions, please try again.",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}
	game.QuestionPicked = questionPicked

	unires := interfaces.UniResponse[interfaces.Question]{
		Message: "This is the question to be answered",
		Data:    questionPicked,
		Status:  "200",
	}
	c.IndentedJSON(http.StatusOK, unires)

}

func GetFifty(c *gin.Context) {
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
	if game.ActivePlayer.Name == game.Participants[0].Name {
		if !game.Participants[0].Lifelines.Fifty {
			game.Participants[0].Lifelines.Fifty = true
			game.QuestionPicked.Points = "1"
			unires := interfaces.UniResponse[string]{
				Message: "50/50 lifeline was succesfully used",
				Data:    game.QuestionPicked.Fifty,
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "This player has not available 50/50 lifeline",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
		}
	} else {
		if !game.Participants[1].Lifelines.Fifty {
			game.Participants[1].Lifelines.Fifty = true
			game.QuestionPicked.Points = "1"
			unires := interfaces.UniResponse[string]{
				Message: "50/50 lifeline was succesfully used",
				Data:    game.QuestionPicked.Fifty,
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "This player has not available 50/50 lifeline",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
		}
	}
}

func GetDouble(c *gin.Context) {
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
	if game.ActivePlayer.Name == game.Participants[0].Name {
		if !game.Participants[0].Lifelines.Double {
			game.Participants[0].Lifelines.Double = true
			points, _ := strconv.Atoi(game.QuestionPicked.Points)
			points = 2 * points
			game.QuestionPicked.Points = strconv.Itoa(points)

			unires := interfaces.UniResponse[string]{
				Message: "2x lifeline was succesfully used",
				Data:    "",
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "This player has not available 2x lifeline",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
		}
	} else {
		if !game.Participants[1].Lifelines.Double {
			game.Participants[1].Lifelines.Double = true
			points, _ := strconv.Atoi(game.QuestionPicked.Points)
			points = 2 * points
			game.QuestionPicked.Points = strconv.Itoa(points)

			unires := interfaces.UniResponse[string]{
				Message: "2x lifeline was succesfully used",
				Data:    game.QuestionPicked.Fifty,
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "This player has not available 2x lifeline",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
		}
	}
}

func GetPhone(c *gin.Context) {
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
	if game.ActivePlayer.Name == game.Participants[0].Name {
		if !game.Participants[0].Lifelines.Phone {
			game.Participants[0].Lifelines.Phone = true

			unires := interfaces.UniResponse[string]{
				Message: "Phone lifeline was succesfully used",
				Data:    "",
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "This player has not available Phone lifeline",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
		}
	} else {
		if !game.Participants[1].Lifelines.Phone {
			game.Participants[1].Lifelines.Phone = true

			unires := interfaces.UniResponse[string]{
				Message: "Phone lifeline was succesfully used",
				Data:    game.QuestionPicked.Fifty,
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "This player has not available Phone lifeline",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
		}
	}
}

func SetCorrectAnswer(c *gin.Context) {
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

	answer, ok := c.GetQuery("answer")
	if !ok {
		unires := interfaces.UniResponse[string]{
			Message: "answer is missing",
			Data:    "",
			Status:  "404",
		}
		c.IndentedJSON(http.StatusOK, unires)
		return
	}

	if game.ActivePlayer.Name == game.Participants[0].Name {
		game.ActivePlayer = game.Participants[1]
		if answer == "y" {
			qpoints, _ := strconv.Atoi(game.QuestionPicked.Points)
			game.Participants[0].Points += qpoints
			unires := interfaces.UniResponse[interfaces.User]{
				Message: "Correct answer",
				Data:    game.Participants[0],
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "Wrong answer",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		}
	} else {
		game.ActivePlayer = game.Participants[0]
		if answer == "y" {
			qpoints, _ := strconv.Atoi(game.QuestionPicked.Points)
			game.Participants[1].Points += qpoints
			unires := interfaces.UniResponse[interfaces.User]{
				Message: "Correct answer",
				Data:    game.Participants[1],
				Status:  "200",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		} else {
			unires := interfaces.UniResponse[string]{
				Message: "Wrong answer",
				Data:    "",
				Status:  "404",
			}
			c.IndentedJSON(http.StatusOK, unires)
			return
		}
	}
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
			questionPicked, _ = chosenQuestion(questionIndex, qPlayed.Category1, qsRemaining.Category1)
			break
		} else if questionIndex < cat2Counter {
			questionPicked, _ = chosenQuestion(questionIndex-cat1Counter, qPlayed.Category2, qsRemaining.Category2)
			break
		} else if questionIndex < cat3Counter {
			questionPicked, _ = chosenQuestion(questionIndex-cat2Counter, qPlayed.Category3, qsRemaining.Category3)
			break
		} else if questionIndex < cat4Counter {
			questionPicked, _ = chosenQuestion(questionIndex-cat3Counter, qPlayed.Category4, qsRemaining.Category4)
			break
		} else if questionIndex < cat5Counter {
			questionPicked, _ = chosenQuestion(questionIndex-cat4Counter, qPlayed.Category5, qsRemaining.Category5)
			break
		} else if questionIndex < cat6Counter {
			questionPicked, _ = chosenQuestion(questionIndex-cat5Counter, qPlayed.Category6, qsRemaining.Category6)
			break
		} else if questionIndex < cat7Counter {
			questionPicked, _ = chosenQuestion(questionIndex-cat6Counter, qPlayed.Category7, qsRemaining.Category7)
			break
		} else if questionIndex < cat8Counter {
			questionPicked, _ = chosenQuestion(questionIndex-cat7Counter, qPlayed.Category8, qsRemaining.Category8)
			break
		} else {
			fmt.Println("This number is bigger than the number of questions left!")
		}
	}
	return questionPicked
}

func chosenQuestion(questionIndex int, questionCategory interfaces.Category, qCategoryTBP interfaces.AllPoints) (interfaces.Question, interfaces.AllPoints) {
	if questionCategory.OnePointers == 0 {
		if questionCategory.TwoPointers == 0 {
			r1 := qCategoryTBP.ThreePointers[questionIndex]
			qCategoryTBP.ThreePointers = remove(qCategoryTBP.ThreePointers, questionIndex)
			return r1, qCategoryTBP
		} else if questionIndex < questionCategory.TwoPointers {
			r1 := qCategoryTBP.TwoPointers[questionIndex]
			qCategoryTBP.TwoPointers = remove(qCategoryTBP.TwoPointers, questionIndex)
			return r1, qCategoryTBP
		} else {
			r1 := qCategoryTBP.ThreePointers[questionIndex-questionCategory.TwoPointers]
			qCategoryTBP.ThreePointers = remove(qCategoryTBP.ThreePointers, questionIndex-questionCategory.TwoPointers)
			return r1, qCategoryTBP
		}
	} else if questionCategory.TwoPointers == 0 {
		if questionIndex < questionCategory.OnePointers {
			r1 := qCategoryTBP.OnePointers[questionIndex]
			qCategoryTBP.OnePointers = remove(qCategoryTBP.OnePointers, questionIndex)
			return r1, qCategoryTBP
		} else {
			r1 := qCategoryTBP.ThreePointers[questionIndex-questionCategory.OnePointers]
			qCategoryTBP.ThreePointers = remove(qCategoryTBP.ThreePointers, questionIndex-questionCategory.OnePointers)
			return r1, qCategoryTBP
		}
	} else {
		if questionIndex < questionCategory.OnePointers {
			r1 := qCategoryTBP.OnePointers[questionIndex]
			qCategoryTBP.OnePointers = remove(qCategoryTBP.OnePointers, questionIndex)
			return r1, qCategoryTBP
		} else if questionIndex < (questionCategory.TwoPointers + questionCategory.OnePointers) {
			r1 := qCategoryTBP.TwoPointers[questionIndex-questionCategory.OnePointers]
			qCategoryTBP.TwoPointers = remove(qCategoryTBP.TwoPointers, questionIndex-questionCategory.OnePointers)
			return r1, qCategoryTBP
		} else {
			r1 := qCategoryTBP.ThreePointers[questionIndex-questionCategory.OnePointers-questionCategory.TwoPointers]
			qCategoryTBP.ThreePointers = remove(qCategoryTBP.ThreePointers, questionIndex-questionCategory.OnePointers-questionCategory.TwoPointers)
			return r1, qCategoryTBP
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

func remove(slice []interfaces.Question, s int) []interfaces.Question {
	return append(slice[:s], slice[s+1:]...)
}

func removeQuestions(p interfaces.Question, a interfaces.Category) interfaces.Category {
	if p.Points == "1" {
		a.OnePointers--
	} else if p.Points == "2" {
		a.TwoPointers--
	} else {
		a.ThreePointers--
	}
	return a
}
