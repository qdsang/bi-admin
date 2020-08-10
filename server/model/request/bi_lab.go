package request

import (
	"bi-admin/model"

	uuid "github.com/satori/go.uuid"
)

type BiLabListSearch struct {
	model.BiLab
	PageInfo
}

type BiLabUpdate struct {
	LabID    uuid.UUID   `json:"lab_id"`
	UserID   uuid.UUID   `json:"user_id"`
	SourceID uuid.UUID   `json:"source_id"`
	Sql      string      `json:"sql"`
	Name     string      `json:"name"`
	Desc     string      `json:"desc"`
	Content  interface{} `json:"content"`
}

type BiLabDelete struct {
	LabID uuid.UUID `json:"lab_id"`
}
