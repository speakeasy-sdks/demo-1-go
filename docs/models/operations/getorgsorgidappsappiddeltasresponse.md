# GetOrgsOrgIDAppsAppIDDeltasResponse


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ContentType`                                                  | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `DeltaResponses`                                               | [][shared.DeltaResponse](../../models/shared/deltaresponse.md) | :heavy_minus_sign:                                             | A possibly empty list of Deployment Deltas.<br/><br/>          |
| `StatusCode`                                                   | *int*                                                          | :heavy_check_mark:                                             | N/A                                                            |
| `RawResponse`                                                  | [*http.Response](https://pkg.go.dev/net/http#Response)         | :heavy_minus_sign:                                             | N/A                                                            |