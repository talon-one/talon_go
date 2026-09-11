# ReviewRisksRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskIds** | Pointer to **[]int64** | The IDs of the risks to move to &#x60;In review&#x60; status. | 

## Methods

### NewReviewRisksRequest

`func NewReviewRisksRequest(riskIds []int64, ) *ReviewRisksRequest`

NewReviewRisksRequest instantiates a new ReviewRisksRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewRisksRequestWithDefaults

`func NewReviewRisksRequestWithDefaults() *ReviewRisksRequest`

NewReviewRisksRequestWithDefaults instantiates a new ReviewRisksRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskIds

`func (o *ReviewRisksRequest) GetRiskIds() []int64`

GetRiskIds returns the RiskIds field if non-nil, zero value otherwise.

### GetRiskIdsOk

`func (o *ReviewRisksRequest) GetRiskIdsOk() (*[]int64, bool)`

GetRiskIdsOk returns a tuple with the RiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskIds

`func (o *ReviewRisksRequest) SetRiskIds(v []int64)`

SetRiskIds sets RiskIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


