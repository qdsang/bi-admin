package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type BiLab struct {
	ID        uint       `json:"-" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
	LabID     uuid.UUID  `json:"lab_id" gorm:"" sql:"index"`
	UserID    uuid.UUID  `json:"user_id"  gorm:"comment:'ID'"`
	SourceID  uuid.UUID  `json:"source_id"  gorm:"comment:''"`
	Content   string     `json:"content" gorm:"type:text;comment:''"`
	Sql       string     `json:"sql" gorm:"type:text;comment:''"`
	Name      string     `json:"name" gorm:"default:'';comment:'名称'" `
	Desc      string     `json:"desc" gorm:"default:'';comment:'描述'" `
	Status    int        `json:"status" gorm:"default:0;comment:'状态'" `
}
