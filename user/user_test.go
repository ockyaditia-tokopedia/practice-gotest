package user

import (
	"testing"
)

func TestInsertUser(t *testing.T) {
	tc := []User{
		User{
			Name:  "User Pertama",
			Email: "user.pertama@gmail.com",
			Phone: "+6282300036200",
		},
	}

	id, err := InsertUser(tc[0])
	if err != nil {
		t.Error(err.Error())
		return
	}

	// delete when finished
	defer DeleteUser(*id)

	//check user email
	user, err := GetUserByID(*id)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	//assert
	if user.Email != tc[0].Email {
		t.Errorf("Error get user expected email: ", tc[0].Email, " but got: ", user.Email)
		return
	}

}

func TestGetUserByEmail(t *testing.T) {
	tc := []User{
		User{
			Name:  "User Pertama",
			Email: "user.pertama@gmail.com",
			Phone: "+6282300036200",
		},
	}

	id, err := InsertUser(tc[0])
	if err != nil {
		t.Error(err.Error())
		return
	}

	// delete when finished
	defer DeleteUser(*id)

	user, err := GetUserByEmail(tc[0].Email)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if user.Email != tc[0].Email {
		t.Errorf("Test failed expected Email: %s, but found %s", user.Email, tc[0].Email)
		return
	}

}
