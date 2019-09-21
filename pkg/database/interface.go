package database

import "time"

type YukariStore struct {
	Id          string            `bson:"_id"`
	Data        []byte            `bson:"data"`
	ContentType string            `bson:"contentType"`
	CreatedOn   time.Time         `bson:"createdOn"`
	Receivers   []YukariReceivers `bson:"receivers"`
}

type YukariReceivers struct {
	Address string `bson:"address"`
	Done    bool   `bson:"done"`
}

type YukariTopic struct {
	ID        string                `bson:"_id"`
	Topic     string                `bson:"topic"`
	Receivers []YukariTopicReceiver `bson:"receivers"`
}

type YukariTopicReceiver struct {
	Address   string    `bson:"address"`
	CreatedOn time.Time `bson:"createdOn"`
}

type YukariDatabase interface {
	Get(id string) (*YukariStore, error)
	Save(store YukariStore) bool

	RegisterTopic(topic string, receivers []string) (*YukariTopic, error)
	GetTopic(topic string) (*YukariTopic, error)
}
