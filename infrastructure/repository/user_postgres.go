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

//Get an user
func (u *UserPostgres) Get(id entity.ID) (*entity.User, error) {
	return getUser(id, u.db)
}

func getUser(id entity.ID, db *sql.DB) (*entity.User, error) {
	var user entity.User

	stmt, err := db.Prepare(
		`SELECT u.id, u.name, u.email, u.paid_date, u.due_date, u.total_month
		FROM users as u WHERE u.id = $1`,
	)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.PaidDate, &user.DueDate, &user.TotalMonth)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
