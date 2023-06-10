# GetOrgsOrgIDInvitationsResponse


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ContentType`                                                            | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `StatusCode`                                                             | *int*                                                                    | :heavy_check_mark:                                                       | N/A                                                                      |
| `RawResponse`                                                            | [*http.Response](https://pkg.go.dev/net/http#Response)                   | :heavy_minus_sign:                                                       | N/A                                                                      |
| `UserInviteResponses`                                                    | [][shared.UserInviteResponse](../../models/shared/userinviteresponse.md) | :heavy_minus_sign:                                                       | The list of the invites issued for the organization.<br/><br/>           |