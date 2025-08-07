# MultipleAudiencesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Name** | Pointer to **string** | The human-friendly display name for this audience. | 
**IntegrationId** | Pointer to **string** | The ID of this audience in the third-party integration. | 
**Status** | Pointer to **string** | Indicates whether the audience is new, updated or unmodified by the request.  | 

## Methods

### GetId

`func (o *MultipleAudiencesItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MultipleAudiencesItem) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *MultipleAudiencesItem) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *MultipleAudiencesItem) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *MultipleAudiencesItem) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MultipleAudiencesItem) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *MultipleAudiencesItem) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *MultipleAudiencesItem) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetName

`func (o *MultipleAudiencesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MultipleAudiencesItem) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *MultipleAudiencesItem) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *MultipleAudiencesItem) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetIntegrationId

`func (o *MultipleAudiencesItem) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *MultipleAudiencesItem) GetIntegrationIdOk() (string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIntegrationId

`func (o *MultipleAudiencesItem) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationId

`func (o *MultipleAudiencesItem) SetIntegrationId(v string)`

SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.

### GetStatus

`func (o *MultipleAudiencesItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MultipleAudiencesItem) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MultipleAudiencesItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MultipleAudiencesItem) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


