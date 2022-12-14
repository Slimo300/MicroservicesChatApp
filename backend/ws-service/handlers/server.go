package handlers

import (
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/auth"
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/events"
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/msgqueue"
	"github.com/Slimo300/MicroservicesChatApp/backend/ws-service/cache"
	"github.com/Slimo300/MicroservicesChatApp/backend/ws-service/database"
	"github.com/Slimo300/MicroservicesChatApp/backend/ws-service/ws"
)

type Server struct {
	DB           database.DBLayer
	Hub          *ws.WSHub
	TokenService auth.TokenClient
	CodeCache    cache.AccessCodeCache
	Emitter      msgqueue.EventEmitter
	Listener     msgqueue.EventListener
	MessageChan  <-chan *ws.Message
	EventChan    chan<- msgqueue.Event
}

func (s *Server) RunHub() {
	go s.ListenToHub()
	s.Hub.Run()
}

func (s *Server) ListenToHub() {
	var msg *ws.Message
	for {
		select {
		case msg = <-s.MessageChan:
			s.Emitter.Emit(events.MessageSentEvent{
				ID:      msg.ID,
				GroupID: msg.Group,
				UserID:  msg.User,
				Nick:    msg.Nick,
				Posted:  msg.When,
				Text:    msg.Message,
			})
		}
	}
}
