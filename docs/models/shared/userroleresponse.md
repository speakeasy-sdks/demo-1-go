# UserRoleResponse

UserRole holds the mapping of role to user for a particular object.


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `CreatedAt`                                               | *string*                                                  | :heavy_check_mark:                                        | The time the user was first registered with Humanitec     | 2020-06-22T09:37:23.523Z                                  |
| `Email`                                                   | **string*                                                 | :heavy_minus_sign:                                        | The email address of the user from the profile            |                                                           |
| `ID`                                                      | *string*                                                  | :heavy_check_mark:                                        | The User ID for this user                                 |                                                           |
| `Invite`                                                  | **string*                                                 | :heavy_minus_sign:                                        | The status of an invitation (If applicable)               |                                                           |
| `Name`                                                    | *string*                                                  | :heavy_check_mark:                                        | The name the user goes by                                 |                                                           |
| `Role`                                                    | *string*                                                  | :heavy_check_mark:                                        | The role that this user holds                             |                                                           |
| `Type`                                                    | *string*                                                  | :heavy_check_mark:                                        | The type of the account. Could be user, service or system |                                                           |
| `User`                                                    | **string*                                                 | :heavy_minus_sign:                                        | The user ID that hold the role                            |                                                           |