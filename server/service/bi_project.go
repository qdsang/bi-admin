package service

import (
	"bi-admin/global"
	"bi-admin/model"
	"bi-admin/model/request"

	uuid "github.com/satori/go.uuid"
)

func GetProjectList(info request.BiProjectListSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BiProject{})
	var itemList []model.BiProject
	// 如果有条件搜索 下方会自动创建搜索语句
	// if info.Name != "" {
	// 	db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	// }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return err, itemList, total
}

func ProjectCreate(item *model.BiProject) (err error) {
	// if (!global.GVA_DB.First(&model.BiProject{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	item.ProjectID = uuid.NewV4()
	err = global.GVA_DB.Create(&item).Error
	return err
}

func ProjectUpdate(item *model.BiProject) (err error) {
	// if (!global.GVA_DB.First(&model.BiProject{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	err = global.GVA_DB.Where("project_id = ?", item.ProjectID).First(&model.BiProject{}).Updates(&item).Error
	return err
}

func ProjectDelete(item *model.BiProject) (err error) {
	err = global.GVA_DB.Delete(&[]model.BiProject{}, "project_id = ?", item.ProjectID).Error
	return err
}

func GetProjectById(projectID string) (err error, item model.BiProject) {
	err = global.GVA_DB.Where("project_id = ?", projectID).First(&item).Error
	return
}
