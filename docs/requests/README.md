# Requests

## Overview

REST APIs for retrieving request information

### Available Operations

* [GenerateRequestPostmanCollection](#generaterequestpostmancollection) - Generate a Postman collection for a particular request.
* [GetRequestFromEventLog](#getrequestfromeventlog) - Get information about a particular request.
* [QueryEventLog](#queryeventlog) - Query the event log to retrieve a list of requests.

## GenerateRequestPostmanCollection

Generates a Postman collection for a particular request. 
Allowing it to be replayed with the same inputs that were captured by the SDK.

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
    req := operations.GenerateRequestPostmanCollectionRequest{
        RequestID: "consequuntur",
    }

    res, err := s.Requests.GenerateRequestPostmanCollection(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.PostmanCollection != nil {
        // handle response
    }
}
```

## GetRequestFromEventLog

Get information about a particular request.

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
    req := operations.GetRequestFromEventLogRequest{
        RequestID: "repellat",
    }

    res, err := s.Requests.GetRequestFromEventLog(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.UnboundedRequest != nil {
        // handle response
    }
}
```

## QueryEventLog

Supports retrieving a list of request captured by the SDK for this workspace.
Allows the filtering of requests on a number of criteria such as ApiID, VersionID, Path, Method, etc.

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
    req := operations.QueryEventLogRequest{
        Filters: &shared.Filters{
            Filters: []shared.Filter{
                shared.Filter{
                    Key: "occaecati",
                    Operator: "numquam",
                    Value: "commodi",
                },
                shared.Filter{
                    Key: "quam",
                    Operator: "molestiae",
                    Value: "velit",
                },
                shared.Filter{
                    Key: "error",
                    Operator: "quia",
                    Value: "quis",
                },
            },
            Limit: 110375,
            Offset: 674752,
            Operator: "animi",
        },
    }

    res, err := s.Requests.QueryEventLog(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BoundedRequests != nil {
        // handle response
    }
}
```
