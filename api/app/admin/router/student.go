package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis/student"
	"go-admin/app/admin/middleware"
	jwt "go-admin/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerStudentRouter)
}

// 需认证的路由代码
func registerStudentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	r := v1.Group("/student").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.POST("/add", student.InsertStudent)
		r.PUT("/update", student.UpdateStudent)
		r.GET("/getPage", student.GetPage)
		r.PUT("/deleteStudent", student.DeleteStudent)
	}
}
