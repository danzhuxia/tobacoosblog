package midlleware

import (
	cors2 "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors2.New(cors2.Config{
			AllowAllOrigins: true,
			AllowMethods: []string{"*"},
			AllowHeaders: []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length", "Authorization"},
			//AllowCredentials: true,
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//},
			MaxAge: 12 * time.Hour,
		})
	}
}
