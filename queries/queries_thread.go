package queries

import (
	"database/sql"
	"dbAPI/models"
	"log"
)

func ThreadInsert(db *sql.DB, thread *models.Thread) error {
	err := db.QueryRow(`
		INSERT INTO threads (author, created, forum, message, title, slug)
		VALUES (
			$1, $2, $3, $4, $5,
			NULLIF($6, '')
		)
		RETURNING *`,
		thread.Author, thread.Created, thread.Forum, thread.Message, thread.Title, thread.Slug).
		Scan(&thread.ID, &thread.Author, &thread.Forum, &thread.Slug, &thread.Title, &thread.Message, &thread.Created, &thread.Votes)

	_, _ = db.Exec(`
		UPDATE forums
		SET threads=threads+1
		WHERE slug=$1
		`,
		thread.Forum)

	return err
}

func ThreadGetBySlug(db *sql.DB, slug string) (*models.Thread, error) {
	thread := models.Thread{}
	err := db.QueryRow(`
		SELECT * FROM threads WHERE slug = $1
		`, slug).
		Scan(&thread.ID, &thread.Author, &thread.Forum, &thread.Slug, &thread.Title, &thread.Message, &thread.Created, &thread.Votes)

	return &thread, err
}

func ThreadGetByID(db *sql.DB, id int) (*models.Thread, error) {
	thread := models.Thread{}
	err := db.QueryRow(`
		SELECT * FROM threads WHERE id = $1
		`, id).
		Scan(&thread.ID, &thread.Author, &thread.Forum, &thread.Slug, &thread.Title, &thread.Message, &thread.Created, &thread.Votes)

	return &thread, err
}

func ThreadVoteIncr(db *sql.DB, thread *models.Thread) error {
	err := db.QueryRow(`
		UPDATE threads
		SET votes = votes + 1
		WHERE id = $1
		RETURNING *
	`, thread.ID).
		Scan(&thread.ID, &thread.Author, &thread.Forum, &thread.Slug, &thread.Title, &thread.Message, &thread.Created, &thread.Votes)

	return err
}

func ThreadVoteDecr(db *sql.DB, thread *models.Thread) error {
	err := db.QueryRow(`
		UPDATE threads
		SET votes = votes - 1
		WHERE id = $1
		RETURNING *
	`, thread.ID).
		Scan(&thread.ID, &thread.Author, &thread.Forum, &thread.Slug, &thread.Title, &thread.Message, &thread.Created, &thread.Votes)

	return err
}

func ThreadQuery(db *sql.DB, forumSlug string, limit int, desc bool, since string) *models.Threads {
	queryStr := `SELECT * FROM threads WHERE forum=$1`
	if since != "" {
		if desc {
			queryStr += " AND created <= $3"
		} else {
			queryStr += " AND created >= $3"
		}
	}
	queryStr += " ORDER BY created"
	if desc {
		queryStr += " DESC"
	}
	queryStr += " LIMIT NULLIF($2, 0)"

	var rows *sql.Rows
	var err error
	if since != "" {
		rows, err = db.Query(queryStr, forumSlug, limit, since)
	} else {
		rows, err = db.Query(queryStr, forumSlug, limit)
	}

	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}

	defer rows.Close()
	threads := models.Threads{}
	for rows.Next() {
		var result models.Thread
		if err := rows.Scan(&result.ID, &result.Author, &result.Forum, &result.Slug, &result.Title, &result.Message, &result.Created, &result.Votes); err != nil {
			// log.Fatal(err)
			log.Println(err)
		}
		if result.ID != 0 {
			threads = append(threads, &result)
		}
	}
	return &threads
}

func ThreadUpdate(db *sql.DB, thread *models.Thread) error {
	err := db.QueryRow(`
		UPDATE threads
		SET message = COALESCE(NULLIF($2, ''), message),
			title = COALESCE(NULLIF($3, ''), title)
		WHERE id=$1
		RETURNING *
		`,
		thread.ID, thread.Message, thread.Title).
		Scan(&thread.ID, &thread.Author, &thread.Forum, &thread.Slug, &thread.Title, &thread.Message, &thread.Created, &thread.Votes)

	return err
}
