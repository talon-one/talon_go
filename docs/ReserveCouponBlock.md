# ReserveCouponBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 

## Methods

### NewReserveCouponBlock

`func NewReserveCouponBlock(type_ string, ) *ReserveCouponBlock`

NewReserveCouponBlock instantiates a new ReserveCouponBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReserveCouponBlockWithDefaults

`func NewReserveCouponBlockWithDefaults() *ReserveCouponBlock`

NewReserveCouponBlockWithDefaults instantiates a new ReserveCouponBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReserveCouponBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReserveCouponBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReserveCouponBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReserveCouponBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReserveCouponBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReserveCouponBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReserveCouponBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *ReserveCouponBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ReserveCouponBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ReserveCouponBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ReserveCouponBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


