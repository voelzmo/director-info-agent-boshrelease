package api

type Login struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	JTI              string `json:"jti"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
