# ExpiringPointsNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Notification name. | 
**Triggers** | Pointer to [**[]ExpiringPointsNotificationTrigger**](ExpiringPointsNotificationTrigger.md) |  | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**BatchSize** | Pointer to **int32** | The required size of each batch of data. This value applies only when &#x60;batchingEnabled&#x60; is &#x60;true&#x60;. | [optional] 

## Methods

### GetName

`func (o *ExpiringPointsNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpiringPointsNotificationPolicy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ExpiringPointsNotificationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ExpiringPointsNotificationPolicy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTriggers

`func (o *ExpiringPointsNotificationPolicy) GetTriggers() []ExpiringPointsNotificationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ExpiringPointsNotificationPolicy) GetTriggersOk() ([]ExpiringPointsNotificationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggers

`func (o *ExpiringPointsNotificationPolicy) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggers

`func (o *ExpiringPointsNotificationPolicy) SetTriggers(v []ExpiringPointsNotificationTrigger)`

SetTriggers gets a reference to the given []ExpiringPointsNotificationTrigger and assigns it to the Triggers field.

### GetBatchingEnabled

`func (o *ExpiringPointsNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *ExpiringPointsNotificationPolicy) GetBatchingEnabledOk() (bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchingEnabled

`func (o *ExpiringPointsNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### SetBatchingEnabled

`func (o *ExpiringPointsNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.

### GetBatchSize

`func (o *ExpiringPointsNotificationPolicy) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *ExpiringPointsNotificationPolicy) GetBatchSizeOk() (int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchSize

`func (o *ExpiringPointsNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### SetBatchSize

`func (o *ExpiringPointsNotificationPolicy) SetBatchSize(v int32)`

SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


