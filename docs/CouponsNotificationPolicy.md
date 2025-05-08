# CouponsNotificationPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Notification name. | 
**Scopes** | Pointer to **[]string** |  | 
**BatchingEnabled** | Pointer to **bool** | Indicates whether batching is activated. | [optional] [default to true]
**IncludeData** | Pointer to **bool** | Indicates whether to include all generated coupons. If &#x60;false&#x60;, only the &#x60;batchId&#x60; of the generated coupons is included. | [optional] 
**BatchSize** | Pointer to **int32** | The required size of each batch of data. This value applies only when &#x60;batchingEnabled&#x60; is &#x60;true&#x60;. | [optional] 

## Methods

### GetName

`func (o *CouponsNotificationPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CouponsNotificationPolicy) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *CouponsNotificationPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *CouponsNotificationPolicy) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetScopes

`func (o *CouponsNotificationPolicy) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CouponsNotificationPolicy) GetScopesOk() ([]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasScopes

`func (o *CouponsNotificationPolicy) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### SetScopes

`func (o *CouponsNotificationPolicy) SetScopes(v []string)`

SetScopes gets a reference to the given []string and assigns it to the Scopes field.

### GetBatchingEnabled

`func (o *CouponsNotificationPolicy) GetBatchingEnabled() bool`

GetBatchingEnabled returns the BatchingEnabled field if non-nil, zero value otherwise.

### GetBatchingEnabledOk

`func (o *CouponsNotificationPolicy) GetBatchingEnabledOk() (bool, bool)`

GetBatchingEnabledOk returns a tuple with the BatchingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchingEnabled

`func (o *CouponsNotificationPolicy) HasBatchingEnabled() bool`

HasBatchingEnabled returns a boolean if a field has been set.

### SetBatchingEnabled

`func (o *CouponsNotificationPolicy) SetBatchingEnabled(v bool)`

SetBatchingEnabled gets a reference to the given bool and assigns it to the BatchingEnabled field.

### GetIncludeData

`func (o *CouponsNotificationPolicy) GetIncludeData() bool`

GetIncludeData returns the IncludeData field if non-nil, zero value otherwise.

### GetIncludeDataOk

`func (o *CouponsNotificationPolicy) GetIncludeDataOk() (bool, bool)`

GetIncludeDataOk returns a tuple with the IncludeData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIncludeData

`func (o *CouponsNotificationPolicy) HasIncludeData() bool`

HasIncludeData returns a boolean if a field has been set.

### SetIncludeData

`func (o *CouponsNotificationPolicy) SetIncludeData(v bool)`

SetIncludeData gets a reference to the given bool and assigns it to the IncludeData field.

### GetBatchSize

`func (o *CouponsNotificationPolicy) GetBatchSize() int32`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *CouponsNotificationPolicy) GetBatchSizeOk() (int32, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchSize

`func (o *CouponsNotificationPolicy) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### SetBatchSize

`func (o *CouponsNotificationPolicy) SetBatchSize(v int32)`

SetBatchSize gets a reference to the given int32 and assigns it to the BatchSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


