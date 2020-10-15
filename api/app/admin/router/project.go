package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/project"
	"go-admin/app/admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerProjectRouter)
}

// 需认证的路由代码
func registerProjectRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/project").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("", project.Create)
		r.DELETE("/:id", project.SoftDelete)
		r.GET("/getPage", project.GetPage)
		r.GET("/getListAll", project.GetListAll)
	}
}
