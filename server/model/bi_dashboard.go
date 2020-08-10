package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type BiDashboard struct {
	ID          uint       `json:"-" gorm:"primary_key"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"-" sql:"index"`
	DashboardID uuid.UUID  `json:"dashboard_id" gorm:"" sql:"index"`
	ProjectID   string     `json:"project_id"  gorm:"comment:'项目ID'"`
	Content     string     `json:"content" gorm:"type:text;comment:''"`
	Name        string     `json:"name" gorm:"default:'';comment:'名称'" `
	Desc        string     `json:"desc" gorm:"default:'';comment:'描述'" `
	Status      int        `json:"status" gorm:"default:0;comment:'状态'" `
	CreatorUID  string     `json:"creator_uid" gorm:"comment:'创建人'"`
}

type BiDashboardChart struct {
	ID               uint       `json:"-" gorm:"primary_key"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"-" sql:"index"`
	DashboardChartID uuid.UUID  `json:"dashboard_chart_id" gorm:"" sql:"index"`
	ProjectID        uuid.UUID  `json:"project_id"  gorm:"comment:'项目ID'" sql:"index"`
	DashboardID      uuid.UUID  `json:"dashboard_id"  gorm:"comment:''" sql:"index"`
	ChartID          uuid.UUID  `json:"chart_id"  gorm:"comment:''" sql:"index"`
	Content          string     `json:"content" gorm:"type:text;comment:''"`
	Status           int        `json:"status" gorm:"default:0;comment:'状态'" `
	CreatorUID       string     `json:"creator_uid" gorm:"comment:'创建人'"`
}
