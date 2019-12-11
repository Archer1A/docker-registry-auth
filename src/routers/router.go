package routers

import (
	"github.com/Archer1A/docker-registry-auth/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/archer/auth",&controllers.AuthController{}) // 设置路由
}
