package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var port = flag.String("port", ":5000", "port of starting server")
var subscriptions []*webpush.Subscription

type serverResponse struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

type sendMessage struct {
	Message string `json:"message"`
}

func main() {
	flag.Parse()

	fmt.Printf("Starting web server in port %v.\n", *port)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading env file: %v \n", err)
	}

	r := gin.Default()

	r.GET("/save-webpush", func(c *gin.Context) {
		s := webpush.Subscription{}
		if err := c.ShouldBindJSON(&s); err != nil {
			c.JSON(500, serverResponse{
				Message: err.Error(),
				Status:  "ERROR",
				Code:    -1,
			})
			return
		}
		subscriptions = append(subscriptions, &s)
		c.JSON(200, serverResponse{
			Message: "Webpush token saved successfully.",
			Status:  "SUCCESS",
			Code:    1,
		})
	})

	r.GET("/send-webpush", func(c *gin.Context) {
		message := sendMessage{}
		if err := c.ShouldBindJSON(&message); err != nil {
			c.JSON(500, serverResponse{
				Message: err.Error(),
				Status:  "ERROR",
				Code:    -1,
			})
			return
		}
		go func() {
			for _, subscription := range subscriptions {
				resp, err := webpush.SendNotification([]byte(message.Message), subscription, &webpush.Options{
					VAPIDPublicKey:  os.Getenv("VAPID_PUBLIC_KEY"),
					VAPIDPrivateKey: os.Getenv("VAPID_PRIVATE_KEY"),
				})
				if err != nil {
					fmt.Printf("Error sending: %v", err)
				}
				defer resp.Body.Close()
			}
		}()
		c.JSON(200, serverResponse{
			Message: "Message sent to queue.",
			Status:  "PROCESSING",
			Code:    1,
		})
	})

	r.Run(*port)
}
