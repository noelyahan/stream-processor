package sproc

import "fmt"

// Store is the interface to wrap basic key value store functionalities
type Store interface {
	Get(key interface{}) (interface{}, error)
	Set(k interface{}, v interface{}) error
	GetAll() ([]interface{}, error)
}

// NewMemStore is the in memory store implementation using thread unsafe map
func NewMemStore(name string) Store {
	return &memStore{name: name, mm: make(map[interface{}]interface{})}
}

type memStore struct {
	name string
	mm   map[interface{}]interface{}
}

func (m memStore) Get(key interface{}) (interface{}, error) {
	v, ok := m.mm[key]
	if !ok {
		return nil, fmt.Errorf("store [%v] not found for key [%v]", m.name, key)
	}
	return v, nil
}
func (m *memStore) Set(key interface{}, val interface{}) error {
	m.mm[key] = val
	return nil
}
func (m memStore) GetAll() ([]interface{}, error) {
	aa := make([]interface{}, 0)
	for _, v := range m.mm {
		aa = append(aa, v)
	}
	return aa, nil
}
