# ModuleDeltasResponse

ModuleDeltas groups the different operations together.


## Fields

| Field                                                                                 | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `Add`                                                                                 | map[string]map[string][ControllerResponse](../../models/shared/controllerresponse.md) | :heavy_check_mark:                                                                    | N/A                                                                                   |
| `Remove`                                                                              | []*string*                                                                            | :heavy_check_mark:                                                                    | N/A                                                                                   |
| `Update`                                                                              | map[string][][UpdateActionResponse](../../models/shared/updateactionresponse.md)      | :heavy_check_mark:                                                                    | N/A                                                                                   |