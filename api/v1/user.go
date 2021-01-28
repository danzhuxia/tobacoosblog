package v1

import (
	"github.com/danzhuxia/ginblog/model"
	"github.com/danzhuxia/ginblog/utils/errmsg"
	"github.com/danzhuxia/ginblog/utils/validator"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var msg string
	_ = c.ShouldBindJSON(&data)
	//将接收的数据进行验证
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})
		return
	}

	code = model.CheckUser(data.Username)
	switch code {
	case errmsg.SUCCESS:
		model.AddUserDao(&data)
	case errmsg.ErrorUsernameUsed:
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询用户列表
func GetUsersList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users, total := model.GetUsersDao(pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	_ = c.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUserDao(id, &data)
	}
	if code == errmsg.ErrorUsernameUsed {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteUser(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	code = model.DeleteUserDao(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
