# ControllerResponse

Controller represents deployment, stateful set etc


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Kind`                                                        | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `Message`                                                     | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `Pods`                                                        | [][PodStateResponse](../../models/shared/podstateresponse.md) | :heavy_check_mark:                                            | N/A                                                           |
| `Replicas`                                                    | *int64*                                                       | :heavy_check_mark:                                            | N/A                                                           |
| `Revision`                                                    | *int64*                                                       | :heavy_check_mark:                                            | N/A                                                           |
| `Status`                                                      | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |