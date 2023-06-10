# test-1

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/demo-1-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"test-1"
	"test-1/pkg/models/operations"
)

func main() {
    s := test1.New()

    ctx := context.Background()
    res, err := s.AccountType.GetOrgsOrgIDResourcesAccountTypes(ctx, operations.GetOrgsOrgIDResourcesAccountTypesRequest{
        OrgID: "corrupti",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccountTypeResponses != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AccountType](docs/sdks/accounttype/README.md)

* [GetOrgsOrgIDResourcesAccountTypes](docs/sdks/accounttype/README.md#getorgsorgidresourcesaccounttypes) - List Resource Account Types available to the organization.

### [ActiveResource](docs/sdks/activeresource/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResID](docs/sdks/activeresource/README.md#deleteorgsorgidappsappidenvsenvidresourcestyperesid) - Delete Active Resources.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDResources](docs/sdks/activeresource/README.md#getorgsorgidappsappidenvsenvidresources) - List Active Resources provisioned in an environment.
* [GetOrgsOrgIDResourcesDefsDefIDResources](docs/sdks/activeresource/README.md#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.

### [Application](docs/sdks/application/README.md)

* [DeleteOrgsOrgIDAppsAppID](docs/sdks/application/README.md#deleteorgsorgidappsappid) - Delete an Application
* [GetOrgsOrgIDApps](docs/sdks/application/README.md#getorgsorgidapps) - List all Applications in an Organization.
* [GetOrgsOrgIDAppsAppID](docs/sdks/application/README.md#getorgsorgidappsappid) - Get an existing Application
* [PostOrgsOrgIDApps](docs/sdks/application/README.md#postorgsorgidapps) - Add a new Application to an Organization

### [Artefact](docs/sdks/artefact/README.md)

* [DeleteOrgsOrgIDArtefactsArtefactID](docs/sdks/artefact/README.md#deleteorgsorgidartefactsartefactid) - Delete Artefact and all related Artefact Versions
* [GetOrgsOrgIDArtefacts](docs/sdks/artefact/README.md#getorgsorgidartefacts) - List all Artefacts.

### [ArtefactVersion](docs/sdks/artefactversion/README.md)

* [GetOrgsOrgIDArtefactVersions](docs/sdks/artefactversion/README.md#getorgsorgidartefactversions) - List all Artefacts Versions.
* [GetOrgsOrgIDArtefactVersionsArtefactVersionID](docs/sdks/artefactversion/README.md#getorgsorgidartefactversionsartefactversionid) - Get an Artefacts Versions.
* [GetOrgsOrgIDArtefactsArtefactIDVersions](docs/sdks/artefactversion/README.md#getorgsorgidartefactsartefactidversions) - List all Artefact Versions of an Artefact.
* [PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionID](docs/sdks/artefactversion/README.md#patchorgsorgidartefactsartefactidversionsversionid) - Update Version of an Artefact.
* [PostOrgsOrgIDArtefactVersions](docs/sdks/artefactversion/README.md#postorgsorgidartefactversions) - Register a new Artefact Version with your organization.

### [AutomationRule](docs/sdks/automationrule/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/sdks/automationrule/README.md#deleteorgsorgidappsappidenvsenvidrulesruleid) - Delete Automation Rule from an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/sdks/automationrule/README.md#getorgsorgidappsappidenvsenvidrules) - List all Automation Rules in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/sdks/automationrule/README.md#getorgsorgidappsappidenvsenvidrulesruleid) - Get a specific Automation Rule for an Environment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/sdks/automationrule/README.md#postorgsorgidappsappidenvsenvidrules) - Create a new Automation Rule for an Environment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/sdks/automationrule/README.md#putorgsorgidappsappidenvsenvidrulesruleid) - Update an existing Automation Rule for an Environment.

### [Delta](docs/sdks/delta/README.md)

* [GetDelta](docs/sdks/delta/README.md#getdelta) - Fetch an existing Delta
* [GetOrgsOrgIDAppsAppIDDeltas](docs/sdks/delta/README.md#getorgsorgidappsappiddeltas) - List Deltas in an Application
* [PatchOrgsOrgIDAppsAppIDDeltasDeltaID](docs/sdks/delta/README.md#patchorgsorgidappsappiddeltasdeltaid) - Update an existing Delta
* [PostOrgsOrgIDAppsAppIDDeltas](docs/sdks/delta/README.md#postorgsorgidappsappiddeltas) - Create a new Delta
* [PutDelta](docs/sdks/delta/README.md#putdelta) - Update an existing Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived](docs/sdks/delta/README.md#putorgsorgidappsappiddeltasdeltaidmetadataarchived) - Mark a Delta as "archived"
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID](docs/sdks/delta/README.md#putorgsorgidappsappiddeltasdeltaidmetadataenvid) - Change the Environment of a Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName](docs/sdks/delta/README.md#putorgsorgidappsappiddeltasdeltaidmetadataname) - Change the name of a Delta

### [Deployment](docs/sdks/deployment/README.md)

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/sdks/deployment/README.md#getorgsorgidappsappidenvsenviddeploys) - List Deployments in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployID](docs/sdks/deployment/README.md#getorgsorgidappsappidenvsenviddeploysdeployid) - Get a specific Deployment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrors](docs/sdks/deployment/README.md#getorgsorgidappsappidenvsenviddeploysdeployiderrors) - List errors that occurred in a Deployment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/sdks/deployment/README.md#postorgsorgidappsappidenvsenviddeploys) - Start a new Deployment.

### [DriverDefinition](docs/sdks/driverdefinition/README.md)

* [DeleteOrgsOrgIDResourcesDriversDriverID](docs/sdks/driverdefinition/README.md#deleteorgsorgidresourcesdriversdriverid) - Delete a Resources Driver.
* [GetOrgsOrgIDResourcesDrivers](docs/sdks/driverdefinition/README.md#getorgsorgidresourcesdrivers) - List Resource Drivers.
* [GetOrgsOrgIDResourcesDriversDriverID](docs/sdks/driverdefinition/README.md#getorgsorgidresourcesdriversdriverid) - Get a Resource Driver.
* [PostOrgsOrgIDResourcesDrivers](docs/sdks/driverdefinition/README.md#postorgsorgidresourcesdrivers) - Register a new Resource Driver.
* [PutOrgsOrgIDResourcesDriversDriverID](docs/sdks/driverdefinition/README.md#putorgsorgidresourcesdriversdriverid) - Update a Resource Driver.

### [Environment](docs/sdks/environment/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvID](docs/sdks/environment/README.md#deleteorgsorgidappsappidenvsenvid) - Delete a specific Environment.
* [GetOrgsOrgIDAppsAppIDEnvs](docs/sdks/environment/README.md#getorgsorgidappsappidenvs) - List all Environments.
* [GetOrgsOrgIDAppsAppIDEnvsEnvID](docs/sdks/environment/README.md#getorgsorgidappsappidenvsenvid) - Get a specific Environment.
* [PostOrgsOrgIDAppsAppIDEnvs](docs/sdks/environment/README.md#postorgsorgidappsappidenvs) - Add a new Environment to an Application.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployID](docs/sdks/environment/README.md#putorgsorgidappsappidenvsenvidfromdeployid) - Rebase to a different Deployment.

### [EnvironmentType](docs/sdks/environmenttype/README.md)

* [DeleteOrgsOrgIDEnvTypesEnvTypeID](docs/sdks/environmenttype/README.md#deleteorgsorgidenvtypesenvtypeid) - Deletes an Environment Type
* [GetOrgsOrgIDEnvTypes](docs/sdks/environmenttype/README.md#getorgsorgidenvtypes) - List all Environment Types
* [GetOrgsOrgIDEnvTypesEnvTypeID](docs/sdks/environmenttype/README.md#getorgsorgidenvtypesenvtypeid) - Get an Environment Type
* [PostOrgsOrgIDEnvTypes](docs/sdks/environmenttype/README.md#postorgsorgidenvtypes) - Add a new Environment Type

### [Event](docs/sdks/event/README.md)

* [DeleteOrgsOrgIDAppsAppIDJobs](docs/sdks/event/README.md#deleteorgsorgidappsappidjobs) - Deletes all Jobs for the Application
* [DeleteOrgsOrgIDAppsAppIDWebhooksJobID](docs/sdks/event/README.md#deleteorgsorgidappsappidwebhooksjobid) - Delete a Webhook
* [GetOrgsOrgIDAppsAppIDWebhooks](docs/sdks/event/README.md#getorgsorgidappsappidwebhooks) - List Webhooks
* [GetOrgsOrgIDAppsAppIDWebhooksJobID](docs/sdks/event/README.md#getorgsorgidappsappidwebhooksjobid) - Get a Webhook
* [GetOrgsOrgIDEvents](docs/sdks/event/README.md#getorgsorgidevents) - List Events
* [PostOrgsOrgIDAppsAppIDWebhooks](docs/sdks/event/README.md#postorgsorgidappsappidwebhooks) - Create a new Webhook
* [PostOrgsOrgIDAppsAppIDWebhooksJobID](docs/sdks/event/README.md#postorgsorgidappsappidwebhooksjobid) - Update a Webhook

### [Image](docs/sdks/image/README.md)

* [GetOrgsOrgIDImages](docs/sdks/image/README.md#getorgsorgidimages) - List all Container Images
* [GetOrgsOrgIDImagesImageID](docs/sdks/image/README.md#getorgsorgidimagesimageid) - Get a specific Image Object
* [GetOrgsOrgIDImagesImageIDBuilds](docs/sdks/image/README.md#getorgsorgidimagesimageidbuilds) - Lists all the Builds of an Image
* [PostOrgsOrgIDImagesImageIDBuilds](docs/sdks/image/README.md#postorgsorgidimagesimageidbuilds) - Add a new Image Build

### [MatchingCriteria](docs/sdks/matchingcriteria/README.md)

* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](docs/sdks/matchingcriteria/README.md#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](docs/sdks/matchingcriteria/README.md#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.

### [Organization](docs/sdks/organization/README.md)

* [GetOrgs](docs/sdks/organization/README.md#getorgs) - List active organizations the user has access to.
* [GetOrgsOrgID](docs/sdks/organization/README.md#getorgsorgid) - Get the specified Organization.

### [Registry](docs/sdks/registry/README.md)

* [DeleteOrgsOrgIDRegistriesRegID](docs/sdks/registry/README.md#deleteorgsorgidregistriesregid) - Deletes an existing registry record and all associated credentials and secrets.
* [GetOrgsOrgIDRegistries](docs/sdks/registry/README.md#getorgsorgidregistries) - Lists available registries for the organization.
* [GetOrgsOrgIDRegistriesRegID](docs/sdks/registry/README.md#getorgsorgidregistriesregid) - Loads a registry record details.
* [GetOrgsOrgIDRegistriesRegIDCreds](docs/sdks/registry/README.md#getorgsorgidregistriesregidcreds) - Returns current account credentials or secret details for the registry.
* [PatchOrgsOrgIDRegistriesRegID](docs/sdks/registry/README.md#patchorgsorgidregistriesregid) - Updates (patches) an existing registry record.
* [PostOrgsOrgIDRegistries](docs/sdks/registry/README.md#postorgsorgidregistries) - Creates a new registry record.

### [ResourceAccount](docs/sdks/resourceaccount/README.md)

* [GetOrgsOrgIDResourcesAccounts](docs/sdks/resourceaccount/README.md#getorgsorgidresourcesaccounts) - List Resource Accounts in the organization.
* [GetOrgsOrgIDResourcesAccountsAccID](docs/sdks/resourceaccount/README.md#getorgsorgidresourcesaccountsaccid) - Get a Resource Account.
* [PatchOrgsOrgIDResourcesAccountsAccID](docs/sdks/resourceaccount/README.md#patchorgsorgidresourcesaccountsaccid) - Update a Resource Account.
* [PostOrgsOrgIDResourcesAccounts](docs/sdks/resourceaccount/README.md#postorgsorgidresourcesaccounts) - Create a new Resource Account in the organization.

### [ResourceDefinition](docs/sdks/resourcedefinition/README.md)

* [DeleteOrgsOrgIDResourcesDefsDefID](docs/sdks/resourcedefinition/README.md#deleteorgsorgidresourcesdefsdefid) - Delete a Resource Definition.
* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](docs/sdks/resourcedefinition/README.md#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [GetOrgsOrgIDResourcesDefs](docs/sdks/resourcedefinition/README.md#getorgsorgidresourcesdefs) - List Resource Definitions.
* [GetOrgsOrgIDResourcesDefsDefID](docs/sdks/resourcedefinition/README.md#getorgsorgidresourcesdefsdefid) - Get a specific Resource Definition.
* [GetOrgsOrgIDResourcesDefsDefIDResources](docs/sdks/resourcedefinition/README.md#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.
* [PatchOrgsOrgIDResourcesDefsDefID](docs/sdks/resourcedefinition/README.md#patchorgsorgidresourcesdefsdefid) - Update a Resource Definition.
* [PostOrgsOrgIDResourcesDefs](docs/sdks/resourcedefinition/README.md#postorgsorgidresourcesdefs) - Create a new Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](docs/sdks/resourcedefinition/README.md#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.
* [PutOrgsOrgIDResourcesDefsDefID](docs/sdks/resourcedefinition/README.md#putorgsorgidresourcesdefsdefid) - Update a Resource Definition.

### [ResourceType](docs/sdks/resourcetype/README.md)

* [GetOrgsOrgIDResourcesTypes](docs/sdks/resourcetype/README.md#getorgsorgidresourcestypes) - List Resource Types.

### [RuntimeInfo](docs/sdks/runtimeinfo/README.md)

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntime](docs/sdks/runtimeinfo/README.md#getorgsorgidappsappidenvsenvidruntime) - Get Runtime information about the environment.
* [GetOrgsOrgIDAppsAppIDRuntime](docs/sdks/runtimeinfo/README.md#getorgsorgidappsappidruntime) - Get Runtime information about specific environments.
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicas](docs/sdks/runtimeinfo/README.md#patchorgsorgidappsappidenvsenvidruntimereplicas) - Set number of replicas for an environment's modules.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePaused](docs/sdks/runtimeinfo/README.md#putorgsorgidappsappidenvsenvidruntimepaused) - Pause / Resume an environment.

### [Set](docs/sdks/set/README.md)

* [GetSets](docs/sdks/set/README.md#getsets) - Get all Deployment Sets
* [GetOrgsOrgIDAppsAppIDSetsSetID](docs/sdks/set/README.md#getorgsorgidappsappidsetssetid) - Get a Deployment Set
* [GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID](docs/sdks/set/README.md#getorgsorgidappsappidsetssetiddiffsourcesetid) - Get the difference between 2 Deployment Sets
* [PostOrgsOrgIDAppsAppIDSetsSetID](docs/sdks/set/README.md#postorgsorgidappsappidsetssetid) - Apply a Deployment Delta to a Deployment Set

### [UserInvite](docs/sdks/userinvite/README.md)

* [GetOrgsOrgIDInvitations](docs/sdks/userinvite/README.md#getorgsorgidinvitations) - List the invites issued for the organization.

### [UserProfile](docs/sdks/userprofile/README.md)

* [DeleteTokensTokenID](docs/sdks/userprofile/README.md#deletetokenstokenid) - DEPRECATED
* [GetCurrentUser](docs/sdks/userprofile/README.md#getcurrentuser) - Gets the extended profile of the current user
* [GetTokens](docs/sdks/userprofile/README.md#gettokens) - DEPRECATED
* [GetUsersMe](docs/sdks/userprofile/README.md#getusersme) - DEPRECATED
* [PatchCurrentUser](docs/sdks/userprofile/README.md#patchcurrentuser) - Updates the extended profile of the current user.
* [PostOrgsOrgIDUsers](docs/sdks/userprofile/README.md#postorgsorgidusers) - Creates a new service user.

### [UserRole](docs/sdks/userrole/README.md)

* [DeleteOrgsOrgIDAppsAppIDUsersUserID](docs/sdks/userrole/README.md#deleteorgsorgidappsappidusersuserid) - Remove the role of a User on an Application
* [DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/sdks/userrole/README.md#deleteorgsorgidenvtypesenvtypeusersuserid) - Remove the role of a User on an Environment Type
* [DeleteOrgsOrgIDUsersUserID](docs/sdks/userrole/README.md#deleteorgsorgidusersuserid) - Remove the role of a User on an Organization
* [GetOrgsOrgIDAppsAppIDUsers](docs/sdks/userrole/README.md#getorgsorgidappsappidusers) - List Users with roles in an App
* [GetOrgsOrgIDAppsAppIDUsersUserID](docs/sdks/userrole/README.md#getorgsorgidappsappidusersuserid) - Get the role of a User on an Application
* [GetOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/sdks/userrole/README.md#getorgsorgidenvtypesenvtypeusersuserid) - Get the role of a User on an Environment Type
* [GetOrgsOrgIDUsers](docs/sdks/userrole/README.md#getorgsorgidusers) - List Users with roles in an Organization
* [GetOrgsOrgIDUsersUserID](docs/sdks/userrole/README.md#getorgsorgidusersuserid) - Get the role of a User on an Organization
* [PatchOrgsOrgIDAppsAppIDUsersUserID](docs/sdks/userrole/README.md#patchorgsorgidappsappidusersuserid) - Update the role of a User on an Application
* [PatchOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/sdks/userrole/README.md#patchorgsorgidenvtypesenvtypeusersuserid) - Update the role of a User on an Environment Type
* [PatchOrgsOrgIDUsersUserID](docs/sdks/userrole/README.md#patchorgsorgidusersuserid) - Update the role of a User on an Organization
* [PostOrgsOrgIDAppsAppIDUsers](docs/sdks/userrole/README.md#postorgsorgidappsappidusers) - Adds a User to an Application with a Role
* [PostOrgsOrgIDEnvTypesEnvTypeUsers](docs/sdks/userrole/README.md#postorgsorgidenvtypesenvtypeusers) - Adds a User to an Environment Type with a Role
* [PostOrgsOrgIDInvitations](docs/sdks/userrole/README.md#postorgsorgidinvitations) - Invites a user to an Organization with a specified role.

### [Value](docs/sdks/value/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/sdks/value/README.md#deleteorgsorgidappsappidenvsenvidvalues) - Delete all Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/sdks/value/README.md#deleteorgsorgidappsappidenvsenvidvalueskey) - Delete Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDValues](docs/sdks/value/README.md#deleteorgsorgidappsappidvalues) - Delete all Shared Value for an App
* [DeleteOrgsOrgIDAppsAppIDValuesKey](docs/sdks/value/README.md#deleteorgsorgidappsappidvalueskey) - Delete Shared Value for an Application
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/sdks/value/README.md#getorgsorgidappsappidenvsenvidvalues) - List Shared Values in an Environment
* [GetOrgsOrgIDAppsAppIDValues](docs/sdks/value/README.md#getorgsorgidappsappidvalues) - List Shared Values in an Application
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/sdks/value/README.md#patchorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PatchOrgsOrgIDAppsAppIDValuesKey](docs/sdks/value/README.md#patchorgsorgidappsappidvalueskey) - Update Shared Value for an Application
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/sdks/value/README.md#postorgsorgidappsappidenvsenvidvalues) - Create a Shared Value for an Environment
* [PostOrgsOrgIDAppsAppIDValues](docs/sdks/value/README.md#postorgsorgidappsappidvalues) - Create a Shared Value for an Application
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/sdks/value/README.md#putorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PutOrgsOrgIDAppsAppIDValuesKey](docs/sdks/value/README.md#putorgsorgidappsappidvalueskey) - Update Shared Value for an Application

### [ValueSetVersion](docs/sdks/valuesetversion/README.md)

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersions](docs/sdks/valuesetversion/README.md#getorgsorgidappsappidenvsenvidvaluesetversions) - List Value Set Versions in an Environment of an App
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionID](docs/sdks/valuesetversion/README.md#getorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionid) - Get a single Value Set Version in an Environment of an App
* [GetOrgsOrgIDAppsAppIDValueSetVersions](docs/sdks/valuesetversion/README.md#getorgsorgidappsappidvaluesetversions) - List Value Set Versions in the App
* [GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionID](docs/sdks/valuesetversion/README.md#getorgsorgidappsappidvaluesetversionsvaluesetversionid) - Get a single Value Set Version from the App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKey](docs/sdks/valuesetversion/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Environment Version history.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestore](docs/sdks/valuesetversion/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKey](docs/sdks/valuesetversion/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKey](docs/sdks/valuesetversion/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Version history.
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestore](docs/sdks/valuesetversion/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKey](docs/sdks/valuesetversion/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an App

### [WorkloadProfile](docs/sdks/workloadprofile/README.md)

* [DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersion](docs/sdks/workloadprofile/README.md#deleteorgsorgidworkloadprofilesprofileidversionsversion) - Delete a Workload Profile Version
* [DeleteOrgsOrgIDWorkloadProfilesProfileQid](docs/sdks/workloadprofile/README.md#deleteorgsorgidworkloadprofilesprofileqid) - Delete a Workload Profile
* [GetOrgsOrgIDWorkloadProfiles](docs/sdks/workloadprofile/README.md#getorgsorgidworkloadprofiles) - List workload profiles available to the organization.
* [GetOrgsOrgIDWorkloadProfilesProfileQid](docs/sdks/workloadprofile/README.md#getorgsorgidworkloadprofilesprofileqid) - Get a Workload Profile
* [GetOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/sdks/workloadprofile/README.md#getorgsorgidworkloadprofilesprofileqidversions) - List versions of the given workload profile with optional constraint.
* [PostOrgsOrgIDWorkloadProfiles](docs/sdks/workloadprofile/README.md#postorgsorgidworkloadprofiles) - Create new Workload Profile
* [PostOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/sdks/workloadprofile/README.md#postorgsorgidworkloadprofilesprofileqidversions) - Add new Version of the Workload Profile

### [Public](docs/sdks/public/README.md)

* [DeleteOrgsOrgIDAppsAppID](docs/sdks/public/README.md#deleteorgsorgidappsappid) - Delete an Application
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvID](docs/sdks/public/README.md#deleteorgsorgidappsappidenvsenvid) - Delete a specific Environment.
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResID](docs/sdks/public/README.md#deleteorgsorgidappsappidenvsenvidresourcestyperesid) - Delete Active Resources.
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/sdks/public/README.md#deleteorgsorgidappsappidenvsenvidrulesruleid) - Delete Automation Rule from an Environment.
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/sdks/public/README.md#deleteorgsorgidappsappidenvsenvidvalues) - Delete all Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/sdks/public/README.md#deleteorgsorgidappsappidenvsenvidvalueskey) - Delete Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDJobs](docs/sdks/public/README.md#deleteorgsorgidappsappidjobs) - Deletes all Jobs for the Application
* [DeleteOrgsOrgIDAppsAppIDUsersUserID](docs/sdks/public/README.md#deleteorgsorgidappsappidusersuserid) - Remove the role of a User on an Application
* [DeleteOrgsOrgIDAppsAppIDValues](docs/sdks/public/README.md#deleteorgsorgidappsappidvalues) - Delete all Shared Value for an App
* [DeleteOrgsOrgIDAppsAppIDValuesKey](docs/sdks/public/README.md#deleteorgsorgidappsappidvalueskey) - Delete Shared Value for an Application
* [DeleteOrgsOrgIDAppsAppIDWebhooksJobID](docs/sdks/public/README.md#deleteorgsorgidappsappidwebhooksjobid) - Delete a Webhook
* [DeleteOrgsOrgIDArtefactsArtefactID](docs/sdks/public/README.md#deleteorgsorgidartefactsartefactid) - Delete Artefact and all related Artefact Versions
* [DeleteOrgsOrgIDEnvTypesEnvTypeID](docs/sdks/public/README.md#deleteorgsorgidenvtypesenvtypeid) - Deletes an Environment Type
* [DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/sdks/public/README.md#deleteorgsorgidenvtypesenvtypeusersuserid) - Remove the role of a User on an Environment Type
* [DeleteOrgsOrgIDRegistriesRegID](docs/sdks/public/README.md#deleteorgsorgidregistriesregid) - Deletes an existing registry record and all associated credentials and secrets.
* [DeleteOrgsOrgIDResourcesDefsDefID](docs/sdks/public/README.md#deleteorgsorgidresourcesdefsdefid) - Delete a Resource Definition.
* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](docs/sdks/public/README.md#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [DeleteOrgsOrgIDResourcesDriversDriverID](docs/sdks/public/README.md#deleteorgsorgidresourcesdriversdriverid) - Delete a Resources Driver.
* [DeleteOrgsOrgIDUsersUserID](docs/sdks/public/README.md#deleteorgsorgidusersuserid) - Remove the role of a User on an Organization
* [DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersion](docs/sdks/public/README.md#deleteorgsorgidworkloadprofilesprofileidversionsversion) - Delete a Workload Profile Version
* [DeleteOrgsOrgIDWorkloadProfilesProfileQid](docs/sdks/public/README.md#deleteorgsorgidworkloadprofilesprofileqid) - Delete a Workload Profile
* [DeleteTokensTokenID](docs/sdks/public/README.md#deletetokenstokenid) - DEPRECATED
* [GetDelta](docs/sdks/public/README.md#getdelta) - Fetch an existing Delta
* [GetSets](docs/sdks/public/README.md#getsets) - Get all Deployment Sets
* [GetCurrentUser](docs/sdks/public/README.md#getcurrentuser) - Gets the extended profile of the current user
* [GetOrgs](docs/sdks/public/README.md#getorgs) - List active organizations the user has access to.
* [GetOrgsOrgID](docs/sdks/public/README.md#getorgsorgid) - Get the specified Organization.
* [GetOrgsOrgIDApps](docs/sdks/public/README.md#getorgsorgidapps) - List all Applications in an Organization.
* [GetOrgsOrgIDAppsAppID](docs/sdks/public/README.md#getorgsorgidappsappid) - Get an existing Application
* [GetOrgsOrgIDAppsAppIDDeltas](docs/sdks/public/README.md#getorgsorgidappsappiddeltas) - List Deltas in an Application
* [GetOrgsOrgIDAppsAppIDEnvs](docs/sdks/public/README.md#getorgsorgidappsappidenvs) - List all Environments.
* [GetOrgsOrgIDAppsAppIDEnvsEnvID](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvid) - Get a specific Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/sdks/public/README.md#getorgsorgidappsappidenvsenviddeploys) - List Deployments in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployID](docs/sdks/public/README.md#getorgsorgidappsappidenvsenviddeploysdeployid) - Get a specific Deployment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrors](docs/sdks/public/README.md#getorgsorgidappsappidenvsenviddeploysdeployiderrors) - List errors that occurred in a Deployment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDResources](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvidresources) - List Active Resources provisioned in an environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvidrules) - List all Automation Rules in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvidrulesruleid) - Get a specific Automation Rule for an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntime](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvidruntime) - Get Runtime information about the environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersions](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvidvaluesetversions) - List Value Set Versions in an Environment of an App
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionID](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionid) - Get a single Value Set Version in an Environment of an App
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/sdks/public/README.md#getorgsorgidappsappidenvsenvidvalues) - List Shared Values in an Environment
* [GetOrgsOrgIDAppsAppIDRuntime](docs/sdks/public/README.md#getorgsorgidappsappidruntime) - Get Runtime information about specific environments.
* [GetOrgsOrgIDAppsAppIDSetsSetID](docs/sdks/public/README.md#getorgsorgidappsappidsetssetid) - Get a Deployment Set
* [GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID](docs/sdks/public/README.md#getorgsorgidappsappidsetssetiddiffsourcesetid) - Get the difference between 2 Deployment Sets
* [GetOrgsOrgIDAppsAppIDUsers](docs/sdks/public/README.md#getorgsorgidappsappidusers) - List Users with roles in an App
* [GetOrgsOrgIDAppsAppIDUsersUserID](docs/sdks/public/README.md#getorgsorgidappsappidusersuserid) - Get the role of a User on an Application
* [GetOrgsOrgIDAppsAppIDValueSetVersions](docs/sdks/public/README.md#getorgsorgidappsappidvaluesetversions) - List Value Set Versions in the App
* [GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionID](docs/sdks/public/README.md#getorgsorgidappsappidvaluesetversionsvaluesetversionid) - Get a single Value Set Version from the App
* [GetOrgsOrgIDAppsAppIDValues](docs/sdks/public/README.md#getorgsorgidappsappidvalues) - List Shared Values in an Application
* [GetOrgsOrgIDAppsAppIDWebhooks](docs/sdks/public/README.md#getorgsorgidappsappidwebhooks) - List Webhooks
* [GetOrgsOrgIDAppsAppIDWebhooksJobID](docs/sdks/public/README.md#getorgsorgidappsappidwebhooksjobid) - Get a Webhook
* [GetOrgsOrgIDArtefactVersions](docs/sdks/public/README.md#getorgsorgidartefactversions) - List all Artefacts Versions.
* [GetOrgsOrgIDArtefactVersionsArtefactVersionID](docs/sdks/public/README.md#getorgsorgidartefactversionsartefactversionid) - Get an Artefacts Versions.
* [GetOrgsOrgIDArtefacts](docs/sdks/public/README.md#getorgsorgidartefacts) - List all Artefacts.
* [GetOrgsOrgIDArtefactsArtefactIDVersions](docs/sdks/public/README.md#getorgsorgidartefactsartefactidversions) - List all Artefact Versions of an Artefact.
* [GetOrgsOrgIDEnvTypes](docs/sdks/public/README.md#getorgsorgidenvtypes) - List all Environment Types
* [GetOrgsOrgIDEnvTypesEnvTypeID](docs/sdks/public/README.md#getorgsorgidenvtypesenvtypeid) - Get an Environment Type
* [GetOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/sdks/public/README.md#getorgsorgidenvtypesenvtypeusersuserid) - Get the role of a User on an Environment Type
* [GetOrgsOrgIDEvents](docs/sdks/public/README.md#getorgsorgidevents) - List Events
* [GetOrgsOrgIDImages](docs/sdks/public/README.md#getorgsorgidimages) - List all Container Images
* [GetOrgsOrgIDImagesImageID](docs/sdks/public/README.md#getorgsorgidimagesimageid) - Get a specific Image Object
* [GetOrgsOrgIDImagesImageIDBuilds](docs/sdks/public/README.md#getorgsorgidimagesimageidbuilds) - Lists all the Builds of an Image
* [GetOrgsOrgIDInvitations](docs/sdks/public/README.md#getorgsorgidinvitations) - List the invites issued for the organization.
* [GetOrgsOrgIDRegistries](docs/sdks/public/README.md#getorgsorgidregistries) - Lists available registries for the organization.
* [GetOrgsOrgIDRegistriesRegID](docs/sdks/public/README.md#getorgsorgidregistriesregid) - Loads a registry record details.
* [GetOrgsOrgIDRegistriesRegIDCreds](docs/sdks/public/README.md#getorgsorgidregistriesregidcreds) - Returns current account credentials or secret details for the registry.
* [GetOrgsOrgIDResourcesAccountTypes](docs/sdks/public/README.md#getorgsorgidresourcesaccounttypes) - List Resource Account Types available to the organization.
* [GetOrgsOrgIDResourcesAccounts](docs/sdks/public/README.md#getorgsorgidresourcesaccounts) - List Resource Accounts in the organization.
* [GetOrgsOrgIDResourcesAccountsAccID](docs/sdks/public/README.md#getorgsorgidresourcesaccountsaccid) - Get a Resource Account.
* [GetOrgsOrgIDResourcesDefs](docs/sdks/public/README.md#getorgsorgidresourcesdefs) - List Resource Definitions.
* [GetOrgsOrgIDResourcesDefsDefID](docs/sdks/public/README.md#getorgsorgidresourcesdefsdefid) - Get a specific Resource Definition.
* [GetOrgsOrgIDResourcesDefsDefIDResources](docs/sdks/public/README.md#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.
* [GetOrgsOrgIDResourcesDrivers](docs/sdks/public/README.md#getorgsorgidresourcesdrivers) - List Resource Drivers.
* [GetOrgsOrgIDResourcesDriversDriverID](docs/sdks/public/README.md#getorgsorgidresourcesdriversdriverid) - Get a Resource Driver.
* [GetOrgsOrgIDResourcesTypes](docs/sdks/public/README.md#getorgsorgidresourcestypes) - List Resource Types.
* [GetOrgsOrgIDUsers](docs/sdks/public/README.md#getorgsorgidusers) - List Users with roles in an Organization
* [GetOrgsOrgIDUsersUserID](docs/sdks/public/README.md#getorgsorgidusersuserid) - Get the role of a User on an Organization
* [GetOrgsOrgIDWorkloadProfiles](docs/sdks/public/README.md#getorgsorgidworkloadprofiles) - List workload profiles available to the organization.
* [GetOrgsOrgIDWorkloadProfilesProfileQid](docs/sdks/public/README.md#getorgsorgidworkloadprofilesprofileqid) - Get a Workload Profile
* [GetOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/sdks/public/README.md#getorgsorgidworkloadprofilesprofileqidversions) - List versions of the given workload profile with optional constraint.
* [GetTokens](docs/sdks/public/README.md#gettokens) - DEPRECATED
* [GetUsersMe](docs/sdks/public/README.md#getusersme) - DEPRECATED
* [PatchCurrentUser](docs/sdks/public/README.md#patchcurrentuser) - Updates the extended profile of the current user.
* [PatchOrgsOrgIDAppsAppIDDeltasDeltaID](docs/sdks/public/README.md#patchorgsorgidappsappiddeltasdeltaid) - Update an existing Delta
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicas](docs/sdks/public/README.md#patchorgsorgidappsappidenvsenvidruntimereplicas) - Set number of replicas for an environment's modules.
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/sdks/public/README.md#patchorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PatchOrgsOrgIDAppsAppIDUsersUserID](docs/sdks/public/README.md#patchorgsorgidappsappidusersuserid) - Update the role of a User on an Application
* [PatchOrgsOrgIDAppsAppIDValuesKey](docs/sdks/public/README.md#patchorgsorgidappsappidvalueskey) - Update Shared Value for an Application
* [PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionID](docs/sdks/public/README.md#patchorgsorgidartefactsartefactidversionsversionid) - Update Version of an Artefact.
* [PatchOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/sdks/public/README.md#patchorgsorgidenvtypesenvtypeusersuserid) - Update the role of a User on an Environment Type
* [PatchOrgsOrgIDRegistriesRegID](docs/sdks/public/README.md#patchorgsorgidregistriesregid) - Updates (patches) an existing registry record.
* [PatchOrgsOrgIDResourcesAccountsAccID](docs/sdks/public/README.md#patchorgsorgidresourcesaccountsaccid) - Update a Resource Account.
* [PatchOrgsOrgIDResourcesDefsDefID](docs/sdks/public/README.md#patchorgsorgidresourcesdefsdefid) - Update a Resource Definition.
* [PatchOrgsOrgIDUsersUserID](docs/sdks/public/README.md#patchorgsorgidusersuserid) - Update the role of a User on an Organization
* [PostOrgsOrgIDApps](docs/sdks/public/README.md#postorgsorgidapps) - Add a new Application to an Organization
* [PostOrgsOrgIDAppsAppIDDeltas](docs/sdks/public/README.md#postorgsorgidappsappiddeltas) - Create a new Delta
* [PostOrgsOrgIDAppsAppIDEnvs](docs/sdks/public/README.md#postorgsorgidappsappidenvs) - Add a new Environment to an Application.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/sdks/public/README.md#postorgsorgidappsappidenvsenviddeploys) - Start a new Deployment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/sdks/public/README.md#postorgsorgidappsappidenvsenvidrules) - Create a new Automation Rule for an Environment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKey](docs/sdks/public/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Environment Version history.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestore](docs/sdks/public/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKey](docs/sdks/public/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/sdks/public/README.md#postorgsorgidappsappidenvsenvidvalues) - Create a Shared Value for an Environment
* [PostOrgsOrgIDAppsAppIDSetsSetID](docs/sdks/public/README.md#postorgsorgidappsappidsetssetid) - Apply a Deployment Delta to a Deployment Set
* [PostOrgsOrgIDAppsAppIDUsers](docs/sdks/public/README.md#postorgsorgidappsappidusers) - Adds a User to an Application with a Role
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKey](docs/sdks/public/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Version history.
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestore](docs/sdks/public/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKey](docs/sdks/public/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an App
* [PostOrgsOrgIDAppsAppIDValues](docs/sdks/public/README.md#postorgsorgidappsappidvalues) - Create a Shared Value for an Application
* [PostOrgsOrgIDAppsAppIDWebhooks](docs/sdks/public/README.md#postorgsorgidappsappidwebhooks) - Create a new Webhook
* [PostOrgsOrgIDAppsAppIDWebhooksJobID](docs/sdks/public/README.md#postorgsorgidappsappidwebhooksjobid) - Update a Webhook
* [PostOrgsOrgIDArtefactVersions](docs/sdks/public/README.md#postorgsorgidartefactversions) - Register a new Artefact Version with your organization.
* [PostOrgsOrgIDEnvTypes](docs/sdks/public/README.md#postorgsorgidenvtypes) - Add a new Environment Type
* [PostOrgsOrgIDEnvTypesEnvTypeUsers](docs/sdks/public/README.md#postorgsorgidenvtypesenvtypeusers) - Adds a User to an Environment Type with a Role
* [PostOrgsOrgIDImagesImageIDBuilds](docs/sdks/public/README.md#postorgsorgidimagesimageidbuilds) - Add a new Image Build
* [PostOrgsOrgIDInvitations](docs/sdks/public/README.md#postorgsorgidinvitations) - Invites a user to an Organization with a specified role.
* [PostOrgsOrgIDRegistries](docs/sdks/public/README.md#postorgsorgidregistries) - Creates a new registry record.
* [PostOrgsOrgIDResourcesAccounts](docs/sdks/public/README.md#postorgsorgidresourcesaccounts) - Create a new Resource Account in the organization.
* [PostOrgsOrgIDResourcesDefs](docs/sdks/public/README.md#postorgsorgidresourcesdefs) - Create a new Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](docs/sdks/public/README.md#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.
* [PostOrgsOrgIDResourcesDrivers](docs/sdks/public/README.md#postorgsorgidresourcesdrivers) - Register a new Resource Driver.
* [PostOrgsOrgIDUsers](docs/sdks/public/README.md#postorgsorgidusers) - Creates a new service user.
* [PostOrgsOrgIDWorkloadProfiles](docs/sdks/public/README.md#postorgsorgidworkloadprofiles) - Create new Workload Profile
* [PostOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/sdks/public/README.md#postorgsorgidworkloadprofilesprofileqidversions) - Add new Version of the Workload Profile
* [PutDelta](docs/sdks/public/README.md#putdelta) - Update an existing Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived](docs/sdks/public/README.md#putorgsorgidappsappiddeltasdeltaidmetadataarchived) - Mark a Delta as "archived"
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID](docs/sdks/public/README.md#putorgsorgidappsappiddeltasdeltaidmetadataenvid) - Change the Environment of a Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName](docs/sdks/public/README.md#putorgsorgidappsappiddeltasdeltaidmetadataname) - Change the name of a Delta
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployID](docs/sdks/public/README.md#putorgsorgidappsappidenvsenvidfromdeployid) - Rebase to a different Deployment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/sdks/public/README.md#putorgsorgidappsappidenvsenvidrulesruleid) - Update an existing Automation Rule for an Environment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePaused](docs/sdks/public/README.md#putorgsorgidappsappidenvsenvidruntimepaused) - Pause / Resume an environment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/sdks/public/README.md#putorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PutOrgsOrgIDAppsAppIDValuesKey](docs/sdks/public/README.md#putorgsorgidappsappidvalueskey) - Update Shared Value for an Application
* [PutOrgsOrgIDResourcesDefsDefID](docs/sdks/public/README.md#putorgsorgidresourcesdefsdefid) - Update a Resource Definition.
* [PutOrgsOrgIDResourcesDriversDriverID](docs/sdks/public/README.md#putorgsorgidresourcesdriversdriverid) - Update a Resource Driver.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
