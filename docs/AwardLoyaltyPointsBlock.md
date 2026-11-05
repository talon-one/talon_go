# AwardLoyaltyPointsBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Name** | Pointer to **string** | The human-readable label attached to the awarded points. | 
**Program** | Pointer to [**AwardLoyaltyPointsBlockProgram**](AwardLoyaltyPointsBlock_program.md) |  | 
**Recipient** | Pointer to **string** | The customer profile that receives the points. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Subledger** | Pointer to **string** | The name of the subledger to add points to. Can be empty if this block adds points to the loyalty program&#39;s main ledger instead of a subledger. | 
**Target** | Pointer to [**map[string]interface{}**](.md) | Identifies the scope over which loyalty points are awarded. The &#x60;type&#x60; field selects the target variant. | 
**Value** | Pointer to [**map[string]interface{}**](.md) | Number of points to award. Either a numeric scalar or a &#x60;{{expression}}&#x60; string that resolves to a number at evaluation time. | 
**Partial** | Pointer to **bool** | When &#x60;true&#x60;, applies a partial points reward when the requested value exceeds the configured budget. | [optional] 
**AwaitsActivation** | Pointer to **bool** | When &#x60;true&#x60;, the awarded points require manual or delayed activation before becoming active. Mutually exclusive with &#x60;startDate&#x60;. | [optional] 
**StartDate** | Pointer to [**map[string]interface{}**](.md) | Timestamp at which the awarded points become active. Mutually exclusive with &#x60;awaitsActivation&#x60;. | [optional] 
**ValidityDuration** | Pointer to **string** | Relative duration (e.g. &#x60;30D&#x60;) after which the awarded points expire. Mutually exclusive with &#x60;expiryDate&#x60;. | [optional] 
**ExpiryDate** | Pointer to [**map[string]interface{}**](.md) | Timestamp at which the awarded points expire. Mutually exclusive with &#x60;validityDuration&#x60;. | [optional] 
**PendingDuration** | Pointer to **string** | Relative duration (e.g. &#x60;3D&#x60;) the awarded points remain pending before activation. | [optional] 

## Methods

### NewAwardLoyaltyPointsBlock

`func NewAwardLoyaltyPointsBlock(type_ string, name string, program AwardLoyaltyPointsBlockProgram, recipient string, subledger string, target map[string]interface{}, value map[string]interface{}, ) *AwardLoyaltyPointsBlock`

NewAwardLoyaltyPointsBlock instantiates a new AwardLoyaltyPointsBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardLoyaltyPointsBlockWithDefaults

`func NewAwardLoyaltyPointsBlockWithDefaults() *AwardLoyaltyPointsBlock`

NewAwardLoyaltyPointsBlockWithDefaults instantiates a new AwardLoyaltyPointsBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AwardLoyaltyPointsBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AwardLoyaltyPointsBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AwardLoyaltyPointsBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AwardLoyaltyPointsBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AwardLoyaltyPointsBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AwardLoyaltyPointsBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AwardLoyaltyPointsBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *AwardLoyaltyPointsBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AwardLoyaltyPointsBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AwardLoyaltyPointsBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AwardLoyaltyPointsBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *AwardLoyaltyPointsBlock) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwardLoyaltyPointsBlock) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwardLoyaltyPointsBlock) SetName(v string)`

SetName sets Name field to given value.


### GetProgram

`func (o *AwardLoyaltyPointsBlock) GetProgram() AwardLoyaltyPointsBlockProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *AwardLoyaltyPointsBlock) GetProgramOk() (*AwardLoyaltyPointsBlockProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *AwardLoyaltyPointsBlock) SetProgram(v AwardLoyaltyPointsBlockProgram)`

SetProgram sets Program field to given value.


### GetRecipient

`func (o *AwardLoyaltyPointsBlock) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AwardLoyaltyPointsBlock) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AwardLoyaltyPointsBlock) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSubledger

`func (o *AwardLoyaltyPointsBlock) GetSubledger() string`

GetSubledger returns the Subledger field if non-nil, zero value otherwise.

### GetSubledgerOk

`func (o *AwardLoyaltyPointsBlock) GetSubledgerOk() (*string, bool)`

