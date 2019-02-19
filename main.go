package main

import (
	"fmt"
	"os"

	"github.com/ywoonder/practice-gotest/data"
)

func main() {
	os.Exit(0)
}

type (
	paymentRequest struct {
		userID int64
		txAmt  int64
	}
	obj struct{}

	FGetUserData func(id int64) (data.User, error)
)

// alternatively can also use interface{}
var getUserData FGetUserData

func init() {
	getUserData = data.GetUserByID
}

/*
Mocking as testing pattern of dependency injection
each layer is tested separately
	1. Data layer
	2. Logic/Business Layer
	3. API Layer
*/

func DoPaymentWithOVO(req paymentRequest) error {
	//1. get user Data (saldo OVO)
	userData, err := getUserData(req.userID)
	if err != nil {
		return err
	}

	fmt.Println(userData)
	//2. validate
	//3. reduce saldo
	//4. return sucess
	return nil
}
