# CreateResourceAccountRequestRequest

CreateResourceAccountRequest describes the request to create a new security account.


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `Credentials`                                                                   | map[string]*interface{}*                                                        | :heavy_minus_sign:                                                              | Credentials associated with the account.                                        |
| `ID`                                                                            | **string*                                                                       | :heavy_minus_sign:                                                              | Unique identifier for the account (in scope of the organization it belongs to). |
| `Name`                                                                          | **string*                                                                       | :heavy_minus_sign:                                                              | Display name.                                                                   |
| `Type`                                                                          | **string*                                                                       | :heavy_minus_sign:                                                              | The type of the account                                                         |