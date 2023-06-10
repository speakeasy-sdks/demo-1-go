# ModuleDeltasRequest

ModuleDeltas groups the different operations together.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Add`                                                                               | map[string]map[string][ControllerRequest](../../models/shared/controllerrequest.md) | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `Remove`                                                                            | []*string*                                                                          | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `Update`                                                                            | map[string][][UpdateActionRequest](../../models/shared/updateactionrequest.md)      | :heavy_minus_sign:                                                                  | N/A                                                                                 |