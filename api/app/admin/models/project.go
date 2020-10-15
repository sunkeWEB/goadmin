package models

import (
	"errors"
	orm "go-admin/common/global"
)

// Project 项目
type Project struct {
	ProjectId  int    `gorm:"primary_key;AUTO_INCREMENT" json:"project_id"`
	Title      string `json:"title"`
	Status     int    `json:"status"`
	CreateBy   string `json:"create_by"`
	UpdateBy   string `json:"update_by"`
	ViewStatus int    `json:"view_status"`
	BaseModel
}

func (Project) TableName() string {
	return "tbl_project"
}

func GetProjectByTitle(title string) (p Project) {
	orm.Eloquent.Table(p.TableName()).Select("tbl_project.*").
		Where("title=?", title)
	return
}

/**
  添加
*/
func (p *Project) Insert() (id int, err error) {

	var count int64
	orm.Eloquent.Table(p.TableName()).Where("title=?", p.Title).Count(&count)
	if count > 0 {
		err = errors.New(p.Title + "已经存在")
		return
	}
	if err = orm.Eloquent.Table(p.TableName()).Create(&p).Error; err != nil {
		return
	}
	id = p.ProjectId
	return
}

/**
  软删除
*/
func (p *Project) SoftDelete() (Result bool, err error) {

	if err = orm.Eloquent.Table(p.TableName()).Model(&p).Updates(Project{
		ViewStatus: p.ViewStatus, UpdateBy: p.UpdateBy,
	}).Error; err != nil {
		return false, err
	}
	return true, err

}

/**
获取所以项目
*/
func (p *Project) GetListAll() ([]Project, error) {
	var doc []Project
	table := orm.Eloquent.Table(p.TableName())
	if err := table.Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

/**
分页获取
*/
func (p *Project) GetPage(pageIndex int, pageSize int) ([]Project, int, error) {
	var doc []Project

	table := orm.Eloquent.Table(p.TableName())

	// 数据权限控制
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = tools.StringToInt(p.DataScope)
	//table, err := dataPermission.GetDataScope("sys_post", table)
	//if err != nil {
	//	return nil, 0, err
	//}
	var count int64

	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Offset(-1).Limit(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	//table.Where("`deleted_at` IS NULL").Count(&count)
	return doc, int(count), nil
}
