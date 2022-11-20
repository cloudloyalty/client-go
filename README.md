MAXMA Client
============

Usage example:
```go
package example

import "github.com/cloudloyalty/client-go"

client := cloudloyalty_client.New(&cloudloyalty_client.Config{
    BaseURL:       "https://api-test.maxma.com",
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

See https://docs.maxma.com/api/