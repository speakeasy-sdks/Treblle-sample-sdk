<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	trebllesamplesdk "github.com/speakeasy-sdks/Treblle-sample-sdk"
	"github.com/speakeasy-sdks/Treblle-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := trebllesamplesdk.New(
		trebllesamplesdk.WithSecurity(""),
	)

	var drinkType *shared.DrinkType = shared.DrinkTypeSpirit

	ctx := context.Background()
	res, err := s.Drinks.ListDrinks(ctx, drinkType)
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->