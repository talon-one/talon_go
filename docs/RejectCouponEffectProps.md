# RejectCouponEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The coupon code that was rejected. | 
**RejectionReason** | Pointer to **string** | The reason why the code was rejected.  - &#x60;CampaignLimitReached&#x60;: The campaign-wide coupon code redemption limit has been reached. - &#x60;CouponExpired&#x60;: The coupon is expired. - &#x60;CouponLimitReached&#x60;: The coupon redemption limit or a campaign budget was reached. - &#x60;CouponNotFound&#x60;: The coupon code is incorrect. - &#x60;CouponPartOfNotRunningCampaign&#x60;: The campaign the coupon belongs to is currently not active. The campaignId field contains the ID of that campaign. - &#x60;CouponRecipientDoesNotMatch&#x60;: The given coupon value does not match the recipient or the coupon is linked to a &#x60;recipientIntegrationID&#x60; but there is no profile in the session. - &#x60;CouponRejectedByCondition&#x60;: Other conditions failed in the rule or all conditions passed but the &#x60;Coupon code is valid&#x60; condition is not present. - &#x60;CouponStartDateInFuture&#x60;: The coupon isn&#39;t active yet. - &#x60;EffectCouldNotBeApplied&#x60;: One of the effects in the campaign wasn&#39;t applied because a limit for that effect was reached (most common use case will be &#x60;setDiscount&#x60; cannot be applied because a discount limit is reached). - &#x60;ProfileLimitReached&#x60;: The profile-specific coupon redemption limit has been reached. - &#x60;CouponPartOfNotTriggeredCampaign&#x60;: The campaign the coupon belongs to was not triggered during evaluation (an exclusive or stackable campaign). The &#x60;campaignId&#x60; field contains the ID of that campaign. - &#x60;CouponReservationRequired&#x60;: The coupon&#39;s &#x60;isReservationMandatory&#x60; property is &#x60;true&#x60;, but the profile does not have a reservation. - &#x60;ProfileRequired&#x60;: The coupon&#39;s &#x60;isReservationMandatory&#x60; property is &#x60;true&#x60; or a [campaign profile budget](https://docs.talon.one/docs/product/campaigns/settings/manage-campaign-budgets) was set, but no profile exists in the session. | 
**ConditionIndex** | Pointer to **int64** | The index of the condition that caused the rejection of the coupon. | [optional] 
**EffectIndex** | Pointer to **int64** | The index of the effect that caused the rejection of the coupon. | [optional] 
**Details** | Pointer to **string** | More details about the failure. | [optional] 
**CampaignExclusionReason** | Pointer to **string** | The reason why the campaign the coupon belongs to was excluded during [campaign evaluation](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation), when &#x60;rejectionReason&#x60; was &#x60;CouponPartOfNotTriggeredCampaign&#x60;. Its possible values are:  - &#x60;CampaignGaveLowerDiscount&#x60;: The required campaign and coupon conditions were met, but another campaign in a [Highest discount value](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation#set-campaign-evaluation-mode) group offered a higher discount value. - &#x60;CampaignIsNotFirst&#x60;: The campaign was not evaluated because another campaign in a [First campaign](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation#set-campaign-evaluation-mode) group was picked and evaluated first. - &#x60;CampaignNotInEvaluationSet&#x60;: The campaign did not meet other evaluation requirements, for example, because the coupon is part of an archived campaign. | [optional] 

## Methods

### NewRejectCouponEffectProps

`func NewRejectCouponEffectProps(value string, rejectionReason string, ) *RejectCouponEffectProps`

NewRejectCouponEffectProps instantiates a new RejectCouponEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectCouponEffectPropsWithDefaults

`func NewRejectCouponEffectPropsWithDefaults() *RejectCouponEffectProps`

NewRejectCouponEffectPropsWithDefaults instantiates a new RejectCouponEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RejectCouponEffectProps) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RejectCouponEffectProps) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RejectCouponEffectProps) SetValue(v string)`

SetValue sets Value field to given value.


### GetRejectionReason

`func (o *RejectCouponEffectProps) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *RejectCouponEffectProps) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *RejectCouponEffectProps) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.


### GetConditionIndex

`func (o *RejectCouponEffectProps) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *RejectCouponEffectProps) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *RejectCouponEffectProps) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *RejectCouponEffectProps) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetEffectIndex

`func (o *RejectCouponEffectProps) GetEffectIndex() int64`

GetEffectIndex returns the EffectIndex field if non-nil, zero value otherwise.

### GetEffectIndexOk

`func (o *RejectCouponEffectProps) GetEffectIndexOk() (*int64, bool)`

GetEffectIndexOk returns a tuple with the EffectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectIndex

`func (o *RejectCouponEffectProps) SetEffectIndex(v int64)`

SetEffectIndex sets EffectIndex field to given value.

### HasEffectIndex

`func (o *RejectCouponEffectProps) HasEffectIndex() bool`

HasEffectIndex returns a boolean if a field has been set.

### GetDetails

`func (o *RejectCouponEffectProps) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RejectCouponEffectProps) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RejectCouponEffectProps) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RejectCouponEffectProps) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCampaignExclusionReason

`func (o *RejectCouponEffectProps) GetCampaignExclusionReason() string`

GetCampaignExclusionReason returns the CampaignExclusionReason field if non-nil, zero value otherwise.

### GetCampaignExclusionReasonOk

`func (o *RejectCouponEffectProps) GetCampaignExclusionReasonOk() (*string, bool)`

GetCampaignExclusionReasonOk returns a tuple with the CampaignExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignExclusionReason

`func (o *RejectCouponEffectProps) SetCampaignExclusionReason(v string)`

SetCampaignExclusionReason sets CampaignExclusionReason field to given value.

### HasCampaignExclusionReason

`func (o *RejectCouponEffectProps) HasCampaignExclusionReason() bool`

HasCampaignExclusionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


