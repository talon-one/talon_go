# TierWillDowngradeNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the notification. | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**Triggers** | Pointer to [**[]TierWillDowngradeNotificationTrigger**](TierWillDowngradeNotificationTrigger.md) |  | 

## Methods

### GetName

`func (o *TierWillDowngradeNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierWillDowngradeNotificationPolicy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TierWillDowngradeNotificationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TierWillDowngradeNotificationPolicy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetBatchingEnabled

`func (o *TierWillDowngradeNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *TierWillDowngradeNotificationPolicy) GetBatchingEnabledOk() (bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchingEnabled

`func (o *TierWillDowngradeNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### SetBatchingEnabled

`func (o *TierWillDowngradeNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.

### GetTriggers

`func (o *TierWillDowngradeNotificationPolicy) GetTriggers() []TierWillDowngradeNotificationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *TierWillDowngradeNotificationPolicy) GetTriggersOk() ([]TierWillDowngradeNotificationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggers

`func (o *TierWillDowngradeNotificationPolicy) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggers

`func (o *TierWillDowngradeNotificationPolicy) SetTriggers(v []TierWillDowngradeNotificationTrigger)`

SetTriggers gets a reference to the given []TierWillDowngradeNotificationTrigger and assigns it to the Triggers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


