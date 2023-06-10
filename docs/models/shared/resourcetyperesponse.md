# ResourceTypeResponse

Resources Types define the technology that Applications can have dependencies on.

Each Resource Type also defines a set of input parameters (`inputs_schema`), and a set of output data (`outputs_schema`). When provisioning a resource, these are treated as inputs and outputs respectively.


## Fields

| Field                                                                              | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `Category`                                                                         | *string*                                                                           | :heavy_check_mark:                                                                 | Category name (used to group similar resources on the UI).                         |
| `InputsSchema`                                                                     | map[string]*interface{}*                                                           | :heavy_check_mark:                                                                 | A JSON Schema specifying the type-specific parameters for the driver (input).      |
| `Name`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | Display name.                                                                      |
| `OutputsSchema`                                                                    | map[string]*interface{}*                                                           | :heavy_check_mark:                                                                 | A JSON Schema specifying the type-specific data passed to the deployment (output). |
| `Type`                                                                             | *string*                                                                           | :heavy_check_mark:                                                                 | Unique resource type identifier (system-wide, across all organizations).           |