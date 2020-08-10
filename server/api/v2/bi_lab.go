package v2

import (
	"bi-admin/global/response"
	"bi-admin/model"
	"bi-admin/model/request"
	resp "bi-admin/model/response"
	"bi-admin/service"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func LabList(c *gin.Context) {
	var pageInfo request.BiLabListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID
	pageInfo.UserID = waitUse.UUID

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	err, list, total := service.GetLabList(pageInfo)
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

func LabCreate(c *gin.Context) {
	var mItem request.BiLabUpdate
	_ = c.ShouldBindJSON(&mItem)

	// content, err := json.Marshal(mItem.Content)
	// if err != nil {
	// 	fmt.Println("json", err)
	// }
	// idn, _ := strconv.Atoi(mItem.ID)

	info := model.BiLab{}

	claims, _ := c.Get("claims")
	waitUse := claims.(*request.CustomClaims)
	info.UserID = waitUse.UUID

	info.Name = mItem.Name
	info.Desc = mItem.Desc

	if err := service.LabCreate(&info); err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkDetailed(info, "成功", c)
	}
}

func LabUpdate(c *gin.Context) {
	var mItem request.BiLabUpdate
	_ = c.ShouldBindJSON(&mItem)

	content, err := json.Marshal(mItem.Content)
	if err != nil {
		fmt.Println("json", err)
	}

	dashboard := model.BiLab{}
	dashboard.LabID = mItem.LabID
	dashboard.SourceID = mItem.SourceID
	dashboard.Name = mItem.Name
	dashboard.Desc = mItem.Desc
	dashboard.Sql = mItem.Sql
	dashboard.Content = string(content)
	fmt.Println(dashboard)
	if err := service.LabUpdateByID(&dashboard); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(dashboard, "成功", c)
	}
}

func LabDelete(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiLabDelete
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiLab{}
	item.LabID = mItem.LabID

	if err := service.LabDelete(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}
