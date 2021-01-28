package router

import (
	v1 "github.com/danzhuxia/ginblog/api/v1"
	"github.com/danzhuxia/ginblog/midlleware"
	"github.com/danzhuxia/ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {

	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(midlleware.Logger())
	r.Use(midlleware.Cors())
	r.Use(gin.Recovery())

	//指定模板文件引用静态文件位置
	//r.Static("/static","static")
	//指定html文件位置
	//r.LoadHTMLGlob("/view/*")

	auth := r.Group("api/v1")
	auth.Use(midlleware.JwtMid())
	{
		//用户模块的路由接口
		//3.编辑用户
		auth.PUT("user/:id", v1.EditUser)
		//4.删除用户
		auth.DELETE("user/:id", v1.DeleteUser)
		//文章模块的路由接口
		//1.添加文章
		auth.POST("article/add", v1.AddArticle)
		//4.编辑文章
		auth.PUT("article/:id", v1.EditArticle)
		//5.删除文章
		auth.DELETE("article/:id", v1.DeleteArticle)
		//分类模块的路由接口
		//1. 添加分类
		auth.POST("category/add", v1.AddCategory)
		//4. 编辑分类
		auth.PUT("category/:id", v1.EditCategory)
		//5. 删除分类
		auth.DELETE("category/:id", v1.DeleteCategory)
		//上传文件
		auth.POST("upload", v1.Upload)
	}

	router := r.Group("api/v1")
	{
		//登录模块的路由
		router.POST("login", v1.Login)
		//1.添加用户
		router.POST("user/add", v1.AddUser)
		//2.查询用户列表
		router.GET("users", v1.GetUsersList)
		//2.查询文章列表
		router.GET("article", v1.GetArticlesList)
		//3.查询单个文章
		router.GET("article/:id", v1.GetArticle)
		//2. 查询单个分类下的文章
		router.GET("category/article/:id", v1.GetCateArt)
		//3. 查询分类列表
		router.GET("category", v1.GetCategoryslist)
	}

	_ = r.Run(utils.HttpPort)
}
