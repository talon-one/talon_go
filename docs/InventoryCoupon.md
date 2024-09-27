# InventoryCoupon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | 
**Value** | Pointer to **string** | The coupon code. | 
**UsageLimit** | Pointer to **int32** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int32** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured.  | [optional] 
**UsageCounter** | Pointer to **int32** | The number of times the coupon has been successfully redeemed. | 
**DiscountCounter** | Pointer to **float32** | The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon. | [optional] 
**DiscountRemainder** | Pointer to **float32** | The remaining discount this coupon can give. | [optional] 
**ReservationCounter** | Pointer to **float32** | The number of times this coupon has been reserved. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Custom attributes associated with this coupon. | [optional] 
**ReferralId** | Pointer to **int32** | The integration ID of the referring customer (if any) for whom this coupon was created as an effect. | [optional] 
**RecipientIntegrationId** | Pointer to **string** | The Integration ID of the customer that is allowed to redeem this coupon. | [optional] 
**ImportId** | Pointer to **int32** | The ID of the Import which created this coupon. | [optional] 
**Reservation** | Pointer to **bool** | Defines the reservation type: - &#x60;true&#x60;: The coupon can be reserved for multiple customers. - &#x60;false&#x60;: The coupon can be reserved only for one customer. It is a personal code.  | [optional] [default to true]
**BatchId** | Pointer to **string** | The id of the batch the coupon belongs to. | [optional] 
**IsReservationMandatory** | Pointer to **bool** | An indication of whether the code can be redeemed only if it has been reserved first. | [optional] [default to false]
**ImplicitlyReserved** | Pointer to **bool** | An indication of whether the coupon is implicitly reserved for all customers. | [optional] 
**ProfileRedemptionCount** | Pointer to **int32** | The number of times the coupon was redeemed by the profile. | 
**State** | Pointer to **string** | Can be:  - &#x60;active&#x60;: The coupon can be used. It is a reserved coupon that is not pending, used, or expired, and it has a non-exhausted limit counter.    **Note:** This coupon state is returned for [scheduled campaigns](https://docs.talon.one/docs/product/campaigns/settings/managing-campaign-schedule), but the coupon cannot be used until the campaign is **running**. - &#x60;used&#x60;: The coupon has been redeemed and cannot be used again. It is not pending and has reached its redemption limit or was redeemed by the profile before expiration. - &#x60;expired&#x60;: The coupon was never redeemed, and it is now expired. It is non-pending, non-active, and non-used by the profile. - &#x60;pending&#x60;: The coupon will be usable in the future. - &#x60;disabled&#x60;: The coupon is part of a non-active campaign.  | 

## Methods

### GetId

`func (o *InventoryCoupon) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryCoupon) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *InventoryCoupon) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *InventoryCoupon) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *InventoryCoupon) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *InventoryCoupon) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *InventoryCoupon) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *InventoryCoupon) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCampaignId

`func (o *InventoryCoupon) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *InventoryCoupon) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *InventoryCoupon) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *InventoryCoupon) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetValue

`func (o *InventoryCoupon) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *InventoryCoupon) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *InventoryCoupon) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *InventoryCoupon) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.

### GetUsageLimit

`func (o *InventoryCoupon) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *InventoryCoupon) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *InventoryCoupon) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *InventoryCoupon) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetDiscountLimit

`func (o *InventoryCoupon) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *InventoryCoupon) GetDiscountLimitOk() (float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountLimit

`func (o *InventoryCoupon) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### SetDiscountLimit

`func (o *InventoryCoupon) SetDiscountLimit(v float32)`

SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.

### GetReservationLimit

`func (o *InventoryCoupon) GetReservationLimit() int32`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *InventoryCoupon) GetReservationLimitOk() (int32, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservationLimit

`func (o *InventoryCoupon) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### SetReservationLimit

`func (o *InventoryCoupon) SetReservationLimit(v int32)`

SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.

### GetStartDate

`func (o *InventoryCoupon) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *InventoryCoupon) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *InventoryCoupon) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *InventoryCoupon) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *InventoryCoupon) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *InventoryCoupon) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *InventoryCoupon) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *InventoryCoupon) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetLimits

`func (o *InventoryCoupon) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *InventoryCoupon) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *InventoryCoupon) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *InventoryCoupon) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetUsageCounter

`func (o *InventoryCoupon) GetUsageCounter() int32`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *InventoryCoupon) GetUsageCounterOk() (int32, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageCounter

`func (o *InventoryCoupon) HasUsageCounter() bool`

HasUsageCounter returns a boolean if a field has been set.

### SetUsageCounter

`func (o *InventoryCoupon) SetUsageCounter(v int32)`

SetUsageCounter gets a reference to the given int32 and assigns it to the UsageCounter field.

### GetDiscountCounter

`func (o *InventoryCoupon) GetDiscountCounter() float32`

GetDiscountCounter returns the DiscountCounter field if non-nil, zero value otherwise.

### GetDiscountCounterOk

`func (o *InventoryCoupon) GetDiscountCounterOk() (float32, bool)`

GetDiscountCounterOk returns a tuple with the DiscountCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountCounter

`func (o *InventoryCoupon) HasDiscountCounter() bool`

HasDiscountCounter returns a boolean if a field has been set.

### SetDiscountCounter

`func (o *InventoryCoupon) SetDiscountCounter(v float32)`

SetDiscountCounter gets a reference to the given float32 and assigns it to the DiscountCounter field.

### GetDiscountRemainder

`func (o *InventoryCoupon) GetDiscountRemainder() float32`

GetDiscountRemainder returns the DiscountRemainder field if non-nil, zero value otherwise.

### GetDiscountRemainderOk

`func (o *InventoryCoupon) GetDiscountRemainderOk() (float32, bool)`

GetDiscountRemainderOk returns a tuple with the DiscountRemainder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountRemainder

`func (o *InventoryCoupon) HasDiscountRemainder() bool`

HasDiscountRemainder returns a boolean if a field has been set.

### SetDiscountRemainder

`func (o *InventoryCoupon) SetDiscountRemainder(v float32)`

SetDiscountRemainder gets a reference to the given float32 and assigns it to the DiscountRemainder field.

### GetReservationCounter

`func (o *InventoryCoupon) GetReservationCounter() float32`

GetReservationCounter returns the ReservationCounter field if non-nil, zero value otherwise.

### GetReservationCounterOk

`func (o *InventoryCoupon) GetReservationCounterOk() (float32, bool)`

GetReservationCounterOk returns a tuple with the ReservationCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservationCounter

`func (o *InventoryCoupon) HasReservationCounter() bool`

HasReservationCounter returns a boolean if a field has been set.

### SetReservationCounter

`func (o *InventoryCoupon) SetReservationCounter(v float32)`

SetReservationCounter gets a reference to the given float32 and assigns it to the ReservationCounter field.

### GetAttributes

`func (o *InventoryCoupon) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InventoryCoupon) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *InventoryCoupon) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *InventoryCoupon) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetReferralId

