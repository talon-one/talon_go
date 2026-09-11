# NewEventV3Entity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The unique ID of the event. Only one event with this ID can be registered.  | 

## Methods

### NewNewEventV3Entity

`func NewNewEventV3Entity(integrationId string, ) *NewEventV3Entity`

NewNewEventV3Entity instantiates a new NewEventV3Entity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewEventV3EntityWithDefaults

`func NewNewEventV3EntityWithDefaults() *NewEventV3Entity`

NewNewEventV3EntityWithDefaults instantiates a new NewEventV3Entity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *NewEventV3Entity) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *NewEventV3Entity) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *NewEventV3Entity) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


