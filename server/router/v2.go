package router

import (
	v2 "bi-admin/api/v2"
	"bi-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitV2Router(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("v2")
	{
		UserRouter := ApiRouter.Group("user")
		{
			UserRouter.POST("signup", v2.Signup)
			UserRouter.POST("login", v2.Login2)
		}
		UserRouterAuth := UserRouter.Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			UserRouterAuth.GET("info", v2.UserInfo)
		}

		ProjectRouter := ApiRouter.Group("project").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			ProjectRouter.GET("list", v2.ProjectList)
			ProjectRouter.POST("create", v2.ProjectCreate)
			ProjectRouter.POST("update", v2.ProjectUpdate)
			ProjectRouter.POST("delete", v2.ProjectDelete)
		}

		DashboardRouter := ApiRouter.Group("dashboard").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			DashboardRouter.GET("list", v2.DashboardList)
			DashboardRouter.POST("create", v2.DashboardCreate)
			DashboardRouter.POST("update", v2.DashboardUpdate)
			DashboardRouter.POST("delete", v2.DashboardDelete)
			DashboardRouter.GET("get/:dashboardID", v2.DashboardGetById)
		}

		ChartboardRouter := ApiRouter.Group("chartboard").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			ChartboardRouter.POST("map", v2.ChartboardMap)
			ChartboardRouter.POST("unmap", v2.ChartboardUnMap)
		}
		ChartboardMapRouter := ApiRouter.Group("chartboardmap").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			ChartboardMapRouter.GET("chartbydashboard", v2.ChartboardByDashboard)
			ChartboardMapRouter.GET("boardbychart", v2.ChartboardBoardByChart)
		}

		ChartRouter := ApiRouter.Group("chart").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			ChartRouter.GET("list", v2.ChartList)
			ChartRouter.POST("create", v2.ChartCreate)
			ChartRouter.POST("update", v2.ChartUpdate)
			ChartRouter.POST("delete", v2.ChartDelete)
			ChartRouter.GET("get/:chartID", v2.ChartGetById)
		}

		SourceRouter := ApiRouter.Group("source").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			SourceRouter.GET("list", v2.SourceList)
			SourceRouter.POST("create", v2.SourceCreate)
			SourceRouter.POST("update", v2.SourceUpdate)
			SourceRouter.POST("delete", v2.SourceDelete)
			SourceRouter.GET("tables/:sourceID/linked", v2.SourceTablesLinked)
			SourceRouter.GET("tables/:sourceID", v2.SourceTables)
		}

		TaskRouter := ApiRouter.Group("task").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			TaskRouter.GET("list", v2.TaskList)
			TaskRouter.POST("create", v2.TaskCreate)
			TaskRouter.POST("update", v2.TaskUpdate)
			TaskRouter.POST("delete", v2.TaskDelete)
		}

		LabRouter := ApiRouter.Group("lab").Use(middleware.JWTAuth()) //.Use(middleware.CasbinHandler())
		{
			LabRouter.GET("list", v2.LabList)
			LabRouter.POST("create", v2.LabCreate)
			LabRouter.POST("update", v2.LabUpdate)
			LabRouter.POST("delete", v2.LabDelete)
		}

		ApiRouterAuth := ApiRouter.Use(middleware.JWTAuth())
		{
			ApiRouterAuth.POST("exesql", v2.SourceExesql)
		}

	}

}
