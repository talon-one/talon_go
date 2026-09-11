# CreateReferralBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**CampaignId** | Pointer to [**map[string]interface{}**](.md) | The ID of the campaign in which the referral code is created. Either a numeric scalar or a &#x60;{{expression}}&#x60; string that resolves to a number at evaluation time. | 
**FriendId** | Pointer to **string** | An optional integration ID of the friend&#39;s profile. | 
**StoreInSession** | Pointer to **bool** | When &#x60;true&#x60;, the referral code is stored in the session. | 
**UsageLimit** | Pointer to [**map[string]interface{}**](.md) | The number of times the referral code code can be redeemed. &#x60;0&#x60; means unlimited redemptions, but any campaign usage limits still apply. Either a numeric scalar or a &#x60;{{expression}}&#x60; string that resolves to a number at evaluation time.  | [optional] 
**StartDate** | Pointer to [**map[string]interface{}**](.md) | Timestamp at which point the referral code becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**map[string]interface{}**](.md) | Expiration date of the referral code. Referral code never expires if this is omitted. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Custom attributes associated with this referral code. | [optional] 
**ValidCharacters** | Pointer to **string** | Characters used to generate the random parts of a code. | [optional] 
**Pattern** | Pointer to **string** | The pattern used to generate codes, such as coupon codes, referral codes, and loyalty cards. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### NewCreateReferralBlock

`func NewCreateReferralBlock(type_ string, campaignId map[string]interface{}, friendId string, storeInSession bool, ) *CreateReferralBlock`

NewCreateReferralBlock instantiates a new CreateReferralBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReferralBlockWithDefaults

`func NewCreateReferralBlockWithDefaults() *CreateReferralBlock`

NewCreateReferralBlockWithDefaults instantiates a new CreateReferralBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateReferralBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateReferralBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateReferralBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateReferralBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CreateReferralBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateReferralBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateReferralBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CreateReferralBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateReferralBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateReferralBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateReferralBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCampaignId

`func (o *CreateReferralBlock) GetCampaignId() map[string]interface{}`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CreateReferralBlock) GetCampaignIdOk() (*map[string]interface{}, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CreateReferralBlock) SetCampaignId(v map[string]interface{})`

SetCampaignId sets CampaignId field to given value.


### GetFriendId

`func (o *CreateReferralBlock) GetFriendId() string`

GetFriendId returns the FriendId field if non-nil, zero value otherwise.

### GetFriendIdOk

`func (o *CreateReferralBlock) GetFriendIdOk() (*string, bool)`

GetFriendIdOk returns a tuple with the FriendId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendId

`func (o *CreateReferralBlock) SetFriendId(v string)`

SetFriendId sets FriendId field to given value.


### GetStoreInSession

`func (o *CreateReferralBlock) GetStoreInSession() bool`

GetStoreInSession returns the StoreInSession field if non-nil, zero value otherwise.

### GetStoreInSessionOk

`func (o *CreateReferralBlock) GetStoreInSessionOk() (*bool, bool)`

GetStoreInSessionOk returns a tuple with the StoreInSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInSession

`func (o *CreateReferralBlock) SetStoreInSession(v bool)`

SetStoreInSession sets StoreInSession field to given value.


### GetUsageLimit

`func (o *CreateReferralBlock) GetUsageLimit() map[string]interface{}`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *CreateReferralBlock) GetUsageLimitOk() (*map[string]interface{}, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *CreateReferralBlock) SetUsageLimit(v map[string]interface{})`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *CreateReferralBlock) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateReferralBlock) GetStartDate() map[string]interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateReferralBlock) GetStartDateOk() (*map[string]interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateReferralBlock) SetStartDate(v map[string]interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateReferralBlock) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *CreateReferralBlock) GetExpiryDate() map[string]interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CreateReferralBlock) GetExpiryDateOk() (*map[string]interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CreateReferralBlock) SetExpiryDate(v map[string]interface{})`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CreateReferralBlock) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetAttributes

`func (o *CreateReferralBlock) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateReferralBlock) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateReferralBlock) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateReferralBlock) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetValidCharacters

`func (o *CreateReferralBlock) GetValidCharacters() string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *CreateReferralBlock) GetValidCharactersOk() (*string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *CreateReferralBlock) SetValidCharacters(v string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *CreateReferralBlock) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetPattern

`func (o *CreateReferralBlock) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateReferralBlock) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateReferralBlock) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateReferralBlock) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


