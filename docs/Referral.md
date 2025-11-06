# Referral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the referral code becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the referral code. Referral never expires if this is omitted. | [optional] 
**UsageLimit** | Pointer to **int64** | The number of times a referral code can be used. &#x60;0&#x60; means no limit but any campaign usage limits will still apply.  | 
**CampaignId** | Pointer to **int64** | ID of the campaign from which the referral received the referral code. | 
**AdvocateProfileIntegrationId** | Pointer to **string** | The Integration ID of the Advocate&#39;s Profile. | 
**FriendProfileIntegrationId** | Pointer to **string** | An optional Integration ID of the Friend&#39;s Profile. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this item. | [optional] 
**ImportId** | Pointer to **int64** | The ID of the Import which created this referral. | [optional] 
**Code** | Pointer to **string** | The referral code. | 
**UsageCounter** | Pointer to **int64** | The number of times this referral code has been successfully used. | 
**BatchId** | Pointer to **string** | The ID of the batch the referrals belong to. | [optional] 

## Methods

### NewReferral

`func NewReferral(id int64, created time.Time, usageLimit int64, campaignId int64, advocateProfileIntegrationId string, code string, usageCounter int64, ) *Referral`

NewReferral instantiates a new Referral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferralWithDefaults

`func NewReferralWithDefaults() *Referral`

NewReferralWithDefaults instantiates a new Referral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Referral) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Referral) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Referral) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Referral) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Referral) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Referral) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetStartDate

`func (o *Referral) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Referral) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Referral) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Referral) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *Referral) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *Referral) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *Referral) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *Referral) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetUsageLimit

`func (o *Referral) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *Referral) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *Referral) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetCampaignId

`func (o *Referral) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *Referral) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *Referral) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetAdvocateProfileIntegrationId

`func (o *Referral) GetAdvocateProfileIntegrationId() string`

GetAdvocateProfileIntegrationId returns the AdvocateProfileIntegrationId field if non-nil, zero value otherwise.

### GetAdvocateProfileIntegrationIdOk

`func (o *Referral) GetAdvocateProfileIntegrationIdOk() (*string, bool)`

GetAdvocateProfileIntegrationIdOk returns a tuple with the AdvocateProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvocateProfileIntegrationId

`func (o *Referral) SetAdvocateProfileIntegrationId(v string)`

SetAdvocateProfileIntegrationId sets AdvocateProfileIntegrationId field to given value.


### GetFriendProfileIntegrationId

`func (o *Referral) GetFriendProfileIntegrationId() string`

GetFriendProfileIntegrationId returns the FriendProfileIntegrationId field if non-nil, zero value otherwise.

### GetFriendProfileIntegrationIdOk

`func (o *Referral) GetFriendProfileIntegrationIdOk() (*string, bool)`

GetFriendProfileIntegrationIdOk returns a tuple with the FriendProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendProfileIntegrationId

`func (o *Referral) SetFriendProfileIntegrationId(v string)`

SetFriendProfileIntegrationId sets FriendProfileIntegrationId field to given value.

### HasFriendProfileIntegrationId

`func (o *Referral) HasFriendProfileIntegrationId() bool`

HasFriendProfileIntegrationId returns a boolean if a field has been set.

### GetAttributes

`func (o *Referral) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Referral) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Referral) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Referral) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetImportId

`func (o *Referral) GetImportId() int64`

GetImportId returns the ImportId field if non-nil, zero value otherwise.

### GetImportIdOk

`func (o *Referral) GetImportIdOk() (*int64, bool)`

GetImportIdOk returns a tuple with the ImportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportId

`func (o *Referral) SetImportId(v int64)`

SetImportId sets ImportId field to given value.

### HasImportId

`func (o *Referral) HasImportId() bool`

HasImportId returns a boolean if a field has been set.

### GetCode

`func (o *Referral) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Referral) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Referral) SetCode(v string)`

SetCode sets Code field to given value.


### GetUsageCounter

`func (o *Referral) GetUsageCounter() int64`

GetUsageCounter returns the UsageCounter field if non-nil, zero value otherwise.

### GetUsageCounterOk

`func (o *Referral) GetUsageCounterOk() (*int64, bool)`

GetUsageCounterOk returns a tuple with the UsageCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageCounter

`func (o *Referral) SetUsageCounter(v int64)`

SetUsageCounter sets UsageCounter field to given value.


### GetBatchId

`func (o *Referral) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *Referral) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *Referral) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *Referral) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


