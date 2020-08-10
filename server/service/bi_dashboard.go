package service

import (
	"bi-admin/global"
	"bi-admin/model"
	"bi-admin/model/request"

	uuid "github.com/satori/go.uuid"
)

func GetDashboardList(info request.BiDashboardListSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BiDashboard{})
	var itemList []model.BiDashboard
	// 如果有条件搜索 下方会自动创建搜索语句
	// if info.Name != "" {
	// 	db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	// }
	// if info.Type != "" {
	// 	db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	// }
	// if info.Status != nil {
	// 	db = db.Where("`status` = ?", info.Status)
	// }
	// if info.Desc != "" {
	// 	db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	// }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return err, itemList, total
}

func DashboardCreate(item *model.BiDashboard) (err error) {
	item.DashboardID = uuid.NewV4()
	err = global.GVA_DB.Create(&item).Error
	return err
}

func GetDashboardById(dashboardID uuid.UUID) (err error, chart model.BiDashboard) {
	err = global.GVA_DB.Where("dashboard_id = ?", dashboardID).First(&chart).Error
	return
}

func DashboardUpdateByID(item *model.BiDashboard) (err error) {
	err = global.GVA_DB.Where("dashboard_id = ?", item.DashboardID).First(&model.BiDashboard{}).Updates(&item).Error
	return
}

func DashboardDelete(item *model.BiDashboard) (err error) {
	err = global.GVA_DB.Delete(&[]model.BiDashboard{}, "dashboard_id = ?", item.DashboardID).Error
	return err
}

func DashboardChartCreate(item *model.BiDashboardChart) (err error) {
	item.DashboardChartID = uuid.NewV4()
	err = global.GVA_DB.Create(&item).Error
	return err
}

func DashboardChartDelete(item *model.BiDashboardChart) (err error) {
	err = global.GVA_DB.Delete(&[]model.BiDashboardChart{}, "dashboard_id = ? and chart_id = ?", item.DashboardID, item.ChartID).Error
	return err
}

func ChartboardByDashboard(info request.BiChartboardByDashboard) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.BiChart{})
	var itemList []model.BiChart
	db = db.Joins("JOIN bi_dashboard_charts dc ON dc.chart_id = bi_charts.chart_id")
	db = db.Where("`dc`.`dashboard_id` = ?", info.DashboardID)
	db = db.Where("`dc`.`deleted_at` IS NULL")
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return err, itemList, total
}
