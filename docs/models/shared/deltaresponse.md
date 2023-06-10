# DeltaResponse

A Deployment Delta (or just "Delta") describes the changes that must be applied to one Deployment Set to generate another Deployment Set. Deployment Deltas are the only way to create new Deployment Sets.

Deployment Deltas can be created fully formed or combined together via PATCHing. They can also be generated from the difference between two Deployment Sets.

**Basic Structure**

```
 {
   "id": <ID of the Deployment Delta.>,
   "metadata": {
     <Properties such as who created the Delta, which Environment it is associated to and a Human-friendly name>
   }
   "modules" : {
     "add" : {
       <ID of Module to add to the Deployment Set> : {
         <An entire Modules object>
       }
     },
     "remove": [
       <An array of Module IDs that should be removed from the Deployment Set>
     ],
    "update": {
       <ID of Module already in the Set to be updated> : [
         <An array of JSON Patch (Search Results (RFC 6902) objects scoped to the module>
       ]
     }
   }
 }
```


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `ID`                                                                  | *string*                                                              | :heavy_check_mark:                                                    | A unique ID for the Delta                                             |
| `Metadata`                                                            | [DeltaMetadataResponse](../../models/shared/deltametadataresponse.md) | :heavy_check_mark:                                                    | N/A                                                                   |
| `Modules`                                                             | [ModuleDeltasResponse](../../models/shared/moduledeltasresponse.md)   | :heavy_check_mark:                                                    | ModuleDeltas groups the different operations together.                |
| `Shared`                                                              | [][UpdateActionResponse](../../models/shared/updateactionresponse.md) | :heavy_check_mark:                                                    | N/A                                                                   |