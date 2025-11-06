# IntegrationCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of the coupon. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time the coupon was created. | 
**CampaignId** | Pointer to **int64** | The ID of the campaign that owns this entity. | 
**Value** | Pointer to **string** | The coupon code. | 
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int64** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.  | [optional] 
**UsageCounter** | Pointer to **int64** | The number of times the coupon has been successfully redeemed. | 
**DiscountCounter** | Pointer to **float32** | The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon. | [optional] 
**DiscountRemainder** | Pointer to **float32** | The remaining discount this coupon can give. | [optional] 
**ReservationCounter** | Pointer to **float32** | The number of times this coupon has been reserved. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Custom attributes associated with this coupon. | [optional] 
**ReferralId** | Pointer to **int64** | The integration ID of the referring customer (if any) for whom this coupon was created as an effect. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The Integration ID of the customer that is allowed to redeem this coupon. | [optional] 
**ImportId** | Pointer to **int64** | The ID of the Import which created this coupon. | [optional] 
**Reservation** | Pointer to **bool** | Defines the reservation type: - &#x60;true&#x60;: The coupon can be reserved for multiple customers. - &#x60;false&#x60;: The coupon can be reserved only for one customer. It is a personal code.  | [optional] [default to true]
**BatchId** | Pointer to **string** | The id of the batch the coupon belongs to. | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]
**ImplicitlyReserved** | Pointer to **bool** | An indication of whether the coupon is implicitly reserved for all customers. | [optional] 
**ProfileRedemptionCount** | Pointer to **int64** | The number of times the coupon was redeemed by the profile. | 

## Methods

### NewIntegrationCoupon

`func NewIntegrationCoupon(id int64, created time.Time, campaignId int64, value string, usageLimit int64, usageCounter int64, profileRedemptionCount int64, ) *IntegrationCoupon`

NewIntegrationCoupon instantiates a new IntegrationCoupon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationCouponWithDefaults

`func NewIntegrationCouponWithDefaults() *IntegrationCoupon`

NewIntegrationCouponWithDefaults instantiates a new IntegrationCoupon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationCoupon) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationCoupon) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationCoupon) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *IntegrationCoupon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *IntegrationCoupon) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *IntegrationCoupon) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCampaignId

`func (o *IntegrationCoupon) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *IntegrationCoupon) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *IntegrationCoupon) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetValue

`func (o *IntegrationCoupon) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IntegrationCoupon) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IntegrationCoupon) SetValue(v string)`

SetValue sets Value field to given value.


### GetUsageLimit

`func (o *IntegrationCoupon) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *IntegrationCoupon) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *IntegrationCoupon) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *IntegrationCoupon) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *IntegrationCoupon) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *IntegrationCoupon) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *IntegrationCoupon) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *IntegrationCoupon) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *IntegrationCoupon) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *IntegrationCoupon) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *IntegrationCoupon) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *IntegrationCoupon) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *IntegrationCoupon) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *IntegrationCoupon) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *IntegrationCoupon) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *IntegrationCoupon) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *IntegrationCoupon) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *IntegrationCoupon) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *IntegrationCoupon) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetLimits

`func (o *IntegrationCoupon) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *IntegrationCoupon) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *IntegrationCoupon) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *IntegrationCoupon) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetUsageCounter

`func (o *IntegrationCoupon) GetUsageCounter() int64`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *IntegrationCoupon) GetUsageCounterOk() (*int64, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *IntegrationCoupon) SetUsageCounter(v int64)`

SetUsageCounter sets UsageCounter field to given value.


### GetDiscountCounter

`func (o *IntegrationCoupon) GetDiscountCounter() float32`

GetDiscountCounter returns the DiscountCounter field if non-nil, zero value otherwise.

### GetDiscountCounterOk

`func (o *IntegrationCoupon) GetDiscountCounterOk() (*float32, bool)`

GetDiscountCounterOk returns a tuple with the DiscountCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountCounter

`func (o *IntegrationCoupon) SetDiscountCounter(v float32)`

SetDiscountCounter sets DiscountCounter field to given value.

### HasDiscountCounter

`func (o *IntegrationCoupon) HasDiscountCounter() bool`

HasDiscountCounter returns a boolean if a field has been set.

