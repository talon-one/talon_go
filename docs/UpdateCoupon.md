# UpdateCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | [optional] 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int64** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.  | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The integration ID for this coupon&#39;s beneficiary&#39;s profile. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]
**ImplicitlyReserved** | Pointer to **bool** | An indication of whether the coupon is implicitly reserved for all customers. | [optional] 

## Methods

### NewUpdateCoupon

`func NewUpdateCoupon() *UpdateCoupon`

NewUpdateCoupon instantiates a new UpdateCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCouponWithDefaults

`func NewUpdateCouponWithDefaults() *UpdateCoupon`

NewUpdateCouponWithDefaults instantiates a new UpdateCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsageLimit

`func (o *UpdateCoupon) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *UpdateCoupon) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *UpdateCoupon) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *UpdateCoupon) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetDiscountLimit

`func (o *UpdateCoupon) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *UpdateCoupon) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *UpdateCoupon) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *UpdateCoupon) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *UpdateCoupon) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *UpdateCoupon) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *UpdateCoupon) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *UpdateCoupon) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateCoupon) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateCoupon) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateCoupon) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateCoupon) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *UpdateCoupon) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *UpdateCoupon) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *UpdateCoupon) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *UpdateCoupon) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetLimits

`func (o *UpdateCoupon) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateCoupon) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *UpdateCoupon) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *UpdateCoupon) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetRecipientIntegrationId

`func (o *UpdateCoupon) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *UpdateCoupon) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *UpdateCoupon) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.

### HasRecipientIntegrationId

`func (o *UpdateCoupon) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateCoupon) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateCoupon) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateCoupon) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateCoupon) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetIsReservationMandatory

`func (o *UpdateCoupon) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *UpdateCoupon) GetIsReservationMandatoryOk() (*bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReservationMandatory

`func (o *UpdateCoupon) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory sets IsReservationMandatory field to given value.

### HasIsReservationMandatory

`func (o *UpdateCoupon) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### GetImplicitlyReserved

`func (o *UpdateCoupon) GetImplicitlyReserved() bool`

GetImplicitlyReserved returns the ImplicitlyReserved field if non-nil, zero value otherwise.

### GetImplicitlyReservedOk

`func (o *UpdateCoupon) GetImplicitlyReservedOk() (*bool, bool)`

GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitlyReserved

`func (o *UpdateCoupon) SetImplicitlyReserved(v bool)`

SetImplicitlyReserved sets ImplicitlyReserved field to given value.

### HasImplicitlyReserved

`func (o *UpdateCoupon) HasImplicitlyReserved() bool`

HasImplicitlyReserved returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


