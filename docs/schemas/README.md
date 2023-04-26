# Schemas

## Overview

REST APIs for managing Schema entities

### Available Operations

* [DeleteSchema](#deleteschema) - Delete a particular schema revision for an Api.
* [DownloadSchema](#downloadschema) - Download the latest schema for a particular apiID.
* [DownloadSchemaRevision](#downloadschemarevision) - Download a particular schema revision for an Api.
* [GetSchema](#getschema) - Get information about the latest schema.
* [GetSchemaDiff](#getschemadiff) - Get a diff of two schema revisions for an Api.
* [GetSchemaRevision](#getschemarevision) - Get information about a particular schema revision for an Api.
* [GetSchemas](#getschemas) - Get information about all schemas associated with a particular apiID.
* [RegisterSchema](#registerschema) - Register a schema.

## DeleteSchema

Delete a particular schema revision for an Api.

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
    req := operations.DeleteSchemaRequest{
        APIID: "enim",
        RevisionID: "odit",
        VersionID: "quo",
    }

    res, err := s.Schemas.DeleteSchema(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## DownloadSchema

Download the latest schema for a particular apiID.

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
    req := operations.DownloadSchemaRequest{
        APIID: "sequi",
        VersionID: "tenetur",
    }

    res, err := s.Schemas.DownloadSchema(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## DownloadSchemaRevision

Download a particular schema revision for an Api.

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
    req := operations.DownloadSchemaRevisionRequest{
        APIID: "ipsam",
        RevisionID: "id",
        VersionID: "possimus",
    }

    res, err := s.Schemas.DownloadSchemaRevision(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## GetSchema

Returns information about the last uploaded schema for a particular API version. 
This won't include the schema itself, that can be retrieved via the downloadSchema operation.

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
    req := operations.GetSchemaRequest{
        APIID: "aut",
        VersionID: "quasi",
    }

    res, err := s.Schemas.GetSchema(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## GetSchemaDiff

Get a diff of two schema revisions for an Api.

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
    req := operations.GetSchemaDiffRequest{
        APIID: "error",
        BaseRevisionID: "temporibus",
        TargetRevisionID: "laborum",
        VersionID: "quasi",
    }

    res, err := s.Schemas.GetSchemaDiff(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.SchemaDiff != nil {
        // handle response
    }
}
```

## GetSchemaRevision

Returns information about the last uploaded schema for a particular schema revision. 
This won't include the schema itself, that can be retrieved via the downloadSchema operation.

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
    req := operations.GetSchemaRevisionRequest{
        APIID: "reiciendis",
        RevisionID: "voluptatibus",
        VersionID: "vero",
    }

    res, err := s.Schemas.GetSchemaRevision(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schema != nil {
        // handle response
    }
}
```

## GetSchemas

Returns information the schemas associated with a particular apiID. 
This won't include the schemas themselves, they can be retrieved via the downloadSchema operation.

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
    req := operations.GetSchemasRequest{
        APIID: "nihil",
        VersionID: "praesentium",
    }

    res, err := s.Schemas.GetSchemas(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Schemata != nil {
        // handle response
    }
}
```

## RegisterSchema

Allows uploading a schema for a particular API version.
This will be used to populate ApiEndpoints and used as a base for any schema generation if present.

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
    req := operations.RegisterSchemaRequest{
        RequestBody: operations.RegisterSchemaRequestBody{
            File: operations.RegisterSchemaRequestBodyFile{
                Content: []byte("voluptatibus"),
                File: "ipsa",
            },
        },
        APIID: "omnis",
        VersionID: "voluptate",
    }

    res, err := s.Schemas.RegisterSchema(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
