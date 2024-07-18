# CouponCreationJobAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | Pointer to **string** | The batch ID coupons created by this job will bear. | 
**Status** | Pointer to **string** | The current status of this request. Possible values: - &#x60;pending verification&#x60; - &#x60;pending&#x60; - &#x60;completed&#x60; - &#x60;failed&#x60; - &#x60;coupon pattern full&#x60;  | 
**CreatedAmount** | Pointer to **int32** | The number of coupon codes that were already created for this request. | 
**FailCount** | Pointer to **int32** | The number of times this job failed. | 
**Errors** | Pointer to **[]string** | An array of individual problems encountered during the request. | 
**CreatedBy** | Pointer to **int32** | ID of the user who created this effect. | 
**Communicated** | Pointer to **bool** | Whether or not the user that created this job was notified of its final state. | 
**ChunkExecutionCount** | Pointer to **int32** | The number of times an attempt to create a chunk of coupons was made during the processing of the job. | 
**ChunkSize** | Pointer to **int32** | The number of coupons that will be created in a single transactions. Coupons will be created in chunks until arriving at the requested amount. | [optional] 

## Methods

### GetBatchId

`func (o *CouponCreationJobAllOf) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CouponCreationJobAllOf) GetBatchIdOk() (string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchId

`func (o *CouponCreationJobAllOf) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchId

`func (o *CouponCreationJobAllOf) SetBatchId(v string)`

SetBatchId gets a reference to the given string and assigns it to the BatchId field.

### GetStatus

`func (o *CouponCreationJobAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponCreationJobAllOf) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *CouponCreationJobAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *CouponCreationJobAllOf) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetCreatedAmount

`func (o *CouponCreationJobAllOf) GetCreatedAmount() int32`

GetCreatedAmount returns the CreatedAmount field if non-nil, zero value otherwise.

### GetCreatedAmountOk

`func (o *CouponCreationJobAllOf) GetCreatedAmountOk() (int32, bool)`

GetCreatedAmountOk returns a tuple with the CreatedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAmount

`func (o *CouponCreationJobAllOf) HasCreatedAmount() bool`

HasCreatedAmount returns a boolean if a field has been set.

### SetCreatedAmount

`func (o *CouponCreationJobAllOf) SetCreatedAmount(v int32)`

SetCreatedAmount gets a reference to the given int32 and assigns it to the CreatedAmount field.

### GetFailCount

`func (o *CouponCreationJobAllOf) GetFailCount() int32`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *CouponCreationJobAllOf) GetFailCountOk() (int32, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailCount

`func (o *CouponCreationJobAllOf) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### SetFailCount

`func (o *CouponCreationJobAllOf) SetFailCount(v int32)`

SetFailCount gets a reference to the given int32 and assigns it to the FailCount field.

### GetErrors

`func (o *CouponCreationJobAllOf) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CouponCreationJobAllOf) GetErrorsOk() ([]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *CouponCreationJobAllOf) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *CouponCreationJobAllOf) SetErrors(v []string)`

SetErrors gets a reference to the given []string and assigns it to the Errors field.

### GetCreatedBy

`func (o *CouponCreationJobAllOf) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CouponCreationJobAllOf) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *CouponCreationJobAllOf) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *CouponCreationJobAllOf) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetCommunicated

`func (o *CouponCreationJobAllOf) GetCommunicated() bool`

GetCommunicated returns the Communicated field if non-nil, zero value otherwise.

### GetCommunicatedOk

`func (o *CouponCreationJobAllOf) GetCommunicatedOk() (bool, bool)`

GetCommunicatedOk returns a tuple with the Communicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommunicated

`func (o *CouponCreationJobAllOf) HasCommunicated() bool`

HasCommunicated returns a boolean if a field has been set.

### SetCommunicated

`func (o *CouponCreationJobAllOf) SetCommunicated(v bool)`

SetCommunicated gets a reference to the given bool and assigns it to the Communicated field.

### GetChunkExecutionCount

`func (o *CouponCreationJobAllOf) GetChunkExecutionCount() int32`

GetChunkExecutionCount returns the ChunkExecutionCount field if non-nil, zero value otherwise.

### GetChunkExecutionCountOk

`func (o *CouponCreationJobAllOf) GetChunkExecutionCountOk() (int32, bool)`

GetChunkExecutionCountOk returns a tuple with the ChunkExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChunkExecutionCount

`func (o *CouponCreationJobAllOf) HasChunkExecutionCount() bool`

HasChunkExecutionCount returns a boolean if a field has been set.

### SetChunkExecutionCount

`func (o *CouponCreationJobAllOf) SetChunkExecutionCount(v int32)`

SetChunkExecutionCount gets a reference to the given int32 and assigns it to the ChunkExecutionCount field.

### GetChunkSize

`func (o *CouponCreationJobAllOf) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *CouponCreationJobAllOf) GetChunkSizeOk() (int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChunkSize

`func (o *CouponCreationJobAllOf) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSize

`func (o *CouponCreationJobAllOf) SetChunkSize(v int32)`

SetChunkSize gets a reference to the given int32 and assigns it to the ChunkSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


