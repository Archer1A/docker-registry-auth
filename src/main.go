package main

import (
	"github.com/Archer1A/docker-registry-auth/filter"
	_ "github.com/Archer1A/docker-registry-auth/routers"
	"github.com/astaxie/beego"
)

func main() {
	// 认证并，添加user
	beego.InsertFilter("/*",beego.BeforeRouter,filter.SecurityFilter)
	beego.Run()
}

