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

func ChartList(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	var pageInfo request.BiChartListSearch
	_ = c.ShouldBindQuery(&pageInfo)

	pageInfo.Page = 0
	pageInfo.PageSize = 100

	err, list, total := service.GetChartList(pageInfo)
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

func ChartGetById(c *gin.Context) {
	// claims, _ := c.Get("claims")
	// 获取头像文件
	// 这里我们通过断言获取 claims内的所有内容
	// waitUse := claims.(*request.CustomClaims)
	// uuid := waitUse.UUID

	cid, _ := uuid.FromString(c.Param("chartID"))

	err, item := service.GetChartById(cid)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
	// response.OkDetailed(resp.SysUserInfoResponse{UserName: "" + hex.EncodeToString(uuid.Bytes())}, "", c)
}

func ChartCreate(c *gin.Context) {
	var mItem request.BiChartUpdate
	_ = c.ShouldBindJSON(&mItem)

	content, err := json.Marshal(mItem.Content)
	if err != nil {
		fmt.Println("json", err)
	}

	chart := model.BiChart{}
	chart.SourceID = mItem.SourceID
	chart.Name = mItem.ChartName
	chart.Desc = mItem.ChartDesc
	chart.Content = string(content)

	if err := service.ChartCreate(&chart); err != nil {
		response.FailWithMessage(fmt.Sprintf("创建失败，%v", err), c)
	} else {
		response.OkDetailed(chart, "成功", c)
	}
}

func ChartUpdate(c *gin.Context) {
	var mItem request.BiChartUpdate
	_ = c.ShouldBindJSON(&mItem)

	content, err := json.Marshal(mItem.Content)
	if err != nil {
		fmt.Println("json", err)
	}
	// cidn, _ := strconv.Atoi(mItem.ID)

	chart := model.BiChart{}
	// chart.ChartID, _ = uuid.FromString(mItem.ChartID)
	chart.ChartID = mItem.ChartID
	chart.SourceID = mItem.SourceID
	chart.Name = mItem.ChartName
	chart.Desc = mItem.ChartDesc
	chart.Content = string(content)

	if err := service.ChartUpdateByID(&chart); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(chart, "成功", c)
	}
}

func ChartDelete(c *gin.Context) {
	var mItem request.BiChartDelete
	_ = c.ShouldBindJSON(&mItem)

	item := model.BiChart{}
	item.ChartID = mItem.ChartID

	if err := service.ChartDelete(&item); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(item, "成功", c)
	}
}
