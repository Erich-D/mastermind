package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Erich-D/mastermind/entities"
	"github.com/Erich-D/mastermind/service"
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/games", getGames)
	router.GET("/game/:id", getGameByID)
	router.GET("/begin", createGame)
	router.POST("/guess", postGuess)
	router.GET("/rounds/:id", getRounds)

	router.Run("localhost:8080")
}

func getGames(c *gin.Context) {
	tmp, err := service.GetGames()
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "games not found"})
	}
	c.IndentedJSON(http.StatusCreated, tmp)
}

func getGameByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	g, err1 := service.GameById(int64(id))
	if err1 != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "game not found"})
	}
	c.IndentedJSON(http.StatusOK, g)
}

func createGame(c *gin.Context) {
	newGame, err1 := service.NewGame()
	if err1 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
	}
	c.IndentedJSON(http.StatusOK, newGame)
}

func postGuess(c *gin.Context) { //todo
	var newRound entities.Round
	if err := c.BindJSON(&newRound); err != nil {
		fmt.Printf(err.Error())
		return
	}
	rd, err3 := service.Guess(newRound)
	if err3 != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err3.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, rd)
}

func getRounds(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rnds, err := service.GetRounds(int64(id))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No rounds found"})
	}
	c.IndentedJSON(http.StatusOK, rnds)
}
