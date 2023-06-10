# GetSetsResponse


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ContentType`                                              | *string*                                                   | :heavy_check_mark:                                         | N/A                                                        |
| `SetResponses`                                             | [][shared.SetResponse](../../models/shared/setresponse.md) | :heavy_minus_sign:                                         | The Requested Deployment Set.<br/><br/>                    |
| `StatusCode`                                               | *int*                                                      | :heavy_check_mark:                                         | N/A                                                        |
| `RawResponse`                                              | [*http.Response](https://pkg.go.dev/net/http#Response)     | :heavy_minus_sign:                                         | N/A                                                        |
| `GetSets404ApplicationJSONString`                          | **string*                                                  | :heavy_minus_sign:                                         | No Deployment Set with ID `setId` found in Application.<br/><br/> |