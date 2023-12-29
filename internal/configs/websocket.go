package configs

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/websocket/v2"
)

func HandleConnections(c *websocket.Conn) {
	defer c.Close()

	for {
		phoneNumbers, err := phoneNumberService.GetAll()
		if err != nil {
			log.Println("Error fetching phone numbers:", err)
			return
		}

		jsonData, err := json.Marshal(phoneNumbers)
		if err != nil {
			log.Println("Error marshaling phone numbers to JSON:", err)
			return
		}

		err = c.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			log.Println("Error sending data to client:", err)
			return
		}

		time.Sleep(1 * time.Second)
	}
}
