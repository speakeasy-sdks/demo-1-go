# PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvIDRequest


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `RequestBody`                                                            | *string*                                                                 | :heavy_check_mark:                                                       | The new Environment ID. (NOTE: The string must still be JSON encoded.)<br/><br/> |
| `AppID`                                                                  | *string*                                                                 | :heavy_check_mark:                                                       | The Application ID.<br/><br/>                                            |
| `DeltaID`                                                                | *string*                                                                 | :heavy_check_mark:                                                       | ID of the Deployment Delta.<br/><br/>                                    |
| `OrgID`                                                                  | *string*                                                                 | :heavy_check_mark:                                                       | The Organization ID.<br/><br/>                                           |