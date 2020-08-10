package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type BiTask struct {
	ID         uint       `json:"-" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-" sql:"index"`
	TaskID     uuid.UUID  `json:"task_id" gorm:"" sql:"index"`
	ProjectID  string     `json:"project_id"  gorm:"default:'';"`
	Type       string     `json:"type" gorm:"default:'';" `
	Name       string     `json:"name" gorm:"default:'';" `
	Desc       string     `json:"desc" gorm:"default:'';" `
	Cron       string     `json:"cron" gorm:"default:'';" `
	Status     int        `json:"status" gorm:"default:0;" `
	CreatorUID string     `json:"creator_uid"  gorm:"default:'';"`
}
