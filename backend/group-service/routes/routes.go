package routes

import (
	"github.com/Slimo300/MicroservicesChatApp/backend/group-service/handlers"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine, server *handlers.Server) {

	engine.Use(CORSMiddleware())

	api := engine.Group("/api")
	api.Use(limits.RequestSizeLimiter(server.MaxBodyBytes))
	apiAuth := api.Use(server.MustAuth())

	apiAuth.GET("/group", server.GetUserGroups)
	apiAuth.POST("/group", server.CreateGroup)
	apiAuth.DELETE("/group/:groupID", server.DeleteGroup)

	apiAuth.POST("/group/:groupID/image", server.SetGroupProfilePicture)
	apiAuth.DELETE("/group/:groupID/image", server.DeleteGroupProfilePicture)

	apiAuth.DELETE("/group/:groupID/member/:memberID", server.DeleteUserFromGroup)
	apiAuth.PATCH("/group/:groupID/member/:memberID", server.GrantPriv)

	apiAuth.GET("/invites", server.GetUserInvites)
	apiAuth.POST("/invites", server.CreateInvite)
	apiAuth.PUT("/invites/:inviteID", server.RespondGroupInvite)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
