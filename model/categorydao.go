package model

import (
	"github.com/danzhuxia/ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//Category Category
type Category struct {
	ID uint	`gorm:"primary_key; auto_increment" json:"id"`
	Name string `gorm:"type: varchar(20); not null " json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	Db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID >0 {
		return errmsg.ErrorCategorynameUsed
	}
	return errmsg.SUCCESS
}

//新增分类
func AddCategoryDao(data *Category) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类列表
func GetCategoryDao(pageSize, pageNum int) ([]Category, int) {
	var cate []Category
	var total int
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err !=nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cate, total
}


//编辑分类
func EditCategoryDao(id int, c *Category) int {
	var cate = &Category{}
	var data = map[string]interface{}{
		"name": c.Name,
	}
	err := Db.Model(cate).Where("id = ?", id).Update(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return  errmsg.SUCCESS
}

//删除用户
func DeleteCategoryDao(id int) int {
	var cate Category
	err := Db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}