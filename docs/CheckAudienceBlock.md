# CheckAudienceBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | An indicator of how the block compares its elements. | 
**Profile** | Pointer to **string** | The customer profile to check against the audience. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Audience** | Pointer to [**CheckAudienceBlockAudience**](CheckAudienceBlock_audience.md) |  | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewCheckAudienceBlock

`func NewCheckAudienceBlock(type_ string, operator string, profile string, audience CheckAudienceBlockAudience, ) *CheckAudienceBlock`

NewCheckAudienceBlock instantiates a new CheckAudienceBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAudienceBlockWithDefaults

`func NewCheckAudienceBlockWithDefaults() *CheckAudienceBlock`

NewCheckAudienceBlockWithDefaults instantiates a new CheckAudienceBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckAudienceBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckAudienceBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckAudienceBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckAudienceBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckAudienceBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckAudienceBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckAudienceBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckAudienceBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckAudienceBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckAudienceBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckAudienceBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *CheckAudienceBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CheckAudienceBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CheckAudienceBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetProfile

`func (o *CheckAudienceBlock) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CheckAudienceBlock) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CheckAudienceBlock) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetAudience

`func (o *CheckAudienceBlock) GetAudience() CheckAudienceBlockAudience`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CheckAudienceBlock) GetAudienceOk() (*CheckAudienceBlockAudience, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CheckAudienceBlock) SetAudience(v CheckAudienceBlockAudience)`

SetAudience sets Audience field to given value.


### GetOnFailure

`func (o *CheckAudienceBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *CheckAudienceBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *CheckAudienceBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *CheckAudienceBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


