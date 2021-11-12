package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang-jwt/jwt"
)

type Config struct {
	ServiceID string `json:"service_id"`
	KeyID     string `json:"key_id"`
	TeamID    string `json:"team_id"`
}

type AppleClientSecret struct {
	Issuer   string `json:"iss"`
	IssueAt  int64  `json:"iat"`
	Expire   int64  `json:"exp"`
	Audience string `json:"aud"`
	Sub      string `json:"sub"`
}

func (a AppleClientSecret) Valid() error {
	return nil
}

func getConfiguration() (*Config, error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(data, config)
	return config, err
}

func main() {
	var err error
	conf, err := getConfiguration()
	if err != nil {
		fmt.Println(err)
		return
	}

	privateBytes, err := ioutil.ReadFile("apple.p8")
	if err != nil {
		fmt.Println(err)
		return
	}

	privateKey, err := jwt.ParseECPrivateKeyFromPEM(privateBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	token := jwt.New(jwt.GetSigningMethod("ES256"))

	token.Header = map[string]interface{}{
		"kid": conf.KeyID,
		"alg": "ES256",
	}

	payload := AppleClientSecret{
		Issuer:   conf.TeamID,
		IssueAt:  time.Now().Unix(),
		Expire:   time.Now().Add(259200 * time.Minute).Unix(),
		Audience: "https://appleid.apple.com",
		Sub:      conf.ServiceID,
	}
	token.Claims = payload

	signature, err := token.SignedString(privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(signature)
}
