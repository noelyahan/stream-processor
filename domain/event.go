package domain

const PrEventType = "PullRequestEvent"
const CommitEventType = "PushEvent"
const WatchEventType = "WatchEvent"

type Event struct {
	Id      string `csv:"id"`
	Type    string `csv:"type"`
	ActorId string `csv:"actor_id"`
	RepoId  string `csv:"repo_id"`
}
