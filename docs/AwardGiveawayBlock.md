# AwardGiveawayBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**GiveawayPool** | Pointer to [**GiveawayPoolReference**](GiveawayPoolReference.md) |  | 
**Profile** | Pointer to **string** | The customer profile to award the giveaway to. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Blocks evaluated when this block fails or returns false. | [optional] 
**OnError** | Pointer to [**map[string][]map[string]interface{}**](array.md) | Named error handlers evaluated when a specific error occurs. | [optional] 

## Methods

### NewAwardGiveawayBlock

`func NewAwardGiveawayBlock(type_ string, giveawayPool GiveawayPoolReference, profile string, ) *AwardGiveawayBlock`

NewAwardGiveawayBlock instantiates a new AwardGiveawayBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardGiveawayBlockWithDefaults

`func NewAwardGiveawayBlockWithDefaults() *AwardGiveawayBlock`

NewAwardGiveawayBlockWithDefaults instantiates a new AwardGiveawayBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwardGiveawayBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwardGiveawayBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwardGiveawayBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwardGiveawayBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AwardGiveawayBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardGiveawayBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardGiveawayBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *AwardGiveawayBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwardGiveawayBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwardGiveawayBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwardGiveawayBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetGiveawayPool

`func (o *AwardGiveawayBlock) GetGiveawayPool() GiveawayPoolReference`

GetGiveawayPool returns the GiveawayPool field if non-nil, zero value otherwise.

### GetGiveawayPoolOk

`func (o *AwardGiveawayBlock) GetGiveawayPoolOk() (*GiveawayPoolReference, bool)`

GetGiveawayPoolOk returns a tuple with the GiveawayPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayPool

`func (o *AwardGiveawayBlock) SetGiveawayPool(v GiveawayPoolReference)`

SetGiveawayPool sets GiveawayPool field to given value.


### GetProfile

`func (o *AwardGiveawayBlock) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *AwardGiveawayBlock) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *AwardGiveawayBlock) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetOnFailure

`func (o *AwardGiveawayBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *AwardGiveawayBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *AwardGiveawayBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *AwardGiveawayBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.

### GetOnError

`func (o *AwardGiveawayBlock) GetOnError() map[string][]map[string]interface{}`

GetOnError returns the OnError field if non-nil, zero value otherwise.

### GetOnErrorOk

`func (o *AwardGiveawayBlock) GetOnErrorOk() (*map[string][]map[string]interface{}, bool)`

GetOnErrorOk returns a tuple with the OnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnError

`func (o *AwardGiveawayBlock) SetOnError(v map[string][]map[string]interface{})`

SetOnError sets OnError field to given value.

### HasOnError

`func (o *AwardGiveawayBlock) HasOnError() bool`

HasOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


