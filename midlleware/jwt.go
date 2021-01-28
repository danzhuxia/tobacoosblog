package midlleware

import (
	"github.com/danzhuxia/ginblog/utils"
	"github.com/danzhuxia/ginblog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

var code int

// 生成Token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	Claim := MyClaims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, Claim)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

//验证Token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{},
	func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmsg.SUCCESS
	} else {
		return nil, errmsg.ERROR
	}
}

//JwtMid
func JwtMid() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenHandler := context.Request.Header.Get("Authorization")
		//code = errmsg.SUCCESS
		if tokenHandler == "" {
			code = errmsg.ErrorTokenNotExist
			context.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHandler, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenTypeWrong
			context.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		key, Tcode := CheckToken(checkToken[1])
		if Tcode == errmsg.ERROR {
			code = errmsg.ErrorTokenWrong
			context.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			context.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			context.Abort()
			return
		}

		context.Set("username", key.UserName)
		context.Next()
	}
}
