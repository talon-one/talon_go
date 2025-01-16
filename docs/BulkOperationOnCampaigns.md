# BulkOperationOnCampaigns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | The operation to perform on the specified campaign IDs.  | 
**CampaignIds** | Pointer to **[]int32** | The list of campaign IDs on which the operation will be performed. | 
**ActivateAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the revisions are finalized after the &#x60;activate_revision&#x60; operation. The current time is used when left blank.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 

## Methods

### GetOperation

`func (o *BulkOperationOnCampaigns) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkOperationOnCampaigns) GetOperationOk() (string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOperation

`func (o *BulkOperationOnCampaigns) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### SetOperation

`func (o *BulkOperationOnCampaigns) SetOperation(v string)`

SetOperation gets a reference to the given string and assigns it to the Operation field.

### GetCampaignIds

`func (o *BulkOperationOnCampaigns) GetCampaignIds() []int32`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *BulkOperationOnCampaigns) GetCampaignIdsOk() ([]int32, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignIds

`func (o *BulkOperationOnCampaigns) HasCampaignIds() bool`

HasCampaignIds returns a boolean if a field has been set.

### SetCampaignIds

`func (o *BulkOperationOnCampaigns) SetCampaignIds(v []int32)`

SetCampaignIds gets a reference to the given []int32 and assigns it to the CampaignIds field.

### GetActivateAt

`func (o *BulkOperationOnCampaigns) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *BulkOperationOnCampaigns) GetActivateAtOk() (time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasActivateAt

`func (o *BulkOperationOnCampaigns) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.

### SetActivateAt

`func (o *BulkOperationOnCampaigns) SetActivateAt(v time.Time)`

SetActivateAt gets a reference to the given time.Time and assigns it to the ActivateAt field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


