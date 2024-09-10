package database

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

var MQ *amqp.Connection

func ConnectRabbitMQ() bool {
	var err error
	MQ, err = amqp.Dial(os.Getenv("RABBIT_URL"))

	if err != nil {
		return false
	}

	// Reconnecting
	go func() {

		log.Info("RabbitMQ baglantisi koptu: %s", <-MQ.NotifyClose(make(chan *amqp.Error)))

		MQ = nil
		ticker := time.NewTicker(1 * time.Minute)
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:

					log.Info("RabbitMQ tekrar baglanti deneniyor...")

					stopTimer := ConnectRabbitMQ()

					if stopTimer {
						close(quit)
					}

					// do stuff
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()

	}()

	return true

}
