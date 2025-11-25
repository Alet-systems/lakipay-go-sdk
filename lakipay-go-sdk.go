package lakipaygosdk

import (
	"net/http"
	"time"
)

const URL = "https://api.lakipay.co/api/v2"

type LakiPayGoSDK struct {
	Credentials *Credentials
	Client      *http.Client
}

type Credentials struct {
	ApiKey    *string
	JWTSecret *string
}

func NewLakiPaySDKWithAPIKey(apiKey string) *LakiPayGoSDK {
	return &LakiPayGoSDK{
		Credentials: &Credentials{
			ApiKey: &apiKey,
		},
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func NewLakiPaySDKWithJWTSecret(jwtSecret string) *LakiPayGoSDK {
	return &LakiPayGoSDK{
		Credentials: &Credentials{
			JWTSecret: &jwtSecret,
		},
		Client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