### GetDiscountRemainder

`func (o *IntegrationCoupon) GetDiscountRemainder() float32`

GetDiscountRemainder returns the DiscountRemainder field if non-nil, zero value otherwise.

### GetDiscountRemainderOk

`func (o *IntegrationCoupon) GetDiscountRemainderOk() (*float32, bool)`

GetDiscountRemainderOk returns a tuple with the DiscountRemainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountRemainder

`func (o *IntegrationCoupon) SetDiscountRemainder(v float32)`

SetDiscountRemainder sets DiscountRemainder field to given value.

### HasDiscountRemainder

`func (o *IntegrationCoupon) HasDiscountRemainder() bool`

HasDiscountRemainder returns a boolean if a field has been set.

### GetReservationCounter

`func (o *IntegrationCoupon) GetReservationCounter() float32`

GetReservationCounter returns the ReservationCounter field if non-nil, zero value otherwise.

### GetReservationCounterOk

`func (o *IntegrationCoupon) GetReservationCounterOk() (*float32, bool)`

GetReservationCounterOk returns a tuple with the ReservationCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationCounter

`func (o *IntegrationCoupon) SetReservationCounter(v float32)`

SetReservationCounter sets ReservationCounter field to given value.

### HasReservationCounter

`func (o *IntegrationCoupon) HasReservationCounter() bool`

HasReservationCounter returns a boolean if a field has been set.

### GetAttributes

`func (o *IntegrationCoupon) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *IntegrationCoupon) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *IntegrationCoupon) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *IntegrationCoupon) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetReferralId

`func (o *IntegrationCoupon) GetReferralId() int64`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *IntegrationCoupon) GetReferralIdOk() (*int64, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *IntegrationCoupon) SetReferralId(v int64)`

SetReferralId sets ReferralId field to given value.

### HasReferralId

`func (o *IntegrationCoupon) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### GetRecipientIntegrationId

`func (o *IntegrationCoupon) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *IntegrationCoupon) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *IntegrationCoupon) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.

### HasRecipientIntegrationId

`func (o *IntegrationCoupon) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### GetImportId

`func (o *IntegrationCoupon) GetImportId() int64`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *IntegrationCoupon) GetImportIdOk() (*int64, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *IntegrationCoupon) SetImportId(v int64)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *IntegrationCoupon) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetReservation

`func (o *IntegrationCoupon) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *IntegrationCoupon) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *IntegrationCoupon) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *IntegrationCoupon) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetBatchId

`func (o *IntegrationCoupon) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *IntegrationCoupon) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *IntegrationCoupon) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *IntegrationCoupon) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### GetIsReservationMandatory

`func (o *IntegrationCoupon) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *IntegrationCoupon) GetIsReservationMandatoryOk() (*bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReservationMandatory

`func (o *IntegrationCoupon) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory sets IsReservationMandatory field to given value.

### HasIsReservationMandatory

`func (o *IntegrationCoupon) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### GetImplicitlyReserved

`func (o *IntegrationCoupon) GetImplicitlyReserved() bool`

GetImplicitlyReserved returns the ImplicitlyReserved field if non-nil, zero value otherwise.

### GetImplicitlyReservedOk

`func (o *IntegrationCoupon) GetImplicitlyReservedOk() (*bool, bool)`

GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitlyReserved

`func (o *IntegrationCoupon) SetImplicitlyReserved(v bool)`

SetImplicitlyReserved sets ImplicitlyReserved field to given value.

### HasImplicitlyReserved

`func (o *IntegrationCoupon) HasImplicitlyReserved() bool`

HasImplicitlyReserved returns a boolean if a field has been set.

### GetProfileRedemptionCount

`func (o *IntegrationCoupon) GetProfileRedemptionCount() int64`

GetProfileRedemptionCount returns the ProfileRedemptionCount field if non-nil, zero value otherwise.

### GetProfileRedemptionCountOk

`func (o *IntegrationCoupon) GetProfileRedemptionCountOk() (*int64, bool)`

GetProfileRedemptionCountOk returns a tuple with the ProfileRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileRedemptionCount

`func (o *IntegrationCoupon) SetProfileRedemptionCount(v int64)`

SetProfileRedemptionCount sets ProfileRedemptionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


