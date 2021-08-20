package sproc

// Reader is the interface to wrap custom data reader functionality
type Reader interface {
	// Read can used to implement the data streaming
	Read(call func(v interface{}))
}