package model

import (
	"fmt"
	"time"

	"github.com/danzhuxia/ginblog/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB
var err error

//InitDB 初始化数据库
func InitDB() {
	Db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))

	//defer db.Close()

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数", err)
	}

	//SingularTable 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	Db.SingularTable(true)

	//AutoMigrate 自动迁移
	Db.AutoMigrate(&User{}, &Article{}, &Category{})

	// SetMaxIdleCons 设置连接池中的最大闲置连接数
	Db.DB().SetMaxIdleConns(10)

	//SetMaxOpenCons 设置数据库的最大连接数量
	Db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifeTime 设置连接的最大可复用时间
	Db.DB().SetConnMaxLifetime(10 * time.Second)

}
