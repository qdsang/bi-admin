package v2

import (
	"bi-admin/global/response"
	"bi-admin/model"
	"bi-admin/model/request"
	resp "bi-admin/model/response"
	"bi-admin/service"
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func ChartboardMap(c *gin.Context) {
	var mItem request.BiChartBoardMapUpdate
	_ = c.ShouldBindJSON(&mItem)

	dchart := model.BiDashboardChart{}

	dchart.ChartID = mItem.ChartID
	dchart.DashboardID = mItem.DashboardID

	if err := service.DashboardChartCreate(&dchart); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(dchart, "成功", c)
	}

}

func ChartboardUnMap(c *gin.Context) {
	var mItem request.BiChartBoardMapUpdate
	_ = c.ShouldBindJSON(&mItem)

	dchart := model.BiDashboardChart{}

	dchart.ChartID = mItem.ChartID
	dchart.DashboardID = mItem.DashboardID

	if err := service.DashboardChartDelete(&dchart); err != nil {
		response.FailWithMessage(fmt.Sprintf("失败，%v", err), c)
	} else {
		response.OkDetailed(dchart, "成功", c)
	}

}

func ChartboardByDashboard(c *gin.Context) {
	// var mItem request.BiChartboardByDashboard
	// errq := c.ShouldBindQuery(&mItem)
	// if errq != nil {
	// 	fmt.Println("ShouldBindQuery", errq)
	// }

	var mItem request.BiChartboardByDashboard
	cid, _ := uuid.FromString(c.Query("dashboard_id"))
	mItem.DashboardID = cid
	mItem.PageSize = 1000
	fmt.Println(mItem)
	err, list, total := service.ChartboardByDashboard(mItem)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(resp.PageResult{
			List:     list,
			Total:    total,
			Page:     mItem.Page,
			PageSize: mItem.PageSize,
		}, c)
	}
}

func ChartboardBoardByChart(c *gin.Context) {
	response.FailWithMessage("获取数据失败", c)
}
