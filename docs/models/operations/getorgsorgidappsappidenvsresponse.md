# GetOrgsOrgIDAppsAppIDEnvsResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ContentType`                                                              | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `EnvironmentResponses`                                                     | [][shared.EnvironmentResponse](../../models/shared/environmentresponse.md) | :heavy_minus_sign:                                                         | A list of Environments associated with an Application.<br/><br/>           |
| `StatusCode`                                                               | *int*                                                                      | :heavy_check_mark:                                                         | N/A                                                                        |
| `RawResponse`                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                     | :heavy_minus_sign:                                                         | N/A                                                                        |