package queries

import (
	"database/sql"
	"dbAPI/models"
)

func StatusGet(db *sql.DB) (*models.Status, error) {
	status := models.Status{}
	err := db.QueryRow(`
        SELECT count(*) FROM forums
    `).Scan(&status.Forum)

	err = db.QueryRow(`
        SELECT count(*) FROM posts
    `).Scan(&status.Post)

	err = db.QueryRow(`
        SELECT count(*) FROM threads
    `).Scan(&status.Thread)

	err = db.QueryRow(`
        SELECT count(*) FROM users
    `).Scan(&status.User)

	return &status, err
}

func StatusClear(db *sql.DB) error {
	_, err := db.Exec(`
		TRUNCATE
			users,
			forums,
			threads,
			posts,
			votes CASCADE
	`)

	return err
}
