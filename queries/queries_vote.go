package queries

import (
	"database/sql"

	"github.com/GavrilyukAG/dbAPI/models"
)

func VoteInsert(db *sql.DB, vote models.Vote) (res sql.Result, err error) {
	res, err = db.Exec(`
        INSERT INTO votes (nickname, thread, voice) VALUES ($1, $2, $3)
        `,
		vote.Nickname, vote.Thread, vote.Voice)

	return
}

func VoteGetByNickname(db *sql.DB, nickname string) (*models.Vote, error) {
	vote := models.Vote{}
	err := db.QueryRow(`
		SELECT * FROM votes WHERE nickname=$1
		`,
		nickname).
		Scan(&vote.Nickname, &vote.Thread, &vote.Voice)

	return &vote, err
}

func VoteUpdate(db *sql.DB, vote *models.Vote) error {
	err := db.QueryRow(`
        UPDATE votes
        SET voice = $1
        RETURNING voice
    `,
		vote.Voice).
		Scan(&vote.Voice)

	return err
}
