# PriceTypeReferences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceTypeId** | Pointer to **int64** | The ID of the price type. | 
**References** | Pointer to [**[]PriceTypeReferenceDetail**](PriceTypeReferenceDetail.md) | A list of entities that reference the price type, including details about the entities. | [optional] 

## Methods

### NewPriceTypeReferences

`func NewPriceTypeReferences(priceTypeId int64, ) *PriceTypeReferences`

NewPriceTypeReferences instantiates a new PriceTypeReferences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTypeReferencesWithDefaults

`func NewPriceTypeReferencesWithDefaults() *PriceTypeReferences`

NewPriceTypeReferencesWithDefaults instantiates a new PriceTypeReferences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceTypeId

`func (o *PriceTypeReferences) GetPriceTypeId() int64`

GetPriceTypeId returns the PriceTypeId field if non-nil, zero value otherwise.

### GetPriceTypeIdOk

`func (o *PriceTypeReferences) GetPriceTypeIdOk() (*int64, bool)`

GetPriceTypeIdOk returns a tuple with the PriceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceTypeId

`func (o *PriceTypeReferences) SetPriceTypeId(v int64)`

SetPriceTypeId sets PriceTypeId field to given value.


### GetReferences

`func (o *PriceTypeReferences) GetReferences() []PriceTypeReferenceDetail`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PriceTypeReferences) GetReferencesOk() (*[]PriceTypeReferenceDetail, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PriceTypeReferences) SetReferences(v []PriceTypeReferenceDetail)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PriceTypeReferences) HasReferences() bool`

HasReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


