# PostOrgsOrgIDAppsAppIDWebhooksJobIDResponse


## Fields

| Field                                                             | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ContentType`                                                     | *string*                                                          | :heavy_check_mark:                                                | N/A                                                               |
| `StatusCode`                                                      | *int*                                                             | :heavy_check_mark:                                                | N/A                                                               |
| `RawResponse`                                                     | [*http.Response](https://pkg.go.dev/net/http#Response)            | :heavy_minus_sign:                                                | N/A                                                               |
| `WebhookResponse`                                                 | [*shared.WebhookResponse](../../models/shared/webhookresponse.md) | :heavy_minus_sign:                                                | Updated successfully, return the webhook<br/><br/>                |