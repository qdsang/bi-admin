package service

import (
	"bi-admin/global"
	"bi-admin/model"
	"bi-admin/model/request"

	uuid "github.com/satori/go.uuid"
)

func GetTaskList(info request.BiTaskListSearch) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BiTask{})
	var itemList []model.BiTask
	// 如果有条件搜索 下方会自动创建搜索语句
	// if info.Name != "" {
	// 	db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	// }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&itemList).Error
	return err, itemList, total
}

func TaskCreate(item *model.BiTask) (err error) {
	// if (!global.GVA_DB.First(&model.BiTask{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	item.TaskID = uuid.NewV4()
	err = global.GVA_DB.Create(&item).Error
	return err
}

func TaskUpdate(item *model.BiTask) (err error) {
	// if (!global.GVA_DB.First(&model.BiTask{}, "type = ?", item.Type).RecordNotFound()) {
	// 	return errors.New("存在相同的type，不允许创建")
	// }
	err = global.GVA_DB.Where("task_id = ?", item.TaskID).First(&model.BiTask{}).Updates(&item).Error
	return err
}

func TaskDelete(item *model.BiTask) (err error) {
	err = global.GVA_DB.Delete(&[]model.BiTask{}, "task_id = ?", item.TaskID).Error
	return err
}

func GetTaskById(taskID string) (err error, item model.BiTask) {
	err = global.GVA_DB.Where("task_id = ?", taskID).First(&item).Error
	return
}
