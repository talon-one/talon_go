# IntegrationHubEventStatusUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | Pointer to **int64** | The ID of the integration hub event. | 
**Status** | Pointer to **string** | The delivery outcome for the event. | 

## Methods

### NewIntegrationHubEventStatusUpdate

`func NewIntegrationHubEventStatusUpdate(eventId int64, status string, ) *IntegrationHubEventStatusUpdate`

NewIntegrationHubEventStatusUpdate instantiates a new IntegrationHubEventStatusUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationHubEventStatusUpdateWithDefaults

`func NewIntegrationHubEventStatusUpdateWithDefaults() *IntegrationHubEventStatusUpdate`

NewIntegrationHubEventStatusUpdateWithDefaults instantiates a new IntegrationHubEventStatusUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *IntegrationHubEventStatusUpdate) GetEventId() int64`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *IntegrationHubEventStatusUpdate) GetEventIdOk() (*int64, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *IntegrationHubEventStatusUpdate) SetEventId(v int64)`

SetEventId sets EventId field to given value.


### GetStatus

`func (o *IntegrationHubEventStatusUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationHubEventStatusUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationHubEventStatusUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


