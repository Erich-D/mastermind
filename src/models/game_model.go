package models

import (
	"database/sql"
	"fmt"

	"github.com/Erich-D/mastermind/entities"
)

type GameModel struct {
	Db *sql.DB
}

func (db GameModel) Update(obj *entities.Game) (int64, error) {
	result, err := db.Db.Exec("UPDATE game SET answer = ?, inProgress = ? WHERE gameId = ?", obj.Answer, obj.Status, obj.ID)
	if err != nil {
		return 0, fmt.Errorf("updategame: %v", err)
	} else {
		return result.RowsAffected()
	}
}

func (db GameModel) Add(obj *entities.Game) (int64, error) {
	result, err := db.Db.Exec("INSERT INTO game (answer) VALUES (?)", obj.Answer)
	if err != nil {
		return 0, fmt.Errorf("Addgame: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Addgame: %v", err)
	}
	return id, nil
}

func (db GameModel) FindAll() ([]entities.Game, error) {
	var games []entities.Game

	rows, err := db.Db.Query("SELECT * FROM game")
	if err != nil {
		return nil, fmt.Errorf("getAllGames %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var gm entities.Game
		if err := rows.Scan(&gm.ID, &gm.Answer, &gm.Status); err != nil {
			return nil, fmt.Errorf("getAllGames %v", err)
		}
		games = append(games, gm)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllGames %v", err)
	}
	return games, nil
}

func (db GameModel) Find(id int64) (entities.Game, error) {
	var gm entities.Game

	row := db.Db.QueryRow("SELECT * FROM game WHERE gameId = ?", id)
	if err := row.Scan(&gm.ID, &gm.Answer, &gm.Status); err != nil {
		if err == sql.ErrNoRows {
			return gm, fmt.Errorf("GameById %d: no such game", id)
		}
		return gm, fmt.Errorf("GameById %d: %v", id, err)
	}
	return gm, nil
}
