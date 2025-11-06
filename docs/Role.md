# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**CampaignGroupID** | Pointer to **int64** | The ID of the [Campaign Group](https://docs.talon.one/docs/product/account/account-settings/managing-campaign-groups) this role was created for.  | [optional] 
**Name** | Pointer to **string** | Name of the role. | 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Members** | Pointer to **[]int64** | A list of user identifiers assigned to this role. | [optional] 
**Acl** | Pointer to [**map[string]interface{}**](.md) | The &#x60;Access Control List&#x60; json defining the role of the user. This represents the access control on the user level. | 

## Methods

### NewRole

`func NewRole(id int64, created time.Time, modified time.Time, accountId int64, name string, acl map[string]interface{}, ) *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Role) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Role) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Role) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Role) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Role) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Role) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Role) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Role) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Role) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAccountId

`func (o *Role) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Role) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Role) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetCampaignGroupID

`func (o *Role) GetCampaignGroupID() int64`

GetCampaignGroupID returns the CampaignGroupID field if non-nil, zero value otherwise.

### GetCampaignGroupIDOk

`func (o *Role) GetCampaignGroupIDOk() (*int64, bool)`

GetCampaignGroupIDOk returns a tuple with the CampaignGroupID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignGroupID

`func (o *Role) SetCampaignGroupID(v int64)`

SetCampaignGroupID sets CampaignGroupID field to given value.

### HasCampaignGroupID

`func (o *Role) HasCampaignGroupID() bool`

HasCampaignGroupID returns a boolean if a field has been set.

### GetName

`func (o *Role) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Role) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Role) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Role) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Role) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Role) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Role) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMembers

`func (o *Role) GetMembers() []int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Role) GetMembersOk() (*[]int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Role) SetMembers(v []int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Role) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetAcl

`func (o *Role) GetAcl() map[string]interface{}`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *Role) GetAclOk() (*map[string]interface{}, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *Role) SetAcl(v map[string]interface{})`

SetAcl sets Acl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


