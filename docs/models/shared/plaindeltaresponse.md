# PlainDeltaResponse

Similar to the delta response, except the id and metadata properties.

**Basic Structure**

```
 {
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
| `Modules`                                                             | [ModuleDeltasResponse](../../models/shared/moduledeltasresponse.md)   | :heavy_check_mark:                                                    | ModuleDeltas groups the different operations together.                |
| `Shared`                                                              | [][UpdateActionResponse](../../models/shared/updateactionresponse.md) | :heavy_check_mark:                                                    | N/A                                                                   |