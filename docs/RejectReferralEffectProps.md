# RejectReferralEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The referral code that was rejected | 
**RejectionReason** | Pointer to **string** | The reason why the code was rejected.  - &#x60;AdvocateNotFound&#x60;: The advocate was not found. - &#x60;CampaignLimitReached&#x60;: The campaign-wide referral code redemption limit has been reached. - &#x60;EffectCouldNotBeApplied&#x60;: One of the effects in the campaign wasn&#39;t applied because a limit for that effect was reached (most common use case will be &#x60;setDiscount&#x60; can not be applied because a discount limit is reached). - &#x60;ProfileLimitReached&#x60;: The profile-specific referral code redemption limit has been reached. - &#x60;ReferralCustomerAlreadyReferred&#x60;: The friend is already referred. - &#x60;ReferralExpired&#x60;: The transferred referral code is expired. - &#x60;ReferralLimitReached&#x60;: The referral code redemption limit has been reached. - &#x60;ReferralNotFound&#x60;: The transferred referral code is wrong. - &#x60;ReferralPartOfNotRunningCampaign&#x60;: The campaign the referral code belongs to is currently not active. The campaign ID field shows the ID of that campaign. - &#x60;ReferralRecipientDoesNotMatch&#x60;: The given referral code value does not match the recipient. - &#x60;ReferralRecipientIdSameAsAdvocate&#x60;: The recipient (friend) has the same id as the advocate. - &#x60;ReferralRejectedByCondition&#x60;: The referral code is valid and in an active campaign, but there were other conditions in that campaign&#39;s rules that were not met. - &#x60;ReferralStartDateInFuture&#x60;: The transferred referral code isn&#39;t active yet. - &#x60;ReferralPartOfNotTriggeredCampaign&#x60;: The campaign the referral code belongs to was not triggered during evaluation (an exclusive or stackable campaign). The campaign ID field shows the ID of that campaign. | 
**ConditionIndex** | Pointer to **int64** | The index of the condition that caused the rejection of the referral. | [optional] 
**EffectIndex** | Pointer to **int64** | The index of the effect that caused the rejection of the referral. | [optional] 
**Details** | Pointer to **string** | More details about the failure. | [optional] 
**CampaignExclusionReason** | Pointer to **string** | The reason why the campaign the referral belongs to was excluded during [campaign evaluation](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation), when &#x60;rejectionReason&#x60; was &#x60;CouponPartOfNotTriggeredCampaign&#x60;. Its possible values are:  - &#x60;CampaignGaveLowerDiscount&#x60;: The required campaign and referral conditions were met, but another campaign in a [Highest discount value](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation#set-campaign-evaluation-mode) group offered a higher discount value. - &#x60;CampaignIsNotFirst&#x60;: The campaign was not evaluated because another campaign in a [First campaign](https://docs.talon.one/docs/product/applications/manage-campaign-evaluation#set-campaign-evaluation-mode) group was picked and evaluated first. - &#x60;CampaignNotInEvaluationSet&#x60;: The campaign did not meet other evaluation requirements, for example, because the referral is part of an archived campaign. | [optional] 

## Methods

### NewRejectReferralEffectProps

`func NewRejectReferralEffectProps(value string, rejectionReason string, ) *RejectReferralEffectProps`

NewRejectReferralEffectProps instantiates a new RejectReferralEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRejectReferralEffectPropsWithDefaults

`func NewRejectReferralEffectPropsWithDefaults() *RejectReferralEffectProps`

NewRejectReferralEffectPropsWithDefaults instantiates a new RejectReferralEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *RejectReferralEffectProps) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RejectReferralEffectProps) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RejectReferralEffectProps) SetValue(v string)`

SetValue sets Value field to given value.


### GetRejectionReason

`func (o *RejectReferralEffectProps) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *RejectReferralEffectProps) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *RejectReferralEffectProps) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.


### GetConditionIndex

`func (o *RejectReferralEffectProps) GetConditionIndex() int64`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *RejectReferralEffectProps) GetConditionIndexOk() (*int64, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *RejectReferralEffectProps) SetConditionIndex(v int64)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *RejectReferralEffectProps) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetEffectIndex

`func (o *RejectReferralEffectProps) GetEffectIndex() int64`

GetEffectIndex returns the EffectIndex field if non-nil, zero value otherwise.

### GetEffectIndexOk

`func (o *RejectReferralEffectProps) GetEffectIndexOk() (*int64, bool)`

GetEffectIndexOk returns a tuple with the EffectIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectIndex

`func (o *RejectReferralEffectProps) SetEffectIndex(v int64)`

SetEffectIndex sets EffectIndex field to given value.

### HasEffectIndex

`func (o *RejectReferralEffectProps) HasEffectIndex() bool`

HasEffectIndex returns a boolean if a field has been set.

### GetDetails

`func (o *RejectReferralEffectProps) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RejectReferralEffectProps) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RejectReferralEffectProps) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RejectReferralEffectProps) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCampaignExclusionReason

`func (o *RejectReferralEffectProps) GetCampaignExclusionReason() string`

GetCampaignExclusionReason returns the CampaignExclusionReason field if non-nil, zero value otherwise.

### GetCampaignExclusionReasonOk

`func (o *RejectReferralEffectProps) GetCampaignExclusionReasonOk() (*string, bool)`

GetCampaignExclusionReasonOk returns a tuple with the CampaignExclusionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignExclusionReason

`func (o *RejectReferralEffectProps) SetCampaignExclusionReason(v string)`

SetCampaignExclusionReason sets CampaignExclusionReason field to given value.

### HasCampaignExclusionReason

`func (o *RejectReferralEffectProps) HasCampaignExclusionReason() bool`

HasCampaignExclusionReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


