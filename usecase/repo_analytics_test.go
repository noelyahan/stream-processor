package usecase

import (
	"log"
	"os"
	"testing"

	sproc "github.com/noelyahan/stream-processor"
	"github.com/noelyahan/stream-processor/domain"
)

func TestRepoAnalytics_TopRepositoriesByCommit(t *testing.T) {
	f1, err := os.Open("../data/repos.csv")
	if err != nil {
		log.Fatal(err)
	}
	repoStream := sproc.NewCSVReader(f1, new(domain.Repo))

	f2, err := os.Open("../data/events.csv")
	if err != nil {
		log.Fatal(err)
	}
	eventStream := sproc.NewCSVReader(f2, new(domain.Event))
	usecase := RepoAnalytics{
		RepoStream: repoStream,
		EventSteam: eventStream,
	}
	usecase.TopRepositoriesByCommit()
}

func TestRepoAnalytics_TopRepositoriesByWatchEvent(t *testing.T) {
	f1, err := os.Open("../data/repos.csv")
	if err != nil {
		log.Fatal(err)
	}
	repoStream := sproc.NewCSVReader(f1, new(domain.Repo))

	f2, err := os.Open("../data/events.csv")
	if err != nil {
		log.Fatal(err)
	}
	eventStream := sproc.NewCSVReader(f2, new(domain.Event))
	usecase := RepoAnalytics{
		RepoStream: repoStream,
		EventSteam: eventStream,
	}
	usecase.TopRepositoriesByWatchEvent()
}
