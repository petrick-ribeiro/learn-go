package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Message struct {
  Greeting string `json:"greeting"`
}

var (
  wsUpgrader = websocket.Upgrader {
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
  }

  wsConn *websocket.Conn
)

func WsEndpoint(w http.ResponseWriter, r *http.Request)  {
  wsUpgrader.CheckOrigin = func(r *http.Request) bool {
    // check http.Request
    return true
  }

  var err error

  wsConn, err = wsUpgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Printf("could not upgrade: %s\n", err.Error())
    return
  }

  defer wsConn.Close()

  // event loop
  for {
    var msg Message

    err := wsConn.ReadJSON(&msg)
    if err != nil {
      fmt.Printf("error reading JSON: %s\n", err.Error())
      break
    }

    fmt.Printf("Message Received: %s\n", msg.Greeting)
    if strings.ToLower(msg.Greeting) == "ping" {
      SendMessage("Pong!")
    } else {
      SendMessage("Hello, client!")
    }
  }
}

func SendMessage(msg string)  {
  err := wsConn.WriteMessage(websocket.TextMessage, []byte(msg))
  if err != nil {
    fmt.Printf("error sending message: %s\n", err.Error())
  }
}

func main()  {
  router := mux.NewRouter()

  router.HandleFunc("/socket", WsEndpoint)

  log.Fatal(http.ListenAndServe(":9100", router))
}
