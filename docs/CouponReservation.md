# CouponReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponId** | Pointer to **int64** | The internal ID of the coupon that was reserved. | 
**RecipientIntegrationId** | Pointer to **string** | The integration identifier of the customer for whom this coupon was reserved. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the coupon reservation was created. | [optional] 

## Methods

### NewCouponReservation

`func NewCouponReservation(couponId int64, recipientIntegrationId string, ) *CouponReservation`

NewCouponReservation instantiates a new CouponReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponReservationWithDefaults

`func NewCouponReservationWithDefaults() *CouponReservation`

NewCouponReservationWithDefaults instantiates a new CouponReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponId

`func (o *CouponReservation) GetCouponId() int64`

GetCouponId returns the CouponId field if non-nil, zero value otherwise.

### GetCouponIdOk

`func (o *CouponReservation) GetCouponIdOk() (*int64, bool)`

GetCouponIdOk returns a tuple with the CouponId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponId

`func (o *CouponReservation) SetCouponId(v int64)`

SetCouponId sets CouponId field to given value.


### GetRecipientIntegrationId

`func (o *CouponReservation) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *CouponReservation) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *CouponReservation) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.


### GetCreatedAt

`func (o *CouponReservation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CouponReservation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CouponReservation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CouponReservation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


