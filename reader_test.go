package sproc

import (
	"github.com/noelyahan/stream-processor/domain"
	"io"
	"strings"
	"testing"
)

func mockStringReader() io.Reader  {
	return strings.NewReader("{\"Id\":\"8422699\",\"Name\":\"Apexal\"}#{\"Id\":\"53201765\",\"Name\":\"ArturoCamacho0\"}#{\"Id\":\"2631623\",\"Name\":\"onosendi\"}#{\"Id\":\"52553915\",\"Name\":\"anggi1234\"}")
}

func TestStringReader_Read(t *testing.T) {
	data := mockStringReader()
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
