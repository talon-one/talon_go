# UpdateLoyaltyCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Status of the loyalty card. Can be &#x60;active&#x60; or &#x60;inactive&#x60;.  | 
**BlockReason** | Pointer to **string** | Reason for transferring and blocking the loyalty card.  | [optional] 

## Methods

### GetStatus

`func (o *UpdateLoyaltyCard) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateLoyaltyCard) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *UpdateLoyaltyCard) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *UpdateLoyaltyCard) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetBlockReason

`func (o *UpdateLoyaltyCard) GetBlockReason() string`

GetBlockReason returns the BlockReason field if non-nil, zero value otherwise.

### GetBlockReasonOk

`func (o *UpdateLoyaltyCard) GetBlockReasonOk() (string, bool)`

GetBlockReasonOk returns a tuple with the BlockReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBlockReason

`func (o *UpdateLoyaltyCard) HasBlockReason() bool`

HasBlockReason returns a boolean if a field has been set.

### SetBlockReason

`func (o *UpdateLoyaltyCard) SetBlockReason(v string)`

SetBlockReason gets a reference to the given string and assigns it to the BlockReason field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


