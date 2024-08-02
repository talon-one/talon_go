# EventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Description** | Pointer to **string** | A description of what the event represents.  | [optional] 
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Name** | Pointer to **string** | The integration name for this event type. This will be used in URLs and cannot be changed after an event type has been created. | 
**Title** | Pointer to **string** | The human-friendly name for this event type. | 

## Methods

### GetCreated

`func (o *EventType) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EventType) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *EventType) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *EventType) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetDescription

`func (o *EventType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventType) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *EventType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *EventType) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetId

`func (o *EventType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EventType) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *EventType) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *EventType) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetName

`func (o *EventType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EventType) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *EventType) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *EventType) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *EventType) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventType) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *EventType) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *EventType) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


