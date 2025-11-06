# ApplicationCustomerEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **int64** | The globally unique Talon.One ID of the customer that created this entity. | [optional] 

## Methods

### NewApplicationCustomerEntity

`func NewApplicationCustomerEntity() *ApplicationCustomerEntity`

NewApplicationCustomerEntity instantiates a new ApplicationCustomerEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCustomerEntityWithDefaults

`func NewApplicationCustomerEntityWithDefaults() *ApplicationCustomerEntity`

NewApplicationCustomerEntityWithDefaults instantiates a new ApplicationCustomerEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *ApplicationCustomerEntity) GetProfileId() int64`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *ApplicationCustomerEntity) GetProfileIdOk() (*int64, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *ApplicationCustomerEntity) SetProfileId(v int64)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *ApplicationCustomerEntity) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


