# CheckReferralBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Redeem** | Pointer to **bool** | When &#x60;true&#x60;, the referral code is redeemed. | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Promotion blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewCheckReferralBlock

`func NewCheckReferralBlock(type_ string, redeem bool, ) *CheckReferralBlock`

NewCheckReferralBlock instantiates a new CheckReferralBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckReferralBlockWithDefaults

`func NewCheckReferralBlockWithDefaults() *CheckReferralBlock`

NewCheckReferralBlockWithDefaults instantiates a new CheckReferralBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckReferralBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckReferralBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckReferralBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CheckReferralBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CheckReferralBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CheckReferralBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CheckReferralBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *CheckReferralBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CheckReferralBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CheckReferralBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CheckReferralBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetRedeem

`func (o *CheckReferralBlock) GetRedeem() bool`

GetRedeem returns the Redeem field if non-nil, zero value otherwise.

### GetRedeemOk

`func (o *CheckReferralBlock) GetRedeemOk() (*bool, bool)`

GetRedeemOk returns a tuple with the Redeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeem

`func (o *CheckReferralBlock) SetRedeem(v bool)`

SetRedeem sets Redeem field to given value.


### GetOnFailure

`func (o *CheckReferralBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *CheckReferralBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *CheckReferralBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *CheckReferralBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


