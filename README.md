# LakiPay Go SDK

The LakiPay Go SDK provides a convenient way to interact with the LakiPay API from your Go applications.

## Installation

To use the LakiPay Go SDK, you first need to have Go installed and set up on your machine. Then, you can install the SDK using the following command:

```bash
go get github.com/Alet-systems/lakipay-go-sdk
```

## Initialization

The SDK can be initialized using either an API key or a JWT secret.

### With API Key

```go
import "github.com/Alet-systems/lakipay-go-sdk"

sdk := lakipaygosdk.NewLakiPaySDKWithAPIKey("YOUR_API_KEY")
```

### With JWT Secret

```go
import "github.com/Alet-systems/lakipay-go-sdk"

sdk := lakipaygosdk.NewLakiPaySDKWithJWTSecret("YOUR_JWT_SECRET")
```

## Usage

The LakiPay Go SDK provides methods for handling payments, withdrawals, and transaction details.

### Checkout

The `Checkout` method allows you to create a new payment checkout.

```go
params := &lakipaygosdk.CheckoutParams{
    Amount:      100,
    CallbackURL: "https://example.com/callback",
    Currency:    "ETB",
    Description: "Test payment",
    PhoneNumber: "251911111111",
    Redirects: lakipaygosdk.Redirects{
        Failed:  "https://example.com/failed",
        Success: "https://example.com/success",
    },
    Reference:        "Test payment",
    SupportedMediums: []string{"sms", "email"},
}

response, err := sdk.Checkout(params)
if err != nil {
    // Handle error
}
// Process response
```

### Direct Payment

The `DirectPayment` method allows you to process a direct payment.

```go
params := &lakipaygosdk.DirectPaymentParams{
    Amount:      1.4,
    CallbackURL: "https://webhook.site/callback",
    Currency:    "ETB",
    Description: "Test payment",
    Details: lakipaygosdk.DirectDetails{
        Phone: "251712345678",
    },
    Medium:      "MPESA",
    PhoneNumber: "251711111111",
    Redirects: lakipaygosdk.Redirects{
        Failed:  "https://webhook.site/failed",
        Success: "https://webhook.site/success",
    },
    Reference: "gibberish",
}

response, err := sdk.DirectPayment(params)
if err != nil {
    // Handle error
}
// Process response
```

### Withdrawal

The `WithDrawal` method allows you to process a withdrawal.

```go
params := &lakipaygosdk.WithDrawalParams{
    Amount:      100,
    CallbackURL: "https://example.com/callback",
    Currency:    "NGN",
    Medium:      "sms",
    PhoneNumber: "",
    Reference:   "Test payment",
}

response, err := sdk.WithDrawal(params)
if err != nil {
    // Handle error
}
// Process response
```

### Get Transaction Details

The `GetTransactionDetails` method allows you to retrieve details of a specific transaction.

```go
transactionID := "YOUR_TRANSACTION_ID"
response, err := sdk.GetTransactionDetails(transactionID)
if err != nil {
    // Handle error
}
// Process response
```

## Webhooks

Webhooks allow you to receive real-time notifications about the status of your transactions. You can specify a `callback_url` in your requests to receive these notifications.

The following is the structure of a webhook request:

```go
type WebhookRequest struct {
    CallbackURL   string `json:"callback_url"`
    MerchantID    string `json:"merchant_id"`
    Message       string `json:"message"`
    ProviderData  string `json:"provider_data"`
    ProviderTxID  string `json:"provider_tx_id"`
    Status        string `json:"status"`
    Timestamp     string `json:"timestamp"`
    TransactionID string `json:"transaction_id"`
    Type          string `json:"type"`
}
```

## Data Models

The SDK uses the following data models for requests and responses.

### Request Parameters

- `CheckoutParams`
- `DirectPaymentParams`
- `WithDrawalParams`

### Response Objects

- `APIResponse`
- `TransactionResponse`

For detailed information on the fields for each model, please refer to the `models.go` file.
