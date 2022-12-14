package handlers

import (
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/auth"
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/msgqueue"
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/storage"
	"github.com/Slimo300/MicroservicesChatApp/backend/user-service/database"
	"github.com/Slimo300/MicroservicesChatApp/backend/user-service/email"
)

type Server struct {
	DB           database.DBLayer
	Emitter      msgqueue.EventEmitter
	ImageStorage storage.StorageLayer
	TokenService auth.TokenClient
	EmailService email.EmailService
	Domain       string
	MaxBodyBytes int64
}
