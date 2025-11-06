# RoleV2PermissionSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the permission set. | 
**LogicalOperations** | Pointer to **[]string** | List of logical operations in the permission set. Each logical operation must be shown under the &#x60;x-permission&#x60; tag on an endpoint level.  | 

## Methods

### NewRoleV2PermissionSet

`func NewRoleV2PermissionSet(name string, logicalOperations []string, ) *RoleV2PermissionSet`

NewRoleV2PermissionSet instantiates a new RoleV2PermissionSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleV2PermissionSetWithDefaults

`func NewRoleV2PermissionSetWithDefaults() *RoleV2PermissionSet`

NewRoleV2PermissionSetWithDefaults instantiates a new RoleV2PermissionSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleV2PermissionSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleV2PermissionSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleV2PermissionSet) SetName(v string)`

SetName sets Name field to given value.


### GetLogicalOperations

`func (o *RoleV2PermissionSet) GetLogicalOperations() []string`

GetLogicalOperations returns the LogicalOperations field if non-nil, zero value otherwise.

### GetLogicalOperationsOk

`func (o *RoleV2PermissionSet) GetLogicalOperationsOk() (*[]string, bool)`

GetLogicalOperationsOk returns a tuple with the LogicalOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalOperations

`func (o *RoleV2PermissionSet) SetLogicalOperations(v []string)`

SetLogicalOperations sets LogicalOperations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


