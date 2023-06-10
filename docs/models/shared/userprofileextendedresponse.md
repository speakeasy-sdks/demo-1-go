# UserProfileExtendedResponse

UserProfileExtended holds the profile information of a user including properties only accessible to the user.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `CreatedAt`                                               | *string*                                                  | :heavy_check_mark:                                        | The time the user was first registered with Humanitec     | 2020-06-22T09:37:23.523Z                                  |
| `Email`                                                   | **string*                                                 | :heavy_minus_sign:                                        | The email address of the user from the profile            |                                                           |
| `ID`                                                      | *string*                                                  | :heavy_check_mark:                                        | The User ID for this user                                 |                                                           |
| `Name`                                                    | *string*                                                  | :heavy_check_mark:                                        | The name the user goes by                                 |                                                           |
| `Properties`                                              | map[string]*interface{}*                                  | :heavy_check_mark:                                        | N/A                                                       |                                                           |
| `Roles`                                                   | map[string]*string*                                       | :heavy_check_mark:                                        | N/A                                                       |                                                           |
| `Type`                                                    | *string*                                                  | :heavy_check_mark:                                        | The type of the account. Could be user, service or system |                                                           |