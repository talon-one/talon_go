# OktaEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type defining an action. | 
**Target** | Pointer to [**[]OktaEventTarget**](OktaEventTarget.md) |  | 

## Methods

### NewOktaEvent

`func NewOktaEvent(eventType string, target []OktaEventTarget, ) *OktaEvent`

NewOktaEvent instantiates a new OktaEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOktaEventWithDefaults

`func NewOktaEventWithDefaults() *OktaEvent`

NewOktaEventWithDefaults instantiates a new OktaEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *OktaEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *OktaEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *OktaEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetTarget

`func (o *OktaEvent) GetTarget() []OktaEventTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *OktaEvent) GetTargetOk() (*[]OktaEventTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *OktaEvent) SetTarget(v []OktaEventTarget)`

SetTarget sets Target field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


