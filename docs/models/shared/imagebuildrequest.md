# ImageBuildRequest

DEPRECATED: This type exists for historical compatibility and should not be used. Please use the [Artefact API](https://api-docs.humanitec.com/#tag/Artefact) instead.

Holds the metadata associated withe a Container Image Build


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Branch`                                                              | **string*                                                             | :heavy_minus_sign:                                                    | The branch name of the branch the build was built on                  |
| `Commit`                                                              | **string*                                                             | :heavy_minus_sign:                                                    | The commit ID that this build was built from.                         |
| `Image`                                                               | **string*                                                             | :heavy_minus_sign:                                                    | The fully qualified Image URL including registry, repository and tag. |
| `Tags`                                                                | []*string*                                                            | :heavy_minus_sign:                                                    | The tag that the build was built from.                                |