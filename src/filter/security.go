package filter

import (
	"context"
	beegoCtx "github.com/astaxie/beego/context"
	"net/http"
)

func SecurityFilter(ctx *beegoCtx.Context)  {
	if ctx == nil {
		return
	}
	req := ctx.Request
	if req == nil {
		return
	}
	Modify(ctx)
}

// 校验User 和password
func Modify(ctx *beegoCtx.Context){
	userName,secret,ok := ctx.Request.BasicAuth()
	if !ok {
		return
	}
	// 对client传输的用户名密码进行校验
	if userName != "user" && secret != "123456" {
		setUserNameToContext("",ctx.Request)
		return
	}
	// 放于容器中
	setUserNameToContext("user",ctx.Request)

}

func setUserNameToContext(userName string ,r *http.Request)  {
	*r = *(r.WithContext(context.WithValue(r.Context(),"User",userName)))
}