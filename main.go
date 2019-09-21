package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
	yukari "yukari/pkg"
	"yukari/pkg/database"
)

func init() {
	// Init log
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	yukariDatabase := database.CreateMongoStore()

	log.Info("Running Yukari.")
	e := echo.New()

	// runner := yukari.CreateYukari()
	topicHandler := yukari.NewYukariTopicHandler(yukariDatabase)

	// Paths
	runnerGroup := e.Group("/run")
	runnerGroup.Use(middleware.AddTrailingSlash())
	runnerGroup.GET("/", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "text/plain", []byte("hello"))
	})

	topicGroup := e.Group("/topics")
	topicGroup.Use(middleware.AddTrailingSlash())
	topicGroup.POST("/", func (c echo.Context) error {
		var createTopicRequest yukari.CreateTopicRequest
		err := c.Bind(&createTopicRequest)

		if err != nil {
			return ReturnError(c, err)
		}

		topic, err := topicHandler.CreateTopic(createTopicRequest.Name)

		if err != nil {
			return ReturnError(c, err)
		}

		return c.JSON(http.StatusOK, topic)
	})

	err := e.Start("0.0.0.0:5555")
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Info("Running on 0.0.0.0:5555")
	}
}

func ReturnError(c echo.Context, err error) error {
	response := yukari.ErrorResponse{Error: err.Error()}
	return c.JSON(http.StatusBadRequest, response)
}
