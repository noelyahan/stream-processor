package usecase

import (
	"github.com/noelyahan/stream-processor/domain"
	"testing"
)

func TestUserAnalytics_ActiveUsersByCommitPRs(t *testing.T) {

	userStream := mockReader{}
	userStream.initUsers()

	eventStream := mockReader{}
	eventStream.initEvents(domain.CommitEventType)

	usecase := UserAnalytics{
		UserStream: userStream,
		EventSteam: eventStream,
	}
	usecase.ActiveUsersByCommitPRs()
}
