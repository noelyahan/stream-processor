package sproc

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jszwec/csvutil"
)

// Reader is the interface to wrap custom data reader functionality
type Reader interface {
	// Read can used to implement the data streaming
	Read(call func(v interface{}))
}

func NewCSVReader(r io.Reader, e interface{}) Reader {
	return csvReader{
		reader:  r,
		encoder: e,
	}
}

type csvReader struct {
	reader  io.Reader
	encoder interface{}
}

func (c csvReader) Read(call func(v interface{})) {
	csvReader := csv.NewReader(c.reader)
	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		log.Fatal(err)
	}
	for {
		if err := dec.Decode(&c.encoder); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		call(c.encoder)
	}
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
