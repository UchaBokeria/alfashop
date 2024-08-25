package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
    conn *websocket.Conn
    mu   sync.Mutex
}

var mu sync.Mutex
var clients = make(map[*Client]struct{})

func main() {
    fmt.Println("🔥 " + "\033[36m" + "Hot Reload Has Started")
    http.HandleFunc("/", listen)
    http.HandleFunc("/trigger", trigger)
    log.Fatal(http.ListenAndServe(":9999", nil))
}

func listen(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("🔥 " + "\033[36m" + "We Got Loaded")
        return
    }

    client := &Client{conn: conn}

    mu.Lock()
    clients[client] = struct{}{}
    mu.Unlock()

    defer func() {
        mu.Lock()
        defer mu.Unlock()
        delete(clients, client)
        client.conn.Close()
    }()

    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("🔥 " + "\033[36m" + "We Got Loaded")
            return
        }
        log.Printf("Received message: %s\n", message)

        err = conn.WriteMessage(messageType, message)
        if err != nil {
            // fmt.Println("🔥 " + "\033[36m" + "We Are Loading")
            return
        }
    }
}

func trigger(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    for client := range clients {
        client.mu.Lock()
        defer client.mu.Unlock()
        err := client.conn.WriteMessage(websocket.TextMessage, []byte("reload"))
        if err != nil {
            fmt.Println("🔥 " + "\033[36m" + "Wait")
            return
        }
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Reloaded successfully"))
}