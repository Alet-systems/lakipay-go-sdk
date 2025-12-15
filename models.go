package lakipaygosdk

type WebhookRequest struct {
	CallbackURL   string `json:"callback_url"`
	MerchantID    string `json:"merchant_id"`
	Message       string `json:"message"`
	ProviderData  string `json:"provider_data"`
	ProviderTxID  string `json:"provider_tx_id"`
	Status        string `json:"status"`
	Timestamp     string `json:"timestamp"`
	TransactionID string `json:"transaction_id"`
	// TODO: type need to be enum
	Type string `json:"type"`
}

type TransactionDetails struct {
	Amount      float64 `json:"amount"`
	CreatedAt   string  `json:"created_at"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	Details     any     `json:"details"`
	FeeAmount   float64 `json:"fee_amount"`
	ID          string  `json:"id"`
	Medium      string  `json:"medium"`
	Reference   string  `json:"reference"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
	Type        string  `json:"type"`
	UpdatedAt   string  `json:"updated_at"`
}

type CheckoutParams struct {
	Amount           float64   `json:"amount"`
	CallbackURL      string    `json:"callback_url"`
	Currency         string    `json:"currency"`
	Description      string    `json:"description"`
	PhoneNumber      string    `json:"phone_number"`
	Redirects        Redirects `json:"redirects"`
	Reference        string    `json:"reference"`
	SupportedMediums []string  `json:"supported_mediums"`
}

type WithDrawalParams struct {
	Amount      float64 `json:"amount"`
	CallbackURL string  `json:"callback_url"`
	Currency    string  `json:"currency"`
	Medium      string  `json:"medium"`
	PhoneNumber string  `json:"phone_number"`
	Reference   string  `json:"reference"`
}

type DirectDetails struct {
	Phone string `json:"phone"`
}

type Redirects struct {
	Failed  string `json:"failed"`
	Success string `json:"success"`
}

type DirectPaymentParams struct {
	Amount      float64       `json:"amount"`
	CallbackURL string        `json:"callback_url"`
	Currency    string        `json:"currency"`
	Description string        `json:"description"`
	Details     DirectDetails `json:"details"`
	Medium      string        `json:"medium"`
	PhoneNumber string        `json:"phone_number"`
	Redirects   Redirects     `json:"redirects"`
	Reference   string        `json:"reference"`
}

type SuccessResponse struct {
	LakiPayTransactionID string `json:"lakipay_transaction_id"`
	Message              string `json:"message"`
	PaymentURL           string `json:"payment_url"`
	ReferenceID          string `json:"reference_id"`
	Status               string `json:"status"`
	Success              bool   `json:"success"`
}

type ErrorType struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type ErrorResponse struct {
	Success bool      `json:"success"`
	Error   ErrorType `json:"error"`
}

type APIResponse struct {
	Success bool             `json:"success"`
	Data    *SuccessResponse `json:"data,omitempty"`
	Error   *ErrorResponse   `json:"error,omitempty"`
}

type TransactionResponse struct {
	Data  *TransactionDetails
	Error *ErrorResponse
}
