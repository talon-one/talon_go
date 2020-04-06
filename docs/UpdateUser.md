# UpdateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Your name. | [optional] 
**Policy** | Pointer to **string** | a blob of acl json | [optional] 
**State** | Pointer to **string** | New state (\&quot;deactivated\&quot; or \&quot;active\&quot;) for the user. Only usable by admins for the user. | [optional] 
**ReleaseUpdate** | Pointer to **bool** | Update the user via email | [optional] 
**LatestFeature** | Pointer to **string** | The latest feature you&#39;ve been notified. | [optional] 
**Roles** | Pointer to **[]int32** | Update | [optional] 
**ApplicationNotificationSubscriptions** | Pointer to [**map[string]interface{}**](.md) |  | [optional] 

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

### GetReleaseUpdate

`func (o *UpdateUser) GetReleaseUpdate() bool`

GetReleaseUpdate returns the ReleaseUpdate field if non-nil, zero value otherwise.

### GetReleaseUpdateOk

`func (o *UpdateUser) GetReleaseUpdateOk() (bool, bool)`

GetReleaseUpdateOk returns a tuple with the ReleaseUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReleaseUpdate

`func (o *UpdateUser) HasReleaseUpdate() bool`

HasReleaseUpdate returns a boolean if a field has been set.

### SetReleaseUpdate

`func (o *UpdateUser) SetReleaseUpdate(v bool)`

SetReleaseUpdate gets a reference to the given bool and assigns it to the ReleaseUpdate field.

### GetLatestFeature

`func (o *UpdateUser) GetLatestFeature() string`

GetLatestFeature returns the LatestFeature field if non-nil, zero value otherwise.

### GetLatestFeatureOk

`func (o *UpdateUser) GetLatestFeatureOk() (string, bool)`

GetLatestFeatureOk returns a tuple with the LatestFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLatestFeature

`func (o *UpdateUser) HasLatestFeature() bool`

HasLatestFeature returns a boolean if a field has been set.

### SetLatestFeature

`func (o *UpdateUser) SetLatestFeature(v string)`

SetLatestFeature gets a reference to the given string and assigns it to the LatestFeature field.

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

`func (o *UpdateUser) GetApplicationNotificationSubscriptions() map[string]interface{}`

GetApplicationNotificationSubscriptions returns the ApplicationNotificationSubscriptions field if non-nil, zero value otherwise.

### GetApplicationNotificationSubscriptionsOk

`func (o *UpdateUser) GetApplicationNotificationSubscriptionsOk() (map[string]interface{}, bool)`

GetApplicationNotificationSubscriptionsOk returns a tuple with the ApplicationNotificationSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationNotificationSubscriptions

`func (o *UpdateUser) HasApplicationNotificationSubscriptions() bool`

HasApplicationNotificationSubscriptions returns a boolean if a field has been set.

### SetApplicationNotificationSubscriptions

`func (o *UpdateUser) SetApplicationNotificationSubscriptions(v map[string]interface{})`

SetApplicationNotificationSubscriptions gets a reference to the given map[string]interface{} and assigns it to the ApplicationNotificationSubscriptions field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


