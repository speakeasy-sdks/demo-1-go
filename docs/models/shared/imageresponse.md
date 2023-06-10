# ImageResponse

DEPRECATED: This type exists for historical compatibility and should not be used. Please use the [Artefact API](https://api-docs.humanitec.com/#tag/Artefact) instead.

Container Images (known simply as Images) can be registered with Humanitec. Continuous Integration (CI) pipelines can then notify Humanitec when a new build of a Container Image becomes available. Humanitec tracks the Image along with metadata about how it was built.


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `AddedAt`                                                            | *string*                                                             | :heavy_check_mark:                                                   | The time the first build of this Image was added to the organization | 2020-06-22T09:37:23.523Z                                             |
| `Builds`                                                             | [][ImageBuildResponse](../../models/shared/imagebuildresponse.md)    | :heavy_check_mark:                                                   | A list of Image Builds ordered by addition date.                     |                                                                      |
| `ID`                                                                 | *string*                                                             | :heavy_check_mark:                                                   | The ID used to group different builds of the same Image together.    |                                                                      |
| `Source`                                                             | *string*                                                             | :heavy_check_mark:                                                   | The Image Source that this Image is added via                        |                                                                      |