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


### [AccountType](docs/accounttype/README.md)

* [GetOrgsOrgIDResourcesAccountTypes](docs/accounttype/README.md#getorgsorgidresourcesaccounttypes) - List Resource Account Types available to the organization.

### [ActiveResource](docs/activeresource/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResID](docs/activeresource/README.md#deleteorgsorgidappsappidenvsenvidresourcestyperesid) - Delete Active Resources.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDResources](docs/activeresource/README.md#getorgsorgidappsappidenvsenvidresources) - List Active Resources provisioned in an environment.
* [GetOrgsOrgIDResourcesDefsDefIDResources](docs/activeresource/README.md#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.

### [Application](docs/application/README.md)

* [DeleteOrgsOrgIDAppsAppID](docs/application/README.md#deleteorgsorgidappsappid) - Delete an Application
* [GetOrgsOrgIDApps](docs/application/README.md#getorgsorgidapps) - List all Applications in an Organization.
* [GetOrgsOrgIDAppsAppID](docs/application/README.md#getorgsorgidappsappid) - Get an existing Application
* [PostOrgsOrgIDApps](docs/application/README.md#postorgsorgidapps) - Add a new Application to an Organization

### [Artefact](docs/artefact/README.md)

* [DeleteOrgsOrgIDArtefactsArtefactID](docs/artefact/README.md#deleteorgsorgidartefactsartefactid) - Delete Artefact and all related Artefact Versions
* [GetOrgsOrgIDArtefacts](docs/artefact/README.md#getorgsorgidartefacts) - List all Artefacts.

### [ArtefactVersion](docs/artefactversion/README.md)

* [GetOrgsOrgIDArtefactVersions](docs/artefactversion/README.md#getorgsorgidartefactversions) - List all Artefacts Versions.
* [GetOrgsOrgIDArtefactVersionsArtefactVersionID](docs/artefactversion/README.md#getorgsorgidartefactversionsartefactversionid) - Get an Artefacts Versions.
* [GetOrgsOrgIDArtefactsArtefactIDVersions](docs/artefactversion/README.md#getorgsorgidartefactsartefactidversions) - List all Artefact Versions of an Artefact.
* [PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionID](docs/artefactversion/README.md#patchorgsorgidartefactsartefactidversionsversionid) - Update Version of an Artefact.
* [PostOrgsOrgIDArtefactVersions](docs/artefactversion/README.md#postorgsorgidartefactversions) - Register a new Artefact Version with your organization.

### [AutomationRule](docs/automationrule/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/automationrule/README.md#deleteorgsorgidappsappidenvsenvidrulesruleid) - Delete Automation Rule from an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/automationrule/README.md#getorgsorgidappsappidenvsenvidrules) - List all Automation Rules in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/automationrule/README.md#getorgsorgidappsappidenvsenvidrulesruleid) - Get a specific Automation Rule for an Environment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/automationrule/README.md#postorgsorgidappsappidenvsenvidrules) - Create a new Automation Rule for an Environment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/automationrule/README.md#putorgsorgidappsappidenvsenvidrulesruleid) - Update an existing Automation Rule for an Environment.

### [Delta](docs/delta/README.md)

* [GetDelta](docs/delta/README.md#getdelta) - Fetch an existing Delta
* [GetOrgsOrgIDAppsAppIDDeltas](docs/delta/README.md#getorgsorgidappsappiddeltas) - List Deltas in an Application
* [PatchOrgsOrgIDAppsAppIDDeltasDeltaID](docs/delta/README.md#patchorgsorgidappsappiddeltasdeltaid) - Update an existing Delta
* [PostOrgsOrgIDAppsAppIDDeltas](docs/delta/README.md#postorgsorgidappsappiddeltas) - Create a new Delta
* [PutDelta](docs/delta/README.md#putdelta) - Update an existing Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived](docs/delta/README.md#putorgsorgidappsappiddeltasdeltaidmetadataarchived) - Mark a Delta as "archived"
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID](docs/delta/README.md#putorgsorgidappsappiddeltasdeltaidmetadataenvid) - Change the Environment of a Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName](docs/delta/README.md#putorgsorgidappsappiddeltasdeltaidmetadataname) - Change the name of a Delta

### [Deployment](docs/deployment/README.md)

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/deployment/README.md#getorgsorgidappsappidenvsenviddeploys) - List Deployments in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployID](docs/deployment/README.md#getorgsorgidappsappidenvsenviddeploysdeployid) - Get a specific Deployment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrors](docs/deployment/README.md#getorgsorgidappsappidenvsenviddeploysdeployiderrors) - List errors that occurred in a Deployment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/deployment/README.md#postorgsorgidappsappidenvsenviddeploys) - Start a new Deployment.

### [DriverDefinition](docs/driverdefinition/README.md)

* [DeleteOrgsOrgIDResourcesDriversDriverID](docs/driverdefinition/README.md#deleteorgsorgidresourcesdriversdriverid) - Delete a Resources Driver.
* [GetOrgsOrgIDResourcesDrivers](docs/driverdefinition/README.md#getorgsorgidresourcesdrivers) - List Resource Drivers.
* [GetOrgsOrgIDResourcesDriversDriverID](docs/driverdefinition/README.md#getorgsorgidresourcesdriversdriverid) - Get a Resource Driver.
* [PostOrgsOrgIDResourcesDrivers](docs/driverdefinition/README.md#postorgsorgidresourcesdrivers) - Register a new Resource Driver.
* [PutOrgsOrgIDResourcesDriversDriverID](docs/driverdefinition/README.md#putorgsorgidresourcesdriversdriverid) - Update a Resource Driver.

### [Environment](docs/environment/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvID](docs/environment/README.md#deleteorgsorgidappsappidenvsenvid) - Delete a specific Environment.
* [GetOrgsOrgIDAppsAppIDEnvs](docs/environment/README.md#getorgsorgidappsappidenvs) - List all Environments.
* [GetOrgsOrgIDAppsAppIDEnvsEnvID](docs/environment/README.md#getorgsorgidappsappidenvsenvid) - Get a specific Environment.
* [PostOrgsOrgIDAppsAppIDEnvs](docs/environment/README.md#postorgsorgidappsappidenvs) - Add a new Environment to an Application.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployID](docs/environment/README.md#putorgsorgidappsappidenvsenvidfromdeployid) - Rebase to a different Deployment.

### [EnvironmentType](docs/environmenttype/README.md)

* [DeleteOrgsOrgIDEnvTypesEnvTypeID](docs/environmenttype/README.md#deleteorgsorgidenvtypesenvtypeid) - Deletes an Environment Type
* [GetOrgsOrgIDEnvTypes](docs/environmenttype/README.md#getorgsorgidenvtypes) - List all Environment Types
* [GetOrgsOrgIDEnvTypesEnvTypeID](docs/environmenttype/README.md#getorgsorgidenvtypesenvtypeid) - Get an Environment Type
* [PostOrgsOrgIDEnvTypes](docs/environmenttype/README.md#postorgsorgidenvtypes) - Add a new Environment Type

### [Event](docs/event/README.md)

* [DeleteOrgsOrgIDAppsAppIDJobs](docs/event/README.md#deleteorgsorgidappsappidjobs) - Deletes all Jobs for the Application
* [DeleteOrgsOrgIDAppsAppIDWebhooksJobID](docs/event/README.md#deleteorgsorgidappsappidwebhooksjobid) - Delete a Webhook
* [GetOrgsOrgIDAppsAppIDWebhooks](docs/event/README.md#getorgsorgidappsappidwebhooks) - List Webhooks
* [GetOrgsOrgIDAppsAppIDWebhooksJobID](docs/event/README.md#getorgsorgidappsappidwebhooksjobid) - Get a Webhook
* [GetOrgsOrgIDEvents](docs/event/README.md#getorgsorgidevents) - List Events
* [PostOrgsOrgIDAppsAppIDWebhooks](docs/event/README.md#postorgsorgidappsappidwebhooks) - Create a new Webhook
* [PostOrgsOrgIDAppsAppIDWebhooksJobID](docs/event/README.md#postorgsorgidappsappidwebhooksjobid) - Update a Webhook

### [Image](docs/image/README.md)

* [GetOrgsOrgIDImages](docs/image/README.md#getorgsorgidimages) - List all Container Images
* [GetOrgsOrgIDImagesImageID](docs/image/README.md#getorgsorgidimagesimageid) - Get a specific Image Object
* [GetOrgsOrgIDImagesImageIDBuilds](docs/image/README.md#getorgsorgidimagesimageidbuilds) - Lists all the Builds of an Image
* [PostOrgsOrgIDImagesImageIDBuilds](docs/image/README.md#postorgsorgidimagesimageidbuilds) - Add a new Image Build

### [MatchingCriteria](docs/matchingcriteria/README.md)

* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](docs/matchingcriteria/README.md#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](docs/matchingcriteria/README.md#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.

### [Organization](docs/organization/README.md)

* [GetOrgs](docs/organization/README.md#getorgs) - List active organizations the user has access to.
* [GetOrgsOrgID](docs/organization/README.md#getorgsorgid) - Get the specified Organization.

### [Registry](docs/registry/README.md)

* [DeleteOrgsOrgIDRegistriesRegID](docs/registry/README.md#deleteorgsorgidregistriesregid) - Deletes an existing registry record and all associated credentials and secrets.
* [GetOrgsOrgIDRegistries](docs/registry/README.md#getorgsorgidregistries) - Lists available registries for the organization.
* [GetOrgsOrgIDRegistriesRegID](docs/registry/README.md#getorgsorgidregistriesregid) - Loads a registry record details.
* [GetOrgsOrgIDRegistriesRegIDCreds](docs/registry/README.md#getorgsorgidregistriesregidcreds) - Returns current account credentials or secret details for the registry.
* [PatchOrgsOrgIDRegistriesRegID](docs/registry/README.md#patchorgsorgidregistriesregid) - Updates (patches) an existing registry record.
* [PostOrgsOrgIDRegistries](docs/registry/README.md#postorgsorgidregistries) - Creates a new registry record.

### [ResourceAccount](docs/resourceaccount/README.md)

* [GetOrgsOrgIDResourcesAccounts](docs/resourceaccount/README.md#getorgsorgidresourcesaccounts) - List Resource Accounts in the organization.
* [GetOrgsOrgIDResourcesAccountsAccID](docs/resourceaccount/README.md#getorgsorgidresourcesaccountsaccid) - Get a Resource Account.
* [PatchOrgsOrgIDResourcesAccountsAccID](docs/resourceaccount/README.md#patchorgsorgidresourcesaccountsaccid) - Update a Resource Account.
* [PostOrgsOrgIDResourcesAccounts](docs/resourceaccount/README.md#postorgsorgidresourcesaccounts) - Create a new Resource Account in the organization.

### [ResourceDefinition](docs/resourcedefinition/README.md)

* [DeleteOrgsOrgIDResourcesDefsDefID](docs/resourcedefinition/README.md#deleteorgsorgidresourcesdefsdefid) - Delete a Resource Definition.
* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](docs/resourcedefinition/README.md#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [GetOrgsOrgIDResourcesDefs](docs/resourcedefinition/README.md#getorgsorgidresourcesdefs) - List Resource Definitions.
* [GetOrgsOrgIDResourcesDefsDefID](docs/resourcedefinition/README.md#getorgsorgidresourcesdefsdefid) - Get a specific Resource Definition.
* [GetOrgsOrgIDResourcesDefsDefIDResources](docs/resourcedefinition/README.md#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.
* [PatchOrgsOrgIDResourcesDefsDefID](docs/resourcedefinition/README.md#patchorgsorgidresourcesdefsdefid) - Update a Resource Definition.
* [PostOrgsOrgIDResourcesDefs](docs/resourcedefinition/README.md#postorgsorgidresourcesdefs) - Create a new Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](docs/resourcedefinition/README.md#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.
* [PutOrgsOrgIDResourcesDefsDefID](docs/resourcedefinition/README.md#putorgsorgidresourcesdefsdefid) - Update a Resource Definition.

### [ResourceType](docs/resourcetype/README.md)

* [GetOrgsOrgIDResourcesTypes](docs/resourcetype/README.md#getorgsorgidresourcestypes) - List Resource Types.

### [RuntimeInfo](docs/runtimeinfo/README.md)

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntime](docs/runtimeinfo/README.md#getorgsorgidappsappidenvsenvidruntime) - Get Runtime information about the environment.
* [GetOrgsOrgIDAppsAppIDRuntime](docs/runtimeinfo/README.md#getorgsorgidappsappidruntime) - Get Runtime information about specific environments.
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicas](docs/runtimeinfo/README.md#patchorgsorgidappsappidenvsenvidruntimereplicas) - Set number of replicas for an environment's modules.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePaused](docs/runtimeinfo/README.md#putorgsorgidappsappidenvsenvidruntimepaused) - Pause / Resume an environment.

### [Set](docs/set/README.md)

* [GetSets](docs/set/README.md#getsets) - Get all Deployment Sets
* [GetOrgsOrgIDAppsAppIDSetsSetID](docs/set/README.md#getorgsorgidappsappidsetssetid) - Get a Deployment Set
* [GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID](docs/set/README.md#getorgsorgidappsappidsetssetiddiffsourcesetid) - Get the difference between 2 Deployment Sets
* [PostOrgsOrgIDAppsAppIDSetsSetID](docs/set/README.md#postorgsorgidappsappidsetssetid) - Apply a Deployment Delta to a Deployment Set

### [UserInvite](docs/userinvite/README.md)

* [GetOrgsOrgIDInvitations](docs/userinvite/README.md#getorgsorgidinvitations) - List the invites issued for the organization.

### [UserProfile](docs/userprofile/README.md)

* [DeleteTokensTokenID](docs/userprofile/README.md#deletetokenstokenid) - DEPRECATED
* [GetCurrentUser](docs/userprofile/README.md#getcurrentuser) - Gets the extended profile of the current user
* [GetTokens](docs/userprofile/README.md#gettokens) - DEPRECATED
* [GetUsersMe](docs/userprofile/README.md#getusersme) - DEPRECATED
* [PatchCurrentUser](docs/userprofile/README.md#patchcurrentuser) - Updates the extended profile of the current user.
* [PostOrgsOrgIDUsers](docs/userprofile/README.md#postorgsorgidusers) - Creates a new service user.

### [UserRole](docs/userrole/README.md)

* [DeleteOrgsOrgIDAppsAppIDUsersUserID](docs/userrole/README.md#deleteorgsorgidappsappidusersuserid) - Remove the role of a User on an Application
* [DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/userrole/README.md#deleteorgsorgidenvtypesenvtypeusersuserid) - Remove the role of a User on an Environment Type
* [DeleteOrgsOrgIDUsersUserID](docs/userrole/README.md#deleteorgsorgidusersuserid) - Remove the role of a User on an Organization
* [GetOrgsOrgIDAppsAppIDUsers](docs/userrole/README.md#getorgsorgidappsappidusers) - List Users with roles in an App
* [GetOrgsOrgIDAppsAppIDUsersUserID](docs/userrole/README.md#getorgsorgidappsappidusersuserid) - Get the role of a User on an Application
* [GetOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/userrole/README.md#getorgsorgidenvtypesenvtypeusersuserid) - Get the role of a User on an Environment Type
* [GetOrgsOrgIDUsers](docs/userrole/README.md#getorgsorgidusers) - List Users with roles in an Organization
* [GetOrgsOrgIDUsersUserID](docs/userrole/README.md#getorgsorgidusersuserid) - Get the role of a User on an Organization
* [PatchOrgsOrgIDAppsAppIDUsersUserID](docs/userrole/README.md#patchorgsorgidappsappidusersuserid) - Update the role of a User on an Application
* [PatchOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/userrole/README.md#patchorgsorgidenvtypesenvtypeusersuserid) - Update the role of a User on an Environment Type
* [PatchOrgsOrgIDUsersUserID](docs/userrole/README.md#patchorgsorgidusersuserid) - Update the role of a User on an Organization
* [PostOrgsOrgIDAppsAppIDUsers](docs/userrole/README.md#postorgsorgidappsappidusers) - Adds a User to an Application with a Role
* [PostOrgsOrgIDEnvTypesEnvTypeUsers](docs/userrole/README.md#postorgsorgidenvtypesenvtypeusers) - Adds a User to an Environment Type with a Role
* [PostOrgsOrgIDInvitations](docs/userrole/README.md#postorgsorgidinvitations) - Invites a user to an Organization with a specified role.

### [Value](docs/value/README.md)

* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/value/README.md#deleteorgsorgidappsappidenvsenvidvalues) - Delete all Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/value/README.md#deleteorgsorgidappsappidenvsenvidvalueskey) - Delete Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDValues](docs/value/README.md#deleteorgsorgidappsappidvalues) - Delete all Shared Value for an App
* [DeleteOrgsOrgIDAppsAppIDValuesKey](docs/value/README.md#deleteorgsorgidappsappidvalueskey) - Delete Shared Value for an Application
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/value/README.md#getorgsorgidappsappidenvsenvidvalues) - List Shared Values in an Environment
* [GetOrgsOrgIDAppsAppIDValues](docs/value/README.md#getorgsorgidappsappidvalues) - List Shared Values in an Application
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/value/README.md#patchorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PatchOrgsOrgIDAppsAppIDValuesKey](docs/value/README.md#patchorgsorgidappsappidvalueskey) - Update Shared Value for an Application
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/value/README.md#postorgsorgidappsappidenvsenvidvalues) - Create a Shared Value for an Environment
* [PostOrgsOrgIDAppsAppIDValues](docs/value/README.md#postorgsorgidappsappidvalues) - Create a Shared Value for an Application
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/value/README.md#putorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PutOrgsOrgIDAppsAppIDValuesKey](docs/value/README.md#putorgsorgidappsappidvalueskey) - Update Shared Value for an Application

### [ValueSetVersion](docs/valuesetversion/README.md)

* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersions](docs/valuesetversion/README.md#getorgsorgidappsappidenvsenvidvaluesetversions) - List Value Set Versions in an Environment of an App
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionID](docs/valuesetversion/README.md#getorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionid) - Get a single Value Set Version in an Environment of an App
* [GetOrgsOrgIDAppsAppIDValueSetVersions](docs/valuesetversion/README.md#getorgsorgidappsappidvaluesetversions) - List Value Set Versions in the App
* [GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionID](docs/valuesetversion/README.md#getorgsorgidappsappidvaluesetversionsvaluesetversionid) - Get a single Value Set Version from the App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKey](docs/valuesetversion/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Environment Version history.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestore](docs/valuesetversion/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKey](docs/valuesetversion/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKey](docs/valuesetversion/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Version history.
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestore](docs/valuesetversion/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKey](docs/valuesetversion/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an App

### [WorkloadProfile](docs/workloadprofile/README.md)

* [DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersion](docs/workloadprofile/README.md#deleteorgsorgidworkloadprofilesprofileidversionsversion) - Delete a Workload Profile Version
* [DeleteOrgsOrgIDWorkloadProfilesProfileQid](docs/workloadprofile/README.md#deleteorgsorgidworkloadprofilesprofileqid) - Delete a Workload Profile
* [GetOrgsOrgIDWorkloadProfiles](docs/workloadprofile/README.md#getorgsorgidworkloadprofiles) - List workload profiles available to the organization.
* [GetOrgsOrgIDWorkloadProfilesProfileQid](docs/workloadprofile/README.md#getorgsorgidworkloadprofilesprofileqid) - Get a Workload Profile
* [GetOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/workloadprofile/README.md#getorgsorgidworkloadprofilesprofileqidversions) - List versions of the given workload profile with optional constraint.
* [PostOrgsOrgIDWorkloadProfiles](docs/workloadprofile/README.md#postorgsorgidworkloadprofiles) - Create new Workload Profile
* [PostOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/workloadprofile/README.md#postorgsorgidworkloadprofilesprofileqidversions) - Add new Version of the Workload Profile

### [Public](docs/public/README.md)

* [DeleteOrgsOrgIDAppsAppID](docs/public/README.md#deleteorgsorgidappsappid) - Delete an Application
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvID](docs/public/README.md#deleteorgsorgidappsappidenvsenvid) - Delete a specific Environment.
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDResourcesTypeResID](docs/public/README.md#deleteorgsorgidappsappidenvsenvidresourcestyperesid) - Delete Active Resources.
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/public/README.md#deleteorgsorgidappsappidenvsenvidrulesruleid) - Delete Automation Rule from an Environment.
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/public/README.md#deleteorgsorgidappsappidenvsenvidvalues) - Delete all Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/public/README.md#deleteorgsorgidappsappidenvsenvidvalueskey) - Delete Shared Value for an Environment
* [DeleteOrgsOrgIDAppsAppIDJobs](docs/public/README.md#deleteorgsorgidappsappidjobs) - Deletes all Jobs for the Application
* [DeleteOrgsOrgIDAppsAppIDUsersUserID](docs/public/README.md#deleteorgsorgidappsappidusersuserid) - Remove the role of a User on an Application
* [DeleteOrgsOrgIDAppsAppIDValues](docs/public/README.md#deleteorgsorgidappsappidvalues) - Delete all Shared Value for an App
* [DeleteOrgsOrgIDAppsAppIDValuesKey](docs/public/README.md#deleteorgsorgidappsappidvalueskey) - Delete Shared Value for an Application
* [DeleteOrgsOrgIDAppsAppIDWebhooksJobID](docs/public/README.md#deleteorgsorgidappsappidwebhooksjobid) - Delete a Webhook
* [DeleteOrgsOrgIDArtefactsArtefactID](docs/public/README.md#deleteorgsorgidartefactsartefactid) - Delete Artefact and all related Artefact Versions
* [DeleteOrgsOrgIDEnvTypesEnvTypeID](docs/public/README.md#deleteorgsorgidenvtypesenvtypeid) - Deletes an Environment Type
* [DeleteOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/public/README.md#deleteorgsorgidenvtypesenvtypeusersuserid) - Remove the role of a User on an Environment Type
* [DeleteOrgsOrgIDRegistriesRegID](docs/public/README.md#deleteorgsorgidregistriesregid) - Deletes an existing registry record and all associated credentials and secrets.
* [DeleteOrgsOrgIDResourcesDefsDefID](docs/public/README.md#deleteorgsorgidresourcesdefsdefid) - Delete a Resource Definition.
* [DeleteOrgsOrgIDResourcesDefsDefIDCriteriaCriteriaID](docs/public/README.md#deleteorgsorgidresourcesdefsdefidcriteriacriteriaid) - Delete a Matching Criteria from a Resource Definition.
* [DeleteOrgsOrgIDResourcesDriversDriverID](docs/public/README.md#deleteorgsorgidresourcesdriversdriverid) - Delete a Resources Driver.
* [DeleteOrgsOrgIDUsersUserID](docs/public/README.md#deleteorgsorgidusersuserid) - Remove the role of a User on an Organization
* [DeleteOrgsOrgIDWorkloadProfilesProfileIDVersionsVersion](docs/public/README.md#deleteorgsorgidworkloadprofilesprofileidversionsversion) - Delete a Workload Profile Version
* [DeleteOrgsOrgIDWorkloadProfilesProfileQid](docs/public/README.md#deleteorgsorgidworkloadprofilesprofileqid) - Delete a Workload Profile
* [DeleteTokensTokenID](docs/public/README.md#deletetokenstokenid) - DEPRECATED
* [GetDelta](docs/public/README.md#getdelta) - Fetch an existing Delta
* [GetSets](docs/public/README.md#getsets) - Get all Deployment Sets
* [GetCurrentUser](docs/public/README.md#getcurrentuser) - Gets the extended profile of the current user
* [GetOrgs](docs/public/README.md#getorgs) - List active organizations the user has access to.
* [GetOrgsOrgID](docs/public/README.md#getorgsorgid) - Get the specified Organization.
* [GetOrgsOrgIDApps](docs/public/README.md#getorgsorgidapps) - List all Applications in an Organization.
* [GetOrgsOrgIDAppsAppID](docs/public/README.md#getorgsorgidappsappid) - Get an existing Application
* [GetOrgsOrgIDAppsAppIDDeltas](docs/public/README.md#getorgsorgidappsappiddeltas) - List Deltas in an Application
* [GetOrgsOrgIDAppsAppIDEnvs](docs/public/README.md#getorgsorgidappsappidenvs) - List all Environments.
* [GetOrgsOrgIDAppsAppIDEnvsEnvID](docs/public/README.md#getorgsorgidappsappidenvsenvid) - Get a specific Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/public/README.md#getorgsorgidappsappidenvsenviddeploys) - List Deployments in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployID](docs/public/README.md#getorgsorgidappsappidenvsenviddeploysdeployid) - Get a specific Deployment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDDeploysDeployIDErrors](docs/public/README.md#getorgsorgidappsappidenvsenviddeploysdeployiderrors) - List errors that occurred in a Deployment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDResources](docs/public/README.md#getorgsorgidappsappidenvsenvidresources) - List Active Resources provisioned in an environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/public/README.md#getorgsorgidappsappidenvsenvidrules) - List all Automation Rules in an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/public/README.md#getorgsorgidappsappidenvsenvidrulesruleid) - Get a specific Automation Rule for an Environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDRuntime](docs/public/README.md#getorgsorgidappsappidenvsenvidruntime) - Get Runtime information about the environment.
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersions](docs/public/README.md#getorgsorgidappsappidenvsenvidvaluesetversions) - List Value Set Versions in an Environment of an App
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionID](docs/public/README.md#getorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionid) - Get a single Value Set Version in an Environment of an App
* [GetOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/public/README.md#getorgsorgidappsappidenvsenvidvalues) - List Shared Values in an Environment
* [GetOrgsOrgIDAppsAppIDRuntime](docs/public/README.md#getorgsorgidappsappidruntime) - Get Runtime information about specific environments.
* [GetOrgsOrgIDAppsAppIDSetsSetID](docs/public/README.md#getorgsorgidappsappidsetssetid) - Get a Deployment Set
* [GetOrgsOrgIDAppsAppIDSetsSetIDDiffSourceSetID](docs/public/README.md#getorgsorgidappsappidsetssetiddiffsourcesetid) - Get the difference between 2 Deployment Sets
* [GetOrgsOrgIDAppsAppIDUsers](docs/public/README.md#getorgsorgidappsappidusers) - List Users with roles in an App
* [GetOrgsOrgIDAppsAppIDUsersUserID](docs/public/README.md#getorgsorgidappsappidusersuserid) - Get the role of a User on an Application
* [GetOrgsOrgIDAppsAppIDValueSetVersions](docs/public/README.md#getorgsorgidappsappidvaluesetversions) - List Value Set Versions in the App
* [GetOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionID](docs/public/README.md#getorgsorgidappsappidvaluesetversionsvaluesetversionid) - Get a single Value Set Version from the App
* [GetOrgsOrgIDAppsAppIDValues](docs/public/README.md#getorgsorgidappsappidvalues) - List Shared Values in an Application
* [GetOrgsOrgIDAppsAppIDWebhooks](docs/public/README.md#getorgsorgidappsappidwebhooks) - List Webhooks
* [GetOrgsOrgIDAppsAppIDWebhooksJobID](docs/public/README.md#getorgsorgidappsappidwebhooksjobid) - Get a Webhook
* [GetOrgsOrgIDArtefactVersions](docs/public/README.md#getorgsorgidartefactversions) - List all Artefacts Versions.
* [GetOrgsOrgIDArtefactVersionsArtefactVersionID](docs/public/README.md#getorgsorgidartefactversionsartefactversionid) - Get an Artefacts Versions.
* [GetOrgsOrgIDArtefacts](docs/public/README.md#getorgsorgidartefacts) - List all Artefacts.
* [GetOrgsOrgIDArtefactsArtefactIDVersions](docs/public/README.md#getorgsorgidartefactsartefactidversions) - List all Artefact Versions of an Artefact.
* [GetOrgsOrgIDEnvTypes](docs/public/README.md#getorgsorgidenvtypes) - List all Environment Types
* [GetOrgsOrgIDEnvTypesEnvTypeID](docs/public/README.md#getorgsorgidenvtypesenvtypeid) - Get an Environment Type
* [GetOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/public/README.md#getorgsorgidenvtypesenvtypeusersuserid) - Get the role of a User on an Environment Type
* [GetOrgsOrgIDEvents](docs/public/README.md#getorgsorgidevents) - List Events
* [GetOrgsOrgIDImages](docs/public/README.md#getorgsorgidimages) - List all Container Images
* [GetOrgsOrgIDImagesImageID](docs/public/README.md#getorgsorgidimagesimageid) - Get a specific Image Object
* [GetOrgsOrgIDImagesImageIDBuilds](docs/public/README.md#getorgsorgidimagesimageidbuilds) - Lists all the Builds of an Image
* [GetOrgsOrgIDInvitations](docs/public/README.md#getorgsorgidinvitations) - List the invites issued for the organization.
* [GetOrgsOrgIDRegistries](docs/public/README.md#getorgsorgidregistries) - Lists available registries for the organization.
* [GetOrgsOrgIDRegistriesRegID](docs/public/README.md#getorgsorgidregistriesregid) - Loads a registry record details.
* [GetOrgsOrgIDRegistriesRegIDCreds](docs/public/README.md#getorgsorgidregistriesregidcreds) - Returns current account credentials or secret details for the registry.
* [GetOrgsOrgIDResourcesAccountTypes](docs/public/README.md#getorgsorgidresourcesaccounttypes) - List Resource Account Types available to the organization.
* [GetOrgsOrgIDResourcesAccounts](docs/public/README.md#getorgsorgidresourcesaccounts) - List Resource Accounts in the organization.
* [GetOrgsOrgIDResourcesAccountsAccID](docs/public/README.md#getorgsorgidresourcesaccountsaccid) - Get a Resource Account.
* [GetOrgsOrgIDResourcesDefs](docs/public/README.md#getorgsorgidresourcesdefs) - List Resource Definitions.
* [GetOrgsOrgIDResourcesDefsDefID](docs/public/README.md#getorgsorgidresourcesdefsdefid) - Get a specific Resource Definition.
* [GetOrgsOrgIDResourcesDefsDefIDResources](docs/public/README.md#getorgsorgidresourcesdefsdefidresources) - List Active Resources provisioned via a specific Resource Definition.
* [GetOrgsOrgIDResourcesDrivers](docs/public/README.md#getorgsorgidresourcesdrivers) - List Resource Drivers.
* [GetOrgsOrgIDResourcesDriversDriverID](docs/public/README.md#getorgsorgidresourcesdriversdriverid) - Get a Resource Driver.
* [GetOrgsOrgIDResourcesTypes](docs/public/README.md#getorgsorgidresourcestypes) - List Resource Types.
* [GetOrgsOrgIDUsers](docs/public/README.md#getorgsorgidusers) - List Users with roles in an Organization
* [GetOrgsOrgIDUsersUserID](docs/public/README.md#getorgsorgidusersuserid) - Get the role of a User on an Organization
* [GetOrgsOrgIDWorkloadProfiles](docs/public/README.md#getorgsorgidworkloadprofiles) - List workload profiles available to the organization.
* [GetOrgsOrgIDWorkloadProfilesProfileQid](docs/public/README.md#getorgsorgidworkloadprofilesprofileqid) - Get a Workload Profile
* [GetOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/public/README.md#getorgsorgidworkloadprofilesprofileqidversions) - List versions of the given workload profile with optional constraint.
* [GetTokens](docs/public/README.md#gettokens) - DEPRECATED
* [GetUsersMe](docs/public/README.md#getusersme) - DEPRECATED
* [PatchCurrentUser](docs/public/README.md#patchcurrentuser) - Updates the extended profile of the current user.
* [PatchOrgsOrgIDAppsAppIDDeltasDeltaID](docs/public/README.md#patchorgsorgidappsappiddeltasdeltaid) - Update an existing Delta
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDRuntimeReplicas](docs/public/README.md#patchorgsorgidappsappidenvsenvidruntimereplicas) - Set number of replicas for an environment's modules.
* [PatchOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/public/README.md#patchorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PatchOrgsOrgIDAppsAppIDUsersUserID](docs/public/README.md#patchorgsorgidappsappidusersuserid) - Update the role of a User on an Application
* [PatchOrgsOrgIDAppsAppIDValuesKey](docs/public/README.md#patchorgsorgidappsappidvalueskey) - Update Shared Value for an Application
* [PatchOrgsOrgIDArtefactsArtefactIDVersionsVersionID](docs/public/README.md#patchorgsorgidartefactsartefactidversionsversionid) - Update Version of an Artefact.
* [PatchOrgsOrgIDEnvTypesEnvTypeUsersUserID](docs/public/README.md#patchorgsorgidenvtypesenvtypeusersuserid) - Update the role of a User on an Environment Type
* [PatchOrgsOrgIDRegistriesRegID](docs/public/README.md#patchorgsorgidregistriesregid) - Updates (patches) an existing registry record.
* [PatchOrgsOrgIDResourcesAccountsAccID](docs/public/README.md#patchorgsorgidresourcesaccountsaccid) - Update a Resource Account.
* [PatchOrgsOrgIDResourcesDefsDefID](docs/public/README.md#patchorgsorgidresourcesdefsdefid) - Update a Resource Definition.
* [PatchOrgsOrgIDUsersUserID](docs/public/README.md#patchorgsorgidusersuserid) - Update the role of a User on an Organization
* [PostOrgsOrgIDApps](docs/public/README.md#postorgsorgidapps) - Add a new Application to an Organization
* [PostOrgsOrgIDAppsAppIDDeltas](docs/public/README.md#postorgsorgidappsappiddeltas) - Create a new Delta
* [PostOrgsOrgIDAppsAppIDEnvs](docs/public/README.md#postorgsorgidappsappidenvs) - Add a new Environment to an Application.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDDeploys](docs/public/README.md#postorgsorgidappsappidenvsenviddeploys) - Start a new Deployment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDRules](docs/public/README.md#postorgsorgidappsappidenvsenvidrules) - Create a new Automation Rule for an Environment.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDPurgeKey](docs/public/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Environment Version history.
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestore](docs/public/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValueSetVersionsValueSetVersionIDRestoreKey](docs/public/README.md#postorgsorgidappsappidenvsenvidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an Environment of an App
* [PostOrgsOrgIDAppsAppIDEnvsEnvIDValues](docs/public/README.md#postorgsorgidappsappidenvsenvidvalues) - Create a Shared Value for an Environment
* [PostOrgsOrgIDAppsAppIDSetsSetID](docs/public/README.md#postorgsorgidappsappidsetssetid) - Apply a Deployment Delta to a Deployment Set
* [PostOrgsOrgIDAppsAppIDUsers](docs/public/README.md#postorgsorgidappsappidusers) - Adds a User to an Application with a Role
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDPurgeKey](docs/public/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidpurgekey) - Purge the value of a specific Shared Value from the App Version history.
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestore](docs/public/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestore) - Restore a Value Set Version in an App
* [PostOrgsOrgIDAppsAppIDValueSetVersionsValueSetVersionIDRestoreKey](docs/public/README.md#postorgsorgidappsappidvaluesetversionsvaluesetversionidrestorekey) - Restore a specific key from the Value Set Version in an App
* [PostOrgsOrgIDAppsAppIDValues](docs/public/README.md#postorgsorgidappsappidvalues) - Create a Shared Value for an Application
* [PostOrgsOrgIDAppsAppIDWebhooks](docs/public/README.md#postorgsorgidappsappidwebhooks) - Create a new Webhook
* [PostOrgsOrgIDAppsAppIDWebhooksJobID](docs/public/README.md#postorgsorgidappsappidwebhooksjobid) - Update a Webhook
* [PostOrgsOrgIDArtefactVersions](docs/public/README.md#postorgsorgidartefactversions) - Register a new Artefact Version with your organization.
* [PostOrgsOrgIDEnvTypes](docs/public/README.md#postorgsorgidenvtypes) - Add a new Environment Type
* [PostOrgsOrgIDEnvTypesEnvTypeUsers](docs/public/README.md#postorgsorgidenvtypesenvtypeusers) - Adds a User to an Environment Type with a Role
* [PostOrgsOrgIDImagesImageIDBuilds](docs/public/README.md#postorgsorgidimagesimageidbuilds) - Add a new Image Build
* [PostOrgsOrgIDInvitations](docs/public/README.md#postorgsorgidinvitations) - Invites a user to an Organization with a specified role.
* [PostOrgsOrgIDRegistries](docs/public/README.md#postorgsorgidregistries) - Creates a new registry record.
* [PostOrgsOrgIDResourcesAccounts](docs/public/README.md#postorgsorgidresourcesaccounts) - Create a new Resource Account in the organization.
* [PostOrgsOrgIDResourcesDefs](docs/public/README.md#postorgsorgidresourcesdefs) - Create a new Resource Definition.
* [PostOrgsOrgIDResourcesDefsDefIDCriteria](docs/public/README.md#postorgsorgidresourcesdefsdefidcriteria) - Add a new Matching Criteria to a Resource Definition.
* [PostOrgsOrgIDResourcesDrivers](docs/public/README.md#postorgsorgidresourcesdrivers) - Register a new Resource Driver.
* [PostOrgsOrgIDUsers](docs/public/README.md#postorgsorgidusers) - Creates a new service user.
* [PostOrgsOrgIDWorkloadProfiles](docs/public/README.md#postorgsorgidworkloadprofiles) - Create new Workload Profile
* [PostOrgsOrgIDWorkloadProfilesProfileQidVersions](docs/public/README.md#postorgsorgidworkloadprofilesprofileqidversions) - Add new Version of the Workload Profile
* [PutDelta](docs/public/README.md#putdelta) - Update an existing Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataArchived](docs/public/README.md#putorgsorgidappsappiddeltasdeltaidmetadataarchived) - Mark a Delta as "archived"
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataEnvID](docs/public/README.md#putorgsorgidappsappiddeltasdeltaidmetadataenvid) - Change the Environment of a Delta
* [PutOrgsOrgIDAppsAppIDDeltasDeltaIDMetadataName](docs/public/README.md#putorgsorgidappsappiddeltasdeltaidmetadataname) - Change the name of a Delta
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDFromDeployID](docs/public/README.md#putorgsorgidappsappidenvsenvidfromdeployid) - Rebase to a different Deployment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRulesRuleID](docs/public/README.md#putorgsorgidappsappidenvsenvidrulesruleid) - Update an existing Automation Rule for an Environment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDRuntimePaused](docs/public/README.md#putorgsorgidappsappidenvsenvidruntimepaused) - Pause / Resume an environment.
* [PutOrgsOrgIDAppsAppIDEnvsEnvIDValuesKey](docs/public/README.md#putorgsorgidappsappidenvsenvidvalueskey) - Update Shared Value for an Environment
* [PutOrgsOrgIDAppsAppIDValuesKey](docs/public/README.md#putorgsorgidappsappidvalueskey) - Update Shared Value for an Application
* [PutOrgsOrgIDResourcesDefsDefID](docs/public/README.md#putorgsorgidresourcesdefsdefid) - Update a Resource Definition.
* [PutOrgsOrgIDResourcesDriversDriverID](docs/public/README.md#putorgsorgidresourcesdriversdriverid) - Update a Resource Driver.
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
