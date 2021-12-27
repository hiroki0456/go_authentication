package repository

import (
	"database/sql"
	"practice/auth/model"

	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository() *Repository {
	db, err := sql.Open("mysql", "root@/auth")
	if err != nil {
		panic(err)
	}
	return &Repository{
		DB: db,
	}
}

func SelectUsers(r *Repository) (model.Users, error) {
	users := model.Users{}
	rows, err := r.DB.Query(`SELECT * FROM Users`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := new(model.User)
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func InsertUser(r *Repository, user *model.User) error {
	stmt, err := r.DB.Prepare(`INSERT INTO Users(name, email, password) VALUES(?, ?, ?)`)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(&user.Name, &user.Email, &user.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = model.UserId(id)
	return nil
}

func SearchForEmail(r *Repository, email string) (*model.User, error) {
	rows, err := r.DB.Query(`SELECT * FROM Users WHERE email = ?`, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}
