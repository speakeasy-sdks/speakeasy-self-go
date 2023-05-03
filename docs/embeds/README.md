# Embeds

## Overview

REST APIs for managing embeds

### Available Operations

* [GetEmbedAccessToken](#getembedaccesstoken) - Get an embed access token for the current workspace.
* [GetValidEmbedAccessTokens](#getvalidembedaccesstokens) - Get all valid embed access tokens for the current workspace.
* [RevokeEmbedAccessToken](#revokeembedaccesstoken) - Revoke an embed access EmbedToken.

## GetEmbedAccessToken

Returns an embed access token for the current workspace. This can be used to authenticate access to externally embedded content.
Filters can be applied allowing views to be filtered to things like particular customerIds.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"registry"
	"registry/pkg/models/operations"
	"registry/pkg/models/shared"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Embeds.GetEmbedAccessToken(ctx, operations.GetEmbedAccessTokenRequest{
        Description: sdk.String("laborum"),
        Duration: sdk.Int64(170909),
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "corporis",
                    Operator: "explicabo",
                    Value: "nobis",
                },
            },
            Limit: 315428,
            Offset: 607831,
            Operator: "nemo",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EmbedAccessTokenResponse != nil {
        // handle response
    }
}
```

## GetValidEmbedAccessTokens

Get all valid embed access tokens for the current workspace.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"registry"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            APIKey: "YOUR_API_KEY_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Embeds.GetValidEmbedAccessTokens(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.EmbedTokens != nil {
        // handle response
    }
}
```

## RevokeEmbedAccessToken

Revoke an embed access EmbedToken.

### Example Usage

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
    res, err := s.Embeds.RevokeEmbedAccessToken(ctx, operations.RevokeEmbedAccessTokenRequest{
        TokenID: "minima",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
