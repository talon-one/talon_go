# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the user. | [optional] 
**State** | Pointer to **string** | The state of the user.   - &#x60;deactivated&#x60;: The user has been deactivated.   - &#x60;active&#x60;: The user is active.  **Note**: Only &#x60;admin&#x60; users can update the state of another user.  | [optional] 
**IsAdmin** | Pointer to **bool** | Indicates whether the user is an &#x60;admin&#x60;. | [optional] 
**Policy** | Pointer to **string** | Indicates the access level of the user. | [optional] 
**Roles** | Pointer to **[]int32** | A list of the IDs of the roles assigned to the user.  **Note**: To find the ID of a role, use the [List roles](/management-api#tag/Roles/operation/listAllRolesV2) endpoint.  | [optional] 
**ApplicationNotificationSubscriptions** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Application notifications that the user is subscribed to. | [optional] 

## Methods

### GetName

`func (o *UpdateUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUser) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateUser) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateUser) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetState

`func (o *UpdateUser) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateUser) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *UpdateUser) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *UpdateUser) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.

### GetIsAdmin

`func (o *UpdateUser) GetIsAdmin() bool`

GetIsAdmin returns the IsAdmin field if non-nil, zero value otherwise.

### GetIsAdminOk

`func (o *UpdateUser) GetIsAdminOk() (bool, bool)`

GetIsAdminOk returns a tuple with the IsAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsAdmin

`func (o *UpdateUser) HasIsAdmin() bool`

HasIsAdmin returns a boolean if a field has been set.

### SetIsAdmin

`func (o *UpdateUser) SetIsAdmin(v bool)`

SetIsAdmin gets a reference to the given bool and assigns it to the IsAdmin field.

### GetPolicy

`func (o *UpdateUser) GetPolicy() string`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *UpdateUser) GetPolicyOk() (string, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPolicy

`func (o *UpdateUser) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### SetPolicy

`func (o *UpdateUser) SetPolicy(v string)`

SetPolicy gets a reference to the given string and assigns it to the Policy field.

### GetRoles

`func (o *UpdateUser) GetRoles() []int32`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateUser) GetRolesOk() ([]int32, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRoles

`func (o *UpdateUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRoles

`func (o *UpdateUser) SetRoles(v []int32)`

SetRoles gets a reference to the given []int32 and assigns it to the Roles field.

### GetApplicationNotificationSubscriptions

`func (o *UpdateUser) GetApplicationNotificationSubscriptions() map[string]map[string]interface{}`

GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field if non-nil, zero value otherwise.

### GetApplicationNotificationSubscriptionsOk

`func (o *UpdateUser) GetApplicationNotificationSubscriptionsOk() (map[string]map[string]interface{}, bool)`

GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationNotificationSubscriptions

`func (o *UpdateUser) HasApplicationNotificationSubscriptions() bool`

HasApplicationNotificationSubscriptions returns a boolean if a field has been set.

### SetApplicationNotificationSubscriptions

`func (o *UpdateUser) SetApplicationNotificationSubscriptions(v map[string]map[string]interface{})`

SetApplicationNotificationSubscriptions gets a reference to the given map[string]map[string]interface{} and assigns it to the ApplicationNotificationSubscriptions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


