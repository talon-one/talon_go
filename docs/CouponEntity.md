# CouponEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The internal ID of the coupon. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time the coupon was created. | 

## Methods

### GetId

`func (o *CouponEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponEntity) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CouponEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CouponEntity) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CouponEntity) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CouponEntity) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CouponEntity) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CouponEntity) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


