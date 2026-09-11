# IntegrationUnlockRewardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID to assign to the created customer reward unlock. | 
**ProfileIntegrationId** | Pointer to **string** | The integration ID of the customer profile unlocking the reward. | 
**CardIdentifier** | Pointer to **string** | The identifier of the loyalty card, which must match the regular expression &#x60;^[A-Za-z0-9._%+@-]+$&#x60;.  | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | The ID of the loyalty program from which points will be deducted. Required when the reward has &#x60;pointsRequired&#x60; configured. | [optional] 
**SubledgerId** | Pointer to **string** | The ID of the subledger from which points will be deducted. Required when the reward has &#x60;pointsRequired&#x60; configured.  To specify the main ledger, provide an empty string (\&quot;\&quot;).  | [optional] 
**ResponseContent** | Pointer to **[]string** | Determines which data is included in the response. Add any of the following optional values to the array to get that data in the response: &#x60;customerProfile&#x60;, &#x60;ruleFailureReasons&#x60;, &#x60;loyalty&#x60;. &#x60;effects&#x60; is always returned regardless of whether it is included here. | [optional] 

## Methods

### NewIntegrationUnlockRewardRequest

`func NewIntegrationUnlockRewardRequest(integrationId string, profileIntegrationId string, ) *IntegrationUnlockRewardRequest`

NewIntegrationUnlockRewardRequest instantiates a new IntegrationUnlockRewardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationUnlockRewardRequestWithDefaults

`func NewIntegrationUnlockRewardRequestWithDefaults() *IntegrationUnlockRewardRequest`

NewIntegrationUnlockRewardRequestWithDefaults instantiates a new IntegrationUnlockRewardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *IntegrationUnlockRewardRequest) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *IntegrationUnlockRewardRequest) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *IntegrationUnlockRewardRequest) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetProfileIntegrationId

`func (o *IntegrationUnlockRewardRequest) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *IntegrationUnlockRewardRequest) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *IntegrationUnlockRewardRequest) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.


### GetCardIdentifier

`func (o *IntegrationUnlockRewardRequest) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *IntegrationUnlockRewardRequest) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *IntegrationUnlockRewardRequest) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *IntegrationUnlockRewardRequest) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *IntegrationUnlockRewardRequest) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *IntegrationUnlockRewardRequest) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *IntegrationUnlockRewardRequest) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *IntegrationUnlockRewardRequest) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetSubledgerId

`func (o *IntegrationUnlockRewardRequest) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *IntegrationUnlockRewardRequest) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *IntegrationUnlockRewardRequest) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.

### HasSubledgerId

`func (o *IntegrationUnlockRewardRequest) HasSubledgerId() bool`

HasSubledgerId returns a boolean if a field has been set.

### GetResponseContent

`func (o *IntegrationUnlockRewardRequest) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *IntegrationUnlockRewardRequest) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *IntegrationUnlockRewardRequest) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *IntegrationUnlockRewardRequest) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


