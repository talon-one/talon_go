# MutableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 

## Methods

### NewMutableEntity

`func NewMutableEntity(modified time.Time, ) *MutableEntity`

NewMutableEntity instantiates a new MutableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutableEntityWithDefaults

`func NewMutableEntityWithDefaults() *MutableEntity`

NewMutableEntityWithDefaults instantiates a new MutableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModified

`func (o *MutableEntity) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *MutableEntity) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *MutableEntity) SetModified(v time.Time)`

SetModified sets Modified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


