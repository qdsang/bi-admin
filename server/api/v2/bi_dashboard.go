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
	uuid "github.com/satori/go.uuid"
)

func DashboardList(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	var pageInfo request.BiDashboardListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	err, list, total := service.GetDashboardList(pageInfo)
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

func DashboardCreate(c *gin.Context) {
	var mItem request.BiDashboardUpdate
	_ = c.ShouldBindJSON(&mItem)

	// content, err := json.Marshal(mItem.Content)
	// if err != nil {
	// 	fmt.Println("json", err)
	// }
	// idn, _ := strconv.Atoi(mItem.ID)

	dashboard := model.BiDashboard{}
	// dashboard.ID = uint(idn)
	dashboard.Name = mItem.Name
	dashboard.Desc = mItem.Desc

	if err := service.DashboardCreate(&dashboard); err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkDetailed(dashboard, "成功", c)
	}
}

func DashboardUpdate(c *gin.Context) {
	var mItem request.BiDashboardUpdate
	_ = c.ShouldBindJSON(&mItem)

	content, err := json.Marshal(mItem.Content)
	if err != nil {
		fmt.Println("json", err)
	}
	// idn, _ := strconv.Atoi(mItem.ID)

	dashboard := model.BiDashboard{}
	dashboard.DashboardID = mItem.DashboardID
	dashboard.Name = mItem.Name
	dashboard.Desc = mItem.Desc
	dashboard.Content = string(content)
	fmt.Println(dashboard)
	if err := service.DashboardUpdateByID(&dashboard); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(dashboard, "成功", c)
	}
}

func DashboardDelete(c *gin.Context) {
	// response.OkWithMessage("创建成功", c)
	var mItem request.BiDashboardDelete
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiDashboard{}
	item.DashboardID = mItem.DashboardID

	if err := service.DashboardDelete(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}

func DashboardGetById(c *gin.Context) {
	did, _ := uuid.FromString(c.Param("dashboardID"))

	err, item := service.GetDashboardById(did)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}
