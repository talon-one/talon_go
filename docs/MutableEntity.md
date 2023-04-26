# MutableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 

## Methods

### GetModified

`func (o *MutableEntity) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *MutableEntity) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *MutableEntity) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *MutableEntity) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


