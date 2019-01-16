package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wendyadi/go_test_example/data"
)

func getUserDummy(id int64) (data.User, error) {
	return data.User{
		ID:      1,
		Name:    "Rita",
		Email:   "rita@tokopedia.com",
		Phone:   "0000",
		OVOCash: 250000,
	}, nil
}

func getUserDummyErr(id int64) (data.User, error) {
	return data.User{
		ID:      1,
		Name:    "Rita",
		Email:   "rita@tokopedia.com",
		Phone:   "0000",
		OVOCash: 250000,
	}, errors.New("simple err")
}

func TestDoPaymentWithOVO(t *testing.T) {
	getUserData = getUserDummy
	testcase := []struct {
		args     paymentRequest
		err      bool
		name     string
		funcInit func()
	}{
		{
			name: "tc1- no err",
			args: paymentRequest{
				userID: 1,
				txAmt:  100000,
			},
			err: false,
			funcInit: func() {
				getUserData = getUserDummy
			},
		},
		{
			name: "tc2- err",
			args: paymentRequest{
				userID: 1,
				txAmt:  100000,
			},
			err: true,
			funcInit: func() {
				getUserData = getUserDummyErr
			},
		},
	}

	for _, tc := range testcase {
		tc.funcInit()
		err := DoPaymentWithOVO(tc.args)
		assert.Equal(t, (err != nil), tc.err, "error doesnt match")
	}
}
