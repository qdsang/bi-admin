package initialize

import (
	"bi-admin/global"
	"bi-admin/model"
)

// 注册数据库表专用
func DBTables() {
	db := global.GVA_DB
	// if !db.HasTable(&model.BiChart{}) {
	// 	db.CreateTable(&model.BiChart{})
	// }

	db.AutoMigrate(model.SysUser{},
		model.SysAuthority{},
		model.JwtBlacklist{},
		model.SysOperationRecord{},
		model.BiProject{},
		model.BiChart{},
		model.BiSource{},
		model.BiDashboard{},
		model.BiDashboardChart{},
		model.BiTask{},
		model.BiLab{},
	)

	global.GVA_LOG.Debug("register table success")
}
