# StrikethroughTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the event that triggered the strikethrough labeling. | 
**Type** | Pointer to **string** | The type of event that triggered the strikethrough labeling. | 
**TriggeredAt** | Pointer to [**time.Time**](time.Time.md) | The creation time of the event that triggered the strikethrough labeling. | 
**TotalAffectedItems** | Pointer to **int32** | The total number of items affected by the event that triggered the strikethrough labeling. | 
**Payload** | Pointer to [**map[string]interface{}**](.md) | The arbitrary properties associated with this trigger type. | 

## Methods

### NewStrikethroughTrigger

`func NewStrikethroughTrigger(id int64, type_ string, triggeredAt time.Time, totalAffectedItems int32, payload map[string]interface{}, ) *StrikethroughTrigger`

NewStrikethroughTrigger instantiates a new StrikethroughTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrikethroughTriggerWithDefaults

`func NewStrikethroughTriggerWithDefaults() *StrikethroughTrigger`

NewStrikethroughTriggerWithDefaults instantiates a new StrikethroughTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StrikethroughTrigger) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrikethroughTrigger) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StrikethroughTrigger) SetId(v int64)`

SetId sets Id field to given value.


### GetType

`func (o *StrikethroughTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StrikethroughTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StrikethroughTrigger) SetType(v string)`

SetType sets Type field to given value.


### GetTriggeredAt

`func (o *StrikethroughTrigger) GetTriggeredAt() time.Time`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *StrikethroughTrigger) GetTriggeredAtOk() (*time.Time, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredAt

`func (o *StrikethroughTrigger) SetTriggeredAt(v time.Time)`

SetTriggeredAt sets TriggeredAt field to given value.


### GetTotalAffectedItems

`func (o *StrikethroughTrigger) GetTotalAffectedItems() int32`

GetTotalAffectedItems returns the TotalAffectedItems field if non-nil, zero value otherwise.

### GetTotalAffectedItemsOk

`func (o *StrikethroughTrigger) GetTotalAffectedItemsOk() (*int32, bool)`

GetTotalAffectedItemsOk returns a tuple with the TotalAffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAffectedItems

`func (o *StrikethroughTrigger) SetTotalAffectedItems(v int32)`

SetTotalAffectedItems sets TotalAffectedItems field to given value.


### GetPayload

`func (o *StrikethroughTrigger) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *StrikethroughTrigger) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *StrikethroughTrigger) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


