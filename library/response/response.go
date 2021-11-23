package response

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, data interface{}) {
	err := r.Response.WriteJson(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	e := r.Response.WriteJson(JsonResponse{
		Code:    err,
		Message: msg,
		Data:    responseData,
	})

	if e != nil {
		fmt.Println(e.Error())
	}

	r.Exit()
}
