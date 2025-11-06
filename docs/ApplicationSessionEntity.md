# ApplicationSessionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionId** | Pointer to **int64** | The globally unique Talon.One ID of the session where this entity was created. | 

## Methods

### NewApplicationSessionEntity

`func NewApplicationSessionEntity(sessionId int64, ) *ApplicationSessionEntity`

NewApplicationSessionEntity instantiates a new ApplicationSessionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSessionEntityWithDefaults

`func NewApplicationSessionEntityWithDefaults() *ApplicationSessionEntity`

NewApplicationSessionEntityWithDefaults instantiates a new ApplicationSessionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionId

`func (o *ApplicationSessionEntity) GetSessionId() int64`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *ApplicationSessionEntity) GetSessionIdOk() (*int64, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *ApplicationSessionEntity) SetSessionId(v int64)`

SetSessionId sets SessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


