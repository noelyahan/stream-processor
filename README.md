## Stream processor for analytics

### Build
    `go build cmd/main.go`

### Usage

```
Usage of ./main:
  -analyze string
        [-analyze users] (only users and repos supported) (default "users")
  -data_dir string
        [-data_dir ../data] (default "./data")
  -filters string
        [-filter commit] (only commit and watch supported) (default "commit")

```

- Top 10 active users sorted by amount of PRs created and commits pushed
  
    `./main -analyze users -filters commit -data_dir ./data`
  

- Top 10 repositories sorted by amount of commits pushed

  `./main -analyze repos -filters commit -data_dir ./data`
  

- Top 10 repositories sorted by amount of watch events

  `./main -analyze repos -filters watch -data_dir ./data`
  