package models

import (
	"encoding/json"
)

type (
	// HTTPResult http的响应结构体 用于统一返回结果
	HTTPResult struct {
		Code    string      `json:"code"`
		Msg     string      `json:"msg"`
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}
)

var (
	RESOURCE_NOT_EXITS = &HTTPResult{Code: "404", Msg: "资源不存在"}
	METHOD_NOT_ALLOWED = &HTTPResult{Code: "405", Msg: "请求的姿势错误"}
	NOT_AUTH           = &HTTPResult{Code: "401", Msg: "请先登录"}
	AUTH_EXPIRE        = &HTTPResult{Code: "401", Msg: "登录过期，请重新登录"}
	UNKNOW_ERROR       = &HTTPResult{Code: "500", Msg: "未知异常，请稍后重试"}
	PARAM_BIND_ERROR   = &HTTPResult{Code: "100400", Msg: "参数校验失败，请检查输入参数"}
	LOGIN_ERROR        = &HTTPResult{Code: "100401", Msg: "登陆失败，请检查用户名和密码"}
	USER_EXITS_ERROR   = &HTTPResult{Code: "100402", Msg: "用户已经存在"}
	INVALID_EMAIL      = &HTTPResult{Code: "100403", Msg: "无效的Email"}
)

// Success 普通的成功返回
func Success(data interface{}) *HTTPResult {
	return &HTTPResult{Code: "200", Msg: "成功", Data: data, Success: true}
}

// Bytes Bytes
func (r *HTTPResult) Bytes() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
