package e

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400
	//UnauthorizedTokenError      = 401

	NOT_FOUND      = 404

	ERROR_EXIST_TAG       = 10001
	ERROR_EXIST_TAG_FAIL  = 10002
	ERROR_NOT_EXIST_TAG   = 10003
	ERROR_GET_TAGS_FAIL   = 10004
	ERROR_COUNT_TAG_FAIL  = 10005
	ERROR_ADD_TAG_FAIL    = 10006
	ERROR_EDIT_TAG_FAIL   = 10007
	ERROR_DELETE_TAG_FAIL = 10008
	ERROR_EXPORT_TAG_FAIL = 10009
	ERROR_IMPORT_TAG_FAIL = 10010

	ERROR_CREATE_FAIL     					= 20000
	ERROR_USER_CREATE_FAIL      			=  200001

	ERROR_AUTH_CHECK_TOKEN_TIMEOUT			=  400001
	ERROR_AUTH_CHECK_TOKEN_FAIL				=  400002
)

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(100000, "服务内部错误")
	InvalidParams             = NewError(100001, "参数错误")
	NotFound                  = NewError(100002, "找不到")
	UnauthorizedAuthNotExist  = NewError(100003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(100004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(100005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(100006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(100007, "请求过多")
	BcryptPasswordError       = NewError(100008, "密码加密失败")
	UnauthorizedFail    	  = NewError(100009, "AuthorizedFail")
	ReferralCodeAuthError     = NewError(200001, "ReferralCodeAuthError.")
	ReferralCodeAuthFail      = NewError(200002, "ReferralCodeAuthFail.")
	ReferralCodeUsed      	  = NewError(200003, "ReferralCodeUsed.")
	UserIDParseError		  = NewError(200004, "UserIDParseError.")

	ERROR_GET_USER_FAIL      = NewError(200005, "获取用户数据失败")

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = NewError(300001, "ERROR_UPLOAD_SAVE_IMAGE_FAIL")
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = NewError(300002, "ERROR_UPLOAD_CHECK_IMAGE_FAIL")
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = NewError(300003, "ERROR_UPLOAD_CHECK_IMAGE_FORMAT")



)