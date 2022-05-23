package models

import (
	"database/sql"
	"fmt"

	"github.com/Erich-D/mastermind/entities"
)

type RoundModel struct {
	Db *sql.DB
}

func (db RoundModel) Update(obj *entities.Round) (int64, error) {
	result, err := db.Db.Exec("UPDATE round set guess = ?, guessTime = ?, guessResult = ?, gameNum = ? WHERE roundId = ?", obj.Guess, obj.GuessTime, obj.GuessResult, obj.GameID, obj.ID)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (db RoundModel) Add(obj *entities.Round) (int64, error) {
	result, err := db.Db.Exec("INSERT INTO round (guess, guessResult, gameNum) VALUES (?,?,?)", obj.Guess, obj.GuessResult, obj.GameID)
	if err != nil {
		return 0, fmt.Errorf("AddRound: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddRound: %v", err)
	}
	return id, nil
}

func (db RoundModel) Find(id int64) (entities.Round, error) {
	var obj entities.Round

	row := db.Db.QueryRow("SELECT * FROM round WHERE roundId = ?", id)
	if err := row.Scan(&obj.ID, &obj.Guess, &obj.GuessTime, &obj.GuessResult, &obj.GameID); err != nil {
		if err == sql.ErrNoRows {
			return obj, fmt.Errorf("GameById %d: no such game", id)
		}
		return obj, fmt.Errorf("GameById %d: %v", id, err)
	}
	return obj, nil
}

func (db RoundModel) FindAllFor(id int64) ([]entities.Round, error) {
	var objs []entities.Round

	rows, err := db.Db.Query("SELECT * FROM round WHERE gameNum = ?", id)
	if err != nil {
		return nil, fmt.Errorf("FindRoundsfor %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var obj entities.Round
		if err := rows.Scan(&obj.ID, &obj.Guess, &obj.GuessTime, &obj.GuessResult, &obj.GameID); err != nil {
			return nil, fmt.Errorf("FindRoundsFor %v", err)
		}
		objs = append(objs, obj)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindRoundsFor %v", err)
	}
	return objs, nil
}
