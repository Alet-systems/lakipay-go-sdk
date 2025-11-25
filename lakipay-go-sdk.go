package lakipaygosdk

import (
	"bytes"
	"encoding/json"
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

func (sdk *LakiPayGoSDK) DirectPayment(params *DirectPaymentParams) (*DirectPaymentSuccessResponse, *DirectPaymentErrorResponse, error) {
	url := URL + "/direct-payment"
	body, err := json.Marshal(params)
	if err != nil {
		return nil, nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	if sdk.Credentials.ApiKey != nil {
		req.Header.Set("X-Api-Key", *sdk.Credentials.ApiKey)
	}

	if sdk.Credentials.JWTSecret != nil {
		req.Header.Set("Authorization", "Bearer "+*sdk.Credentials.JWTSecret)
	}

	resp, err := sdk.Client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var response DirectPaymentSuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, nil, err
	}

	var errorResponse DirectPaymentErrorResponse
	err = json.NewDecoder(resp.Body).Decode(&errorResponse)
	if err != nil {
		return nil, nil, err
	}

	if errorResponse.Success {
		return &response, nil, nil
	} else {
		return nil, &errorResponse, nil
	}
}
