package usecase

import (
	"github.com/noelyahan/stream-processor/domain"
	"testing"
)

func TestRepoAnalytics_TopRepositoriesByCommit(t *testing.T) {

	repoStream := mockReader{}
	repoStream.initRepos()

	eventStream := mockReader{}
	eventStream.initEvents(domain.CommitEventType)

	usecase := RepoAnalytics{
		RepoStream: repoStream,
		EventSteam: eventStream,
	}
	usecase.TopRepositoriesByCommit()
}

func TestRepoAnalytics_TopRepositoriesByWatchEvent(t *testing.T) {
	repoStream := mockReader{}
	repoStream.initRepos()

	eventStream := mockReader{}
	eventStream.initEvents(domain.WatchEventType)

	usecase := RepoAnalytics{
		RepoStream: repoStream,
		EventSteam: eventStream,
	}
	usecase.TopRepositoriesByWatchEvent()
}
