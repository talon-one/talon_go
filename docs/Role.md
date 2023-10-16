# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**CampaignGroupID** | Pointer to **int32** | The ID of the [Campaign Group](https://docs.talon.one/docs/product/account/account-settings/managing-campaign-groups) this role was created for.  | [optional] 
**Name** | Pointer to **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Members** | Pointer to **[]int32** | A list of user identifiers assigned to this role. | [optional] 
**Acl** | Pointer to [**map[string]interface{}**](.md) | The &#x60;Access Control List&#x60; json defining the role of the user. This represents the access control on the user level. | 

## Methods

### GetId

`func (o *Role) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Role) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Role) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Role) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Role) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Role) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Role) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *Role) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Role) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *Role) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *Role) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetAccountId

`func (o *Role) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Role) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Role) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Role) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetCampaignGroupID

`func (o *Role) GetCampaignGroupID() int32`

GetCampaignGroupID returns the CampaignGroupID field if non-nil, zero value otherwise.

### GetCampaignGroupIDOk

`func (o *Role) GetCampaignGroupIDOk() (int32, bool)`

GetCampaignGroupIDOk returns a tuple with the CampaignGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignGroupID

`func (o *Role) HasCampaignGroupID() bool`

HasCampaignGroupID returns a boolean if a field has been set.

### SetCampaignGroupID

`func (o *Role) SetCampaignGroupID(v int32)`

SetCampaignGroupID gets a reference to the given int32 and assigns it to the CampaignGroupID field.

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Role) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetMembers

`func (o *Role) GetMembers() []int32`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Role) GetMembersOk() ([]int32, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMembers

`func (o *Role) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### SetMembers

`func (o *Role) SetMembers(v []int32)`

SetMembers gets a reference to the given []int32 and assigns it to the Members field.

### GetAcl

`func (o *Role) GetAcl() map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *Role) GetAclOk() (map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAcl

`func (o *Role) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### SetAcl

`func (o *Role) SetAcl(v map[string]interface{})`

SetAcl gets a reference to the given map[string]interface{} and assigns it to the Acl field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


