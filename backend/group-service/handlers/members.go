package handlers

import (
	"log"
	"net/http"

	"github.com/Slimo300/MicroservicesChatApp/backend/group-service/models"
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/apperrors"
	"github.com/Slimo300/MicroservicesChatApp/backend/lib/events"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GrantPriv(c *gin.Context) {
	userID := c.GetString("userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid ID"})
		return
	}
	groupID := c.Param("groupID")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid member ID"})
		return
	}
	memberID := c.Param("memberID")
	memberUUID, err := uuid.Parse(memberID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid member ID"})
		return
	}

	var rights models.MemberRights
	if err := c.ShouldBindJSON(&rights); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "bad request, all 3 fields must be present"})
		return
	}

	_, err = s.DB.GrantRights(userUUID, groupUUID, memberUUID, rights)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{"err": err.Error()})
	}

	// emit MemberUpdate

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (s *Server) DeleteUserFromGroup(c *gin.Context) {
	userID := c.GetString("userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid ID"})
		return
	}
	groupID := c.Param("groupID")
	groupUUID, err := uuid.Parse(groupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid member ID"})
		return
	}
	memberID := c.Param("memberID")
	memberUUID, err := uuid.Parse(memberID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid member ID"})
		return
	}

	if err := s.DB.DeleteMember(userUUID, groupUUID, memberUUID); err != nil {
		c.JSON(apperrors.Status(err), gin.H{"err": err.Error()})
		return
	}

	if err := s.Emitter.Emit(events.MemberDeletedEvent{ID: memberUUID}); err != nil {
		log.Printf("Emitter failed: %s", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"message": "member deleted"})
}
