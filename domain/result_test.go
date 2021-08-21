package domain

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	aa := make([]ResultMat, 0)
	for i := 1; i < 10; i++ {
		aa = append(aa, ResultMat{
			Id:    fmt.Sprintf("user%d", i),
			Name:  fmt.Sprintf("user%d", i),
			Count: time.Now().Second()%i,
		})
	}
	sort.Sort(ByCount(aa))
	last := aa[0].Count
	for _, a := range aa {
		if last < a.Count {
			t.Fatal("error in sorting")
		}
		last = a.Count
	}
}
