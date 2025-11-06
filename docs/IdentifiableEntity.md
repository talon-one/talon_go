# IdentifiableEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints. | 

## Methods

### NewIdentifiableEntity

`func NewIdentifiableEntity(id int64, ) *IdentifiableEntity`

NewIdentifiableEntity instantiates a new IdentifiableEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentifiableEntityWithDefaults

`func NewIdentifiableEntityWithDefaults() *IdentifiableEntity`

NewIdentifiableEntityWithDefaults instantiates a new IdentifiableEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdentifiableEntity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentifiableEntity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentifiableEntity) SetId(v int64)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


