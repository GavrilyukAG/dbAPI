package queries

import (
	"database/sql"
	"dbAPI/models"
	"log"
)

func UserInsert(db *sql.DB, user models.User) (res sql.Result, err error) {
	res, err = db.Exec(`
		INSERT INTO users (about, email, fullname, nickname) VALUES ($1, $2, $3, $4)
		`,
		user.About, user.Email, user.Fullname, user.Nickname)

	return
}

func UserGetAll(db *sql.DB, nickname string, email string) (*models.Users, error) {
	rows, err := db.Query(`
		SELECT * FROM users WHERE nickname=$1 OR email=$2
		`,
		nickname, email)

	if err != nil {
		// log.Println(err)
		return nil, err
	}

	defer rows.Close()
	users := models.Users{}
	for rows.Next() {
		var result models.User
		if err := rows.Scan(&result.Nickname, &result.About, &result.Email, &result.Fullname); err != nil {
			// log.Fatal(err)
		}
		users = append(users, &result)
	}

	return &users, nil
}

func UserGetByNickname(db *sql.DB, nickname string) (*models.User, error) {
	user := models.User{}
	err := db.QueryRow(`
		SELECT * FROM users WHERE nickname=$1
		`,
		nickname).
		Scan(&user.Nickname, &user.About, &user.Email, &user.Fullname)

	return &user, err
}

func UserUpdate(db *sql.DB, user *models.User) error {
	err := db.QueryRow(`
		UPDATE users
		SET about = COALESCE(NULLIF($2, ''), about),
			email = COALESCE(NULLIF($3, ''), email),
			fullname = COALESCE(NULLIF($4, ''), fullname)
		WHERE nickname=$1
		RETURNING about, email, fullname
		`,
		user.Nickname, user.About, user.Email, user.Fullname).
		Scan(&user.About, &user.Email, &user.Fullname)

	return err
}

func UserGetAllBySlug(db *sql.DB, forumSlug, since string, limit int, desc bool) (*models.Users, error) {
	queryStr := `
		SELECT DISTINCT u.nickname COLLATE "ucs_basic", u.about, u.email, u.fullname
		FROM users u
		LEFT JOIN threads t ON t.author=u.nickname
		LEFT JOIN posts p ON p.author=u.nickname
		WHERE (t.forum=$1 OR p.forum=$1)
	`
	if since != "" {
		if desc {
			queryStr += ` AND u.nickname COLLATE "ucs_basic" < $3 COLLATE "ucs_basic"`
		} else {
			queryStr += ` AND u.nickname COLLATE "ucs_basic" > $3 COLLATE "ucs_basic"`
		}
	}

	if desc {
		queryStr += `
			ORDER BY u.nickname COLLATE "ucs_basic" DESC
			LIMIT $2
		`
	} else {
		queryStr += `
			ORDER BY u.nickname COLLATE "ucs_basic"
			LIMIT $2
		`
	}

	var (
		rows *sql.Rows
		err  error
	)
	if since != "" {
		rows, err = db.Query(queryStr, forumSlug, limit, since)
	} else {
		rows, err = db.Query(queryStr, forumSlug, limit)
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	users := models.Users{}
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.Nickname, &user.About, &user.Email, &user.Fullname); err != nil {
			log.Println(err)
		}
		users = append(users, &user)
	}

	return &users, err
}
