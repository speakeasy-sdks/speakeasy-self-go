<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"registry"
	"registry/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Apis.GetApis(ctx, operations.GetApisRequest{
        Metadata: map[string][]string{
            "provident": []string{
                "quibusdam",
                "unde",
                "nulla",
            },
            "corrupti": []string{
                "vel",
                "error",
                "deserunt",
                "suscipit",
            },
            "iure": []string{
                "debitis",
                "ipsa",
            },
        },
        Op: &operations.GetApisOp{
            And: false,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Apis != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->