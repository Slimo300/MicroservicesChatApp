package events

import (
	"github.com/google/uuid"
)

type MemberCreatedEvent struct {
	ID               uuid.UUID `json:"ID" mapstructure:"ID"`
	GroupID          uuid.UUID `json:"groupID" mapstructure:"groupID"`
	UserID           uuid.UUID `json:"userID" mapstructure:"userID"`
	User             User      `json:"User" mapstructure:"User"`
	Adding           bool      `json:"adding" mapstructure:"adding"`
	DeletingMembers  bool      `json:"deletingMembers" mapstructure:"deletingMembers"`
	DeletingMessages bool      `json:"deletingMessages" mapstructure:"deletingMessages"`
	Admin            bool      `json:"admin" mapstructure:"admin"`
	Creator          bool      `json:"creator" mapstructure:"creator"`
}

func (MemberCreatedEvent) EventName() string {
	return "groups.membercreated"
}
