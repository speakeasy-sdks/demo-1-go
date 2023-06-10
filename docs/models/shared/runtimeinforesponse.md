# RuntimeInfoResponse

RuntimeInfo object returned by the runtime endpoint. Represents a list post statuses grouped by modules and controllers (deployments and stateful sets).


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Modules`                                                                             | map[string]map[string][ControllerResponse](../../models/shared/controllerresponse.md) | :heavy_check_mark:                                                                    | N/A                                                                                   |
| `Namespace`                                                                           | *string*                                                                              | :heavy_check_mark:                                                                    | N/A                                                                                   |