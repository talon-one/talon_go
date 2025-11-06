# CustomerProfileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of the customer profile. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time the customer profile was created. | 

## Methods

### NewCustomerProfileEntity

`func NewCustomerProfileEntity(id int64, created time.Time, ) *CustomerProfileEntity`

NewCustomerProfileEntity instantiates a new CustomerProfileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerProfileEntityWithDefaults

`func NewCustomerProfileEntityWithDefaults() *CustomerProfileEntity`

NewCustomerProfileEntityWithDefaults instantiates a new CustomerProfileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerProfileEntity) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerProfileEntity) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerProfileEntity) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *CustomerProfileEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CustomerProfileEntity) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CustomerProfileEntity) SetCreated(v time.Time)`

SetCreated sets Created field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


