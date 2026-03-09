# RoleV2Readonly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsReadonly** | Pointer to **bool** | Identifies if the role is read-only. For read-only roles, you can only assign or unassign users. You cannot edit any other properties, such as the name, description, or permissions. The &#39;isReadonly&#39; property cannot be set for new or existing roles. It is reserved for predefined roles, such as the Talon.One support role. | [optional] [default to false]

## Methods

### NewRoleV2Readonly

`func NewRoleV2Readonly() *RoleV2Readonly`

NewRoleV2Readonly instantiates a new RoleV2Readonly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleV2ReadonlyWithDefaults

`func NewRoleV2ReadonlyWithDefaults() *RoleV2Readonly`

NewRoleV2ReadonlyWithDefaults instantiates a new RoleV2Readonly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsReadonly

`func (o *RoleV2Readonly) GetIsReadonly() bool`

GetIsReadonly returns the IsReadonly field if non-nil, zero value otherwise.

### GetIsReadonlyOk

`func (o *RoleV2Readonly) GetIsReadonlyOk() (*bool, bool)`

GetIsReadonlyOk returns a tuple with the IsReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReadonly

`func (o *RoleV2Readonly) SetIsReadonly(v bool)`

SetIsReadonly sets IsReadonly field to given value.

### HasIsReadonly

`func (o *RoleV2Readonly) HasIsReadonly() bool`

HasIsReadonly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


