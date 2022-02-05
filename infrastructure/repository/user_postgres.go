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

//Delete an user
func (u *UserPostgres) Delete(id entity.ID) error {
	stmt, err := u.db.Prepare(`DELETE FROM users WHERE id = $1`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

//Update an user
func (u *UserPostgres) Update(e *entity.User) error {
	stmt, err := u.db.Prepare(`UPDATE users SET name=$1, email= $2, paid_date= $3, due_date= $4, total_month= $5 WHERE id = $6`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Email, e.PaidDate, e.DueDate, e.TotalMonth, e.ID)

	if err != nil {
		return err
	}

	return nil
}

//List users
func (u *UserPostgres) List() ([]*entity.User, error) {
	var users []*entity.User

	stmt, err := u.db.Prepare(
		`SELECT id, name, email, paid_date, due_date, total_month
		FROM users`,
	)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	for rows.Next() {
		var u entity.User

		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.PaidDate, &u.DueDate, &u.TotalMonth)
		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil

}
