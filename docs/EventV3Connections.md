# EventV3Connections

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectedSessionId** | Pointer to **string** | The ID of the session to reference. The session must be in &#x60;closed&#x60; state. Otherwise, the API call will fail. | [optional] 

## Methods

### NewEventV3Connections

`func NewEventV3Connections() *EventV3Connections`

NewEventV3Connections instantiates a new EventV3Connections object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventV3ConnectionsWithDefaults

`func NewEventV3ConnectionsWithDefaults() *EventV3Connections`

NewEventV3ConnectionsWithDefaults instantiates a new EventV3Connections object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectedSessionId

`func (o *EventV3Connections) GetConnectedSessionId() string`

GetConnectedSessionId returns the ConnectedSessionId field if non-nil, zero value otherwise.

### GetConnectedSessionIdOk

`func (o *EventV3Connections) GetConnectedSessionIdOk() (*string, bool)`

GetConnectedSessionIdOk returns a tuple with the ConnectedSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSessionId

`func (o *EventV3Connections) SetConnectedSessionId(v string)`

SetConnectedSessionId sets ConnectedSessionId field to given value.

### HasConnectedSessionId

`func (o *EventV3Connections) HasConnectedSessionId() bool`

HasConnectedSessionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


