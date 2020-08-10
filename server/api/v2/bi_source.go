package v2

import (
	"bi-admin/global/response"
	"bi-admin/model"
	"bi-admin/model/request"
	resp "bi-admin/model/response"
	"bi-admin/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SourceList(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	var pageInfo request.BiSourceListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	err, list, total := service.GetSourceList(pageInfo)
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

func SourceCreate(c *gin.Context) {
	var mItem model.BiSource
	_ = c.ShouldBindJSON(&mItem)
	err := service.SourceCreate(&mItem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkDetailed(mItem, "创建成功", c)

	}
}

func SourceUpdate(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiSourceUpdate
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiSource{}
	// item.SourceID, _ = uuid.FromString(mItem.SourceID)
	item.SourceID = mItem.SourceID
	item.ProjectID = mItem.ProjectID
	item.Type = mItem.Type
	item.BaseAlias = mItem.BaseAlias
	item.Database = mItem.Database
	item.Host = mItem.Host
	item.Username = mItem.Username
	item.Password = mItem.Password
	portNum, _ := strconv.Atoi(mItem.Port)
	item.Port = portNum

	fmt.Println(item.Port)

	if err := service.SourceUpdate(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}

func SourceDelete(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiSourceDelete
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiSource{}
	item.SourceID = mItem.SourceID

	if err := service.SourceDelete(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}

func SourceTables(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	var pageInfo request.BiSourceListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	sid := c.Param("sourceID")
	err, list, total := service.GetSourceTables(sid)
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
}

func SourceTablesLinked(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	var pageInfo request.BiSourceListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	sid := c.Param("sourceID")
	err, list, total := service.GetSourceTables(sid)
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
}
