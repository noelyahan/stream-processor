package usecase

import (
	"fmt"
	sproc "github.com/noelyahan/stream-processor"
	"github.com/noelyahan/stream-processor/domain"
	"sort"
)



type UserAnalytics struct {
	UserStream  sproc.Reader
	EventSteam  sproc.Reader
	userStore   sproc.Store
	resultStore sproc.Store
}

func (u *UserAnalytics) ActiveUsersByCommitPRs() {
	u.userStore = sproc.NewMemStore("user-store")

	sproc.NewStream(u.UserStream, u.userStore).
		Transform(func(val interface{}) (k interface{}, v interface{}) {
			u := val.(*domain.User)
			newU := domain.User{
				Id:   u.Id,
				Name: u.Name,
			}
			return newU.Id, newU
		}).
		ToStore()

	u.resultStore = sproc.NewMemStore("result-store")

	sproc.NewStream(u.EventSteam, u.resultStore).
		Filter(func(val interface{}) bool {
			e := val.(*domain.Event)
			return e.Type == domain.PrEventType || e.Type == domain.CommitEventType
		}).
		Transform(func(val interface{}) (k interface{}, v interface{}) {
			e := val.(*domain.Event)
			o, _ := u.userStore.Get(e.ActorId)
			user := o.(domain.User)
			o, _ = u.resultStore.Get(e.ActorId)
			c := 0
			if o != nil {
				m := o.(domain.ResultMat)
				c = m.Count
			}
			c++

			return e.ActorId, domain.ResultMat{
				Id:    e.ActorId,
				Name:  user.Name,
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
