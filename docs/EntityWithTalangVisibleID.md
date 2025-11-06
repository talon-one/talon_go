# EntityWithTalangVisibleID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 

## Methods

### NewEntityWithTalangVisibleID

`func NewEntityWithTalangVisibleID(id int64, created time.Time, ) *EntityWithTalangVisibleID`

NewEntityWithTalangVisibleID instantiates a new EntityWithTalangVisibleID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityWithTalangVisibleIDWithDefaults

`func NewEntityWithTalangVisibleIDWithDefaults() *EntityWithTalangVisibleID`

NewEntityWithTalangVisibleIDWithDefaults instantiates a new EntityWithTalangVisibleID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityWithTalangVisibleID) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityWithTalangVisibleID) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityWithTalangVisibleID) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *EntityWithTalangVisibleID) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntityWithTalangVisibleID) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntityWithTalangVisibleID) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


