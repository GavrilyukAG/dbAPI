package queries

import (
	"database/sql"
	"dbAPI/models"
	"log"
)

func PostInsert(db *sql.DB, post *models.Post) error {
	err := db.QueryRow(`
        INSERT INTO posts (author, forum, thread, message, created, isEdited, parent, path)
        VALUES (
			$1, $2, $3, $4, $5, $6, $7,
			((SELECT path FROM posts p WHERE p.id=$7) || (SELECT currval('posts_id_seq')::integer))
		)
        RETURNING id
        `,
		post.Author, post.Forum, post.Thread, post.Message, post.Created, post.IsEdited, post.Parent).
		Scan(&post.ID)
	// Scan(&post.ID, &post.Author, &post.Forum, &post.Thread, &post.Message, &post.Created, &post.IsEdited, &post.Parent)

	_, _ = db.Exec(`
		UPDATE forums
			SET posts=posts+1
		WHERE slug=$1
		`,
		post.Forum)

	return err
}

func PostGetFlat(db *sql.DB, threadID, limit int, desc bool, since int) *models.Posts {
	queryStr := `SELECT * FROM posts WHERE thread=$1`
	if since != 0 {
		if desc {
			queryStr += " AND id < $3"
		} else {
			queryStr += " AND id > $3"
		}
	}

	if desc {
		queryStr += " ORDER BY id DESC LIMIT NULLIF($2, 0)"
	} else {
		queryStr += " ORDER BY id LIMIT NULLIF($2, 0)"
	}

	var rows *sql.Rows
	var err error
	if since != 0 {
		rows, err = db.Query(queryStr, threadID, limit, since)
	} else {
		rows, err = db.Query(queryStr, threadID, limit)
	}

	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}

	defer rows.Close()
	posts := models.Posts{}
	for rows.Next() {
		var post models.Post
		var path []byte
		if err := rows.Scan(&post.ID, &post.Author, &post.Forum, &post.Thread, &post.Message, &post.Created, &post.IsEdited, &post.Parent, &path); err != nil {
			// log.Fatal(err)
			log.Println("Error with posts: ", err)
		}
		if post.ID != 0 {
			posts = append(posts, &post)
		}
	}
	return &posts
}

func PostGetTree(db *sql.DB, threadID int, limit int, desc bool, since int) *models.Posts {
	queryStr := `
	WITH p AS (
		SELECT *, row_number() OVER (ORDER BY path) FROM posts
		WHERE thread=$1
	)
	SELECT id, author, forum, thread, message, created, isEdited, parent FROM p
	`

	if since != 0 {
		if desc {
			queryStr += ` WHERE p.row_number < (SELECT row_number FROM p WHERE p.id=$3)`
		} else {
			queryStr += ` WHERE p.row_number > (SELECT row_number FROM p WHERE p.id=$3)`
		}
	}

	if desc {
		queryStr += " ORDER BY path DESC, id DESC LIMIT NULLIF($2, 0)"
	} else {
		queryStr += " ORDER BY path, id LIMIT NULLIF($2, 0)"
	}

	var rows *sql.Rows
	var err error
	if since != 0 {
		rows, err = db.Query(queryStr, threadID, limit, since)
	} else {
		rows, err = db.Query(queryStr, threadID, limit)
	}

	if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}

	defer rows.Close()
	posts := models.Posts{}
	for rows.Next() {
		var post models.Post
		// var path []byte
		if err := rows.Scan(&post.ID, &post.Author, &post.Forum, &post.Thread, &post.Message, &post.Created, &post.IsEdited, &post.Parent /*, &path*/); err != nil {
			// log.Fatal(err)
			log.Println("Error with posts: ", err)
		}
		if post.ID != 0 {
			posts = append(posts, &post)
		}
	}
	return &posts
}

func PostGetParentTree(db *sql.DB, threadID int, limit int, desc bool, since int) *models.Posts {
	queryStr := `
	SELECT p.id, p.author, p.forum, p.thread, p.message, p.created, p.isEdited, p.parent
	FROM posts p
	JOIN (
		SELECT id FROM posts
		WHERE thread=$1 AND parent=0
	`

	if since != 0 {
		if desc {
			queryStr += `
			AND array[path[1]] && array(
				SELECT path[1] FROM posts WHERE path[1] < (
					SELECT path[1] FROM posts WHERE path && array[$3]::integer[] AND thread=$1
					ORDER BY path[1] DESC, path
					LIMIT 1
				)
			)
			`
		} else {
			queryStr += `
			AND array[path[1]] && array(
				SELECT path[1] FROM posts WHERE path[1] > (
					SELECT path[1] FROM posts WHERE path && array[$3]::integer[] AND thread=$1
					LIMIT 1
				)
			)
			`
		}
	}

	if desc {
		queryStr += `
		ORDER BY path[1] DESC, path
		LIMIT NULLIF($2, 0)
		) AS t ON t.id=path[1]
		ORDER BY path[1] DESC, path
		`
	} else {
		queryStr += `
		LIMIT NULLIF($2, 0)
		) AS t ON t.id=path[1]
		ORDER BY path
		`
	}

	var rows *sql.Rows
	var err error
	if since != 0 {
		rows, err = db.Query(queryStr, threadID, limit, since)
	} else {
		rows, err = db.Query(queryStr, threadID, limit)
	}

	if err != nil {
		// log.Fatal(err)
		log.Println("Parent tree query ERROR", err)
	}

	defer rows.Close()
	posts := models.Posts{}
	for rows.Next() {
		var post models.Post
		// var path []byte
		if err := rows.Scan(&post.ID, &post.Author, &post.Forum, &post.Thread, &post.Message, &post.Created, &post.IsEdited, &post.Parent /*, &path*/); err != nil {
			// log.Fatal(err)
			log.Println("Error with posts: ", err)
		}
		if post.ID != 0 {
			posts = append(posts, &post)
		}
	}
	return &posts
}

func PostGetDetails(db *sql.DB, postID int) (*models.PostFull, error) {
	postDetails := models.PostFull{}

	post := models.Post{}
	err := db.QueryRow(`
		SELECT id, author, forum, thread, message, created, isEdited, parent
		FROM posts
		WHERE id=$1
	`,
		postID).
		Scan(&post.ID, &post.Author, &post.Forum, &post.Thread, &post.Message, &post.Created, &post.IsEdited, &post.Parent)

	postDetails.Post = &post

	return &postDetails, err
}

func PostUpdate(db *sql.DB, postID int, message string) (*models.Post, error) {
	post := models.Post{}
	err := db.QueryRow(`
		UPDATE posts
		SET message=$2, isEdited=true
		WHERE id=$1
		RETURNING id, author, forum, thread, message, created, isEdited, parent
	`,
		postID, message).
		Scan(&post.ID, &post.Author, &post.Forum, &post.Thread, &post.Message, &post.Created, &post.IsEdited, &post.Parent)

	return &post, err
}

func PostGetById(db *sql.DB, postID int64) (*models.Post, error) {
	post := models.Post{}
	err := db.QueryRow(`
		SELECT id, author, forum, thread, message, created, isEdited, parent
		FROM posts
		WHERE id=$1
	`,
		postID).
		Scan(&post.ID, &post.Author, &post.Forum, &post.Thread, &post.Message, &post.Created, &post.IsEdited, &post.Parent)

	return &post, err
}
