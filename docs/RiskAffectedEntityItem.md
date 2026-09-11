# RiskAffectedEntityItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The integration ID of the affected entity. | 
**ActivityValue** | Pointer to **float64** | The observed value of the monitored activity metric for this entity. | 
**Threshold** | Pointer to **float64** | The anomaly threshold computed for the entity&#39;s Application group. | 
**SeverityRatio** | Pointer to **float64** | The ratio of the observed value to the threshold. | 
**Criticality** | Pointer to **string** | The critical classification bucket of this entity. | 

## Methods

### NewRiskAffectedEntityItem

`func NewRiskAffectedEntityItem(entityId string, activityValue float64, threshold float64, severityRatio float64, criticality string, ) *RiskAffectedEntityItem`

NewRiskAffectedEntityItem instantiates a new RiskAffectedEntityItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAffectedEntityItemWithDefaults

`func NewRiskAffectedEntityItemWithDefaults() *RiskAffectedEntityItem`

NewRiskAffectedEntityItemWithDefaults instantiates a new RiskAffectedEntityItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *RiskAffectedEntityItem) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *RiskAffectedEntityItem) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *RiskAffectedEntityItem) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.


### GetActivityValue

`func (o *RiskAffectedEntityItem) GetActivityValue() float64`

GetActivityValue returns the ActivityValue field if non-nil, zero value otherwise.

### GetActivityValueOk

`func (o *RiskAffectedEntityItem) GetActivityValueOk() (*float64, bool)`

GetActivityValueOk returns a tuple with the ActivityValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityValue

`func (o *RiskAffectedEntityItem) SetActivityValue(v float64)`

SetActivityValue sets ActivityValue field to given value.


### GetThreshold

`func (o *RiskAffectedEntityItem) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *RiskAffectedEntityItem) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *RiskAffectedEntityItem) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetSeverityRatio

`func (o *RiskAffectedEntityItem) GetSeverityRatio() float64`

GetSeverityRatio returns the SeverityRatio field if non-nil, zero value otherwise.

### GetSeverityRatioOk

`func (o *RiskAffectedEntityItem) GetSeverityRatioOk() (*float64, bool)`

GetSeverityRatioOk returns a tuple with the SeverityRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityRatio

`func (o *RiskAffectedEntityItem) SetSeverityRatio(v float64)`

SetSeverityRatio sets SeverityRatio field to given value.


### GetCriticality

`func (o *RiskAffectedEntityItem) GetCriticality() string`

GetCriticality returns the Criticality field if non-nil, zero value otherwise.

### GetCriticalityOk

`func (o *RiskAffectedEntityItem) GetCriticalityOk() (*string, bool)`

GetCriticalityOk returns a tuple with the Criticality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticality

`func (o *RiskAffectedEntityItem) SetCriticality(v string)`

SetCriticality sets Criticality field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


