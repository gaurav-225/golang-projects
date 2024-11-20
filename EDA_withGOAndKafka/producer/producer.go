package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {

	/*
	envs, err := godotenv.Read(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }

	FIBER_LISTEN_ADDRESS := envs["FIBER_LISTEN_ADDRESS"]
	*/

	godotenv.Load(".env")
	FIBER_LISTEN_ADDRESS := os.Getenv("FIBER_LISTEN_ADDRESS")
	app := fiber.New()

	api := app.Group("/api/v1")
	api.Post("/comments", createComment)
	app.Listen(FIBER_LISTEN_ADDRESS)
}

func ConnectProducer(brokersURL []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersURL, config)

	if err != nil {
		return nil, err
	}

	return conn, nil
}

func PushCommentToQueue(topic string, message []byte) error{

	KAFKA_ADDRESS := os.Getenv("KAFKA_ADDRESS")
	brokersURL := []string{KAFKA_ADDRESS}
	producer, err := ConnectProducer(brokersURL)

	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic : topic,
		Value : sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Msg is stored in topic(%s)/parition(%d)/offset(%d)", topic, partition, offset)

	return nil

}

func createComment(c *fiber.Ctx) error {
	cmt := new(Comment)

	if err := c.BodyParser(cmt); err != nil {
		log.Println(err)
		c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})

		return err
	}

	cmtInBytes, err := json.Marshal(cmt)
	PushCommentToQueue("comments", cmtInBytes)

	err = c.JSON(&fiber.Map{
		"success": true,
		"message": "Comment pushed sucssfully",
		"comment": cmt,
	})

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"sucess":  false,
			"msssage": "Error creating product",
		})

		return err
	}

	return err

}
