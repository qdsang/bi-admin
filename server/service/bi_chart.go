package service

import (
	"bi-admin/global"
	"bi-admin/model"
	"bi-admin/model/request"

	uuid "github.com/satori/go.uuid"
)

func GetChartList(info request.BiChartListSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BiChart{})
	var itemList []model.BiChart
	// 如果有条件搜索 下方会自动创建搜索语句
	// if info.Name != "" {
	// 	db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	// }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return err, itemList, total
}

func ChartCreate(item *model.BiChart) (err error) {
	// if (!global.GVA_DB.First(&model.BiSource{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	item.ChartID = uuid.NewV4()
	err = global.GVA_DB.Create(&item).Error
	return err
}

func GetChartById(chartID uuid.UUID) (err error, chart model.BiChart) {
	err = global.GVA_DB.Where("chart_id = ?", chartID).First(&chart).Error
	return
}

func ChartUpdateByID(item *model.BiChart) (err error) {
	// if (!global.GVA_DB.First(&model.BiSource{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	err = global.GVA_DB.Where("chart_id = ?", item.ChartID).First(&model.BiChart{}).Updates(&item).Error
	return err
}

func ChartDelete(item *model.BiChart) (err error) {
	err = global.GVA_DB.Delete(&[]model.BiChart{}, "chart_id = ?", item.ChartID).Error
	return err
}
