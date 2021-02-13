package myerror

import (
	"errors"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 2 || n > 500 {
		return nil, errors.New("param error ")
	}
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret, nil
}

func TestGetFibonacci(t *testing.T) {
	var fibonacci []int
	var err error

	if fibonacci, err = GetFibonacci(800); err != nil {
		t.Logf("err for get fibonacci list: %v", err)
	}

	t.Log(fibonacci)
}
