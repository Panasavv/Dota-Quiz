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
	router.Run("localhost:8080")
	//game.StartGame(*QuestionsPlayed)

}
