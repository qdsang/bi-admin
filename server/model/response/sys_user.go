package response

import (
	"bi-admin/model"
)

type SysUserResponse struct {
	User model.SysUser `json:"user"`
}

type SysUserInfoResponse struct {
	UserName string `json:"username"`
}
type LoginResponse struct {
	User      model.SysUser `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
