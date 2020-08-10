package request

import (
	"bi-admin/model"

	uuid "github.com/satori/go.uuid"
)

type BiChartListSearch struct {
	model.BiChart
	PageInfo
}

type BiChartUpdate struct {
	ChartID   uuid.UUID   `json:"id"`
	SourceID  string      `json:"source_id" form:"source_id"`
	ChartName string      `json:"chart_name" form:"chart_name"`
	ChartDesc string      `json:"desc" form:"desc"`
	Content   interface{} `json:"content"`
}

type BiChartDelete struct {
	ChartID uuid.UUID `json:"chart_id"`
}
