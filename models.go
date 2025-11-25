package lakipaygosdk

type DirectDetais struct {
	Phone string `json:"phone"`
}

type Redirects struct {
	Failed  string `json:"failed"`
	Success string `json:"success"`
}

type DirectPaymentParams struct {
	Amount      float64      `json:"amount"`
	CallbackURL string       `json:"callback_url"`
	Currency    string       `json:"currency"`
	Description string       `json:"description"`
	Details     DirectDetais `json:"details"`
	Medium      string       `json:"medium"`
	PhoneNumber string       `json:"phone_number"`
	Redirects   Redirects    `json:"redirects"`
	Reference   string       `json:"reference"`
}

type DirectPaymentSuccessResponse struct {
	LakiPayTransactionID string `json:"lakipay_transaction_id"`
	Message              string `json:"message"`
	PaymentURL           string `json:"payment_url"`
	ReferenceID          string `json:"reference_id"`
	Status               string `json:"status"`
	Success              bool   `json:"success"`
}

type DirectPaymentErrorType struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

type DirectPaymentErrorResponse struct {
	Success bool                   `json:"success"`
	Error   DirectPaymentErrorType `json:"error"`
}
