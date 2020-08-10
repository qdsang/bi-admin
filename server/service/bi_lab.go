package service

import (
	"bi-admin/global"
	"bi-admin/model"
	"bi-admin/model/request"

	uuid "github.com/satori/go.uuid"
)

func GetLabList(info request.BiLabListSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BiLab{})
	var itemList []model.BiLab
	db = db.Where("`user_id` = ?", info.UserID)
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return err, itemList, total
}

func LabCreate(item *model.BiLab) (err error) {
	item.LabID = uuid.NewV4()
	err = global.GVA_DB.Create(&item).Error
	return err
}

func GetLabById(labID string) (err error, chart model.BiLab) {
	err = global.GVA_DB.Where("lab_id = ?", labID).First(&chart).Error
	return
}

func LabUpdateByID(item *model.BiLab) (err error) {
	err = global.GVA_DB.Where("lab_id = ?", item.LabID).First(&model.BiLab{}).Updates(&item).Error
	return
}

func LabDelete(item *model.BiLab) (err error) {
	err = global.GVA_DB.Delete(&[]model.BiLab{}, "lab_id = ?", item.LabID).Error
	return err
}
