# UserInvite

### Available Operations

* [GetOrgsOrgIDInvitations](#getorgsorgidinvitations) - List the invites issued for the organization.

## GetOrgsOrgIDInvitations

List the invites issued for the organization.

### Example Usage

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
    res, err := s.UserInvite.GetOrgsOrgIDInvitations(ctx, operations.GetOrgsOrgIDInvitationsRequest{
        OrgID: "vel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UserInviteResponses != nil {
        // handle response
    }
}
```
