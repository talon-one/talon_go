# BulkOperationOnCampaigns

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | The operation to perform on the specified campaign IDs.  | 
**CampaignIds** | Pointer to **[]int64** | The list of campaign IDs on which the operation will be performed. | 
**ActivateAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the revisions are finalized after the &#x60;activate_revision&#x60; operation. The current time is used when left blank.  **Note:** It must be an RFC3339 timestamp string.  | [optional] 

## Methods

### NewBulkOperationOnCampaigns

`func NewBulkOperationOnCampaigns(operation string, campaignIds []int64, ) *BulkOperationOnCampaigns`

NewBulkOperationOnCampaigns instantiates a new BulkOperationOnCampaigns object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationOnCampaignsWithDefaults

`func NewBulkOperationOnCampaignsWithDefaults() *BulkOperationOnCampaigns`

NewBulkOperationOnCampaignsWithDefaults instantiates a new BulkOperationOnCampaigns object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *BulkOperationOnCampaigns) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *BulkOperationOnCampaigns) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *BulkOperationOnCampaigns) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetCampaignIds

`func (o *BulkOperationOnCampaigns) GetCampaignIds() []int64`

GetCampaignIds returns the CampaignIds field if non-nil, zero value otherwise.

### GetCampaignIdsOk

`func (o *BulkOperationOnCampaigns) GetCampaignIdsOk() (*[]int64, bool)`

GetCampaignIdsOk returns a tuple with the CampaignIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIds

`func (o *BulkOperationOnCampaigns) SetCampaignIds(v []int64)`

SetCampaignIds sets CampaignIds field to given value.


### GetActivateAt

`func (o *BulkOperationOnCampaigns) GetActivateAt() time.Time`

GetActivateAt returns the ActivateAt field if non-nil, zero value otherwise.

### GetActivateAtOk

`func (o *BulkOperationOnCampaigns) GetActivateAtOk() (*time.Time, bool)`

GetActivateAtOk returns a tuple with the ActivateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateAt

`func (o *BulkOperationOnCampaigns) SetActivateAt(v time.Time)`

SetActivateAt sets ActivateAt field to given value.

### HasActivateAt

`func (o *BulkOperationOnCampaigns) HasActivateAt() bool`

HasActivateAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


