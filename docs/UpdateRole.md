# UpdateRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the role. | [optional] 
**Description** | Pointer to **string** | Description of the role. | [optional] 
**Acl** | Pointer to **string** | The &#x60;Access Control List&#x60; json defining the role of the user. This represents the access control on the user level. | [optional] 
**Members** | Pointer to **[]int64** | An array of user identifiers. | [optional] 

## Methods

### NewUpdateRole

`func NewUpdateRole() *UpdateRole`

NewUpdateRole instantiates a new UpdateRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleWithDefaults

`func NewUpdateRoleWithDefaults() *UpdateRole`

NewUpdateRoleWithDefaults instantiates a new UpdateRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAcl

`func (o *UpdateRole) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *UpdateRole) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *UpdateRole) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *UpdateRole) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetMembers

`func (o *UpdateRole) GetMembers() []int64`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *UpdateRole) GetMembersOk() (*[]int64, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *UpdateRole) SetMembers(v []int64)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *UpdateRole) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


