package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var AllRooms RoomMap

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	roomID := AllRooms.CreateRoom()

	type respond struct {
		RoomID string `json:"roomID"`
	}

	_ = json.NewEncoder(w).Encode(respond{RoomID: roomID})
}

var upgrade = websocket.Upgrader{
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type broadCastMsg struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

var broadcast = make(chan broadCastMsg)

func broadCaster() {
	for {
		msg := <-broadcast

		for _, clients := range AllRooms.Map[msg.RoomID] {
			if clients.Conn != msg.Client {
				err := clients.Conn.WriteJSON(msg.Message)

				if err != nil {
					slog.Error("failed to write message", err)
					_ = clients.Conn.Close()
				}
			}
		}
	}
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomIDParams, ok := r.URL.Query()["roomID"]

	if !ok {
		slog.Error("roomID not found in url params")
		return
	}

	roomID := roomIDParams[0]

	ws, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("failed to upgrade connection", err)
		return
	}

	AllRooms.InsertIntoRoom(roomID, false, ws)

	go broadCaster()

	for {
		var msg broadCastMsg

		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			slog.Error("read error", err)
		}

		msg.Client = ws
		msg.RoomID = roomID

		broadcast <- msg
	}
}
