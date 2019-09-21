package yukari

import "yukari/pkg/database"

type Event interface {
	Announce() error
}

type Runner interface {
	Announce() error
}

type TopicHandler interface {
	CreateTopic(name string) (*database.YukariTopic, error)
	GetTopic(name string) (*database.YukariTopic, error)
}

type HttpCaller interface {
	GET()
	POST(url string, contentType string, data []byte) int
}