package queries

import (
	"database/sql"

	"github.com/GavrilyukAG/dbAPI/models"
)

func ForumInsert(db *sql.DB, forum *models.Forum) error {
	err := db.QueryRow(`
        INSERT INTO forums (slug, title, "user")
        VALUES ($1, $2,
            (SELECT nickname FROM users WHERE nickname=$3))
        RETURNING *
    `,
		forum.Slug, forum.Title, forum.User).
		Scan(&forum.Slug, &forum.Posts, &forum.Threads, &forum.Title, &forum.User)

	return err
}

func ForumGetByUsername(db *sql.DB, username string) (*models.Forum, error) {
	forum := models.Forum{}
	err := db.QueryRow(`
        SELECT * FROM forums WHERE "user"=$1
        `, username).
		Scan(&forum.Slug, &forum.Posts, &forum.Threads, &forum.Title, &forum.User)

	return &forum, err
}

func ForumGetBySlug(db *sql.DB, slug string) (*models.Forum, error) {
	forum := models.Forum{}
	err := db.QueryRow(`
        SELECT * FROM forums
        WHERE slug=$1
        `, slug).
		Scan(&forum.Slug, &forum.Posts, &forum.Threads, &forum.Title, &forum.User)

	return &forum, err
}
