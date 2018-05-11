package omisego

import (
	"encoding/base64"
)

type Authorization interface {
	CreateAuthorizationHeader() string
}

type AdminClientAuth struct {
	ApiKeyId string
	ApiKey   string
}

type AdminUserAuth struct {
	ApiKeyId      string
	ApiKey        string
	UserId        string
	UserAuthToken string
}

type ServerAuth struct {
	AccessKey string
	SecretKey string
}

type ClientAuth struct {
	ApiKey             string
	AuthenicationToken string
}

func (a *AdminClientAuth) CreateAuthorizationHeader() string {
	data := []byte(a.ApiKeyId + ":" + a.ApiKey)
	str := base64.StdEncoding.EncodeToString(data)
	return "OMGAdmin " + str
}

func (a *AdminUserAuth) CreateAuthorizationHeader() string {
	data := []byte(a.ApiKeyId + ":" + a.ApiKey + ":" + a.UserId + ":" + a.UserAuthToken)
	str := base64.StdEncoding.EncodeToString(data)
	return "OMGAdmin " + str
}

func (s *ServerAuth) CreateAuthorizationHeader() string {
	data := []byte(s.AccessKey + ":" + s.SecretKey)
	str := base64.StdEncoding.EncodeToString(data)
	return "OMGServer " + str
}

func (c *ClientAuth) CreateAuthorizationHeader() string {
	data := []byte(c.ApiKey + ":" + c.AuthenicationToken)
	str := base64.StdEncoding.EncodeToString(data)
	return "OMGClient " + str
}
