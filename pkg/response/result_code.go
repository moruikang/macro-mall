// @Author moruikang
// @Date 2025/3/15 21:40:00
// @Desc 返回码

package response

var (
	SuccessCode       = Response{Code: 200, Message: "操作成功"}
	FailCode          = Response{Code: 500, Message: "操作失败"}
	InternalErrorCode = Response{Code: 500, Message: "内部错误"}
	ValidateErrorCode = Response{Code: 400, Message: "参数校验错误"}
	NotFoundCode      = Response{Code: 404, Message: "资源不存在"}
	UnauthorizedCode  = Response{Code: 401, Message: "暂未登录或token已过期"}
	ForbiddenCode     = Response{Code: 403, Message: "没有相关权限"}
)
