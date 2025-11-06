# ReferralRejectionReason

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int64** |  | 
**ReferralId** | Pointer to **int64** |  | 
**Reason** | Pointer to **string** |  | 

## Methods

### NewReferralRejectionReason

`func NewReferralRejectionReason(campaignId int64, referralId int64, reason string, ) *ReferralRejectionReason`

NewReferralRejectionReason instantiates a new ReferralRejectionReason object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralRejectionReasonWithDefaults

`func NewReferralRejectionReasonWithDefaults() *ReferralRejectionReason`

NewReferralRejectionReasonWithDefaults instantiates a new ReferralRejectionReason object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *ReferralRejectionReason) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ReferralRejectionReason) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ReferralRejectionReason) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetReferralId

`func (o *ReferralRejectionReason) GetReferralId() int64`

GetReferralId returns the ReferralId field if non-nil, zero value otherwise.

### GetReferralIdOk

`func (o *ReferralRejectionReason) GetReferralIdOk() (*int64, bool)`

GetReferralIdOk returns a tuple with the ReferralId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralId

`func (o *ReferralRejectionReason) SetReferralId(v int64)`

SetReferralId sets ReferralId field to given value.


### GetReason

`func (o *ReferralRejectionReason) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ReferralRejectionReason) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ReferralRejectionReason) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


