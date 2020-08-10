package request

import (
	"bi-admin/model"

	uuid "github.com/satori/go.uuid"
)

type BiDashboardListSearch struct {
	model.BiDashboard
	PageInfo
}

type BiDashboardUpdate struct {
	DashboardID uuid.UUID   `json:"dashboard_id"`
	Name        string      `json:"name" form:"name"`
	Desc        string      `json:"desc" form:"desc"`
	Content     interface{} `json:"content"`
}

type BiDashboardDelete struct {
	DashboardID uuid.UUID `json:"dashboard_id"`
}

type BiChartBoardMapUpdate struct {
	DashboardID uuid.UUID `json:"dashboard_id" `
	ChartID     uuid.UUID `json:"chart_id" `
}

type BiChartboardByDashboard struct {
	PageInfo
	DashboardID uuid.UUID `json:"dashboard_id" form:"dashboard_id"`
}
