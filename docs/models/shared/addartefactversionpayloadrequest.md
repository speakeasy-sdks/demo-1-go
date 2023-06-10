# AddArtefactVersionPayloadRequest

AddArtefactVersionPayload describes the payload for a new ArtefactVersion request.


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `Commit`                                                    | **string*                                                   | :heavy_minus_sign:                                          | (Optional) The commit ID the Artefact Version was built on. |
| `Digest`                                                    | **string*                                                   | :heavy_minus_sign:                                          | (Optional) The Artefact Version digest.                     |
| `Name`                                                      | *string*                                                    | :heavy_check_mark:                                          | The Artefact name.                                          |
| `Ref`                                                       | **string*                                                   | :heavy_minus_sign:                                          | (Optional) The ref the Artefact Version was built from.     |
| `Type`                                                      | *string*                                                    | :heavy_check_mark:                                          | The Artefact Version type.                                  |
| `Version`                                                   | **string*                                                   | :heavy_minus_sign:                                          | (Optional) The Artefact Version.                            |