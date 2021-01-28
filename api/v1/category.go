package v1

import (
	"github.com/danzhuxia/ginblog/model"
	"github.com/danzhuxia/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	switch code {
	case errmsg.SUCCESS:
		model.AddCategoryDao(&data)
	case errmsg.ErrorCategorynameUsed:
		code = errmsg.ErrorCategorynameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询分类列表
func GetCategoryslist(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	cate, total := model.GetCategoryDao(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    cate,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCategory(c *gin.Context) {
	var data model.Category
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.EditCategoryDao(id, &data)
	}
	if code == errmsg.ErrorCategorynameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCategory(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	code = model.DeleteCategoryDao(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
