# RoleV2RolesGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**map[string]RoleV2ApplicationDetails**](RoleV2ApplicationDetails.md) | A map of the link between the Application, campaign, or draft campaign-related permission set and the Application ID the permissions apply to. | [optional] 
**LoyaltyPrograms** | Pointer to **map[string]string** | A map of the link between the loyalty program-related permission set and the Application ID the permissions apply to. | [optional] 
**CampaignAccessGroups** | Pointer to **map[string]string** | A map of the link between the campaign access group-related permission set and the Application ID the permissions apply to. | [optional] 
**Account** | Pointer to **string** | Name of the account-level permission set | [optional] 

## Methods

### NewRoleV2RolesGroup

`func NewRoleV2RolesGroup() *RoleV2RolesGroup`

NewRoleV2RolesGroup instantiates a new RoleV2RolesGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleV2RolesGroupWithDefaults

`func NewRoleV2RolesGroupWithDefaults() *RoleV2RolesGroup`

NewRoleV2RolesGroupWithDefaults instantiates a new RoleV2RolesGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *RoleV2RolesGroup) GetApplications() map[string]RoleV2ApplicationDetails`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RoleV2RolesGroup) GetApplicationsOk() (*map[string]RoleV2ApplicationDetails, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *RoleV2RolesGroup) SetApplications(v map[string]RoleV2ApplicationDetails)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *RoleV2RolesGroup) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetLoyaltyPrograms

`func (o *RoleV2RolesGroup) GetLoyaltyPrograms() map[string]string`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *RoleV2RolesGroup) GetLoyaltyProgramsOk() (*map[string]string, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPrograms

`func (o *RoleV2RolesGroup) SetLoyaltyPrograms(v map[string]string)`

SetLoyaltyPrograms sets LoyaltyPrograms field to given value.

### HasLoyaltyPrograms

`func (o *RoleV2RolesGroup) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### GetCampaignAccessGroups

`func (o *RoleV2RolesGroup) GetCampaignAccessGroups() map[string]string`

GetCampaignAccessGroups returns the CampaignAccessGroups field if non-nil, zero value otherwise.

### GetCampaignAccessGroupsOk

`func (o *RoleV2RolesGroup) GetCampaignAccessGroupsOk() (*map[string]string, bool)`

GetCampaignAccessGroupsOk returns a tuple with the CampaignAccessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignAccessGroups

`func (o *RoleV2RolesGroup) SetCampaignAccessGroups(v map[string]string)`

SetCampaignAccessGroups sets CampaignAccessGroups field to given value.

### HasCampaignAccessGroups

`func (o *RoleV2RolesGroup) HasCampaignAccessGroups() bool`

HasCampaignAccessGroups returns a boolean if a field has been set.

### GetAccount

`func (o *RoleV2RolesGroup) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *RoleV2RolesGroup) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *RoleV2RolesGroup) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *RoleV2RolesGroup) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


