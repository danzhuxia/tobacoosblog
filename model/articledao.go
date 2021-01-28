package model

import (
	"github.com/danzhuxia/ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

//Article Article
type Article struct {
	gorm.Model
	Category Category	`gorm:"foreignkey:Cid"`
	Title    string `gorm:"type: varchar(100); not null" json:"title,omitempty"`
	Cid      int    `gorm:"type: int;not null" json:"cid,omitempty"`
	Desc     string `gorm:"type: varchar(200)" json:"desc,omitempty"`
	Content  string `gorm:"type: longtext" json:"content,omitempty"`
	Img      string `gorm:"type: varchar(100)" json:"img,omitempty"`
}

//新增分类
func AddArtDao(data *Article) int {
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询文章列表
func GetArtDao(pageSize, pageNum int) ([]Article, int, int) {
	var art []Article
	var total int
	err := Db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Count(&total).Error
	if err !=nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR, 0
	}
	return art, errmsg.SUCCESS, total
}

// 查询分类下的所有文章
func GetCateArtDao(id, pageSize, pageNum int) ([]Article, int, int) {
	var arts []Article
	var total int
	err := Db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Where("cid = ?",id).Find(&arts).Count(&total).Error
	if err != nil {
		return nil, errmsg.ErrorCategoryNotExit,0
	}
	//if arts == nil {
	//	return nil, errmsg.ERROR_CATE_ARTICLE_NOT_EXIT
	//}
	return arts, errmsg.SUCCESS,total
}

// 查询单个文章
func GetArtInfoDao(id int) (Article, int) {
	var art Article
	err := Db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art,errmsg.ErrorArticleNotExit
	}
	return art, errmsg.SUCCESS
}

//编辑文章
func EditArtDao(id int, a *Article) int {
	var art = &Article{}
	var data = map[string]interface{}{
		"title": a.Title,
		"cid":a.Cid,
		"desc":a.Desc,
		"content":a.Content,
		"img":a.Img,
	}
	err := Db.Model(art).Where("id = ?", id).Update(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return  errmsg.SUCCESS
}

//删除文章
func DeleteArtDao(id int) int {
	var art Article
	err := Db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}