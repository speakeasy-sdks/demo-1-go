# GetOrgsOrgIDAppsResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ApplicationResponses`                                                     | [][shared.ApplicationResponse](../../models/shared/applicationresponse.md) | :heavy_minus_sign:                                                         | Possibly empty list of Applications.                                       |
| `ContentType`                                                              | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `StatusCode`                                                               | *int*                                                                      | :heavy_check_mark:                                                         | N/A                                                                        |
| `RawResponse`                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                     | :heavy_minus_sign:                                                         | N/A                                                                        |