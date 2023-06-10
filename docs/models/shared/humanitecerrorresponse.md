# HumanitecErrorResponse

HumanitecError represents a standard Humanitec Error


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `Details`                                           | map[string]*interface{}*                            | :heavy_minus_sign:                                  | (Optional) Additional information is enclosed here. |
| `Error`                                             | *string*                                            | :heavy_check_mark:                                  | A short code to help with error identification.     |
| `Message`                                           | *string*                                            | :heavy_check_mark:                                  | A Human readable message about the error.           |