// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package database

import (
	"time"
)

type Campaign struct {
	ID        int32
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

type Objective struct {
	ID          int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description string
	IsActive    bool
	IsComplete  bool
	Number      int32
	QuestID     int32
}

type Quest struct {
	ID                   int32
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Name                 string
	Description          string
	Npc                  string
	IsComplete           bool
	IsActive             bool
	CompletedDescription string
	Number               int32
	CampaignID           int32
}
