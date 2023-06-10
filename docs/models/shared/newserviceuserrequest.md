# NewServiceUserRequest

NewServiceUser holds the definition of a new service user.


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `Email`                                                                             | **string*                                                                           | :heavy_minus_sign:                                                                  | The email address that should get notifications about this service user. (Optional) |
| `Name`                                                                              | **string*                                                                           | :heavy_minus_sign:                                                                  | The name that should be shown for this service user.                                |
| `Role`                                                                              | **string*                                                                           | :heavy_minus_sign:                                                                  | The role that the service user should have on the organization it is created in     |