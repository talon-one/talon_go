# TierDowngradeNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the notification. | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**BatchSize** | Pointer to **int32** | The required size of each batch of data. This value applies only when &#x60;batchingEnabled&#x60; is &#x60;true&#x60;. | [optional] 

## Methods

### GetName

`func (o *TierDowngradeNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TierDowngradeNotificationPolicy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *TierDowngradeNotificationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *TierDowngradeNotificationPolicy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetBatchingEnabled

`func (o *TierDowngradeNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *TierDowngradeNotificationPolicy) GetBatchingEnabledOk() (bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchingEnabled

`func (o *TierDowngradeNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### SetBatchingEnabled

`func (o *TierDowngradeNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.

### GetBatchSize

`func (o *TierDowngradeNotificationPolicy) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *TierDowngradeNotificationPolicy) GetBatchSizeOk() (int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchSize

`func (o *TierDowngradeNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### SetBatchSize

`func (o *TierDowngradeNotificationPolicy) SetBatchSize(v int32)`

SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


