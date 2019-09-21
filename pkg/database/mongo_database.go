package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	log "github.com/sirupsen/logrus"
)

type Mongo struct {
	Client *mongo.Client
	Database *mongo.Database
}

func (m Mongo) Get(id string) (*YukariStore, error) {
	return nil, nil
}

func (m Mongo) Save(store YukariStore) bool {
	return true
}

func (m Mongo) RegisterTopic(topic string, receivers []string) (*YukariTopic, error) {
	yukariTopic := YukariTopic{
		ID: primitive.NewObjectID().Hex(),
		Topic: topic,
	}

	topicReceivers := []YukariTopicReceiver{}
	if len(receivers) > 0 {
		for _, receiver := range receivers {
			topicReceiver := YukariTopicReceiver{
				Address:   receiver,
				CreatedOn: time.Now(),
			}

			topicReceivers = append(topicReceivers, topicReceiver)
		}
	}

	yukariTopic.Receivers = topicReceivers

	result, err := m.
		Database.
		Collection(TOPIC_COLLECTION).
		InsertOne(context.TODO(), yukariTopic)

	if err != nil {
		return nil, err
	}

	if result.InsertedID == nil {
		return nil, errors.New(NEW_TOPIC_INSERT_ERROR)
	} else {
		yukariTopic.ID = result.InsertedID.(string)
	}

	return &yukariTopic, nil
}

func (m Mongo) GetTopic(topic string) (*YukariTopic, error) {
	var yukariTopic YukariTopic
	err :=  m.Database.
		Collection(TOPIC_COLLECTION).
		FindOne(context.TODO(), bson.M{
			"topic": topic,
		}).
		Decode(&yukariTopic)

	if err != nil {
		return nil, err
	}
	return &yukariTopic, nil
}

func CreateMongoStore() YukariDatabase {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return nil
	}

	log.Info("Connected")

	database := client.Database(DATABASE)

	return Mongo{
		Client: client,
		Database: database,
	}
}