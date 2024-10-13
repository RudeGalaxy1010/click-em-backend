package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var (
	availableRooms []RoomData = make([]RoomData, 0)
)

type webSocketHandler struct {
	upgrader websocket.Upgrader
}

func (wsh webSocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := wsh.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Printf("error %s when upgrading connection to websocket", err)
		return
	}

	defer func() {
		log.Println("closing connection")
		c.Close()
	}()

	for {
		mt, message, err := c.ReadMessage()

		if err != nil {
			log.Printf("Error %s when reading message from client", err)
			return
		}

		if mt == websocket.BinaryMessage {
			err = c.WriteMessage(websocket.TextMessage, []byte("server doesn't support binary messages"))

			if err != nil {
				log.Printf("Error %s when sending message to client", err)
			}

			return
		}

		decoder := json.NewDecoder(strings.NewReader(string(message)))
		i := 1

		for {
			response := getResponse(decoder)
			err = c.WriteMessage(websocket.TextMessage, []byte(response))

			if err != nil {
				log.Printf("Error %s when sending message to client", err)
				return
			}

			i = i + 1
			time.Sleep(2 * time.Second)
		}
	}
}

func getResponse(decoder *json.Decoder) string {
	command, err := decoder.Token()

	if err != nil {
		log.Println(err)
		return err.Error()
	}

	encoder := json.NewEncoder()

	switch command {
	case "CreateRoom":
		break
	case "AvailableRooms":
	case "ConnectToRoom":
		break
	case "UserInfo":
		break
	case "GameOver":
		break
	default:
		return string(fmt.Sprintf("Unknown command: %s", command))
	}
}

func main() {
	webSocketHandler := webSocketHandler{
		upgrader: websocket.Upgrader{},
	}

	http.Handle("/", webSocketHandler)
	log.Print("Starting server...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
