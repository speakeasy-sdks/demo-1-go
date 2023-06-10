# GetOrgsResponse


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ContentType`                                                                 | *string*                                                                      | :heavy_check_mark:                                                            | N/A                                                                           |
| `OrganizationResponses`                                                       | [][shared.OrganizationResponse](../../models/shared/organizationresponse.md)  | :heavy_minus_sign:                                                            | A list of available organizations (based on the current user access level).<br/><br/> |
| `StatusCode`                                                                  | *int*                                                                         | :heavy_check_mark:                                                            | N/A                                                                           |
| `RawResponse`                                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)                        | :heavy_minus_sign:                                                            | N/A                                                                           |