# UserInviteResponse

UserInvite stores the invitation details.


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     | Example                                         |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `CreatedAt`                                     | *string*                                        | :heavy_check_mark:                              | The timestamp this invitation was created.      | 2020-06-22T09:37:23.523Z                        |
| `CreatedBy`                                     | *string*                                        | :heavy_check_mark:                              | The ID of the user who created this invitation. |                                                 |
| `Email`                                         | **string*                                       | :heavy_minus_sign:                              | The email address of the user from the profile. |                                                 |
| `ExpiresAt`                                     | *string*                                        | :heavy_check_mark:                              | The timestamp this invitation would expire.     | 2020-06-22T09:37:23.523Z                        |
| `UserID`                                        | *string*                                        | :heavy_check_mark:                              | The User ID for this user.                      |                                                 |