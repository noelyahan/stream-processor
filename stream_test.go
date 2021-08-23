package sproc

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"

	"github.com/noelyahan/stream-processor/domain"
)

func TestNewStream(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringUserReader()
	reader := stringReader{
		reader:  data,
		encoder: new(domain.User),
	}

	NewStream(reader, store).Transform(func(val interface{}) (k interface{}, v interface{}) {
		u := val.(*domain.User)
		return u.Id, u.Name
	}).ToStore()

	aa, _ := store.GetAll()
	if len(aa) == 0 {
		t.Error("expect data from final state store")
	}
}

func TestStream_Transform(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringUserReader()
	reader := stringReader{
		reader:  data,
		encoder: new(domain.User),
	}

	NewStream(reader, store).Transform(func(val interface{}) (k interface{}, v interface{}) {
		u := val.(*domain.User)
		m := domain.ResultMat{
			Id:    u.Id,
			Name:  u.Name,
			Count: time.Now().Second(),
		}
		return u.Id, m
	}).ToStore()

	aa, _ := store.GetAll()
	expectMap := reflect.TypeOf(domain.ResultMat{}).Name()
	for _, a := range aa {
		got := reflect.TypeOf(a).Name()
		if got != expectMap {
			t.Errorf("expect to be mapped [%v] but got [%v]", expectMap, got)
		}
	}
}

func TestStream_Sort(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringUserReader()
	reader := stringReader{
		reader:  data,
		encoder: new(domain.User),
	}

	strm := NewStream(reader, store).Transform(func(val interface{}) (k interface{}, v interface{}) {
		u := val.(*domain.User)
		m := domain.ResultMat{
			Id:    u.Id,
			Name:  u.Name,
			Count: rand.Intn(100),
		}
		return u.Id, m
	}).ToStore().Sort(func(arr []interface{}) []interface{} {
		uu := make([]domain.ResultMat, 0)
		for _, a := range arr {
			uu = append(uu, a.(domain.ResultMat))
		}
		sort.Sort(domain.ByCount(uu))

		aa := make([]interface{}, 0)
		for _, u := range uu {
			aa = append(aa, u)
		}
		return aa
	})

	last := 200
	strm.Print(func(idx, val interface{}) {
		a := val.(domain.ResultMat)
		if last < a.Count {
			t.Fatal("error in sorting")
		}
		last = a.Count
	})
}

func TestStream_Limit(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringUserReader()
	reader := stringReader{
		reader:  data,
		encoder: new(domain.User),
	}

	expectL := 2
	strm := NewStream(reader, store).Transform(func(val interface{}) (k interface{}, v interface{}) {
		u := val.(*domain.User)
		m := domain.ResultMat{
			Id:    u.Id,
			Name:  u.Name,
			Count: rand.Intn(100),
		}
		return u.Id, m
	}).ToStore().Limit(expectL)

	count := 0
	strm.Print(func(idx, val interface{}) {
		count++
	})
	if expectL != count {
		t.Errorf("expect limit [%v] got limit [%v]", expectL, count)
	}
}

func TestStream_Filter(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringUserReader()
	reader := stringReader{
		reader:  data,
		encoder: new(domain.User),
	}

	strm := NewStream(reader, store).Filter(func(val interface{}) bool {
		u := val.(*domain.User)
		return u.Id == "8422699"
	}).Transform(func(val interface{}) (k interface{}, v interface{}) {
		u := val.(*domain.User)
		m := domain.ResultMat{
			Id:    u.Id,
			Name:  u.Name,
			Count: rand.Intn(100),
		}
		return u.Id, m
	}).ToStore()

	expectL := 1
	count := 0
	strm.Print(func(idx, val interface{}) {
		count++
	})
	if expectL != count {
		t.Errorf("expect limit [%v] got limit [%v]", expectL, count)
	}
}
