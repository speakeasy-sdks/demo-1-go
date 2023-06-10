# GetOrgsOrgIDImagesResponse


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ContentType`                                                  | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `ImageResponses`                                               | [][shared.ImageResponse](../../models/shared/imageresponse.md) | :heavy_minus_sign:                                             | Possibly empty list of Container Images.<br/><br/>             |
| `StatusCode`                                                   | *int*                                                          | :heavy_check_mark:                                             | N/A                                                            |
| `RawResponse`                                                  | [*http.Response](https://pkg.go.dev/net/http#Response)         | :heavy_minus_sign:                                             | N/A                                                            |