package request

import "bi-admin/model"

type BiTaskListSearch struct {
	model.BiTask
	PageInfo
}

type BiTaskUpdate struct {
	model.BiTask
}

type BiTaskDelete struct {
	model.BiTask
}
