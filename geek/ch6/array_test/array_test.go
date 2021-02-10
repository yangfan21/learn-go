package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{2, 3, 4, 5, 6, 7}
	arr1[1] = 8
	t.Log(arr, arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 67, 8, 9}
	for _, val := range arr {
		//t.Log(index)
		t.Log(val)
	}
}

func TestMultiArray(t *testing.T) {
	arr := [...][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	t.Log(arr)
	arr2 := [...][3]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}
	for index, val := range arr2 {
		t.Log(index)
		t.Log(val)
	}
}

// 左边包含右边不包含
func TestIndexArray(t *testing.T) {
	arr := [...][3]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}
	arr1 := arr[1:]
	arr2 := arr[:1]
	t.Log(arr1)
	t.Log(arr2)
}
