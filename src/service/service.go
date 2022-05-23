package service

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Erich-D/mastermind/db"
	"github.com/Erich-D/mastermind/entities"
	"github.com/Erich-D/mastermind/models"
)

var gm models.GameModel = models.GameModel{Db: db.Db}

var rm models.RoundModel = models.RoundModel{Db: db.Db}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GetGames() ([]entities.Game, error) {
	gms, err := gm.FindAll()
	if err != nil {
		log.Fatal(err)
		fmt.Errorf("GetGames %v", err)
		return nil, err
	}
	for indx, val := range gms {
		if val.Status {
			gms[indx].Answer = "????"
		}
	}
	return gms, nil
}

func GameById(id int64) (entities.Game, error) {
	g, err := gm.Find(id)
	if err != nil {
		log.Fatal(err)
		fmt.Errorf("GetById %d: %v", id, err)
		return g, err
	}
	if g.Status {
		g.Answer = "????"
	}
	return g, nil
}

func NewGame() (entities.Game, error) {
	var newGame entities.Game
	newGame.Answer = creatAnswer()
	id, err := gm.Add(&newGame)
	if err != nil {
		log.Fatal(err)
		fmt.Errorf("AddNewGame %v", err)
	}
	g, err1 := gm.Find(id)
	if err1 != nil {
		log.Fatal(err)
		return g, fmt.Errorf("GetById %d: %v", id, err)
	}
	g.Answer = "????"
	return g, nil
}

func Guess(rnd entities.Round) (entities.Round, error) {
	id, _ := strconv.Atoi(rnd.GameID)
	thisGame, err1 := gm.Find(int64(id))
	if err1 != nil {
		return rnd, fmt.Errorf(err1.Error())
	}
	rnd.GuessResult = testGuess(thisGame, rnd.Guess)
	rid, err2 := rm.Add(&rnd)
	if err2 != nil {
		return rnd, fmt.Errorf(err2.Error())
	}
	rtn, err3 := rm.Find(rid)
	if err3 != nil {
		return rnd, fmt.Errorf(err3.Error())
	}
	return rtn, nil
}

func GetRounds(id int64) ([]entities.Round, error) {
	rounds, err := rm.FindAllFor(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if len(rounds) < 1 {
		return nil, fmt.Errorf("Nothing Found")
	}
	return rounds, nil
}

//***********************helper functions for NewGame*********************
func creatAnswer() string {

	ansBytes := make([]byte, 4)
	for indx, _ := range ansBytes {
		ansBytes[indx] = answerGen(ansBytes)
	}
	return string(ansBytes[:])
}

func answerGen(nums []byte) byte {
	inc := true
	var b byte
	for next := true; next; next = inc {
		b = byte(randInt(48, 57))
		for _, val := range nums {
			inc = val == b
			if inc {
				break
			}
		}
	}
	return b
}

func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}

//**************************helper functions for Guess ********************************
func testGuess(g entities.Game, guess string) string {
	answ := g.Answer
	e := 0
	p := 0
	for i := 0; i < 4; i++ {
		fmt.Printf("%c\n", answ[i])
		if answ[i] == guess[i] {
			e++
			continue
		}
		if charExist(answ, guess[i]) {
			p++
		}
	}
	if e == 4 {
		g.Status = false
		gm.Update(&g)
	}
	return "e:" + strconv.Itoa(e) + ",p:" + strconv.Itoa(p)
}

func charExist(answ string, guess byte) bool {
	for i := 0; i < 4; i++ {
		if answ[i] == guess {
			return true
		}
	}
	return false
}
