# StrikethroughTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the event that triggered the strikethrough labeling. | 
**Type** | Pointer to **string** | The type of event that triggered the strikethrough labeling. | 
**TriggeredAt** | Pointer to [**time.Time**](time.Time.md) | The creation time of the event that triggered the strikethrough labeling. | 
**TotalAffectedItems** | Pointer to **int32** | The total number of items affected by the event that triggered the strikethrough labeling. | 
**Payload** | Pointer to [**map[string]interface{}**](.md) | The arbitrary properties associated with this trigger type. | 

## Methods

### GetId

`func (o *StrikethroughTrigger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StrikethroughTrigger) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *StrikethroughTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *StrikethroughTrigger) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetType

`func (o *StrikethroughTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StrikethroughTrigger) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *StrikethroughTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *StrikethroughTrigger) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetTriggeredAt

`func (o *StrikethroughTrigger) GetTriggeredAt() time.Time`

GetTriggeredAt returns the TriggeredAt field if non-nil, zero value otherwise.

### GetTriggeredAtOk

`func (o *StrikethroughTrigger) GetTriggeredAtOk() (time.Time, bool)`

GetTriggeredAtOk returns a tuple with the TriggeredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTriggeredAt

`func (o *StrikethroughTrigger) HasTriggeredAt() bool`

HasTriggeredAt returns a boolean if a field has been set.

### SetTriggeredAt

`func (o *StrikethroughTrigger) SetTriggeredAt(v time.Time)`

SetTriggeredAt gets a reference to the given time.Time and assigns it to the TriggeredAt field.

### GetTotalAffectedItems

`func (o *StrikethroughTrigger) GetTotalAffectedItems() int32`

GetTotalAffectedItems returns the TotalAffectedItems field if non-nil, zero value otherwise.

### GetTotalAffectedItemsOk

`func (o *StrikethroughTrigger) GetTotalAffectedItemsOk() (int32, bool)`

GetTotalAffectedItemsOk returns a tuple with the TotalAffectedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTotalAffectedItems

`func (o *StrikethroughTrigger) HasTotalAffectedItems() bool`

HasTotalAffectedItems returns a boolean if a field has been set.

### SetTotalAffectedItems

`func (o *StrikethroughTrigger) SetTotalAffectedItems(v int32)`

SetTotalAffectedItems gets a reference to the given int32 and assigns it to the TotalAffectedItems field.

### GetPayload

`func (o *StrikethroughTrigger) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *StrikethroughTrigger) GetPayloadOk() (map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPayload

`func (o *StrikethroughTrigger) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayload

`func (o *StrikethroughTrigger) SetPayload(v map[string]interface{})`

SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


