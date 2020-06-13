package models

import (
	"encoding/json"
	"errors"
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
	Error404          = &HTTPResult{Code: "404", Msg: "资源不存在"}
	Error405          = &HTTPResult{Code: "405", Msg: "请求的姿势错误"}
	Error401          = &HTTPResult{Code: "401", Msg: "请先登录"}
	ErrorAuthExpire   = &HTTPResult{Code: "401", Msg: "登录过期，请重新登录"}
	Error500          = &HTTPResult{Code: "500", Msg: "未知异常，请稍后重试"}
	ErrorParamBind    = &HTTPResult{Code: "100400", Msg: "参数校验失败，请检查输入参数"}
	ErrorLogin        = &HTTPResult{Code: "100401", Msg: "登陆失败，请检查用户名和密码"}
	ErrorUserExits    = &HTTPResult{Code: "100402", Msg: "用户已经存在"}
	ErrorInvalidEmail = &HTTPResult{Code: "100403", Msg: "无效的Email"}
)

var (
	ErrUserNotExits = errors.New("用户不存在")
	ErrUserExits    = errors.New("用户已存在")
	ErrUserLogin    = errors.New("请检查用户名密码")
	ErrRegister     = errors.New("注册失败")
	ErrUnknow       = errors.New("未知错误")
)

// Success 普通的成功返回
func Success(data interface{}) *HTTPResult {
	return &HTTPResult{Code: "200", Msg: "成功", Data: data, Success: true}
}

// Faild 发现了错误的请求
func Faild(warp error) *HTTPResult {
	return &HTTPResult{Code: "400", Msg: warp.Error(), Data: nil, Success: false}
}

// Bytes Bytes
func (r *HTTPResult) Bytes() []byte {
	bs, _ := json.Marshal(r)
	return bs
}
