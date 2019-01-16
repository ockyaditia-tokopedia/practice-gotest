package user

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const connectionDB = "user=postgres password=postgres dbname=se_training host=127.0.0.1 port=5432 sslmode=disable"

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Connect("postgres", connectionDB)
	if err != nil {
		panic(err.Error())
	}
}

// User struct reflect training_user table
type User struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Phone string `db:"phone"`
}

// GetUserByID selects user by id
func GetUserByID(id int64) (*User, error) {
	query := `SELECT id, name, email, phone FROM training_user WHERE id = $1`

	var user User

	err := db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// InsertUser insert new user
func InsertUser(user User) (*int64, error) {
	query := `INSERT INTO training_user(name, email, phone) VALUES($1, $2, $3) RETURNING id`
	var id int64
	err := db.QueryRow(query, user.Name, user.Email, user.Phone).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

// GetUserByEmail return user by email
func GetUserByEmail(email string) (*User, error) {
	query := `SELECT id, name, email, phone FROM training_user WHERE LOWER(email) = LOWER($1)`

	var user User
	err := db.Get(&user, query, email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser will delete user by ID
func DeleteUser(id int64) error {
	query := `DELETE FROM training_user WHERE id = $1`

	_, err := db.Exec(query, id)
	return err
}
