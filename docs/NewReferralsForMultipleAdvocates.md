# NewReferralsForMultipleAdvocates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the referral code becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the referral code. Referral never expires if this is omitted. | [optional] 
**UsageLimit** | Pointer to **int32** | The number of times a referral code can be used. &#x60;0&#x60; means no limit but any campaign usage limits will still apply.  | 
**CampaignId** | Pointer to **int32** | The ID of the campaign from which the referral received the referral code. | 
**AdvocateProfileIntegrationIds** | Pointer to **[]string** | An array containing all the respective advocate profiles. | 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this referral code. | [optional] 
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the &#x60;[A-Z, 0-9]&#x60; regular expression.  | [optional] 
**ReferralPattern** | Pointer to **string** | The pattern used to generate referrals. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### GetStartDate

`func (o *NewReferralsForMultipleAdvocates) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *NewReferralsForMultipleAdvocates) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *NewReferralsForMultipleAdvocates) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *NewReferralsForMultipleAdvocates) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *NewReferralsForMultipleAdvocates) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *NewReferralsForMultipleAdvocates) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *NewReferralsForMultipleAdvocates) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *NewReferralsForMultipleAdvocates) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetUsageLimit

`func (o *NewReferralsForMultipleAdvocates) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *NewReferralsForMultipleAdvocates) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *NewReferralsForMultipleAdvocates) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *NewReferralsForMultipleAdvocates) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetCampaignId

`func (o *NewReferralsForMultipleAdvocates) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *NewReferralsForMultipleAdvocates) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *NewReferralsForMultipleAdvocates) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *NewReferralsForMultipleAdvocates) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocates) GetAdvocateProfileIntegrationIds() []string`

GetAdvocateProfileIntegrationIds returns the AdvocateProfileIntegrationIds field if non-nil, zero value otherwise.

### GetAdvocateProfileIntegrationIdsOk

`func (o *NewReferralsForMultipleAdvocates) GetAdvocateProfileIntegrationIdsOk() ([]string, bool)`

GetAdvocateProfileIntegrationIdsOk returns a tuple with the AdvocateProfileIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocates) HasAdvocateProfileIntegrationIds() bool`

HasAdvocateProfileIntegrationIds returns a boolean if a field has been set.

### SetAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocates) SetAdvocateProfileIntegrationIds(v []string)`

SetAdvocateProfileIntegrationIds gets a reference to the given []string and assigns it to the AdvocateProfileIntegrationIds field.

### GetAttributes

`func (o *NewReferralsForMultipleAdvocates) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewReferralsForMultipleAdvocates) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewReferralsForMultipleAdvocates) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewReferralsForMultipleAdvocates) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetValidCharacters

`func (o *NewReferralsForMultipleAdvocates) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewReferralsForMultipleAdvocates) GetValidCharactersOk() ([]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidCharacters

`func (o *NewReferralsForMultipleAdvocates) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### SetValidCharacters

`func (o *NewReferralsForMultipleAdvocates) SetValidCharacters(v []string)`

SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.

### GetReferralPattern

`func (o *NewReferralsForMultipleAdvocates) GetReferralPattern() string`

GetReferralPattern returns the ReferralPattern field if non-nil, zero value otherwise.

### GetReferralPatternOk

`func (o *NewReferralsForMultipleAdvocates) GetReferralPatternOk() (string, bool)`

GetReferralPatternOk returns a tuple with the ReferralPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralPattern

`func (o *NewReferralsForMultipleAdvocates) HasReferralPattern() bool`

HasReferralPattern returns a boolean if a field has been set.

### SetReferralPattern

`func (o *NewReferralsForMultipleAdvocates) SetReferralPattern(v string)`

SetReferralPattern gets a reference to the given string and assigns it to the ReferralPattern field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


