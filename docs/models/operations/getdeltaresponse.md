# GetDeltaResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `DeltaResponse`                                               | [*shared.DeltaResponse](../../models/shared/deltaresponse.md) | :heavy_minus_sign:                                            | The requested Deployment Delta.                               |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |
| `GetDelta404ApplicationJSONString`                            | **string*                                                     | :heavy_minus_sign:                                            | No Deployment Delta with ID `deltaId` found in Application.<br/><br/> |