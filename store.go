package sproc

// Store is the interface to wrap basic key value store functionalities
type Store interface {
	Get(key interface{}) (interface{}, error)
	Set(k interface{}, v interface{}) error
	GetAll() ([]interface{}, error)
}
