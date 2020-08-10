package request

import "bi-admin/model"

type BiSourceListSearch struct {
	model.BiSource
	PageInfo
}

type BiSourceSql struct {
	SourceID string `json:"source_id" form:"source_id"`
	Sql      string `json:"sql" form:"sql"`
}

type BiSourceUpdate struct {
	model.BiSource
	Port string `json:"port" gorm:"default:0;" `
}

type BiSourceDelete struct {
	model.BiSource
}
