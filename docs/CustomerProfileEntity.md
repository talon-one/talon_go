# CustomerProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of the customer profile. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time the customer profile was created. | 

## Methods

### GetId

`func (o *CustomerProfileEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerProfileEntity) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CustomerProfileEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CustomerProfileEntity) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CustomerProfileEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerProfileEntity) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CustomerProfileEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CustomerProfileEntity) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