GetSubledgerOk returns a tuple with the Subledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledger

`func (o *AwardLoyaltyPointsBlock) SetSubledger(v string)`

SetSubledger sets Subledger field to given value.


### GetTarget

`func (o *AwardLoyaltyPointsBlock) GetTarget() map[string]interface{}`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AwardLoyaltyPointsBlock) GetTargetOk() (*map[string]interface{}, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AwardLoyaltyPointsBlock) SetTarget(v map[string]interface{})`

SetTarget sets Target field to given value.


### GetValue

`func (o *AwardLoyaltyPointsBlock) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AwardLoyaltyPointsBlock) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AwardLoyaltyPointsBlock) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetPartial

`func (o *AwardLoyaltyPointsBlock) GetPartial() bool`

GetPartial returns the Partial field if non-nil, zero value otherwise.

### GetPartialOk

`func (o *AwardLoyaltyPointsBlock) GetPartialOk() (*bool, bool)`

GetPartialOk returns a tuple with the Partial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartial

`func (o *AwardLoyaltyPointsBlock) SetPartial(v bool)`

SetPartial sets Partial field to given value.

### HasPartial

`func (o *AwardLoyaltyPointsBlock) HasPartial() bool`

HasPartial returns a boolean if a field has been set.

### GetAwaitsActivation

`func (o *AwardLoyaltyPointsBlock) GetAwaitsActivation() bool`

GetAwaitsActivation returns the AwaitsActivation field if non-nil, zero value otherwise.

### GetAwaitsActivationOk

`func (o *AwardLoyaltyPointsBlock) GetAwaitsActivationOk() (*bool, bool)`

GetAwaitsActivationOk returns a tuple with the AwaitsActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwaitsActivation

`func (o *AwardLoyaltyPointsBlock) SetAwaitsActivation(v bool)`

SetAwaitsActivation sets AwaitsActivation field to given value.

### HasAwaitsActivation

`func (o *AwardLoyaltyPointsBlock) HasAwaitsActivation() bool`

HasAwaitsActivation returns a boolean if a field has been set.

### GetStartDate

`func (o *AwardLoyaltyPointsBlock) GetStartDate() map[string]interface{}`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *AwardLoyaltyPointsBlock) GetStartDateOk() (*map[string]interface{}, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *AwardLoyaltyPointsBlock) SetStartDate(v map[string]interface{})`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *AwardLoyaltyPointsBlock) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetValidityDuration

`func (o *AwardLoyaltyPointsBlock) GetValidityDuration() string`

GetValidityDuration returns the ValidityDuration field if non-nil, zero value otherwise.

### GetValidityDurationOk

`func (o *AwardLoyaltyPointsBlock) GetValidityDurationOk() (*string, bool)`

GetValidityDurationOk returns a tuple with the ValidityDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityDuration

`func (o *AwardLoyaltyPointsBlock) SetValidityDuration(v string)`

SetValidityDuration sets ValidityDuration field to given value.

### HasValidityDuration

`func (o *AwardLoyaltyPointsBlock) HasValidityDuration() bool`

HasValidityDuration returns a boolean if a field has been set.

### GetExpiryDate

`func (o *AwardLoyaltyPointsBlock) GetExpiryDate() map[string]interface{}`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *AwardLoyaltyPointsBlock) GetExpiryDateOk() (*map[string]interface{}, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *AwardLoyaltyPointsBlock) SetExpiryDate(v map[string]interface{})`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *AwardLoyaltyPointsBlock) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetPendingDuration

`func (o *AwardLoyaltyPointsBlock) GetPendingDuration() string`

GetPendingDuration returns the PendingDuration field if non-nil, zero value otherwise.

### GetPendingDurationOk

`func (o *AwardLoyaltyPointsBlock) GetPendingDurationOk() (*string, bool)`

GetPendingDurationOk returns a tuple with the PendingDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDuration

`func (o *AwardLoyaltyPointsBlock) SetPendingDuration(v string)`

SetPendingDuration sets PendingDuration field to given value.

### HasPendingDuration

`func (o *AwardLoyaltyPointsBlock) HasPendingDuration() bool`

HasPendingDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


