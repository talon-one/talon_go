# PriceTypeReferenceDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferencingType** | Pointer to **string** | The entity type that references the price type. For example, a campaign or an Application cart item filter. | 
**ReferencingId** | Pointer to **int64** | The ID of the entity that references the price type. | 
**ReferencingName** | Pointer to **string** | The name of the entity that references the price type. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that contains the entity that references the price type. | [optional] 

## Methods

### NewPriceTypeReferenceDetail

`func NewPriceTypeReferenceDetail(referencingType string, referencingId int64, referencingName string, ) *PriceTypeReferenceDetail`

NewPriceTypeReferenceDetail instantiates a new PriceTypeReferenceDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceTypeReferenceDetailWithDefaults

`func NewPriceTypeReferenceDetailWithDefaults() *PriceTypeReferenceDetail`

NewPriceTypeReferenceDetailWithDefaults instantiates a new PriceTypeReferenceDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferencingType

`func (o *PriceTypeReferenceDetail) GetReferencingType() string`

GetReferencingType returns the ReferencingType field if non-nil, zero value otherwise.

### GetReferencingTypeOk

`func (o *PriceTypeReferenceDetail) GetReferencingTypeOk() (*string, bool)`

GetReferencingTypeOk returns a tuple with the ReferencingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingType

`func (o *PriceTypeReferenceDetail) SetReferencingType(v string)`

SetReferencingType sets ReferencingType field to given value.


### GetReferencingId

`func (o *PriceTypeReferenceDetail) GetReferencingId() int64`

GetReferencingId returns the ReferencingId field if non-nil, zero value otherwise.

### GetReferencingIdOk

`func (o *PriceTypeReferenceDetail) GetReferencingIdOk() (*int64, bool)`

GetReferencingIdOk returns a tuple with the ReferencingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingId

`func (o *PriceTypeReferenceDetail) SetReferencingId(v int64)`

SetReferencingId sets ReferencingId field to given value.


### GetReferencingName

`func (o *PriceTypeReferenceDetail) GetReferencingName() string`

GetReferencingName returns the ReferencingName field if non-nil, zero value otherwise.

### GetReferencingNameOk

`func (o *PriceTypeReferenceDetail) GetReferencingNameOk() (*string, bool)`

GetReferencingNameOk returns a tuple with the ReferencingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencingName

`func (o *PriceTypeReferenceDetail) SetReferencingName(v string)`

SetReferencingName sets ReferencingName field to given value.


### GetApplicationId

`func (o *PriceTypeReferenceDetail) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *PriceTypeReferenceDetail) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *PriceTypeReferenceDetail) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *PriceTypeReferenceDetail) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


