# EnvironmentDefinitionRequest

The ID, Name, Type, and Deployment the Environment will be derived from.




## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `FromDeployID`                                                               | *string*                                                                     | :heavy_check_mark:                                                           | Defines the existing Deployment the new Environment will be based on.        |
| `ID`                                                                         | *string*                                                                     | :heavy_check_mark:                                                           | The ID the Environment is referenced as.                                     |
| `Name`                                                                       | *string*                                                                     | :heavy_check_mark:                                                           | The Human-friendly name for the Environment.                                 |
| `Type`                                                                       | *string*                                                                     | :heavy_check_mark:                                                           | The Environment Type. This is used for organizing and managing Environments. |