# NewEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** | ID of the customer profile set by your integration layer.  **Note:** If the customer does not yet have a known &#x60;profileId&#x60;, we recommend you use a guest &#x60;profileId&#x60;.  | [optional] 
**StoreIntegrationId** | Pointer to **string** | The integration ID of the store. You choose this ID when you create a store. | [optional] 
**Type** | Pointer to **string** | A string representing the event. Must not be a reserved event name. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary additional JSON data associated with the event. | 
**SessionId** | Pointer to **string** | The ID of the session that this event occurred in. | 

## Methods

### NewNewEvent

`func NewNewEvent(type_ string, attributes map[string]interface{}, sessionId string, ) *NewEvent`

NewNewEvent instantiates a new NewEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewEventWithDefaults

`func NewNewEventWithDefaults() *NewEvent`

NewNewEventWithDefaults instantiates a new NewEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *NewEvent) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *NewEvent) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *NewEvent) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *NewEvent) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetStoreIntegrationId

`func (o *NewEvent) GetStoreIntegrationId() string`

GetStoreIntegrationId returns the StoreIntegrationId field if non-nil, zero value otherwise.

### GetStoreIntegrationIdOk

`func (o *NewEvent) GetStoreIntegrationIdOk() (*string, bool)`

GetStoreIntegrationIdOk returns a tuple with the StoreIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIntegrationId

`func (o *NewEvent) SetStoreIntegrationId(v string)`

SetStoreIntegrationId sets StoreIntegrationId field to given value.

### HasStoreIntegrationId

`func (o *NewEvent) HasStoreIntegrationId() bool`

HasStoreIntegrationId returns a boolean if a field has been set.

### GetType

`func (o *NewEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewEvent) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *NewEvent) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewEvent) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewEvent) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetSessionId

`func (o *NewEvent) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *NewEvent) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *NewEvent) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


