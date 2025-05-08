# ExpiringCouponsNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Notification name. | 
**Triggers** | Pointer to [**[]ExpiringCouponsNotificationTrigger**](ExpiringCouponsNotificationTrigger.md) |  | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**BatchSize** | Pointer to **int32** | The required size of each batch of data. This value applies only when &#x60;batchingEnabled&#x60; is &#x60;true&#x60;. | [optional] 

## Methods

### GetName

`func (o *ExpiringCouponsNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpiringCouponsNotificationPolicy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *ExpiringCouponsNotificationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *ExpiringCouponsNotificationPolicy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTriggers

`func (o *ExpiringCouponsNotificationPolicy) GetTriggers() []ExpiringCouponsNotificationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ExpiringCouponsNotificationPolicy) GetTriggersOk() ([]ExpiringCouponsNotificationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggers

`func (o *ExpiringCouponsNotificationPolicy) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### SetTriggers

`func (o *ExpiringCouponsNotificationPolicy) SetTriggers(v []ExpiringCouponsNotificationTrigger)`

SetTriggers gets a reference to the given []ExpiringCouponsNotificationTrigger and assigns it to the Triggers field.

### GetBatchingEnabled

`func (o *ExpiringCouponsNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *ExpiringCouponsNotificationPolicy) GetBatchingEnabledOk() (bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchingEnabled

`func (o *ExpiringCouponsNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### SetBatchingEnabled

`func (o *ExpiringCouponsNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.

### GetBatchSize

`func (o *ExpiringCouponsNotificationPolicy) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *ExpiringCouponsNotificationPolicy) GetBatchSizeOk() (int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchSize

`func (o *ExpiringCouponsNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### SetBatchSize

`func (o *ExpiringCouponsNotificationPolicy) SetBatchSize(v int32)`

SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


