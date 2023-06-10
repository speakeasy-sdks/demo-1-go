# PostOrgsOrgIDRegistriesResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `ErrorInfoResponse`                                                   | [*shared.ErrorInfoResponse](../../models/shared/errorinforesponse.md) | :heavy_minus_sign:                                                    | Request parameters or payload are incomplete or invalid.<br/><br/>    |
| `RegistryResponse`                                                    | [*shared.RegistryResponse](../../models/shared/registryresponse.md)   | :heavy_minus_sign:                                                    | A newly created record details.<br/><br/>                             |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |