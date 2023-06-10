# WebhookRequest

Webhook is a special type of a Job, it performs a HTTPS request to a specified URL with specified headers.


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Disabled`                                                    | **bool*                                                       | :heavy_minus_sign:                                            | Defines whether this job is currently disabled.               |
| `Headers`                                                     | map[string]*interface{}*                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `ID`                                                          | **string*                                                     | :heavy_minus_sign:                                            | Job ID, unique within the Organization                        |
| `Payload`                                                     | map[string]*interface{}*                                      | :heavy_minus_sign:                                            | N/A                                                           |
| `Triggers`                                                    | [][EventBaseRequest](../../models/shared/eventbaserequest.md) | :heavy_minus_sign:                                            | A list of Events by which the Job is triggered                |
| `URL`                                                         | **string*                                                     | :heavy_minus_sign:                                            | Thw webhook's URL (without protocol, only HTTPS is supported) |