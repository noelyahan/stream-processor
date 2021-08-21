package sproc

import (
	"github.com/noelyahan/stream-processor/domain"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)


func TestNewStream(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringReader()
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
	data := mockStringReader()
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
		if  got != expectMap {
			t.Errorf("expect to be mapped [%v] but got [%v]", expectMap, got)
		}
	}
}

func TestStream_Sort(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringReader()
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
	strm.Print(func(val interface{}) {
		a := val.(domain.ResultMat)
		if last < a.Count {
			t.Fatal("error in sorting")
		}
		last = a.Count
	})
}

func TestStream_Limit(t *testing.T) {
	store := NewMemStore("my-store")
	data := mockStringReader()
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
	strm.Print(func(val interface{}) {
		count++
	})
	if expectL != count {
		t.Errorf("expect limit [%v] got limit [%v]", expectL, count)
	}
}
