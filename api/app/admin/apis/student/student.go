package student

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/models"
	"go-admin/common/apis"
	"go-admin/tools"
	"go-admin/tools/app"
)

type student struct {
	apis.Api
}

func InsertStudent(c *gin.Context) {
	var data models.Student
	err := c.Bind(&data)

	data.StudentId = tools.GetNewStudentId()

	if data.Name == "" {
		app.Error(c, 200, errors.New(""), "姓名不能为空")
		return
	}

	if data.Grade == 0 {
		app.Error(c, 200, errors.New(""), "年级不能为空")
		return
	}
	if data.ClassBy == 0 {
		app.Error(c, 200, errors.New(""), "班级不能为空")
		return
	}

	data.CreateBy = tools.GetUserIdStr(c)
	tools.HasError(err, "", 500)
	result, err := data.Insert()
	tools.HasError(err, "", -1)
	app.OK(c, result, "")

}

func GetPage(c *gin.Context) {
	var data models.Student
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize, err = tools.StringToInt(size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex, err = tools.StringToInt(index)
	}

	result, count, err := data.GetPage(pageIndex, pageSize)
	tools.HasError(err, "", -1)
	app.PageOK(c, result, count, pageIndex, pageSize, "")

}

func UpdateStudent(c *gin.Context) {
	var data models.Student
	err := c.Bind(&data)

	if data.StudentId == 0 {
		app.Error(c, 200, err, "学号不能为空")
		return
	}
	_, _ = data.UpdateStudent()

	app.OK(c, "", "修改成功")
}

func DeleteStudent(c *gin.Context) {
	var data models.Student
	err := c.Bind(&data)

	if data.StudentId == 0 {
		app.Error(c, 200, err, "学号不能为空")
		return
	}
	_, _ = data.DeleteStudent()

	app.OK(c, "", "删除成功")
}