`func (o *InventoryCoupon) GetReferralId() int32`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *InventoryCoupon) GetReferralIdOk() (int32, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralId

`func (o *InventoryCoupon) HasReferralId() bool`

HasReferralId returns a boolean if a field has been set.

### SetReferralId

`func (o *InventoryCoupon) SetReferralId(v int32)`

SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.

### GetRecipientIntegrationId

`func (o *InventoryCoupon) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *InventoryCoupon) GetRecipientIntegrationIdOk() (string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRecipientIntegrationId

`func (o *InventoryCoupon) HasRecipientIntegrationId() bool`

HasRecipientIntegrationId returns a boolean if a field has been set.

### SetRecipientIntegrationId

`func (o *InventoryCoupon) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.

### GetImportId

`func (o *InventoryCoupon) GetImportId() int32`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *InventoryCoupon) GetImportIdOk() (int32, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImportId

`func (o *InventoryCoupon) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### SetImportId

`func (o *InventoryCoupon) SetImportId(v int32)`

SetImportId gets a reference to the given int32 and assigns it to the ImportId field.

### GetReservation

`func (o *InventoryCoupon) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *InventoryCoupon) GetReservationOk() (bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservation

`func (o *InventoryCoupon) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### SetReservation

`func (o *InventoryCoupon) SetReservation(v bool)`

SetReservation gets a reference to the given bool and assigns it to the Reservation field.

### GetBatchId

`func (o *InventoryCoupon) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *InventoryCoupon) GetBatchIdOk() (string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchId

`func (o *InventoryCoupon) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchId

`func (o *InventoryCoupon) SetBatchId(v string)`

SetBatchId gets a reference to the given string and assigns it to the BatchId field.

### GetIsReservationMandatory

`func (o *InventoryCoupon) GetIsReservationMandatory() bool`

GetIsReservationMandatory returns the IsReservationMandatory field if non-nil, zero value otherwise.

### GetIsReservationMandatoryOk

`func (o *InventoryCoupon) GetIsReservationMandatoryOk() (bool, bool)`

GetIsReservationMandatoryOk returns a tuple with the IsReservationMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasIsReservationMandatory

`func (o *InventoryCoupon) HasIsReservationMandatory() bool`

HasIsReservationMandatory returns a boolean if a field has been set.

### SetIsReservationMandatory

`func (o *InventoryCoupon) SetIsReservationMandatory(v bool)`

SetIsReservationMandatory gets a reference to the given bool and assigns it to the IsReservationMandatory field.

### GetImplicitlyReserved

`func (o *InventoryCoupon) GetImplicitlyReserved() bool`

GetImplicitlyReserved returns the ImplicitlyReserved field if non-nil, zero value otherwise.

### GetImplicitlyReservedOk

`func (o *InventoryCoupon) GetImplicitlyReservedOk() (bool, bool)`

GetImplicitlyReservedOk returns a tuple with the ImplicitlyReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImplicitlyReserved

`func (o *InventoryCoupon) HasImplicitlyReserved() bool`

HasImplicitlyReserved returns a boolean if a field has been set.

### SetImplicitlyReserved

`func (o *InventoryCoupon) SetImplicitlyReserved(v bool)`

SetImplicitlyReserved gets a reference to the given bool and assigns it to the ImplicitlyReserved field.

### GetProfileRedemptionCount

`func (o *InventoryCoupon) GetProfileRedemptionCount() int32`

GetProfileRedemptionCount returns the ProfileRedemptionCount field if non-nil, zero value otherwise.

### GetProfileRedemptionCountOk

`func (o *InventoryCoupon) GetProfileRedemptionCountOk() (int32, bool)`

GetProfileRedemptionCountOk returns a tuple with the ProfileRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasProfileRedemptionCount

`func (o *InventoryCoupon) HasProfileRedemptionCount() bool`

HasProfileRedemptionCount returns a boolean if a field has been set.

### SetProfileRedemptionCount

`func (o *InventoryCoupon) SetProfileRedemptionCount(v int32)`

SetProfileRedemptionCount gets a reference to the given int32 and assigns it to the ProfileRedemptionCount field.

### GetState

`func (o *InventoryCoupon) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *InventoryCoupon) GetStateOk() (string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasState

`func (o *InventoryCoupon) HasState() bool`

HasState returns a boolean if a field has been set.

### SetState

`func (o *InventoryCoupon) SetState(v string)`

SetState gets a reference to the given string and assigns it to the State field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


