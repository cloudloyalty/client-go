CloudLoyalty Client
===================

Usage:
```go
import "github.com/cloudloyalty/client-go"

client := cloudloyalty_client.New(&cloudloyalty_client.Config{
    BaseURL:       config.CLBaseUrl,
    ProcessingKey: "your key",
})

clientQuery := &cloudloyalty_client.GetBalanceQuery{
    PhoneNumber: "phone number",
    Card:        "card",
}

resp, err := client.GetBalance(ctx, clientQuery)
if resp != nil {
    // resp.Client.FullName
}
```

See www.cloudloyalty.ru.