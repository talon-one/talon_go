# CustomerProfileReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the customer reward instance. A customer profile can have multiple instances of the same reward. | 
**IntegrationId** | Pointer to **string** | The integration ID of the customer reward instance. | 
**RewardId** | Pointer to **int64** | The ID of the reward this instance belongs to. | 
**RewardIntegrationId** | Pointer to **string** | The integration ID of the reward this instance belongs to. | 
**RewardName** | Pointer to **string** | The name of the reward. | 
**Description** | Pointer to **string** | The customer-facing description of the reward. | [optional] 
**Rule** | Pointer to [**RuleMetadata**](RuleMetadata.md) |  | [optional] 
**Status** | Pointer to **string** | The status of the customer reward: - &#x60;unlocked&#x60;: The reward is available for use. - &#x60;used&#x60;: The reward has been used.  | 
**UnlockedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the reward was unlocked. | 
**UnlockedByProfileIntegrationId** | Pointer to **string** | The integration ID of the customer profile that unlocked the reward.   For rewards unlocked with a loyalty card, this can be any customer profile  linked to that loyalty card.  | [optional] 
**UsedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the reward was used. | [optional] 
**UsedByProfileIntegrationId** | Pointer to **string** | The integration ID of the customer profile that used the reward.   For rewards unlocked with a loyalty card, this can be any customer profile  linked to that loyalty card.   Only returned when the reward has been used.  | [optional] 
**LoyaltyProgramId** | Pointer to **int64** | The ID of the loyalty program that the loyalty card belongs to. Only returned for rewards unlocked with a loyalty card. | [optional] 
**LoyaltyCardIdentifier** | Pointer to **string** | The identifier of the loyalty card, which must match the regular expression &#x60;^[A-Za-z0-9._%+@-]+$&#x60;.  | [optional] 

## Methods

### NewCustomerProfileReward

`func NewCustomerProfileReward(id int64, integrationId string, rewardId int64, rewardIntegrationId string, rewardName string, status string, unlockedAt time.Time, ) *CustomerProfileReward`

NewCustomerProfileReward instantiates a new CustomerProfileReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerProfileRewardWithDefaults

`func NewCustomerProfileRewardWithDefaults() *CustomerProfileReward`

NewCustomerProfileRewardWithDefaults instantiates a new CustomerProfileReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerProfileReward) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerProfileReward) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerProfileReward) SetId(v int64)`

SetId sets Id field to given value.


### GetIntegrationId

`func (o *CustomerProfileReward) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerProfileReward) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CustomerProfileReward) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetRewardId

`func (o *CustomerProfileReward) GetRewardId() int64`

GetRewardId returns the RewardId field if non-nil, zero value otherwise.

### GetRewardIdOk

`func (o *CustomerProfileReward) GetRewardIdOk() (*int64, bool)`

GetRewardIdOk returns a tuple with the RewardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardId

`func (o *CustomerProfileReward) SetRewardId(v int64)`

SetRewardId sets RewardId field to given value.


### GetRewardIntegrationId

`func (o *CustomerProfileReward) GetRewardIntegrationId() string`

GetRewardIntegrationId returns the RewardIntegrationId field if non-nil, zero value otherwise.

### GetRewardIntegrationIdOk

`func (o *CustomerProfileReward) GetRewardIntegrationIdOk() (*string, bool)`

GetRewardIntegrationIdOk returns a tuple with the RewardIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardIntegrationId

`func (o *CustomerProfileReward) SetRewardIntegrationId(v string)`

SetRewardIntegrationId sets RewardIntegrationId field to given value.


### GetRewardName

`func (o *CustomerProfileReward) GetRewardName() string`

GetRewardName returns the RewardName field if non-nil, zero value otherwise.

### GetRewardNameOk

`func (o *CustomerProfileReward) GetRewardNameOk() (*string, bool)`

GetRewardNameOk returns a tuple with the RewardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardName

`func (o *CustomerProfileReward) SetRewardName(v string)`

SetRewardName sets RewardName field to given value.


### GetDescription

`func (o *CustomerProfileReward) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomerProfileReward) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomerProfileReward) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomerProfileReward) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRule

`func (o *CustomerProfileReward) GetRule() RuleMetadata`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CustomerProfileReward) GetRuleOk() (*RuleMetadata, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CustomerProfileReward) SetRule(v RuleMetadata)`

