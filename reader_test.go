package sproc

import (
	"github.com/noelyahan/stream-processor/domain"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func mockStringUserReader() io.Reader {
	return strings.NewReader("{\"Id\":\"8422699\",\"Name\":\"Apexal\"}#{\"Id\":\"53201765\",\"Name\":\"ArturoCamacho0\"}#{\"Id\":\"2631623\",\"Name\":\"onosendi\"}#{\"Id\":\"52553915\",\"Name\":\"anggi1234\"}")
}

func TestStringReader_Read(t *testing.T) {
	data := mockStringUserReader()
	userStream := stringReader{
		reader:  data,
		encoder: new(domain.User),
	}
	uu := make([]*domain.User, 0)

	userStream.Read(func(v interface{}) {
		uu = append(uu, v.(*domain.User))
	})

	for _, u := range uu {
		if u.Id == "" {
			t.Error("expect user id to be mapped")
		}

		if u.Name == "" {
			t.Error("expect user name to be mapped")
		}
	}
	nExpect := 4
	if len(uu) != nExpect {
		t.Errorf("expect [%v] but got [%v]", nExpect, len(uu))
	}
}

func TestCsvReader_Read(t *testing.T) {
	f1, err := os.Open("./data/actors.csv")
	if err != nil {
		log.Fatal(err)
	}
	c := 0
	NewCSVReader(f1, new(domain.User)).Read(func(v interface{}) {
		c++
	})
	if c == 0 {
		t.Error("csv data expect not to be empty")
	}
	t.Logf("total csv data read: [%v]", c)
}
