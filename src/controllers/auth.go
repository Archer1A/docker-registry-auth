package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Archer1A/docker-registry-auth/service/token"
	"github.com/astaxie/beego"
	"log"
	"net/http"
)

type AuthController struct {
	beego.Controller
}

func (ac *AuthController) Get() {
	request := ac.Ctx.Request
	//service := ac.GetString("service")
	scopes := ac.GetString("scope")
	token, err := token.Creator(request)
	if err != nil {
		ac.CustomAbort(http.StatusUnauthorized,"")

	}
	rs, err := json.Marshal(token)
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Println(scopes)
	fmt.Println(string(rs))
	ac.Data["json"] =token
	ac.ServeJSON()

}

func (ac *AuthController) Auth()  {

}