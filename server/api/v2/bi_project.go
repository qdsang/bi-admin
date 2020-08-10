package v2

import (
	"bi-admin/global/response"
	"bi-admin/model"
	"bi-admin/model/request"
	resp "bi-admin/model/response"
	"bi-admin/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ProjectList(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	var pageInfo request.BiProjectListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	err, list, total := service.GetProjectList(pageInfo)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, c)
	}
	// response.OkDetailed(resp.SysUserInfoResponse{UserName: "" + hex.EncodeToString(uuid.Bytes())}, "", c)
}

func ProjectCreate(c *gin.Context) {
	var mItem model.BiProject
	_ = c.ShouldBindJSON(&mItem)
	err := service.ProjectCreate(&mItem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkDetailed(mItem, "创建成功", c)

	}
}

func ProjectUpdate(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiProjectUpdate
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiProject{}
	item.ProjectID = mItem.ProjectID
	item.Name = mItem.Name
	item.Desc = mItem.Desc
	item.Type = mItem.Type
	item.Status = mItem.Status

	if err := service.ProjectUpdate(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}

func ProjectDelete(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiProjectDelete
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiProject{}
	item.ProjectID = mItem.ProjectID

	if err := service.ProjectDelete(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}
