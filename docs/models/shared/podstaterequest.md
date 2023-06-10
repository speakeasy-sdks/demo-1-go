# PodStateRequest

PodState represents single pod status


## Fields

| Field                      | Type                       | Required                   | Description                |
| -------------------------- | -------------------------- | -------------------------- | -------------------------- |
| `ContainerStatuses`        | []map[string]*interface{}* | :heavy_minus_sign:         | N/A                        |
| `Phase`                    | **string*                  | :heavy_minus_sign:         | N/A                        |
| `PodName`                  | **string*                  | :heavy_minus_sign:         | N/A                        |
| `Revision`                 | **int64*                   | :heavy_minus_sign:         | N/A                        |
| `Status`                   | **string*                  | :heavy_minus_sign:         | N/A                        |