# UpdateCouponBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | [optional] 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int64** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Optional property to set the value of custom coupon attributes. They are defined in the Campaign Manager, see [Managing attributes](https://docs.talon.one/docs/product/account/dev-tools/managing-attributes).  Coupon attributes can also be set to _mandatory_ in your Application [settings](https://docs.talon.one/docs/product/applications/using-attributes#making-attributes-mandatory). If your Application uses mandatory attributes, you must use this property to set their value.  | [optional] 
**BatchID** | Pointer to **string** | The ID of the batch the coupon(s) belong to. | [optional] 

## Methods

### NewUpdateCouponBatch

`func NewUpdateCouponBatch() *UpdateCouponBatch`

NewUpdateCouponBatch instantiates a new UpdateCouponBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCouponBatchWithDefaults

`func NewUpdateCouponBatchWithDefaults() *UpdateCouponBatch`

NewUpdateCouponBatchWithDefaults instantiates a new UpdateCouponBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageLimit

`func (o *UpdateCouponBatch) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *UpdateCouponBatch) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *UpdateCouponBatch) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *UpdateCouponBatch) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetDiscountLimit

`func (o *UpdateCouponBatch) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *UpdateCouponBatch) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *UpdateCouponBatch) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *UpdateCouponBatch) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *UpdateCouponBatch) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *UpdateCouponBatch) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *UpdateCouponBatch) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *UpdateCouponBatch) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateCouponBatch) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateCouponBatch) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateCouponBatch) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateCouponBatch) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *UpdateCouponBatch) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *UpdateCouponBatch) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *UpdateCouponBatch) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *UpdateCouponBatch) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateCouponBatch) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateCouponBatch) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateCouponBatch) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateCouponBatch) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetBatchID

`func (o *UpdateCouponBatch) GetBatchID() string`

GetBatchID returns the BatchID field if non-nil, zero value otherwise.

### GetBatchIDOk

`func (o *UpdateCouponBatch) GetBatchIDOk() (*string, bool)`

GetBatchIDOk returns a tuple with the BatchID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchID

`func (o *UpdateCouponBatch) SetBatchID(v string)`

SetBatchID sets BatchID field to given value.

### HasBatchID

`func (o *UpdateCouponBatch) HasBatchID() bool`

HasBatchID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


