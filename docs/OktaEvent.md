# OktaEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type defining an action. | 
**Target** | Pointer to [**[]OktaEventTarget**](OktaEventTarget.md) |  | 

## Methods

### GetEventType

`func (o *OktaEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OktaEvent) GetEventTypeOk() (string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEventType

`func (o *OktaEvent) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### SetEventType

`func (o *OktaEvent) SetEventType(v string)`

SetEventType gets a reference to the given string and assigns it to the EventType field.

### GetTarget

`func (o *OktaEvent) GetTarget() []OktaEventTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *OktaEvent) GetTargetOk() ([]OktaEventTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTarget

`func (o *OktaEvent) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### SetTarget

`func (o *OktaEvent) SetTarget(v []OktaEventTarget)`

SetTarget gets a reference to the given []OktaEventTarget and assigns it to the Target field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


