# SetResponse

A Deployment Set (or just "Set") defines all of the non-Environment specific configuration for Modules and External Resources. Each of these Modules or External Resources has a unique name.

Deployment Sets are immutable and their ID is a cryptographic hash of their content. This means that a Deployment Set can be globally identified based on its ID. It also means that referencing a Deployment Set by ID will always return the same Deployment Set.

Deployment Sets cannot be created directly, instead they are created by applying a Deployment Delta to an existing Deployment Set.

**Basic Structure**

```
 {
   "id": <ID of the Deployment Set>,
   "modules" : {
     <ID of Module> : {
       "profile": <Defines how the optional "spec" property is interpreted>
       "spec": {
         <Properties that depend on the "profile" property.>
       }
       "externals": {
         <External Resource Name> : {
           "type": <Resource Type>,
           "params": {
             <Properties which parametrize the resource depending on the Resource Type.>
           }
         }
       }
     }
   }
 }
```

For details about how the Humanitec provided profiles work, see (Deployment Set Profiles)[].


## Fields

| Field                                                                                                                    | Type                                                                                                                     | Required                                                                                                                 | Description                                                                                                              |
| ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| `ID`                                                                                                                     | *string*                                                                                                                 | :heavy_check_mark:                                                                                                       | The ID which is a hash of the content of the Deployment Set.                                                             |
| `Modules`                                                                                                                | map[string]map[string][ControllerResponse](../../models/shared/controllerresponse.md)                                    | :heavy_check_mark:                                                                                                       | The Modules that make up the Set                                                                                         |
| `Shared`                                                                                                                 | map[string]*interface{}*                                                                                                 | :heavy_check_mark:                                                                                                       | Resources that are shared across the set                                                                                 |
| `Version`                                                                                                                | *int64*                                                                                                                  | :heavy_check_mark:                                                                                                       | The version of the Deployment Set Schema to use. (Currently, only 0 is supported, and if omitted, version 0 is assumed.) |