# ClusterSecretRequest

ClusterSecret represents Kubernetes secret reference.


## Fields

| Field                                                      | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `Namespace`                                                | *string*                                                   | :heavy_check_mark:                                         | Namespace to look for the Kubernetes secret definition in. |
| `Secret`                                                   | *string*                                                   | :heavy_check_mark:                                         | Name that identifies the Kubernetes secret.                |