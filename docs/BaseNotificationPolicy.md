# BaseNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the notification. | 
**Triggers** | Pointer to [**[]TierWillDowngradeNotificationTrigger**](TierWillDowngradeNotificationTrigger.md) |  | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**Scopes** | Pointer to **[]string** |  | 

## Methods

### GetName

`func (o *BaseNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseNotificationPolicy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *BaseNotificationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *BaseNotificationPolicy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTriggers

`func (o *BaseNotificationPolicy) GetTriggers() []TierWillDowngradeNotificationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *BaseNotificationPolicy) GetTriggersOk() ([]TierWillDowngradeNotificationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggers

`func (o *BaseNotificationPolicy) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggers

`func (o *BaseNotificationPolicy) SetTriggers(v []TierWillDowngradeNotificationTrigger)`

SetTriggers gets a reference to the given []TierWillDowngradeNotificationTrigger and assigns it to the Triggers field.

### GetBatchingEnabled

`func (o *BaseNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *BaseNotificationPolicy) GetBatchingEnabledOk() (bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchingEnabled

`func (o *BaseNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### SetBatchingEnabled

`func (o *BaseNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.

### GetScopes

`func (o *BaseNotificationPolicy) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *BaseNotificationPolicy) GetScopesOk() ([]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScopes

`func (o *BaseNotificationPolicy) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopes

`func (o *BaseNotificationPolicy) SetScopes(v []string)`

SetScopes gets a reference to the given []string and assigns it to the Scopes field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


