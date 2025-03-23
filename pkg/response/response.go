// @Author moruikang
// @Date 2025/3/15 21:41:00
// @Desc

package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) GetCode() int {
	return r.Code
}

func (r Response) GetMessage() string {
	return r.Message
}

func Success(g *gin.Context, data interface{}) {
	g.JSONP(http.StatusOK,
		Response{
			Code:    SuccessCode.GetCode(),
			Message: SuccessCode.GetMessage(),
			Data:    data,
		})
	return
}

func Fail(g *gin.Context, message string) {
	g.JSONP(http.StatusOK,
		Response{
			Code:    FailCode.GetCode(),
			Message: FailCode.GetMessage(),
		})
	return
}

func InternalServerError(g *gin.Context, message string) {
	response := Response{
		Code:    InternalErrorCode.GetCode(),
		Message: InternalErrorCode.GetMessage(),
	}

	if message != "" {
		response.Message = fmt.Sprintf("%s %s", InternalErrorCode.GetMessage(), message)
	}
	//g.JSONP(http.StatusInternalServerError, response)
	g.JSONP(http.StatusOK, response)
	return
}

func ValidateError(g *gin.Context, message string) {
	response := Response{
		Code:    ValidateErrorCode.GetCode(),
		Message: ValidateErrorCode.GetMessage(),
	}

	if message != "" {
		response.Message = fmt.Sprintf("%s %s", ValidateErrorCode.GetMessage(), message)
	}
	g.JSONP(http.StatusOK, response)
	return
}

func Unauthorized(g *gin.Context) {
	response := Response{
		Code:    UnauthorizedCode.GetCode(),
		Message: UnauthorizedCode.GetMessage(),
	}

	g.JSONP(http.StatusOK, response)
	return
}

func Forbidden(g *gin.Context) {
	response := Response{
		Code:    ForbiddenCode.GetCode(),
		Message: ForbiddenCode.GetMessage(),
	}

	g.JSONP(http.StatusOK, response)
	return
}
