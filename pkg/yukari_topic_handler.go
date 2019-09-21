package yukari

import "yukari/pkg/database"

type CreateTopicRequest struct {
	Name string `json:"name"`
}

type YukariTopicHandler struct {
	Store database.YukariDatabase
}

func NewYukariTopicHandler(store database.YukariDatabase) *YukariTopicHandler {
	return &YukariTopicHandler{Store: store}
}

func (y YukariTopicHandler) CreateTopic(name string) (*database.YukariTopic, error) {
	topic, err := y.Store.RegisterTopic(name, []string{})

	return topic, err
}
