# CreateCouponBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**CampaignId** | Pointer to [**map[string]interface{}**](.md) | The ID of the campaign in which the coupon code is created. | 
**RecipientId** | Pointer to **string** | The integration ID of the customer that is allowed to redeem this coupon. | 
**StoreInSession** | Pointer to **bool** | When &#x60;true&#x60;, the coupon is stored in the session. | 
**UsageLimit** | Pointer to [**map[string]interface{}**](.md) | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions, but any campaign usage limits still apply. Either a numeric scalar or a &#x60;{{expression}}&#x60; string that resolves to a number at evaluation time.  | [optional] 
**DiscountLimit** | Pointer to [**map[string]interface{}**](.md) | The total discount value that the code can give. Typically used to represent a gift card value. Either a numeric scalar or a &#x60;{{expression}}&#x60; string that resolves to a number at evaluation time.  | [optional] 
**StartDate** | Pointer to [**map[string]interface{}**](.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**map[string]interface{}**](.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Custom attributes associated with this coupon code. | [optional] 
**ValidCharacters** | Pointer to **string** | Characters used to generate the random parts of a code. | [optional] 
**Pattern** | Pointer to **string** | The pattern used to generate codes, such as coupon codes, referral codes, and loyalty cards. The character &#x60;#&#x60; is a placeholder and is replaced by a random character from the &#x60;validCharacters&#x60; set.  | [optional] 

## Methods

### NewCreateCouponBlock

`func NewCreateCouponBlock(type_ string, campaignId map[string]interface{}, recipientId string, storeInSession bool, ) *CreateCouponBlock`

NewCreateCouponBlock instantiates a new CreateCouponBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCouponBlockWithDefaults

`func NewCreateCouponBlockWithDefaults() *CreateCouponBlock`

NewCreateCouponBlockWithDefaults instantiates a new CreateCouponBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateCouponBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCouponBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCouponBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCouponBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CreateCouponBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCouponBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCouponBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CreateCouponBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateCouponBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateCouponBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateCouponBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCampaignId

`func (o *CreateCouponBlock) GetCampaignId() map[string]interface{}`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CreateCouponBlock) GetCampaignIdOk() (*map[string]interface{}, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CreateCouponBlock) SetCampaignId(v map[string]interface{})`

SetCampaignId sets CampaignId field to given value.


### GetRecipientId

`func (o *CreateCouponBlock) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *CreateCouponBlock) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *CreateCouponBlock) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetStoreInSession

`func (o *CreateCouponBlock) GetStoreInSession() bool`

GetStoreInSession returns the StoreInSession field if non-nil, zero value otherwise.

### GetStoreInSessionOk

`func (o *CreateCouponBlock) GetStoreInSessionOk() (*bool, bool)`

GetStoreInSessionOk returns a tuple with the StoreInSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreInSession

`func (o *CreateCouponBlock) SetStoreInSession(v bool)`

SetStoreInSession sets StoreInSession field to given value.


### GetUsageLimit

`func (o *CreateCouponBlock) GetUsageLimit() map[string]interface{}`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *CreateCouponBlock) GetUsageLimitOk() (*map[string]interface{}, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *CreateCouponBlock) SetUsageLimit(v map[string]interface{})`

SetUsageLimit sets UsageLimit field to given value.

### HasUsageLimit

`func (o *CreateCouponBlock) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### GetDiscountLimit

`func (o *CreateCouponBlock) GetDiscountLimit() map[string]interface{}`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *CreateCouponBlock) GetDiscountLimitOk() (*map[string]interface{}, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *CreateCouponBlock) SetDiscountLimit(v map[string]interface{})`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *CreateCouponBlock) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *CreateCouponBlock) GetStartDate() map[string]interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CreateCouponBlock) GetStartDateOk() (*map[string]interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CreateCouponBlock) SetStartDate(v map[string]interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CreateCouponBlock) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *CreateCouponBlock) GetExpiryDate() map[string]interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CreateCouponBlock) GetExpiryDateOk() (*map[string]interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CreateCouponBlock) SetExpiryDate(v map[string]interface{})`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CreateCouponBlock) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetAttributes

`func (o *CreateCouponBlock) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CreateCouponBlock) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CreateCouponBlock) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CreateCouponBlock) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetValidCharacters

`func (o *CreateCouponBlock) GetValidCharacters() string`

GetValidCharacters returns the ValidCharacters field if non-nil, zero value otherwise.

### GetValidCharactersOk

`func (o *CreateCouponBlock) GetValidCharactersOk() (*string, bool)`

GetValidCharactersOk returns a tuple with the ValidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidCharacters

`func (o *CreateCouponBlock) SetValidCharacters(v string)`

SetValidCharacters sets ValidCharacters field to given value.

### HasValidCharacters

`func (o *CreateCouponBlock) HasValidCharacters() bool`

HasValidCharacters returns a boolean if a field has been set.

### GetPattern

`func (o *CreateCouponBlock) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateCouponBlock) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateCouponBlock) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateCouponBlock) HasPattern() bool`

HasPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


