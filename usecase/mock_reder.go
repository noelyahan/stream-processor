package usecase

import (
	"fmt"
	"github.com/noelyahan/stream-processor/domain"
	"math/rand"
)

type mockReader struct {
	data []interface{}
}

func (m *mockReader) initUsers() {
	m.data = make([]interface{}, 0)
	max := 20
	for i := 0; i < max; i++ {
		id := rand.Intn(max)
		m.data = append(m.data, &domain.User{
			Id:   fmt.Sprintf("%d", id),
			Name: fmt.Sprintf("user-%d", id),
		})
	}
}

func (m *mockReader) initRepos() {
	m.data = make([]interface{}, 0)
	max := 10
	for i := 0; i < max; i++ {
		id := rand.Intn(max)
		m.data = append(m.data, &domain.Repo{
			Id:   fmt.Sprintf("%d", id),
			Name: fmt.Sprintf("repo-%d", id),
		})
	}
}

func (m *mockReader) initEvents(eType string) {
	m.data = make([]interface{}, 0)
	max := 50
	userC := 20
	repoC := 10
	for i := 0; i < max; i++ {
		id := rand.Intn(userC)
		repoId := rand.Intn(repoC)
		m.data = append(m.data, &domain.Event{
			Id:      fmt.Sprintf("%d", id),
			Type:    eType,
			ActorId: fmt.Sprintf("%d", id),
			RepoId:  fmt.Sprintf("%d", repoId),
		})
	}
}

func (m mockReader) Read(call func(v interface{})) {
	for _, v := range m.data {
		call(v)
	}
}