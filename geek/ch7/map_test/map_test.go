package map_test

import (
	"errors"
	"testing"
)

func TestMapInit(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	t.Log(m1, len(m1))
	m2 := make(map[string]string)
	m2["a"] = "A"
	m2["b"] = "B"
	m2["c"] = "C"
	t.Log(m2, len(m2))
	m3 := make(map[int]int, 10)
	t.Log(m3, len(m3))
}

func TestAccessNotExistsKey(t *testing.T) {
	m1 := make(map[string]int, 10)
	m1["a"] = 69
	if v, ok := m1["a"]; ok {
		t.Log(v)
	} else {
		m1["a"] = 58
		t.Log(m1)
	}
}

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"

	if got != want {
		t.Errorf("got '%s' want '%s' ", got, want)
	}
}

func Search(dictionary map[string]string, s string) string {
	if val, ok := dictionary[s]; ok {
		return val
	}
	return ""
}

// 使用Map的Value作为一个函数，需要在用的时候把函数的参数传进去
func TestMapFunc(t *testing.T) {
	m1 := make(map[string]func(int, int) (int, error))
	m1["plus"] = func(one int, two int) (int, error) {
		return one + two, nil
	}
	m1["multi"] = func(one int, two int) (int, error) {
		return one * two, nil
	}
	m1["sub"] = func(one int, two int) (int, error) {
		return one - two, nil
	}
	m1["div"] = func(one int, two int) (int, error) {
		if two == 0 {
			return 0, errors.New("the denominator must not be 0")
		}
		return one / two, nil
	}
	m1["squareAdd"] = func(one int, two int) (int, error) {
		return one*one + two*two, nil
	}
	t.Log(m1["multi"](2, 4))
	t.Log(m1["div"](8, 0))
}

func TestMapToSet(t *testing.T) {
	mapSet := make(map[interface{}]bool)
	data, err := SetSearch(mapSet, "data")
	if err != nil {
		t.Logf("ser search error : %+v", err)
	}
	if !data {
		err := AddMap(mapSet, "data")
		if err != nil {
			t.Logf("add map set error : %+v", err)
		}
	}
}

func AddMap(set map[interface{}]bool, s string) error {
	set[s] = true
	return nil
}

func SetSearch(set map[interface{}]bool, s string) (bool, error) {
	if val, ok := set[s]; ok {
		return val, nil
	}
	return false, nil
}
