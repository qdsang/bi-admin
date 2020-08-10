package request

import "bi-admin/model"

type BiProjectListSearch struct {
	model.BiProject
	PageInfo
}

type BiProjectUpdate struct {
	model.BiProject
}

type BiProjectDelete struct {
	model.BiProject
}
