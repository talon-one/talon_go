# DiscardRisksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskIds** | Pointer to **[]int64** | The IDs of the risks to discard. | 
**Reason** | Pointer to **string** | The reason the risks are being discarded. | 
**Comment** | Pointer to **string** | Free-text description of why the risks are being discarded. Required when &#x60;reason&#x60; is &#x60;other&#x60;, optional for &#x60;expected_behavior&#x60;.  | [optional] 

## Methods

### NewDiscardRisksRequest

`func NewDiscardRisksRequest(riskIds []int64, reason string, ) *DiscardRisksRequest`

NewDiscardRisksRequest instantiates a new DiscardRisksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscardRisksRequestWithDefaults

`func NewDiscardRisksRequestWithDefaults() *DiscardRisksRequest`

NewDiscardRisksRequestWithDefaults instantiates a new DiscardRisksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskIds

`func (o *DiscardRisksRequest) GetRiskIds() []int64`

GetRiskIds returns the RiskIds field if non-nil, zero value otherwise.

### GetRiskIdsOk

`func (o *DiscardRisksRequest) GetRiskIdsOk() (*[]int64, bool)`

GetRiskIdsOk returns a tuple with the RiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskIds

`func (o *DiscardRisksRequest) SetRiskIds(v []int64)`

SetRiskIds sets RiskIds field to given value.


### GetReason

`func (o *DiscardRisksRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *DiscardRisksRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *DiscardRisksRequest) SetReason(v string)`

SetReason sets Reason field to given value.


### GetComment

`func (o *DiscardRisksRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DiscardRisksRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DiscardRisksRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DiscardRisksRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


