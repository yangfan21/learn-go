package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, []int{1, 2, 3, 4, 5}...)
	t.Log(s0, len(s0), cap(s0))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 4)
	t.Log(s2, len(s2), cap(s2))
	//t.Log(s2[0], s2[1], s2[2], s2[3], s2[4])
}

func TestGrowingSlice(t *testing.T) {
	var s []int
	for i := 0; i < 100; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := [12]string{"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December"}
	t.Log(year)
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	Q3 := year[6:9]
	t.Log(Q3, len(Q3), cap(Q3))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknown"
	t.Log(Q2, len(Q2), cap(Q2))
	t.Log(Q3, len(Q3), cap(Q3))
	t.Log(year)
}
