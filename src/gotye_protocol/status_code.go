package gotye_protocol

const (
	//common
	API_SUCCESS               = 10000
	API_SERVER_ERROR          = 10001
	API_PARAM_ERROR           = 10002
	API_EXPIRED_SESSION_ERROR = 10003

	//account
	API_ACCOUNT_NOT_EXISTS_ERROR = 10100
	API_LOGIN_PASSWORD_ERROR     = 10101
	API_ACCOUNT_EXISTS_ERROR     = 10102
	API_PHONE_EXISTS_ERROR       = 10103
	API_PHONE_NOT_EXISTS_ERROR   = 10104
	API_EMAIL_EXISTS_ERROR       = 10105
	API_EMAIL_NOT_EXISTS_ERROR   = 10106
	API_AUTHCODE_ERROR           = 10107

	//liveroom
	API_LIVEROOM_ID_NOT_EXIST_ERROR     = 10200
	API_REPECT_PASSWORD_LIVEROOM_ERROR  = 10201
	API_LIVEROOM_NOT_EXISTS_ERROR       = 10202
	API_INVALID_LIVEROOM_NAME_ERROR     = 10203
	API_INVALID_PASSWORD_LIVEROOM_ERROR = 10204

	//picture
	API_DECODE_HEAD_PIC_ERROR = 10300
)

var ApiStatus = map[int]string{
	API_SUCCESS:               "成功",
	API_SERVER_ERROR:          "系统异常",
	API_PARAM_ERROR:           "参数错误",
	API_EXPIRED_SESSION_ERROR: "Token无效",

	API_ACCOUNT_NOT_EXISTS_ERROR: "账号不存在",
	API_LOGIN_PASSWORD_ERROR:     "登录密码错误",
	API_ACCOUNT_EXISTS_ERROR:     "账号已注册",
	API_PHONE_EXISTS_ERROR:       "手机已注册",
	API_PHONE_NOT_EXISTS_ERROR:   "手机不存在",
	API_EMAIL_EXISTS_ERROR:       "邮箱已注册",
	API_EMAIL_NOT_EXISTS_ERROR:   "邮箱不存在",
	API_AUTHCODE_ERROR:           "验证码错误",

	API_LIVEROOM_ID_NOT_EXIST_ERROR:     "直播室ID不存在",
	API_REPECT_PASSWORD_LIVEROOM_ERROR:  "直播室密码重复",
	API_LIVEROOM_NOT_EXISTS_ERROR:       "直播室不存在",
	API_INVALID_LIVEROOM_NAME_ERROR:     "直播室名称非法",
	API_INVALID_PASSWORD_LIVEROOM_ERROR: "直播室密码非法",

	API_DECODE_HEAD_PIC_ERROR: "头像解码错误",
}
