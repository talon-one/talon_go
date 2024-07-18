# RoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignGroupID** | Pointer to **int32** | The ID of the [Campaign Group](https://docs.talon.one/docs/product/account/account-settings/managing-campaign-groups) this role was created for.  | [optional] 
**Name** | Pointer to **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Members** | Pointer to **[]int32** | A list of user identifiers assigned to this role. | [optional] 
**Acl** | Pointer to [**map[string]interface{}**](.md) | The &#x60;Access Control List&#x60; json defining the role of the user. This represents the access control on the user level. | 

## Methods

### GetCampaignGroupID

`func (o *RoleAllOf) GetCampaignGroupID() int32`

GetCampaignGroupID returns the CampaignGroupID field if non-nil, zero value otherwise.

### GetCampaignGroupIDOk

`func (o *RoleAllOf) GetCampaignGroupIDOk() (int32, bool)`

GetCampaignGroupIDOk returns a tuple with the CampaignGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignGroupID

`func (o *RoleAllOf) HasCampaignGroupID() bool`

HasCampaignGroupID returns a boolean if a field has been set.

### SetCampaignGroupID

`func (o *RoleAllOf) SetCampaignGroupID(v int32)`

SetCampaignGroupID gets a reference to the given int32 and assigns it to the CampaignGroupID field.

### GetName

`func (o *RoleAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleAllOf) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *RoleAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *RoleAllOf) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *RoleAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleAllOf) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *RoleAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *RoleAllOf) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetMembers

`func (o *RoleAllOf) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *RoleAllOf) GetMembersOk() ([]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *RoleAllOf) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *RoleAllOf) SetMembers(v []int32)`

SetMembers gets a reference to the given []int32 and assigns it to the Members field.

### GetAcl

`func (o *RoleAllOf) GetAcl() map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *RoleAllOf) GetAclOk() (map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcl

`func (o *RoleAllOf) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### SetAcl

`func (o *RoleAllOf) SetAcl(v map[string]interface{})`

SetAcl gets a reference to the given map[string]interface{} and assigns it to the Acl field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


