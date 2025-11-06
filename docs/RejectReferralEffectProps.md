# RejectReferralEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The referral code that was rejected. | 
**RejectionReason** | Pointer to **string** | The reason why this referral code was rejected. | 
**ConditionIndex** | Pointer to **int64** | The index of the condition that caused the rejection of the referral. | [optional] 
**EffectIndex** | Pointer to **int64** | The index of the effect that caused the rejection of the referral. | [optional] 
**Details** | Pointer to **string** | More details about the failure. | [optional] 
**CampaignExclusionReason** | Pointer to **string** | The reason why the campaign was not applied. | [optional] 

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


