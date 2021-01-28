package model

import (
	"encoding/base64"
	"github.com/danzhuxia/ginblog/utils/errmsg"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

//User 用户
type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20); not null " json:"username" validate:"required,min=6,max=12" label:"用户名"`
	Password string `gorm:"type: varchar(20); not null " json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type: int; DEFAULT:2" json:"role" validate:"required,gte=1" label:"权限"`
}

//查询用户是否存在
func CheckUser(username string) (code int) {
	var users User
	Db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

//新增用户
func AddUserDao(data *User) int {
	data.Password = ScryptPwd(data.Password)
	err := Db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户列表
func GetUsersDao(pageSize, pageNum int) ([]User, int) {
	var users []User
	var total int
	err := Db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

//编辑用户
func EditUserDao(id int, u *User) int {
	var user = &User{}
	var data = map[string]interface{}{
		"username": u.Username,
		"role":     u.Role,
	}
	err := Db.Model(user).Where("id = ?", id).Update(data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除用户
func DeleteUserDao(id int) int {
	var user User
	err := Db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func ScryptPwd(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

//登录验证
func CheckLogin(username, password string) int {
	var user User
	Db.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ErrorUserNotExist
	}
	if ScryptPwd(password) != user.Password {
		return errmsg.ErrorPasswordWrong
	}
	if user.Role != 1 {
		return errmsg.ErrorUserNoRight
	}
	return errmsg.SUCCESS
}
