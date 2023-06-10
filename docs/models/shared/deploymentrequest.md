# DeploymentRequest

Deployments represent updates to the running state of an Environment.

Deployments are made by applying _Deltas_ to a state defined by an existing Deployment. The Environment’s from_deploy property defines the Deployment. This Deployment is usually but not always in the current Environment. If the Deployment is from another Environment, the state of that Environment will be "cloned" into the current Environment with the option to apply a Delta.


## Fields

| Field                                                                                             | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `Comment`                                                                                         | **string*                                                                                         | :heavy_minus_sign:                                                                                | An optional comment to help communicate the purpose of the Deployment.                            |
| `DeltaID`                                                                                         | **string*                                                                                         | :heavy_minus_sign:                                                                                | ID of the Deployment Delta describing the changes to the current Environment for this Deployment. |
| `ValueSetVersionID`                                                                               | **string*                                                                                         | :heavy_minus_sign:                                                                                | ID of the Value Set Version describe the values to be used for this Deployment.                   |