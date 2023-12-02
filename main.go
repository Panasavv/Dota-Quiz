package main

import (
	"game"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/initGame", game.InitGame)
	router.GET("/getGameByID", game.GetGameByID)
	router.POST("/setUsers", game.SetUsers)
	router.POST("/setQuestions", game.SetQuestions)
	router.GET("/getQuestions", game.GetQuestions)
	router.GET("/getChosenQuestion", game.GetChosenQuestion)
	router.GET("/getFifty", game.GetFifty)
	router.GET("/getDouble", game.GetDouble)
	router.GET("/getPhone", game.GetPhone)
	router.POST("/setCorrectAnswer", game.SetCorrectAnswer)
	router.Run("localhost:8080")
	//game.StartGame(*QuestionsPlayed)

}

/*for i := 0; i < 17; i++ {
	questions.CropImages(interfaces.ImagesTBC1[i])
}
for i := 0; i < 21; i++ {
		questions.CropImages(interfaces.ImagesTBC3[i])
	}
	return
*/
