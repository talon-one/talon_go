# RoleV2PermissionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the permission set. | 
**LogicalOperations** | Pointer to **[]string** | List of logical operations in the permission set. Each logical operation must be shown under the &#x60;x-permission&#x60; tag on an endpoint level.  | 

## Methods

### GetName

`func (o *RoleV2PermissionSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleV2PermissionSet) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *RoleV2PermissionSet) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *RoleV2PermissionSet) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetLogicalOperations

`func (o *RoleV2PermissionSet) GetLogicalOperations() []string`

GetLogicalOperations returns the LogicalOperations field if non-nil, zero value otherwise.

### GetLogicalOperationsOk

`func (o *RoleV2PermissionSet) GetLogicalOperationsOk() ([]string, bool)`

GetLogicalOperationsOk returns a tuple with the LogicalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLogicalOperations

`func (o *RoleV2PermissionSet) HasLogicalOperations() bool`

HasLogicalOperations returns a boolean if a field has been set.

### SetLogicalOperations

`func (o *RoleV2PermissionSet) SetLogicalOperations(v []string)`

SetLogicalOperations gets a reference to the given []string and assigns it to the LogicalOperations field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


