package author

import "context"

type Authentication struct {
	clientId     string
	clientSecret string
}

// NewClientAuthentication 构造凭证
func NewClientAuthentication(clientId, clientSecret string) *Authentication {
	return &Authentication{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

// WithClientCredentials 通过客户端初始化凭证
func (a *Authentication) WithClientCredentials(clientId, clientSecret string) {
	a.clientId = clientId
	a.clientSecret = clientSecret
}

// GetRequestMetadata 从meta中取凭证信息：clientId，clientSecret
func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error) {

	return map[string]string{
		"client_id":     a.clientId,
		"client_secret": a.clientSecret,
	}, nil
}

// RequireTransportSecurity 指示凭据是否需要传输安全性。
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}
