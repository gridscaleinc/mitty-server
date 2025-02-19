package talk

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"mitty.co/mitty-server/app/helpers"
	"mitty.co/mitty-server/app/models"
)

var broadcast = make(chan Message) // broadcast channel
var pubsub = PubSub{
	topicsMap:       make(map[string]map[*websocket.Conn]Client),
	reverseTopicMap: make(map[*websocket.Conn]string),
}

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Message ...
//  package message of various types
type Message struct {
	MessageType   string               `json:"messageType"`
	Topic         string               `json:"topic"`
	Command       string               `json:"command"`
	Sender        Client               `json:"sender"`
	Conversation  models.Conversation  `json:"conversation"`
	Teleportation models.Teleportation `json:"teleportation"`
}

// Client Websocket Client
type Client struct {
	UserID    int    `json:"userId"`
	UserIcon  string `json:"userIcon"`
	UserName  string `json:"userName"`
	Connected bool   `json:"connected"`
}

// WebsocketHandler ...
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
	client := Client{UserID: user.ID, UserIcon: user.Icon, UserName: user.Name, Connected: true}

	logrus.Printf("WebsocketHandler Start handling new client.")

	dbmap := helpers.GetPostgres()

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			logrus.Printf("error: %v", err)
			pubsub.unsubscribe(ws)
			break
		}

		logrus.Println("http connection :来た")
		if msg.Command == "subscribe" {
			pubsub.subscribe(ws, client, msg.Topic)
		} else if msg.Command == "talk" {
			// Send the newly received message to the broadcast channel
			msg.Conversation.SpeakerID = int64(client.UserID)
			msg.Sender = client
			pubsub.publish(msg)
			tx, err := dbmap.Begin()
			if err != nil {
				logrus.Printf("error: %v", err)
			}
			conversaton := msg.Conversation
			if err := conversaton.Insert(*tx); err != nil {
				logrus.Printf("error: %v", err)
				tx.Rollback()
			} else {
				tx.Commit()
			}
		} else if msg.Command == "teleport" {
			msg.Teleportation.MittyID = int(client.UserID)
			msg.Sender = client
			pubsub.publish(msg)
		}
	}
}

// MessageHandler ...
func MessageHandler() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		pubsub.publish(msg)
	}
}

type PubSub struct {
	topicsMap       map[string]map[*websocket.Conn]Client
	reverseTopicMap map[*websocket.Conn]string
}

func (pubsub *PubSub) subscribe(ws *websocket.Conn, client Client, topic string) {
	clients, ok := pubsub.topicsMap[topic]
	logrus.Println(topic)
	if !ok {
		logrus.Println("first client of meeting")
		clients := make(map[*websocket.Conn]Client)
		pubsub.topicsMap[topic] = clients
		clients[ws] = client
		pubsub.reverseTopicMap[ws] = topic
		return
	} else {
		clients[ws] = client
		pubsub.reverseTopicMap[ws] = topic
	}
}

func (pubsub *PubSub) publish(msg Message) {
	logrus.Println(msg.Topic)
	clients, ok := pubsub.topicsMap[msg.Topic]
	if !ok {
		return
	}

	for websocket := range clients {
		err := websocket.WriteJSON(msg)
		if err != nil {
			logrus.Printf("error: %v", err)
			websocket.Close()
			delete(clients, websocket)
			if len(clients) == 0 {
				delete(pubsub.topicsMap, msg.Topic)
			}
		}
	}
}

func (pubsub *PubSub) unsubscribe(ws *websocket.Conn) {
	topic, ok := pubsub.reverseTopicMap[ws]
	if !ok {
		return
	}

	delete(pubsub.reverseTopicMap, ws)
	clients, ok := pubsub.topicsMap[topic]
	if !ok {
		return
	}

	delete(clients, ws)
}

// SendMessage ...
func SendMessage(msg *Message) {
	broadcast <- *msg
}
