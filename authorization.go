package omisego

import (
	"encoding/base64"
)

type Authorization interface {
	CreateAuthorizationHeader() string
}

type AdminAuth struct {
	UserId        string
	UserAuthToken string
}

type ProviderAuth struct {
	AccessKey string
	SecretKey string
}

type ClientAuth struct {
	ApiKey             string
	AuthenicationToken string
}

func (a *AdminAuth) CreateAuthorizationHeader() string {
	data := []byte(a.UserId + ":" + a.UserAuthToken)
	str := base64.StdEncoding.EncodeToString(data)
	return "OMGAdmin " + str
}

func (s *ProviderAuth) CreateAuthorizationHeader() string {
	data := []byte(s.AccessKey + ":" + s.SecretKey)
	str := base64.StdEncoding.EncodeToString(data)
	return "OMGProvider " + str
}

func (c *ClientAuth) CreateAuthorizationHeader() string {
	data := []byte(c.ApiKey + ":" + c.AuthenicationToken)
	str := base64.StdEncoding.EncodeToString(data)
	return "OMGClient " + str
}
