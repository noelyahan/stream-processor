package sproc

import (
	"fmt"
	"testing"
)

func TestNewMemStore(t *testing.T) {
	store := NewMemStore("my-store")

	// Error get test
	v, err := store.Get("foo")
	if v != nil {
		t.Error("should be empty")
	}

	if err == nil {
		t.Error("should throw an error on empty get")
	}

	// Success set
	err = store.Set("foo", "bar")
	if err != nil {
		t.Error("should not throw an error")
	}

	// Success get test
	v, err = store.Get("foo")
	if v == nil {
		t.Error("should be empty but got", v)
	}

	if err != nil {
		t.Error("should throw an error on empty get")
	}

	// Reset store
	store = NewMemStore("my-store")
	expectN := 5
	// Success get all
	for i := 0; i < expectN; i++ {
		k := fmt.Sprintf("foo%v", i)
		v := fmt.Sprintf("bar%v", i)
		store.Set(k, v)
	}

	for i := 0; i < expectN; i++ {
		k := fmt.Sprintf("foo%v", i)
		v := fmt.Sprintf("bar%v", i)
		store.Set(k, v)
	}
	vv, err := store.GetAll()
	if len(vv) != expectN {
		t.Errorf("ecpect only [%v] elements but got [%v]", expectN, len(vv))
	}

}
