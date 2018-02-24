package workers

import (
	"fmt"
	"os"
	"strconv"

	"github.com/benmanns/goworker"
)

const (
	// DefaultConcurrency - how many workers running at once
	DefaultConcurrency = 2

	// DefaultURL - connection point to redis
	DefaultURL = "redis://localhost:6379/"

	// DefaultConnections - number of connections to redis instance
	DefaultConnections = 100

	// DefaultNamespace - namespace of the workers within redis instance
	DefaultNamespace = "resque:"

	// DefaultInterval - polling between checking for enqueued jobs
	DefaultInterval = 5.0
)

// InitWorkers - establishes worker settings and connection
func InitWorkers() {
	settings := goworker.WorkerSettings{
		URI:            setURI(),
		Connections:    setMaxConnections(),
		Queues:         []string{"default"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    setConcurrency(),
		Namespace:      setNamespace(),
		Interval:       DefaultInterval,
	}
	goworker.SetSettings(settings)
	goworker.Register("Default", myFunc)
}

func myFunc(queue string, args ...interface{}) error {
	fmt.Printf("From %s, %v\n", queue, args)
	return nil
}

func setURI() string {
	url := os.Getenv("WORKERS_URL")
	if url == "" {
		return DefaultURL
	}

	return url
}

func setMaxConnections() int {
	conns := os.Getenv("MAX_WORKERS_CONNECTIONS")

	num, err := strconv.Atoi(conns)
	if err != nil || conns == "" {
		return DefaultConnections
	}

	return num
}

func setConcurrency() int {
	concurrency := os.Getenv("WOKERS_CONCURRENCY")

	num, err := strconv.Atoi(concurrency)
	if err != nil || concurrency == "" {
		return DefaultConcurrency
	}

	return num
}

func setNamespace() string {
	namespace := os.Getenv("WORKERS_NAMESPACE")
	if namespace == "" {
		return DefaultNamespace
	}

	return namespace
}
