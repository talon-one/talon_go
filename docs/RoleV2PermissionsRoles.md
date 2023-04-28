# RoleV2PermissionsRoles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**map[string]RoleV2ApplicationDetails**](RoleV2ApplicationDetails.md) |  | [optional] 
**LoyaltyPrograms** | Pointer to **map[string]string** |  | [optional] 
**CampaignAccessGroups** | Pointer to **map[string]string** |  | [optional] 

## Methods

### GetApplications

`func (o *RoleV2PermissionsRoles) GetApplications() map[string]RoleV2ApplicationDetails`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *RoleV2PermissionsRoles) GetApplicationsOk() (map[string]RoleV2ApplicationDetails, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplications

`func (o *RoleV2PermissionsRoles) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### SetApplications

`func (o *RoleV2PermissionsRoles) SetApplications(v map[string]RoleV2ApplicationDetails)`

SetApplications gets a reference to the given map[string]RoleV2ApplicationDetails and assigns it to the Applications field.

### GetLoyaltyPrograms

`func (o *RoleV2PermissionsRoles) GetLoyaltyPrograms() map[string]string`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *RoleV2PermissionsRoles) GetLoyaltyProgramsOk() (map[string]string, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyPrograms

`func (o *RoleV2PermissionsRoles) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### SetLoyaltyPrograms

`func (o *RoleV2PermissionsRoles) SetLoyaltyPrograms(v map[string]string)`

SetLoyaltyPrograms gets a reference to the given map[string]string and assigns it to the LoyaltyPrograms field.

### GetCampaignAccessGroups

`func (o *RoleV2PermissionsRoles) GetCampaignAccessGroups() map[string]string`

GetCampaignAccessGroups returns the CampaignAccessGroups field if non-nil, zero value otherwise.

### GetCampaignAccessGroupsOk

`func (o *RoleV2PermissionsRoles) GetCampaignAccessGroupsOk() (map[string]string, bool)`

GetCampaignAccessGroupsOk returns a tuple with the CampaignAccessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignAccessGroups

`func (o *RoleV2PermissionsRoles) HasCampaignAccessGroups() bool`

HasCampaignAccessGroups returns a boolean if a field has been set.

### SetCampaignAccessGroups

`func (o *RoleV2PermissionsRoles) SetCampaignAccessGroups(v map[string]string)`

SetCampaignAccessGroups gets a reference to the given map[string]string and assigns it to the CampaignAccessGroups field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


