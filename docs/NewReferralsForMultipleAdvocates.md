# NewReferralsForMultipleAdvocates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the referral code becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the referral code. Referral never expires if this is omitted. | [optional] 
**UsageLimit** | Pointer to **int64** | The number of times a referral code can be used. &#x60;0&#x60; means no limit but any campaign usage limits will still apply.  | 
**CampaignId** | Pointer to **int64** | The ID of the campaign from which the referral received the referral code. | 
**AdvocateProfileIntegrationIds** | Pointer to **[]string** | An array containing all the respective advocate profiles. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this referral code. | [optional] 
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the &#x60;[A-Z, 0-9]&#x60; regular expression.  | [optional] 
**ReferralPattern** | Pointer to **string** | The pattern used to generate referrals. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### NewNewReferralsForMultipleAdvocates

`func NewNewReferralsForMultipleAdvocates(usageLimit int64, campaignId int64, advocateProfileIntegrationIds []string, ) *NewReferralsForMultipleAdvocates`

NewNewReferralsForMultipleAdvocates instantiates a new NewReferralsForMultipleAdvocates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewReferralsForMultipleAdvocatesWithDefaults

`func NewNewReferralsForMultipleAdvocatesWithDefaults() *NewReferralsForMultipleAdvocates`

NewNewReferralsForMultipleAdvocatesWithDefaults instantiates a new NewReferralsForMultipleAdvocates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartDate

`func (o *NewReferralsForMultipleAdvocates) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewReferralsForMultipleAdvocates) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *NewReferralsForMultipleAdvocates) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *NewReferralsForMultipleAdvocates) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *NewReferralsForMultipleAdvocates) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewReferralsForMultipleAdvocates) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *NewReferralsForMultipleAdvocates) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *NewReferralsForMultipleAdvocates) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetUsageLimit

`func (o *NewReferralsForMultipleAdvocates) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewReferralsForMultipleAdvocates) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *NewReferralsForMultipleAdvocates) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetCampaignId

`func (o *NewReferralsForMultipleAdvocates) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *NewReferralsForMultipleAdvocates) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *NewReferralsForMultipleAdvocates) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocates) GetAdvocateProfileIntegrationIds() []string`

GetAdvocateProfileIntegrationIds returns the AdvocateProfileIntegrationIds field if non-nil, zero value otherwise.

### GetAdvocateProfileIntegrationIdsOk

`func (o *NewReferralsForMultipleAdvocates) GetAdvocateProfileIntegrationIdsOk() (*[]string, bool)`

GetAdvocateProfileIntegrationIdsOk returns a tuple with the AdvocateProfileIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocates) SetAdvocateProfileIntegrationIds(v []string)`

SetAdvocateProfileIntegrationIds sets AdvocateProfileIntegrationIds field to given value.


### GetAttributes

`func (o *NewReferralsForMultipleAdvocates) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewReferralsForMultipleAdvocates) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewReferralsForMultipleAdvocates) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewReferralsForMultipleAdvocates) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetValidCharacters

`func (o *NewReferralsForMultipleAdvocates) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewReferralsForMultipleAdvocates) GetValidCharactersOk() (*[]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *NewReferralsForMultipleAdvocates) SetValidCharacters(v []string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *NewReferralsForMultipleAdvocates) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetReferralPattern

`func (o *NewReferralsForMultipleAdvocates) GetReferralPattern() string`

GetReferralPattern returns the ReferralPattern field if non-nil, zero value otherwise.

### GetReferralPatternOk

`func (o *NewReferralsForMultipleAdvocates) GetReferralPatternOk() (*string, bool)`

GetReferralPatternOk returns a tuple with the ReferralPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferralPattern

`func (o *NewReferralsForMultipleAdvocates) SetReferralPattern(v string)`

SetReferralPattern sets ReferralPattern field to given value.

### HasReferralPattern

`func (o *NewReferralsForMultipleAdvocates) HasReferralPattern() bool`

HasReferralPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


