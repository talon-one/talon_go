# LoyaltyCardBatchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfCardsGenerated** | Pointer to **int64** | Number of loyalty cards in the batch. | 
**BatchId** | Pointer to **string** | ID of the loyalty card batch. | 

## Methods

### NewLoyaltyCardBatchResponse

`func NewLoyaltyCardBatchResponse(numberOfCardsGenerated int64, batchId string, ) *LoyaltyCardBatchResponse`

NewLoyaltyCardBatchResponse instantiates a new LoyaltyCardBatchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyCardBatchResponseWithDefaults

`func NewLoyaltyCardBatchResponseWithDefaults() *LoyaltyCardBatchResponse`

NewLoyaltyCardBatchResponseWithDefaults instantiates a new LoyaltyCardBatchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfCardsGenerated

`func (o *LoyaltyCardBatchResponse) GetNumberOfCardsGenerated() int64`

GetNumberOfCardsGenerated returns the NumberOfCardsGenerated field if non-nil, zero value otherwise.

### GetNumberOfCardsGeneratedOk

`func (o *LoyaltyCardBatchResponse) GetNumberOfCardsGeneratedOk() (*int64, bool)`

GetNumberOfCardsGeneratedOk returns a tuple with the NumberOfCardsGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCardsGenerated

`func (o *LoyaltyCardBatchResponse) SetNumberOfCardsGenerated(v int64)`

SetNumberOfCardsGenerated sets NumberOfCardsGenerated field to given value.


### GetBatchId

`func (o *LoyaltyCardBatchResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *LoyaltyCardBatchResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *LoyaltyCardBatchResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


