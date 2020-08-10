package v2

import (
	"bi-admin/global/response"
	"bi-admin/model/request"
	"bi-admin/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name    string `json:"name" form:"name"`
	Message string `json:"message" form:"message"`
	Nick    string `json:"nick" form:"nick"`
}

func SourceExesql(c *gin.Context) {
	var param request.BiSourceSql
	_ = c.ShouldBindJSON(&param)

	fmt.Println(param.Sql)

	err, list, _ := service.GetSourceSql(param.SourceID, param.Sql)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithData(list, c)
	}
}
