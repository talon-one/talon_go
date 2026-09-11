# RiskCriticalityUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RiskIds** | Pointer to **[]int64** | The IDs of the risks to reclassify. | 
**Criticality** | Pointer to **string** | The criticality to assign to risks. Only &#x60;not_critical&#x60; is accepted: critical risks can be reclassified as non-critical, but not the other way around.  | 

## Methods

### NewRiskCriticalityUpdate

`func NewRiskCriticalityUpdate(riskIds []int64, criticality string, ) *RiskCriticalityUpdate`

NewRiskCriticalityUpdate instantiates a new RiskCriticalityUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskCriticalityUpdateWithDefaults

`func NewRiskCriticalityUpdateWithDefaults() *RiskCriticalityUpdate`

NewRiskCriticalityUpdateWithDefaults instantiates a new RiskCriticalityUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRiskIds

`func (o *RiskCriticalityUpdate) GetRiskIds() []int64`

GetRiskIds returns the RiskIds field if non-nil, zero value otherwise.

### GetRiskIdsOk

`func (o *RiskCriticalityUpdate) GetRiskIdsOk() (*[]int64, bool)`

GetRiskIdsOk returns a tuple with the RiskIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskIds

`func (o *RiskCriticalityUpdate) SetRiskIds(v []int64)`

SetRiskIds sets RiskIds field to given value.


### GetCriticality

`func (o *RiskCriticalityUpdate) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *RiskCriticalityUpdate) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *RiskCriticalityUpdate) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


