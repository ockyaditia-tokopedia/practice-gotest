package data

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type (
	User struct {
		ID      int64  `db:"id"`
		Name    string `db:"name"`
		Email   string `db:"email"`
		Phone   string `db:"phone"`
		OVOCash int64  `db:"ovo_cash"`
	}
)

const (
	QueryGetUserByID = `
		SELECT
			id,
			name,
			email, 
			phone,
			ovo_cash,
		FROM ws_user
		WHERE id = $1
	`

	QueryGetUsersByName = `
	SELECT
		id,
		name,
		email, 
		phone,
		ovo_cash,
	FROM ws_user
	WHERE name like %$1%
`
)

var db *sqlx.DB

func GetUserByID(id int64) (User, error) {
	var user User
	err := db.Get(&user, QueryGetUserByID, id)
	if err != nil {
		return User{}, err
	}

	fmt.Printf("user data %+v", user)
	return user, nil
}

func GetUserByName(name string) ([]User, error) {
	var users []User
	err := db.Get(&users, QueryGetUsersByName, name)
	if err != nil {
		return nil, err
	}

	return users, nil
}
