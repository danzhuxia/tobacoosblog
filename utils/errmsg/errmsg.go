package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code = 100X...用户模块的错误
	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenNotExist  = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008

	//code = 200X...文章模块的错误
	ErrorArticleNotExit     = 2001
	ErrorCateArticleNotExit = 2002

	//code = 300X...分类模块错误
	ErrorCategorynameUsed = 3001
	ErrorCategoryNotExit  = 3002
)

var codeMsg = map[int]string{
	SUCCESS: "OK",
	ERROR:   "FAIL",

	ErrorUsernameUsed:       "用户名已经被使用！",
	ErrorPasswordWrong:      "密码错误！",
	ErrorUserNotExist:       "用户不存在！",
	ErrorUserNoRight:        "用户无管理权限",
	ErrorTokenNotExist:      "TOKEN不存在！",
	ErrorTokenRuntime:       "TOKEN超时！",
	ErrorTokenWrong:         "TOKEN出错！",
	ErrorTokenTypeWrong:     "TOKEN格式出错！",
	ErrorArticleNotExit:     "文章不存在！",
	ErrorCateArticleNotExit: "该分类下没有文章！",
	ErrorCategorynameUsed:   "分类已使用！",
	ErrorCategoryNotExit:    "分类不存在！",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
