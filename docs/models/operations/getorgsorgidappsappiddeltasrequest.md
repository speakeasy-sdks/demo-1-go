# GetOrgsOrgIDAppsAppIDDeltasRequest


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `AppID`                                                         | *string*                                                        | :heavy_check_mark:                                              | The Application ID.<br/><br/>                                   |
| `Archived`                                                      | **bool*                                                         | :heavy_minus_sign:                                              | If true, return archived Deltas.<br/><br/>                      |
| `Env`                                                           | **string*                                                       | :heavy_minus_sign:                                              | Only return Deltas associated with the specified Environment.<br/><br/> |
| `OrgID`                                                         | *string*                                                        | :heavy_check_mark:                                              | The Organization ID.<br/><br/>                                  |