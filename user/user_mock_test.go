package user

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var mock sqlmock.Sqlmock

func InitMock() {
	var err error
	var dbx *sql.DB
	dbx, mock, err = sqlmock.New()
	if err != nil {
		panic(err.Error())
	}

	db = sqlx.NewDb(dbx, "postgres")

}

// func TestGetUserMock(t *testing.T) {
// 	InitMock()
// 	userID := int64(20)

// 	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone"})
// 	rows.AddRow(20, "User Test Mock", "user.mock@gmail.com", "+72303838388383")

// 	mock.ExpectQuery("SELECT id, name, email, phone FROM training_user").WithArgs(userID).WillReturnRows(rows)

// 	user, err := GetUserByID(userID)
// 	if err != nil {
// 		t.Error(err.Error())
// 		return
// 	}

// 	if user.Email != "user.mock@gmail.com" {
// 		t.Errorf("Test failed expect email %s but got %s", "user.mock@gmail.com", user.Email)
// 	}

// 	// t.Log(user)
// }
