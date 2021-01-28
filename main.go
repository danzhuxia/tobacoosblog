package main

import (
	"github.com/danzhuxia/ginblog/model"
	"github.com/danzhuxia/ginblog/router"
)

func main() {
	//引用数据库并在结束时关闭
	defer model.Db.Close()
	model.InitDB()
	router.InitRouter()
}
