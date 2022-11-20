package handlers

import (
	"net/http"

	"github.com/Slimo300/MicroservicesChatApp/backend/lib/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GetUserInvites(c *gin.Context) {

	userID := c.GetString("userID")
	userUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "Invalid ID"})
		return
	}

	invites, err := s.DB.GetUserInvites(userUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	if len(invites) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, invites)
}

func (s *Server) CreateInvite(c *gin.Context) {
	userID := c.GetString("userID")
	userUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid ID"})
		return
	}

	load := struct {
		GroupID string `json:"group"`
		Target  string `json:"target"`
	}{}

	// getting req body
	if err := c.ShouldBindJSON(&load); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	groupUID, err := uuid.Parse(load.GroupID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid group ID"})
		return
	}
	targetUUID, err := uuid.Parse(load.Target)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid target user ID"})
		return
	}

	_, err = s.DB.AddInvite(userUID, targetUUID, groupUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": "internal database error"})
		return
	}

	// s.actionChan <- &communication.Action{Invite: invite}

	c.JSON(http.StatusCreated, gin.H{"message": "invite sent"})
}

func (s *Server) RespondGroupInvite(c *gin.Context) {
	userID := c.GetString("userID")
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid ID"})
		return
	}
	inviteID := c.Param("inviteID")
	inviteUUID, err := uuid.Parse(inviteID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "invalid invite id"})
		return
	}

	load := struct {
		Answer *bool `json:"answer" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&load); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "answer not specified"})
		return
	}

	invite, group, member, err := s.DB.AnswerInvite(userUUID, inviteUUID, *load.Answer)
	if err != nil {
		c.JSON(apperrors.Status(err), gin.H{"err": err.Error})
		return
	}

	if member != nil {
		// Emit NewMember
	}
	if invite != nil {
		// Emit InviteUpdate
	}

	c.JSON(http.StatusOK, group)
}
