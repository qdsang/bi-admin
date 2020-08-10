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

func TaskList(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	var pageInfo request.BiTaskListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	err, list, total := service.GetTaskList(pageInfo)
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

func TaskCreate(c *gin.Context) {
	var mItem model.BiTask
	_ = c.ShouldBindJSON(&mItem)
	err := service.TaskCreate(&mItem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkDetailed(mItem, "创建成功", c)

	}
}

func TaskUpdate(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiTaskUpdate
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiTask{}
	item.TaskID = mItem.TaskID
	item.ProjectID = mItem.ProjectID
	item.Type = mItem.Type
	item.Name = mItem.Name
	item.Desc = mItem.Desc
	item.Status = mItem.Status

	if err := service.TaskUpdate(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}

func TaskDelete(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiTaskDelete
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiTask{}
	item.TaskID = mItem.TaskID

	if err := service.TaskDelete(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}
