package talk

import (
	"net/http"
    "time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"mitty.co/mitty-server/app/models"
	"mitty.co/mitty-server/app/helpers"
)

var broadcast = make(chan Message)           // broadcast channel
var pubsub = PubSub {
		 topicsMap : make(map[string]map[*websocket.Conn]Client),
	     reverseTopicMap : make(map[*websocket.Conn]int),
}

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	MessageType string `json:"messageType"`
	Topic int `json:"topic"`
	Command string `json:"command"`
	Conversation models.Conversation `json:"Conversation"`
}

// Websocket Client
type Client struct {
	UserID int `json:"userId"`
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
   client := Client{UserID: user.ID,UserName: user.Name,Connected: true}
	
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
		
	    log.Println("http connection :来た")
        if (msg.Command == "subscribe") {
            pubsub.subscribe(ws, msg.Topic)
        }  else if (msg.Command == "talk") {
		    // Send the newly received message to the broadcast channel
		    pubsub.publish(msg)
		     tx, err := dbmap.Begin()
            if err != nil {
		        logrus.Printf("error: %v", err)
	        }
	        conversaton := msg.Conversation
  		    conversaton.SpeakerID = int64(client.UserID)
            if err := conversaton.Insert(*tx); err != nil {
		        logrus.Printf("error: %v", err)
		        tx.Rollback()
	        } else {
	    	    tx.Commit()
	        }
		}
	}
}

func MessageHandler() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
        pubsub.publish(msg)
	}
}

type PubSub struct {
      topicsMap  map[int]map[*websocket.Conn]bool
	  reverseTopicMap map[*websocket.Conn]int
}

func (pubsub *PubSub) subscribe(ws *websocket.Conn, meeting int) {
     clients,ok := pubsub.topicsMap[meeting] 	
     if !ok {
     	   log.Println("first client of meeting")
           clients := make(map[*websocket.Conn]bool)
           pubsub.topicsMap[meeting] = clients
           clients[ws] = true
           pubsub.reverseTopicMap[ws]=meeting
          return 
     } else {
     	   clients[ws] = true
           pubsub.reverseTopicMap[ws]=meeting
     }
}

func (pubsub *PubSub) publish(msg Message) {
	  clients,ok := pubsub.topicsMap[msg.MeetingId] 
	  if (!ok) {
	      return
	  }
	  
	  for client := range clients {
	  	    err := client.WriteJSON(msg)
	  	    if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
				if (len(clients) == 0) {
				    delete(pubsub.topicsMap, msg.MeetingId)
				}
			}
	  }
}

func (pubsub *PubSub) unsubscribe(ws *websocket.Conn) {
	meeting,ok := pubsub.reverseTopicMap[ws]
	if !ok {
		return 
	}
	
	delete (pubsub.reverseTopicMap, ws)
	clients,ok := pubsub.topicsMap[meeting]
	if (!ok) {
	    return 
	}
	
	delete(clients,ws)
}

