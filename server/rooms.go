package server

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	sync.RWMutex
	Map map[string][]Participant
}

func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

func (r *RoomMap) Get(roomID string) []Participant {
	r.Lock()
	defer r.Unlock()

	return r.Map[roomID]
}

func (r *RoomMap) CreateRoom() string {
	r.Lock()
	defer r.Unlock()

	roomID := uuid.New().String()
	r.Map[roomID] = []Participant{}

	return roomID
}

func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Lock()
	defer r.Unlock()

	p := Participant{host, conn}

	r.Map[roomID] = append(r.Map[roomID], p)
}

func (r *RoomMap) DeleteRoom(roomID string) {
	r.Lock()
	defer r.Unlock()

	delete(r.Map, roomID)
}
