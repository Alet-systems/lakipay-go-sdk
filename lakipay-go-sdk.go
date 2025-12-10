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

func (sdk *LakiPayGoSDK) GetTransactionDetails(transactionID string) (*TransactionDetails, *ErrorResponse, error) {
	url := URL + "/payment/transaction/" + transactionID
	req, err := http.NewRequest("GET", url, nil)
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

	var response TransactionDetails
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, nil, err
	}

	var errorResponse ErrorResponse
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

func (sdk *LakiPayGoSDK) Checkout(params *CheckoutParams) (*SuccessResponse, *ErrorResponse, error) {
	url := URL + "/payment/checkout"
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

	var response SuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, nil, err
	}

	var errorResponse ErrorResponse
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

func (sdk *LakiPayGoSDK) WithDrawal(params *WithDrawalParams) (*SuccessResponse, *ErrorResponse, error) {
	url := URL + "/payment/withdrawal"
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

	var response SuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, nil, err
	}

	var errorResponse ErrorResponse
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

func (sdk *LakiPayGoSDK) DirectPayment(params *DirectPaymentParams) (*SuccessResponse, *ErrorResponse, error) {
	url := URL + "/payment/direct"
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

	var response SuccessResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, nil, err
	}

	var errorResponse ErrorResponse
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
