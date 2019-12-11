package token

import (
	"fmt"
	"github.com/Archer1A/docker-registry-auth/models/auth"
	"net/http"
	"strings"
)

func Creator(r *http.Request)(*auth.Token,error)  {
	userName  :=r.Context().Value("User")
	service := r.URL.Query().Get("service")
	scopes := parserScopes(r)
	if userName == nil{
		return nil,fmt.Errorf("")
	}
	user,ok := userName.(string)
	if !ok {
		return nil,fmt.Errorf("")
	}
	var access  []*auth.ResourceActions
	access = filterAccess(scopes)

	return  MakeToken(user,service,access)
}

func parserScopes(r *http.Request) []string  {
	path := r.URL
	var scopes []string
	for _,scope := range path.Query()["scope"]   {
		scopes = append(scopes,strings.Split(scope," ")...)
	}
	return scopes
}
func filterAccess(scopes []string)(res  []*auth.ResourceActions) {
	//var res  []*auth.ResourceActions
	for _,scope := range scopes{
		if scope == "" {
			continue

		}
		result := strings.Split(scope,":")
		typee := ""
		name := ""
		actions := []string{}
		length := len(result)
		if length == 1 {
			typee = result[0]
		}else if length == 2 {
			typee = result[0]
			name = result[1]
		}else {
			typee = result[0]
			name = strings.Join(result[1:length-1],":")
			if len(result[length-1])>0 {
				actions = strings.Split(result[length-1],",")
			}
		}
		res = append(res, &auth.ResourceActions{
			Type:    typee,
			Class:   "",
			Name:    name,
			Actions: actions,
		})

	}
	return
}