# LoyaltyCardBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfCards** | Pointer to **int32** | Number of loyalty cards in the batch. | 
**BatchId** | Pointer to **string** | ID of the loyalty card batch. | [optional] 
**Status** | Pointer to **string** | Status of the loyalty cards in the batch. | [optional] [default to STATUS_ACTIVE]

## Methods

### GetNumberOfCards

`func (o *LoyaltyCardBatch) GetNumberOfCards() int32`

GetNumberOfCards returns the NumberOfCards field if non-nil, zero value otherwise.

### GetNumberOfCardsOk

`func (o *LoyaltyCardBatch) GetNumberOfCardsOk() (int32, bool)`

GetNumberOfCardsOk returns a tuple with the NumberOfCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfCards

`func (o *LoyaltyCardBatch) HasNumberOfCards() bool`

HasNumberOfCards returns a boolean if a field has been set.

### SetNumberOfCards

`func (o *LoyaltyCardBatch) SetNumberOfCards(v int32)`

SetNumberOfCards gets a reference to the given int32 and assigns it to the NumberOfCards field.

### GetBatchId

`func (o *LoyaltyCardBatch) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *LoyaltyCardBatch) GetBatchIdOk() (string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchId

`func (o *LoyaltyCardBatch) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchId

`func (o *LoyaltyCardBatch) SetBatchId(v string)`

SetBatchId gets a reference to the given string and assigns it to the BatchId field.

### GetStatus

`func (o *LoyaltyCardBatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoyaltyCardBatch) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *LoyaltyCardBatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *LoyaltyCardBatch) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


