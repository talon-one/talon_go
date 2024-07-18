# NewReferralsForMultipleAdvocatesAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | Pointer to **int32** | The ID of the campaign from which the referral received the referral code. | 
**AdvocateProfileIntegrationIds** | Pointer to **[]string** | An array containing all the respective advocate profiles. | 
**Attributes** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Arbitrary properties associated with this referral code. | [optional] 
**ValidCharacters** | Pointer to **[]string** | List of characters used to generate the random parts of a code. By default, the list of characters is equivalent to the &#x60;[A-Z, 0-9]&#x60; regular expression.  | [optional] 
**ReferralPattern** | Pointer to **string** | The pattern used to generate referrals. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### GetCampaignId

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *NewReferralsForMultipleAdvocatesAllOf) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *NewReferralsForMultipleAdvocatesAllOf) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetAdvocateProfileIntegrationIds() []string`

GetAdvocateProfileIntegrationIds returns the AdvocateProfileIntegrationIds field if non-nil, zero value otherwise.

### GetAdvocateProfileIntegrationIdsOk

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetAdvocateProfileIntegrationIdsOk() ([]string, bool)`

GetAdvocateProfileIntegrationIdsOk returns a tuple with the AdvocateProfileIntegrationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocatesAllOf) HasAdvocateProfileIntegrationIds() bool`

HasAdvocateProfileIntegrationIds returns a boolean if a field has been set.

### SetAdvocateProfileIntegrationIds

`func (o *NewReferralsForMultipleAdvocatesAllOf) SetAdvocateProfileIntegrationIds(v []string)`

SetAdvocateProfileIntegrationIds gets a reference to the given []string and assigns it to the AdvocateProfileIntegrationIds field.

### GetAttributes

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetAttributes() map[string]map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetAttributesOk() (map[string]map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewReferralsForMultipleAdvocatesAllOf) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewReferralsForMultipleAdvocatesAllOf) SetAttributes(v map[string]map[string]interface{})`

SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.

### GetValidCharacters

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetValidCharacters() []string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetValidCharactersOk() ([]string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValidCharacters

`func (o *NewReferralsForMultipleAdvocatesAllOf) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### SetValidCharacters

`func (o *NewReferralsForMultipleAdvocatesAllOf) SetValidCharacters(v []string)`

SetValidCharacters gets a reference to the given []string and assigns it to the ValidCharacters field.

### GetReferralPattern

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetReferralPattern() string`

GetReferralPattern returns the ReferralPattern field if non-nil, zero value otherwise.

### GetReferralPatternOk

`func (o *NewReferralsForMultipleAdvocatesAllOf) GetReferralPatternOk() (string, bool)`

GetReferralPatternOk returns a tuple with the ReferralPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReferralPattern

`func (o *NewReferralsForMultipleAdvocatesAllOf) HasReferralPattern() bool`

HasReferralPattern returns a boolean if a field has been set.

### SetReferralPattern

`func (o *NewReferralsForMultipleAdvocatesAllOf) SetReferralPattern(v string)`

SetReferralPattern gets a reference to the given string and assigns it to the ReferralPattern field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


