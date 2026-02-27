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

func (sdk *LakiPayGoSDK) GetApiKey() string {
	if sdk.Credentials == nil || sdk.Credentials.ApiKey == nil {
		return ""
	}
	return *sdk.Credentials.ApiKey
}

func (sdk *LakiPayGoSDK) SetApiKey(apiKey string) {
	if sdk.Credentials == nil {
		sdk.Credentials = &Credentials{}
	}
	key := apiKey
	sdk.Credentials.ApiKey = &key
}

func (sdk *LakiPayGoSDK) GetTransactionDetails(transactionID string) (*TransactionResponse, error) {
	url := URL + "/payment/transaction/" + transactionID
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()

	var response TransactionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (sdk *LakiPayGoSDK) Checkout(params *CheckoutParams) (*APIResponse, error) {
	url := URL + "/payment/checkout"
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()

	var response APIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (sdk *LakiPayGoSDK) WithDrawal(params *WithDrawalParams) (*APIResponse, error) {
	url := URL + "/payment/withdrawal"
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()

	var response APIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (sdk *LakiPayGoSDK) DirectPayment(params *DirectPaymentParams) (*APIResponse, error) {
	url := URL + "/payment/direct"
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
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
		return nil, err
	}
	defer resp.Body.Close()

	var response APIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