SetRule sets Rule field to given value.

### HasRule

`func (o *CustomerProfileReward) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerProfileReward) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerProfileReward) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerProfileReward) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUnlockedAt

`func (o *CustomerProfileReward) GetUnlockedAt() time.Time`

GetUnlockedAt returns the UnlockedAt field if non-nil, zero value otherwise.

### GetUnlockedAtOk

`func (o *CustomerProfileReward) GetUnlockedAtOk() (*time.Time, bool)`

GetUnlockedAtOk returns a tuple with the UnlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedAt

`func (o *CustomerProfileReward) SetUnlockedAt(v time.Time)`

SetUnlockedAt sets UnlockedAt field to given value.


### GetUnlockedByProfileIntegrationId

`func (o *CustomerProfileReward) GetUnlockedByProfileIntegrationId() string`

GetUnlockedByProfileIntegrationId returns the UnlockedByProfileIntegrationId field if non-nil, zero value otherwise.

### GetUnlockedByProfileIntegrationIdOk

`func (o *CustomerProfileReward) GetUnlockedByProfileIntegrationIdOk() (*string, bool)`

GetUnlockedByProfileIntegrationIdOk returns a tuple with the UnlockedByProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedByProfileIntegrationId

`func (o *CustomerProfileReward) SetUnlockedByProfileIntegrationId(v string)`

SetUnlockedByProfileIntegrationId sets UnlockedByProfileIntegrationId field to given value.

### HasUnlockedByProfileIntegrationId

`func (o *CustomerProfileReward) HasUnlockedByProfileIntegrationId() bool`

HasUnlockedByProfileIntegrationId returns a boolean if a field has been set.

### GetUsedAt

`func (o *CustomerProfileReward) GetUsedAt() time.Time`

GetUsedAt returns the UsedAt field if non-nil, zero value otherwise.

### GetUsedAtOk

`func (o *CustomerProfileReward) GetUsedAtOk() (*time.Time, bool)`

GetUsedAtOk returns a tuple with the UsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAt

`func (o *CustomerProfileReward) SetUsedAt(v time.Time)`

SetUsedAt sets UsedAt field to given value.

### HasUsedAt

`func (o *CustomerProfileReward) HasUsedAt() bool`

HasUsedAt returns a boolean if a field has been set.

### GetUsedByProfileIntegrationId

`func (o *CustomerProfileReward) GetUsedByProfileIntegrationId() string`

GetUsedByProfileIntegrationId returns the UsedByProfileIntegrationId field if non-nil, zero value otherwise.

### GetUsedByProfileIntegrationIdOk

`func (o *CustomerProfileReward) GetUsedByProfileIntegrationIdOk() (*string, bool)`

GetUsedByProfileIntegrationIdOk returns a tuple with the UsedByProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedByProfileIntegrationId

`func (o *CustomerProfileReward) SetUsedByProfileIntegrationId(v string)`

SetUsedByProfileIntegrationId sets UsedByProfileIntegrationId field to given value.

### HasUsedByProfileIntegrationId

`func (o *CustomerProfileReward) HasUsedByProfileIntegrationId() bool`

HasUsedByProfileIntegrationId returns a boolean if a field has been set.

### GetLoyaltyProgramId

`func (o *CustomerProfileReward) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *CustomerProfileReward) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *CustomerProfileReward) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.

### HasLoyaltyProgramId

`func (o *CustomerProfileReward) HasLoyaltyProgramId() bool`

HasLoyaltyProgramId returns a boolean if a field has been set.

### GetLoyaltyCardIdentifier

`func (o *CustomerProfileReward) GetLoyaltyCardIdentifier() string`

GetLoyaltyCardIdentifier returns the LoyaltyCardIdentifier field if non-nil, zero value otherwise.

### GetLoyaltyCardIdentifierOk

`func (o *CustomerProfileReward) GetLoyaltyCardIdentifierOk() (*string, bool)`

GetLoyaltyCardIdentifierOk returns a tuple with the LoyaltyCardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyCardIdentifier

`func (o *CustomerProfileReward) SetLoyaltyCardIdentifier(v string)`

SetLoyaltyCardIdentifier sets LoyaltyCardIdentifier field to given value.

### HasLoyaltyCardIdentifier

`func (o *CustomerProfileReward) HasLoyaltyCardIdentifier() bool`

HasLoyaltyCardIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


