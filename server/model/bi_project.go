package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type BiProject struct {
	ID         uint       `json:"-" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-" sql:"index"`
	ProjectID  uuid.UUID  `json:"project_id" gorm:"" sql:"index"`
	Type       string     `json:"type" gorm:"default:'';" `
	Name       string     `json:"name" gorm:"default:'';" `
	Desc       string     `json:"desc" gorm:"default:'';" `
	Status     int        `json:"status" gorm:"default:0;" `
	CreatorUID string     `json:"creator_uid"  gorm:"default:'';"`
}
