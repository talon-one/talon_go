# ReserveCouponEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CouponValue** | Pointer to **string** | The value of the coupon currently on scope. | 
**ProfileIntegrationId** | Pointer to **string** | The ID of this customer profile in the third-party integration. | 
**IsNewReservation** | Pointer to **bool** | Indicates whether this is a new coupon reservation or not. | 

## Methods

### NewReserveCouponEffectProps

`func NewReserveCouponEffectProps(couponValue string, profileIntegrationId string, isNewReservation bool, ) *ReserveCouponEffectProps`

NewReserveCouponEffectProps instantiates a new ReserveCouponEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReserveCouponEffectPropsWithDefaults

`func NewReserveCouponEffectPropsWithDefaults() *ReserveCouponEffectProps`

NewReserveCouponEffectPropsWithDefaults instantiates a new ReserveCouponEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCouponValue

`func (o *ReserveCouponEffectProps) GetCouponValue() string`

GetCouponValue returns the CouponValue field if non-nil, zero value otherwise.

### GetCouponValueOk

`func (o *ReserveCouponEffectProps) GetCouponValueOk() (*string, bool)`

GetCouponValueOk returns a tuple with the CouponValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponValue

`func (o *ReserveCouponEffectProps) SetCouponValue(v string)`

SetCouponValue sets CouponValue field to given value.


### GetProfileIntegrationId

`func (o *ReserveCouponEffectProps) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *ReserveCouponEffectProps) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *ReserveCouponEffectProps) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.


### GetIsNewReservation

`func (o *ReserveCouponEffectProps) GetIsNewReservation() bool`

GetIsNewReservation returns the IsNewReservation field if non-nil, zero value otherwise.

### GetIsNewReservationOk

`func (o *ReserveCouponEffectProps) GetIsNewReservationOk() (*bool, bool)`

GetIsNewReservationOk returns a tuple with the IsNewReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNewReservation

`func (o *ReserveCouponEffectProps) SetIsNewReservation(v bool)`

SetIsNewReservation sets IsNewReservation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


