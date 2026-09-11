# RewardWithUnlocks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The unique ID of the reward. | 
**IntegrationId** | Pointer to **string** | A unique identifier used to reference the reward in API integrations. | 
**Name** | Pointer to **string** | The customer-facing name of the reward. | 
**Description** | Pointer to **string** | Customer-facing description of the reward. | [optional] 
**Rule** | Pointer to [**RuleMetadata**](RuleMetadata.md) |  | 
**Unlocked** | Pointer to [**[]CustomerReward**](CustomerReward.md) | The customer profile&#39;s unlocks of this reward that are not yet &#x60;used&#x60;. | [optional] 

## Methods

### NewRewardWithUnlocks

`func NewRewardWithUnlocks(id int64, integrationId string, name string, rule RuleMetadata, ) *RewardWithUnlocks`

NewRewardWithUnlocks instantiates a new RewardWithUnlocks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardWithUnlocksWithDefaults

`func NewRewardWithUnlocksWithDefaults() *RewardWithUnlocks`

NewRewardWithUnlocksWithDefaults instantiates a new RewardWithUnlocks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RewardWithUnlocks) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardWithUnlocks) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardWithUnlocks) SetId(v int64)`

SetId sets Id field to given value.


### GetIntegrationId

`func (o *RewardWithUnlocks) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *RewardWithUnlocks) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *RewardWithUnlocks) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetName

`func (o *RewardWithUnlocks) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RewardWithUnlocks) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RewardWithUnlocks) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RewardWithUnlocks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RewardWithUnlocks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RewardWithUnlocks) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RewardWithUnlocks) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRule

`func (o *RewardWithUnlocks) GetRule() RuleMetadata`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *RewardWithUnlocks) GetRuleOk() (*RuleMetadata, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *RewardWithUnlocks) SetRule(v RuleMetadata)`

SetRule sets Rule field to given value.


### GetUnlocked

`func (o *RewardWithUnlocks) GetUnlocked() []CustomerReward`

GetUnlocked returns the Unlocked field if non-nil, zero value otherwise.

### GetUnlockedOk

`func (o *RewardWithUnlocks) GetUnlockedOk() (*[]CustomerReward, bool)`

GetUnlockedOk returns a tuple with the Unlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlocked

`func (o *RewardWithUnlocks) SetUnlocked(v []CustomerReward)`

SetUnlocked sets Unlocked field to given value.

### HasUnlocked

`func (o *RewardWithUnlocks) HasUnlocked() bool`

HasUnlocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


