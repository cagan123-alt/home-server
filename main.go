package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true 
    },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close() // defer run after function return

    for {
        messageType, p, err := ws.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        log.Printf("Received: %s\n", p)

        if err := ws.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleConnections)

    log.Println("http server started on :6234")
    err := http.ListenAndServe(":6234", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
