# ConfirmRisksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskIds** | Pointer to **[]int64** | The IDs of the risks to confirm. | 
**Comment** | Pointer to **string** | Free-text description of how the risk was resolved. | 

## Methods

### NewConfirmRisksRequest

`func NewConfirmRisksRequest(riskIds []int64, comment string, ) *ConfirmRisksRequest`

NewConfirmRisksRequest instantiates a new ConfirmRisksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmRisksRequestWithDefaults

`func NewConfirmRisksRequestWithDefaults() *ConfirmRisksRequest`

NewConfirmRisksRequestWithDefaults instantiates a new ConfirmRisksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskIds

`func (o *ConfirmRisksRequest) GetRiskIds() []int64`

GetRiskIds returns the RiskIds field if non-nil, zero value otherwise.

### GetRiskIdsOk

`func (o *ConfirmRisksRequest) GetRiskIdsOk() (*[]int64, bool)`

GetRiskIdsOk returns a tuple with the RiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskIds

`func (o *ConfirmRisksRequest) SetRiskIds(v []int64)`

SetRiskIds sets RiskIds field to given value.


### GetComment

`func (o *ConfirmRisksRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfirmRisksRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfirmRisksRequest) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


