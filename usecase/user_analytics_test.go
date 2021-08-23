package usecase

import (
	"log"
	"os"
	"testing"

	sproc "github.com/noelyahan/stream-processor"
	"github.com/noelyahan/stream-processor/domain"
)

func TestUserAnalytics_ActiveUsersByCommitPRs(t *testing.T) {
	f1, err := os.Open("../data/actors.csv")
	if err != nil {
		log.Fatal(err)
	}
	userStream := sproc.NewCSVReader(f1, new(domain.User))

	f2, err := os.Open("../data/events.csv")
	if err != nil {
		log.Fatal(err)
	}
	eventStream := sproc.NewCSVReader(f2, new(domain.Event))
	usecase := UserAnalytics{
		UserStream: userStream,
		EventSteam: eventStream,
	}
	usecase.ActiveUsersByCommitPRs()
}
