package v1

import (
	"github.com/danzhuxia/ginblog/model"
	"github.com/danzhuxia/ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddArticle(c *gin.Context) {
	var art model.Article
	_ = c.ShouldBindJSON(&art)
	code = model.AddArtDao(&art)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询分类下所有文章
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	id, _ := strconv.Atoi(c.Param("id"))

	if arts, code, total := model.GetCateArtDao(id, pageSize, pageNum); arts == nil {
		code = errmsg.ErrorCateArticleNotExit
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    arts,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    arts,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

//查询单个文章
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.GetArtInfoDao(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    art,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticlesList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	articles, code, total := model.GetArtDao(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var art model.Article
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	_ = c.ShouldBindJSON(&art)
	code = model.EditArtDao(id, &art)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	code = model.DeleteArtDao(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
