package models

import (
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"go-admin/common/dto"
	orm "go-admin/common/global"
	"go-admin/tools/excel"
)

type Student struct {
	StudentId       int    `json:"student_id" gorm:"primaryKey;<-:create;unique"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Age             int32  `json:"age"`
	Sex             int32  `json:"sex" gorm:"default:1"`
	Status          int32  `json:"status" gorm:"default:1"`
	Avatar          string `json:"avatar"`
	Grade           int    `json:"grade"`
	ClassBy         int    `json:"class_by"`
	Address         string `json:"address"`
	ContactBy       string `json:"contact_by"`
	ContactPhone    string `json:"contact_phone"`
	ContactRelation string `json:"contact_relation"`
	CreateBy        string `json:"create_by"`
	BaseModel
}

func (s Student) TableName() string {
	return "tbl_student"
}

/**
添加
*/
func (s *Student) Insert() (i int64, err error) {
	var count int64
	table := orm.Eloquent.Table(s.TableName())
	if err := table.Model(&s).Where("student_id=?", s.StudentId).Count(&count).Error; err != nil {
		return 0, err
	}
	if count > 0 {
		return 0, errors.New("学号已经存在！")
	}
	result := table.Create(&s)
	if result.Error != nil {
		err = result.Error
		return
	}
	i = result.RowsAffected
	return
}

func (s *Student) GetPage(pageIndex int, pageSize int) ([]Student, int, error) {
	var doc []Student
	table := orm.Eloquent.Table(s.TableName())

	if s.StudentId != 0 {
		table.Where("student_id=?", s.StudentId)
	}

	if s.Name != "" {
		table.Where("name=?", s.Name)
	}

	if s.Phone != "" {
		table.Where("phone=?", s.Phone)
	}
	var count int64

	if err := table.Order("created_At").Scopes(dto.FindPage(pageSize, pageIndex, &count, &doc)).Error; err != nil {
		return nil, 0, err
	}
	return doc, int(count), nil
}

func (s *Student) allData() ([]map[string]interface{}, error) {
	var doc []map[string]interface{}
	table := orm.Eloquent.Table(s.TableName())
	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

/**
修改
*/
func (s *Student) UpdateStudent() (update Student, err error) {

	if err = orm.Eloquent.Table(s.TableName()).First(&update, s.StudentId).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table(s.TableName()).Model(&update).Updates(&s).Error; err != nil {
		return
	}
	return
}

func (s *Student) DeleteStudent() (update Student, err error) {

	if err = orm.Eloquent.Table(s.TableName()).First(&update, s.StudentId).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table(s.TableName()).Delete(&s).Error; err != nil {
		return
	}
	return
}

/**
导出
*/
func (s *Student) Export() *excelize.File {
	var excelUtil excel.ExcelUtils
	var f = excelUtil.CreateFile()

	title := []map[string]string{
		{"label": "姓名", "key": "name"},
		{"label": "学号", "key": "student_id"},
		{"label": "性别", "key": "sex"},
		{"label": "年级", "key": "grade"},
	}

	data, _ := s.allData()

	f, err := excelUtil.Export(title, data)
	if err != nil {

	}

	return f
}
