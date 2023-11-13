# Config
(*Config*)

### Available Operations

* [SubscribeToWebhooks](#subscribetowebhooks) - Subscribe to webhooks.

## SubscribeToWebhooks

Subscribe to webhooks.

### Example Usage

```go
package main

import(
	"context"
	"log"
	trebllesamplesdk "github.com/speakeasy-sdks/Treblle-sample-sdk"
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/models/operations"
)

func main() {
    s := trebllesamplesdk.New(
        trebllesamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.Config.SubscribeToWebhooks(ctx, []operations.RequestBody{
        operations.RequestBody{},
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]operations.RequestBody](../../.md)                 | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.SubscribeToWebhooksResponse](../../pkg/models/operations/subscribetowebhooksresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.APIError | 5XX                | application/json   |
| sdkerrors.SDKError | 400-600            | */*                |
