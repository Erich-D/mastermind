# mastermind
game API in Go

API has five end points
/begin   GET-Creates new game, stores in database, returns new game in JSON ex.{"id": "3","answer": "????","status": true}The answer is hidden if status is true. 
              
/guess   POST-Takes in guess and game id ex.{"gameId: "5", "guess": "5942"} and return round with results ex.{"id": "17", "guess": "9542", "guessTime": "2022-05-22T18:03:20Z", "guessresult": "e:1,p:0", "gameId": "5"} e = exact matches, p = partial matches
                                                                                                    
/games   GET- returns JSON array representing a list of all game objects. Answers are hidden for games in progress(status = true)

/game/:id  GET- returns JSON game object with matching id. Answer is hidden if status = true. 

/rounds/:id GET- returns JSON list of all rounds for game matching id.
