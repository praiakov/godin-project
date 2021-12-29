package repository

import (
	"database/sql"

	"github.com/praiakov/godin/entity"
)

//User
//Postgres repo
type UserPostgres struct {
	db *sql.DB
}

//NewUserPostgresSQL create new repository
func NewUserMyPostgres(db *sql.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

//Create an user
func (u *UserPostgres) Create(e *entity.User) (entity.ID, error) {
	stmt, err := u.db.Prepare(
		`INSERT INTO users(id,name,email,paid_date,due_date,total_month) 
		 VALUES($1,$2,$3,$4,$5,$6)`)

	if err != nil {
		panic("error to create user")
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, e.Name, e.Email, e.PaidDate, e.DueDate, e.TotalMonth)

	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}
