package usecase

import (
	"fmt"
	"sort"

	sproc "github.com/noelyahan/stream-processor"
	"github.com/noelyahan/stream-processor/domain"
)

type RepoAnalytics struct {
	RepoStream  sproc.Reader
	EventSteam  sproc.Reader
	repoStore   sproc.Store
	resultStore sproc.Store
}

func (u *RepoAnalytics) TopRepositoriesByCommit() {
	u.processRepoStream()
	u.processEventStream(func(val interface{}) bool {
		e := val.(*domain.Event)

		repo, _ := u.repoStore.Get(e.RepoId)
		if repo == nil {
			return false
		}

		return e.Type == domain.CommitEventType
	})
}

func (u *RepoAnalytics) TopRepositoriesByWatchEvent() {
	u.processRepoStream()
	u.processEventStream(func(val interface{}) bool {
		e := val.(*domain.Event)

		repo, _ := u.repoStore.Get(e.RepoId)
		if repo == nil {
			return false
		}

		return e.Type == domain.WatchEventType
	})
}

func (u *RepoAnalytics) processRepoStream() {
	u.repoStore = sproc.NewMemStore("repos-store")
	sproc.NewStream(u.RepoStream, u.repoStore).
		Transform(func(val interface{}) (k interface{}, v interface{}) {
			u := val.(*domain.Repo)
			newU := domain.Repo{
				Id:   u.Id,
				Name: u.Name,
			}
			return newU.Id, newU
		}).
		ToStore()
}

func (u *RepoAnalytics) processEventStream(predicate func(val interface{}) bool) {
	u.resultStore = sproc.NewMemStore("result-store")
	sproc.NewStream(u.EventSteam, u.resultStore).
		Filter(predicate).
		Transform(func(val interface{}) (k interface{}, v interface{}) {
			e := val.(*domain.Event)
			o, _ := u.repoStore.Get(e.RepoId)
			repo := o.(domain.Repo)
			o, _ = u.resultStore.Get(e.RepoId)
			c := 0
			if o != nil {
				m := o.(domain.ResultMat)
				c = m.Count
			}
			c++

			return e.RepoId, domain.ResultMat{
				Id:    e.RepoId,
				Name:  repo.Name,
				Count: c,
			}
		}).
		ToStore().
		Sort(func(arr []interface{}) []interface{} {
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
		}).
		Limit(10).
		Print(func(idx, val interface{}) {
			o := val.(domain.ResultMat)
			fmt.Printf("[%v] \t%v\n", idx.(int)+1, o.Name)
		})
}
