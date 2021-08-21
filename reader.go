package sproc

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"
)

// Reader is the interface to wrap custom data reader functionality
type Reader interface {
	// Read can used to implement the data streaming
	Read(call func(v interface{}))
}

type stringReader struct {
	reader  io.Reader
	encoder interface{}
}

func (c stringReader) Read(call func(v interface{})) {
	b, _ := ioutil.ReadAll(c.reader)
	aa := strings.Split(string(b), "#")
	for _, a := range aa {
		json.Unmarshal([]byte(a), &c.encoder)
		call(c.encoder)
	}
}