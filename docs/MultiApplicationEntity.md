# MultiApplicationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIds** | Pointer to **[]int64** | The IDs of the Applications that are related to this entity. | 

## Methods

### NewMultiApplicationEntity

`func NewMultiApplicationEntity(applicationIds []int64, ) *MultiApplicationEntity`

NewMultiApplicationEntity instantiates a new MultiApplicationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiApplicationEntityWithDefaults

`func NewMultiApplicationEntityWithDefaults() *MultiApplicationEntity`

NewMultiApplicationEntityWithDefaults instantiates a new MultiApplicationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIds

`func (o *MultiApplicationEntity) GetApplicationIds() []int64`

GetApplicationIds returns the ApplicationIds field if non-nil, zero value otherwise.

### GetApplicationIdsOk

`func (o *MultiApplicationEntity) GetApplicationIdsOk() (*[]int64, bool)`

GetApplicationIdsOk returns a tuple with the ApplicationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIds

`func (o *MultiApplicationEntity) SetApplicationIds(v []int64)`

SetApplicationIds sets ApplicationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


