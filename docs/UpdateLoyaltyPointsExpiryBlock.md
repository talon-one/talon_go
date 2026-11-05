# UpdateLoyaltyPointsExpiryBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for this block. | [optional] [readonly] 
**Type** | Pointer to **string** | Identifies the block variant and determines which additional properties are present in it. | 
**Tags** | Pointer to **[]string** | Semantic labels attached to this block. | [optional] [readonly] 
**Operator** | Pointer to **string** | &#x60;setTo&#x60; sets the expiry to an exact date; &#x60;laterBy&#x60; extends the current expiry by a relative duration. | 
**Program** | Pointer to [**UpdateLoyaltyPointsExpiryBlockProgram**](UpdateLoyaltyPointsExpiryBlock_program.md) |  | 
**Recipient** | Pointer to **string** | The customer profile whose points are affected. &#x60;Current&#x60; targets the customer in the current session; &#x60;Advocate&#x60; targets the person who invited their friend via referral program. | 
**Subledger** | Pointer to **string** | The name of the subledger whose points&#39; expiry is changed. Can be empty if this block targets the loyalty program&#39;s main ledger instead of a subledger. | 
**Value** | Pointer to [**map[string]interface{}**](.md) | An absolute expiry date (ISO 8601) when &#x60;operator&#x60; is &#x60;setTo&#x60;, or a relative duration (e.g. &#x60;30D&#x60;) when &#x60;operator&#x60; is &#x60;laterBy&#x60;. | 
**OnFailure** | Pointer to **[]map[string]interface{}** | Blocks evaluated when this block fails or returns false. | [optional] 

## Methods

### NewUpdateLoyaltyPointsExpiryBlock

`func NewUpdateLoyaltyPointsExpiryBlock(type_ string, operator string, program UpdateLoyaltyPointsExpiryBlockProgram, recipient string, subledger string, value map[string]interface{}, ) *UpdateLoyaltyPointsExpiryBlock`

NewUpdateLoyaltyPointsExpiryBlock instantiates a new UpdateLoyaltyPointsExpiryBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoyaltyPointsExpiryBlockWithDefaults

`func NewUpdateLoyaltyPointsExpiryBlockWithDefaults() *UpdateLoyaltyPointsExpiryBlock`

NewUpdateLoyaltyPointsExpiryBlockWithDefaults instantiates a new UpdateLoyaltyPointsExpiryBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateLoyaltyPointsExpiryBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateLoyaltyPointsExpiryBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateLoyaltyPointsExpiryBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateLoyaltyPointsExpiryBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateLoyaltyPointsExpiryBlock) SetType(v string)`

SetType sets Type field to given value.


### GetTags

`func (o *UpdateLoyaltyPointsExpiryBlock) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateLoyaltyPointsExpiryBlock) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateLoyaltyPointsExpiryBlock) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetOperator

`func (o *UpdateLoyaltyPointsExpiryBlock) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *UpdateLoyaltyPointsExpiryBlock) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetProgram

`func (o *UpdateLoyaltyPointsExpiryBlock) GetProgram() UpdateLoyaltyPointsExpiryBlockProgram`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetProgramOk() (*UpdateLoyaltyPointsExpiryBlockProgram, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *UpdateLoyaltyPointsExpiryBlock) SetProgram(v UpdateLoyaltyPointsExpiryBlockProgram)`

SetProgram sets Program field to given value.


### GetRecipient

`func (o *UpdateLoyaltyPointsExpiryBlock) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *UpdateLoyaltyPointsExpiryBlock) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.


### GetSubledger

`func (o *UpdateLoyaltyPointsExpiryBlock) GetSubledger() string`

GetSubledger returns the Subledger field if non-nil, zero value otherwise.

### GetSubledgerOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetSubledgerOk() (*string, bool)`

GetSubledgerOk returns a tuple with the Subledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledger

`func (o *UpdateLoyaltyPointsExpiryBlock) SetSubledger(v string)`

SetSubledger sets Subledger field to given value.


### GetValue

`func (o *UpdateLoyaltyPointsExpiryBlock) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateLoyaltyPointsExpiryBlock) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.


### GetOnFailure

`func (o *UpdateLoyaltyPointsExpiryBlock) GetOnFailure() []map[string]interface{}`

GetOnFailure returns the OnFailure field if non-nil, zero value otherwise.

### GetOnFailureOk

`func (o *UpdateLoyaltyPointsExpiryBlock) GetOnFailureOk() (*[]map[string]interface{}, bool)`

GetOnFailureOk returns a tuple with the OnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnFailure

`func (o *UpdateLoyaltyPointsExpiryBlock) SetOnFailure(v []map[string]interface{})`

SetOnFailure sets OnFailure field to given value.

### HasOnFailure

`func (o *UpdateLoyaltyPointsExpiryBlock) HasOnFailure() bool`

HasOnFailure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


