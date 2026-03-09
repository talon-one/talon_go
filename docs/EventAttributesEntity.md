# EventAttributesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | A string representing the event name. Must not be a reserved event name. You create this value when you [create an attribute](https://docs.talon.one/docs/dev/concepts/entities/events#creating-a-custom-event) of type &#x60;event&#x60; in the Campaign Manager.  | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary additional JSON properties associated with the event. They must be created in the Campaign Manager before setting them with this property. See [creating custom attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes#creating-a-custom-attribute). | [optional] 

## Methods

### NewEventAttributesEntity

`func NewEventAttributesEntity(type_ string, ) *EventAttributesEntity`

NewEventAttributesEntity instantiates a new EventAttributesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventAttributesEntityWithDefaults

`func NewEventAttributesEntityWithDefaults() *EventAttributesEntity`

NewEventAttributesEntityWithDefaults instantiates a new EventAttributesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EventAttributesEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventAttributesEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventAttributesEntity) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *EventAttributesEntity) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *EventAttributesEntity) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *EventAttributesEntity) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *EventAttributesEntity) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


