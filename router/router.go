package router

import (
	"FurbotServer-Go/controllers"
	"FurbotServer-Go/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoadRouter 路由
func LoadRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// 加载中间件
	g.Use(mw...)

	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	api := g.Group("/api")
	apiV2 := api.Group("/v2")
	{
		apiV2.GET("/image/:img", controllers.GetFursuitImage)
	}
	{
		getFursuit := apiV2.Group("")
		getFursuit.Use(middleware.VisitorAuth) // 鉴权
		getFursuit.GET("/getFursuitRand", controllers.GetFursuitRand)
		getFursuit.GET("/getFursuitByID", controllers.GetFursuitByID)
		getFursuit.GET("/getFursuitByName", controllers.GetFursuitByName)
	}
	{
		admin := apiV2.Group("")
		admin.Use(middleware.AdminAuth)

		admin.GET("/authList", controllers.GetAuthList)
		admin.PUT("/auth", controllers.AddAuth)
		admin.POST("/auth", controllers.FixAuth)
		admin.DELETE("/auth/:qq", controllers.DeleteAuth)

		admin.PUT("/fursuit", controllers.AddFursuit)
		admin.DELETE("/fursuit/:fid", controllers.DeleteFursuit)
	}

	return g
}
