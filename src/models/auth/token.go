package auth

type Token struct {
	Token       string `json:"token"`
	ExpiresIn   int    `json:"expires_in"`
	IssuedAt    string `json:"issued_at"`
}

// 存在资源和其操作权限
type ResourceActions struct {
	Type    string   `json:"type"`
	Class   string   `json:"class,omitempty"`
	Name    string   `json:"name"`
	Actions []string `json:"actions"`
}

type ClaimSet struct {
	// Public claims
	Issuer     string `json:"iss"` // 发行人
	Subject    string `json:"sub"`
	Audience   string `json:"aud"` //
	Expiration int64  `json:"exp"`
	NotBefore  int64  `json:"nbf"`
	IssuedAt   int64  `json:"iat"` // 发行时间
	JWTID      string `json:"jti"` // JWT id

	// Private claims
	Access []*ResourceActions `json:"access"`


}

//JWT 的header
type Head struct {
	Type string `json:"typ"`
	Kid string `json:"kid,omitempty"`
	Alg string `json:"alg"`
}