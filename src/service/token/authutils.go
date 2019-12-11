package token

import (
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Archer1A/docker-registry-auth/models/auth"
	"github.com/docker/distribution/registry/auth/token"
	"github.com/docker/libtrust"
	"strings"
	"time"
)

func MakeToken(userName,service string, access []*auth.ResourceActions)(*auth.Token,error)  {
	// 使用docker libtrust 包读取pem
	pk,err := libtrust.LoadKeyFile("C:\\Users\\Vic\\go\\src\\github.com\\Archer1A\\docker-registry-auth\\conf\\private_key.pem")
	if err != nil{
		return nil,err
	}
	expiration := 1800
	issuer := "harbor-token-issuer"
	// 创建token
	tk,expiresIn,issuedAt ,err := makeTokenCore(issuer,userName,service,expiration,access,pk)
	if err != nil {
		return nil,err
	}

	rs := fmt.Sprintf("%s.%s",tk.Raw,base64UrlEncode(tk.Signature))
	return &auth.Token{
		Token:       rs,
		ExpiresIn:   expiresIn,
		IssuedAt:    issuedAt.Format(time.RFC3339),
	},nil
}

func makeTokenCore(issuer, subject, audience string, expiration int,
	access []*auth.ResourceActions, signingKey libtrust.PrivateKey) (t *token.Token, expiresIn int, issuedAt *time.Time, err error) {
		// 生成Header
	joseHeader := &auth.Head{
		Type: "JWT",
		Kid:  signingKey.KeyID(),
		Alg:  "RS256",
	}
	// 生成随机的16位jwt ID
	jwtID ,err := randString(16)
	if err != nil {
		return
	}
	now := time.Now().UTC()
	issuedAt = &now
	expiresIn = expiration
	// 生成claim Set
	claimSet := &auth.ClaimSet{
		Issuer:     issuer,
		Subject:    subject,
		Audience:   audience,
		Expiration: now.Add(time.Duration(expiration)*time.Minute).Unix(),
		NotBefore:  now.Unix(),
		IssuedAt:   now.Unix(),
		JWTID:      jwtID,
		Access:     access,
	}
	// 实体类转json
	var joseHeaderBytes ,claimSetBytes []byte
	if joseHeaderBytes,err = json.Marshal(joseHeader);err != nil {
		return nil,0,nil, fmt.Errorf("Error to generate jwt id: %s", err)
	}
	if claimSetBytes,err = json.Marshal(claimSet);err != nil {
		return nil,0,nil, fmt.Errorf("Error to generate jwt id: %s", err)
	}
	fmt.Println(string(claimSetBytes))
	fmt.Println(string(joseHeaderBytes))
	// 将header 和claim set base64 编码
	encodedJoseHeader := base64UrlEncode(joseHeaderBytes)
	encodedClaimSet := base64UrlEncode(claimSetBytes)
	// 将header 和claim 通过 . 连接
	payload := fmt.Sprintf("%s.%s",encodedJoseHeader,encodedClaimSet)
	var signatureBytes []byte
	// 加入私钥进行加密
	if signatureBytes ,_ ,err = signingKey.Sign(strings.NewReader(payload),crypto.SHA256);err!=nil {
		return
	}
	// 将签名base64 编码
	signature := base64UrlEncode(signatureBytes)
	tokenString := fmt.Sprintf("%s.%s",payload,signature)
	// 通过docker 的工具包生成Token 实体类
	t,err = token.NewToken(tokenString)
	return


}

func randString(length int)(string,error)  {
	const alphanum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rb := make([]byte,length)
	_,err := rand.Read(rb)
	if err != nil {
		return "",err
	}
	for i, b := range rb {
		rb[i] =  alphanum[int(b)%len(alphanum)]
	}
	return string(rb),nil

}

func base64UrlEncode(b []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(b), "=")
}