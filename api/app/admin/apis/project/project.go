package project

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/models"
	"go-admin/common/apis"
	"go-admin/tools"
	"go-admin/tools/app"
	"time"
)

type Project struct {
	apis.Api
}

type ViewStatus int

const (
	ViewStatusDelete ViewStatus = 2 // 软删除
	ViewStatusShow   ViewStatus = 1
)

/**
添加
*/
func Create(c *gin.Context) {
	var data models.Project
	err := c.Bind(&data)

	if data.Title == "" {
		app.Error(c, 200, err, "title 不能为空")
		return
	}

	data.CreateBy = tools.GetUserIdStr(c)
	data.CreatedAt = time.Now()
	tools.HasError(err, "", 500)
	result, err := data.Insert()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

/**
  软删除
*/
func SoftDelete(c *gin.Context) {
	var data models.Project
	deleteId := tools.IdsStrToIdsIntGroup("id", c)
	data.UpdatedAt = time.Now()
	data.UpdateBy = tools.GetUserIdStr(c)
	data.ViewStatus = 2
	data.ProjectId = deleteId[0]

	result, err := data.SoftDelete()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")
}

/**
分页获取
*/
func GetPage(c *gin.Context) {
	var data models.Project
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = tools.StringToInt(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = tools.StringToInt(index)
	}

	result, count, err := data.GetPage(pageSize, pageIndex)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")
}

func GetListAll(c *gin.Context) {
	var data models.Project
	t, _ := data.GetListAll()
	app.OK(c, t, "")
}
