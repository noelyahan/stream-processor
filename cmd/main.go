package main

import (
	"flag"
	"fmt"
	sproc "github.com/noelyahan/stream-processor"
	"github.com/noelyahan/stream-processor/domain"
	"github.com/noelyahan/stream-processor/usecase"
	"io"
	"log"
	"os"
)

func main() {
	supFlags := []string{"commit", "watch", "users", "repos"}
	analyzePtr := flag.String("analyze", "users", "[-analyze users] (only users and repos supported)")
	filtersPtr := flag.String("filters", "commit", `[-filter commit] (only commit and watch supported)`)
	dataLocPtr := flag.String("data_dir", "../data", `[-data_dir ../data]`)
	flag.Parse()

	invalidF1 := true
	invalidF2 := true
	for _, f := range supFlags {
		if f == *analyzePtr {
			invalidF1 = false
		}
		if f == *filtersPtr {
			invalidF2 = false
		}
	}
	if invalidF1 || invalidF2 {
		flag.Usage()
		os.Exit(0)
	}

	if *analyzePtr == "users" {
		userStream := sproc.NewCSVReader(loadFileReader(*dataLocPtr, "actors.csv"), new(domain.User))
		eventStream := sproc.NewCSVReader(loadFileReader(*dataLocPtr, "events.csv"), new(domain.Event))
		usecase := usecase.UserAnalytics{
			UserStream: userStream,
			EventSteam: eventStream,
		}
		usecase.ActiveUsersByCommitPRs()
		return
	}

	repoStream := sproc.NewCSVReader(loadFileReader(*dataLocPtr, "repos.csv"), new(domain.Repo))
	eventStream := sproc.NewCSVReader(loadFileReader(*dataLocPtr, "events.csv"), new(domain.Event))
	usecase := usecase.RepoAnalytics{
		RepoStream: repoStream,
		EventSteam: eventStream,
	}
	if *filtersPtr == "commit" {
		usecase.TopRepositoriesByCommit()
	} else if *filtersPtr == "watch" {
		usecase.TopRepositoriesByWatchEvent()
	}
}

func loadFileReader(dirPath, fName string) io.Reader {
	f, err := os.Open(fmt.Sprintf("%v/%v", dirPath, fName))
	if err != nil {
		log.Fatal(err)
	}
	return f
}
