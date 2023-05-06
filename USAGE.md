<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.AccountType.GetOrgsOrgIDResourcesAccountTypes(ctx, operations.GetOrgsOrgIDResourcesAccountTypesRequest{
        OrgID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountTypeResponses != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->