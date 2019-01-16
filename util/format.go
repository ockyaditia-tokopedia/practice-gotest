package util

import (
	"fmt"
	"strconv"
)

func ToInt(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}

func ArrayToInt(strs []string) []int {
	arr := []int{}
	for _, val := range strs {
		arr = append(arr, ToInt(val))
	}
	return arr
}

func MaskCardNumber(c string) (string, error) {
	l := len([]rune(c))
	if l != 16 {
		return "", fmt.Errorf("invalid card number")
	}

	r := make([]rune, l)
	copy(r, []rune(c))
	for i := 6; i < 12; i++ {
		r[i] = '*'
	}

	return string(r), nil
}
