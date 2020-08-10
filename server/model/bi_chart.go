package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type BiChart struct {
	ID         uint       `json:"-" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-" sql:"index"`
	ChartID    uuid.UUID  `json:"chart_id" sql:"index"`
	ProjectID  string     `json:"project_id" gorm:"default:'';"`
	SourceID   string     `json:"source_id" gorm:"default:'';"`
	Content    string     `json:"content" gorm:"type:text;default:'';"`
	Name       string     `json:"chart_name" gorm:"default:'';" `
	Desc       string     `json:"desc" gorm:"default:'';" `
	Status     int        `json:"status" gorm:"default:0;" `
	CreatorUID string     `json:"creator_uid" gorm:"default:'';"`
}
