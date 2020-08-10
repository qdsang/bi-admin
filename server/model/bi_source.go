package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// base_alias: "13231231"
// created_at: "2020-07-29T06:25:17.546Z"
// creator: 1
// database: "test"
// host: "localhost"
// is_private: true
// password: null
// port: 3306
// source_id: "10f18b13-954d-4730-8168-29ed3ffc3b36"
// status: 1
// updated_at: "2020-07-29T06:25:17.540Z"
// username: "root"

type BiSource struct {
	ID         uint       `json:"-" gorm:"primary_key"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-" sql:"index"`
	SourceID   uuid.UUID  `json:"source_id" gorm:"" sql:"index"`
	ProjectID  string     `json:"project_id"  gorm:"default:'';"`
	Type       string     `json:"type" gorm:"default:'';" `
	Name       string     `json:"name" gorm:"default:'';" `
	BaseAlias  string     `json:"base_alias" gorm:"default:'';" `
	Database   string     `json:"database" gorm:"default:'';"`
	Host       string     `json:"host" gorm:"default:'';"`
	Username   string     `json:"username" gorm:"default:'';" `
	Password   string     `json:"password" gorm:"default:'';" `
	Port       int        `json:"port" gorm:"default:0;" `
	Status     int        `json:"status" gorm:"default:0;" `
	CreatorUID string     `json:"creator_uid"  gorm:"default:'';"`
}
