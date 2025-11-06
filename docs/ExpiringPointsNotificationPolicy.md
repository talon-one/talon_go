# ExpiringPointsNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Notification name. | 
**Triggers** | Pointer to [**[]ExpiringPointsNotificationTrigger**](ExpiringPointsNotificationTrigger.md) |  | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**BatchSize** | Pointer to **int64** | The required size of each batch of data. This value applies only when &#x60;batchingEnabled&#x60; is &#x60;true&#x60;. | [optional] [default to 1000]

## Methods

### NewExpiringPointsNotificationPolicy

`func NewExpiringPointsNotificationPolicy(name string, triggers []ExpiringPointsNotificationTrigger, ) *ExpiringPointsNotificationPolicy`

NewExpiringPointsNotificationPolicy instantiates a new ExpiringPointsNotificationPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpiringPointsNotificationPolicyWithDefaults

`func NewExpiringPointsNotificationPolicyWithDefaults() *ExpiringPointsNotificationPolicy`

NewExpiringPointsNotificationPolicyWithDefaults instantiates a new ExpiringPointsNotificationPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExpiringPointsNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpiringPointsNotificationPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpiringPointsNotificationPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetTriggers

`func (o *ExpiringPointsNotificationPolicy) GetTriggers() []ExpiringPointsNotificationTrigger`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ExpiringPointsNotificationPolicy) GetTriggersOk() (*[]ExpiringPointsNotificationTrigger, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ExpiringPointsNotificationPolicy) SetTriggers(v []ExpiringPointsNotificationTrigger)`

SetTriggers sets Triggers field to given value.


### GetBatchingEnabled

`func (o *ExpiringPointsNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *ExpiringPointsNotificationPolicy) GetBatchingEnabledOk() (*bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchingEnabled

`func (o *ExpiringPointsNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled sets BatchingEnabled field to given value.

### HasBatchingEnabled

`func (o *ExpiringPointsNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### GetBatchSize

`func (o *ExpiringPointsNotificationPolicy) GetBatchSize() int64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *ExpiringPointsNotificationPolicy) GetBatchSizeOk() (*int64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *ExpiringPointsNotificationPolicy) SetBatchSize(v int64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *ExpiringPointsNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


