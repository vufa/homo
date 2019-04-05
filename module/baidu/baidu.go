//
// Copyright (C) 2019 Codist. - All Rights Reserved
// Unauthorized copying of this file, via any medium is strictly prohibited
// Proprietary and confidential
// Written by Codist <i@codist.me>, March 2019
//

package baidu

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const VOICE_AUTH_URL = "https://openapi.baidu.com/oauth/2.0/token"

// Authorizer 用于设置access_token
// 可以通过RESTFul api的方式从百度方获取
// 有效期为一个月，可以存至数据库中然后从数据库中获取
type Authorizer interface {
	Authorize(*Client) error
}

type Client struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	Authorizer   Authorizer
}

type AuthResponse struct {
	AccessToken      string `json:"access_token"`  //要获取的Access Token
	ExpireIn         string `json:"expire_in"`     //Access Token的有效期(秒为单位，一般为1个月)；
	RefreshToken     string `json:"refresh_token"` //以下参数忽略，暂时不用
	Scope            string `json:"scope"`
	SessionKey       string `json:"session_key"`
	SessionSecret    string `json:"session_secret"`
	ERROR            string `json:"error"`             //错误码；关于错误码的详细信息请参考鉴权认证错误码(http://ai.baidu.com/docs#/Auth/top)
	ErrorDescription string `json:"error_description"` //错误描述信息，帮助理解和解决发生的错误。
}

type DefaultAuthorizer struct{}

func (da DefaultAuthorizer) Authorize(client *Client) error {
	resp, err := http.PostForm(VOICE_AUTH_URL, url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {client.ClientID},
		"client_secret": {client.ClientSecret},
	})
	if err != nil {
		return err
	}
	authresponse := new(AuthResponse)
	if err := json.NewDecoder(resp.Body).Decode(&authresponse); err != nil {
		return err
	}
	if authresponse.ERROR != "" || authresponse.AccessToken == "" {
		return errors.New("授权失败:" + authresponse.ErrorDescription)
	}

	client.AccessToken = authresponse.AccessToken
	return nil
}

func (client *Client) Auth() error {
	if client.AccessToken != "" {
		return nil
	}

	if err := client.Authorizer.Authorize(client); err != nil {
		return err
	}
	return nil
}

func (client *Client) SetAuther(auth Authorizer) {
	client.Authorizer = auth
}

func NewClient(ApiKey, secretKey string) *Client {
	return &Client{
		ClientID:     ApiKey,
		ClientSecret: secretKey,
		Authorizer:   DefaultAuthorizer{},
	}
}
