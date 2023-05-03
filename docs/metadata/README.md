# Metadata

## Overview

REST APIs for managing Version Metadata entities

### Available Operations

* [DeleteVersionMetadata](#deleteversionmetadata) - Delete metadata for a particular apiID and versionID.
* [GetVersionMetadata](#getversionmetadata) - Get all metadata for a particular apiID and versionID.
* [InsertVersionMetadata](#insertversionmetadata) - Insert metadata for a particular apiID and versionID.

## DeleteVersionMetadata

Delete metadata for a particular apiID and versionID.

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
    res, err := s.Metadata.DeleteVersionMetadata(ctx, operations.DeleteVersionMetadataRequest{
        APIID: "excepturi",
        MetaKey: "accusantium",
        MetaValue: "iure",
        VersionID: "culpa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetVersionMetadata

Get all metadata for a particular apiID and versionID.

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
    res, err := s.Metadata.GetVersionMetadata(ctx, operations.GetVersionMetadataRequest{
        APIID: "doloribus",
        VersionID: "sapiente",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionMetadata != nil {
        // handle response
    }
}
```

## InsertVersionMetadata

Insert metadata for a particular apiID and versionID.

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
    res, err := s.Metadata.InsertVersionMetadata(ctx, operations.InsertVersionMetadataRequest{
        VersionMetadataInput: shared.VersionMetadataInput{
            MetaKey: "architecto",
            MetaValue: "mollitia",
        },
        APIID: "dolorem",
        VersionID: "culpa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VersionMetadata != nil {
        // handle response
    }
}
```
