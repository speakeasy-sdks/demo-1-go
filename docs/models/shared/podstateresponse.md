# PodStateResponse

PodState represents single pod status


## Fields

| Field                      | Type                       | Required                   | Description                |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `ContainerStatuses`        | []map[string]*interface{}* | :heavy_check_mark:         | N/A                        |
| `Phase`                    | *string*                   | :heavy_check_mark:         | N/A                        |
| `PodName`                  | *string*                   | :heavy_check_mark:         | N/A                        |
| `Revision`                 | *int64*                    | :heavy_check_mark:         | N/A                        |
| `Status`                   | *string*                   | :heavy_check_mark:         | N/A                        |