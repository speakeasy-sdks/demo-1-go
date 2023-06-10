# PostOrgsOrgIDAppsAppIDWebhooksResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `ErrorInfoResponse`                                                   | [*shared.ErrorInfoResponse](../../models/shared/errorinforesponse.md) | :heavy_minus_sign:                                                    | Some of the values supplied are invalid.<br/><br/>                    |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |
| `WebhookResponse`                                                     | [*shared.WebhookResponse](../../models/shared/webhookresponse.md)     | :heavy_minus_sign:                                                    | A created webhook.<br/><br/>                                          |