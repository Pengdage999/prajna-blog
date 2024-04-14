package errmsg

// 在 utils 文件夹下，另外起一个包

const (
	SUCCESS = 200
	ERROR   = 500

	// 1000... 用户模块的错误

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	ERROR_TOKEN_EXIST      = 1009

	// 2000... 文章模块的错误

	ERROR_ART_NOT_EXIST = 2001

	// 3000... 分类模块的错误

	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
)

// 字典

var codeMsg = map[int]string{
	SUCCESS:                "成功!",
	ERROR:                  "错误!",
	ERROR_USERNAME_USED:    "用户已存在",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "Token不存在",
	ERROR_TOKEN_RUNTIME:    "Token已过期",
	ERROR_TOKEN_WRONG:      "Token不正确",
	ERROR_TOKEN_TYPE_WRONG: "Token格式错误",
	ERROR_USER_NO_RIGHT:    "用户无权限",
	ERROR_TOKEN_EXIST:      "Token已存在",

	ERROR_ART_NOT_EXIST: "文章不存在",

	ERROR_CATENAME_USED:  "分类已经存在",
	ERROR_CATE_NOT_EXIST: "分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
