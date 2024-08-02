# ExpiringCouponsNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**Name** | Pointer to **string** | Notification name. | 
**Triggers** | Pointer to [**[]ExpiringCouponsNotificationTrigger**](ExpiringCouponsNotificationTrigger.md) |  | 

## Methods

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


