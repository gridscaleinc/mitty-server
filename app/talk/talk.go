package talk

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"mitty.co/mitty-server/app/models"
)

var clients = make(map[*websocket.Conn]Client) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	MeetingID int64 `json:"meetingId"`
	ReplyToID int64 `json:"replyToId"`
	Speaking  string `json:"speaking"`
	SpeakerID String `json:"speakerId"`
	SpeakTime String `json:"speakTime"`
}

// Websocket Client
type Client struct {
	UserID string `json:"userId"`
	UserName string `json:userName`
	Connected bool `json:connected`
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("X-Mitty-AccessToken")
	user, err := models.GetUserByAccessToken(accessToken)
	if err != nil || user == nil {
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
		return
	}
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
   client := Client{
		UserID: user.ID
		UserName: user.Name
		Connected: true
	}
	
	clients[ws] = client
	logrus.Printf("WebsocketHandler Start handling new client.")
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			logrus.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		msg.SpeakerID = client.UserID
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func MessageHandler() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		logrus.Printf("New message:%v", msg)
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				logrus.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
