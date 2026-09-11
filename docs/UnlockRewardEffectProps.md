# UnlockRewardEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID assigned to the customer reward unlock. | 
**RewardId** | Pointer to **int64** | The internal ID of the reward that was unlocked. | 
**ApplicationId** | Pointer to **int64** | The internal ID of the application the reward belongs to. | 
**ProfileIntegrationId** | Pointer to **string** | The integration ID of the customer profile that unlocked the reward. | 
**UnlockedAt** | Pointer to [**time.Time**](time.Time.md) | The time the reward was unlocked. | 
**CardIdentifier** | Pointer to **string** | The identifier of the loyalty card, which must match the regular expression &#x60;^[A-Za-z0-9._%+@-]+$&#x60;.  | [optional] 

## Methods

### NewUnlockRewardEffectProps

`func NewUnlockRewardEffectProps(integrationId string, rewardId int64, applicationId int64, profileIntegrationId string, unlockedAt time.Time, ) *UnlockRewardEffectProps`

NewUnlockRewardEffectProps instantiates a new UnlockRewardEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlockRewardEffectPropsWithDefaults

`func NewUnlockRewardEffectPropsWithDefaults() *UnlockRewardEffectProps`

NewUnlockRewardEffectPropsWithDefaults instantiates a new UnlockRewardEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *UnlockRewardEffectProps) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *UnlockRewardEffectProps) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *UnlockRewardEffectProps) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetRewardId

`func (o *UnlockRewardEffectProps) GetRewardId() int64`

GetRewardId returns the RewardId field if non-nil, zero value otherwise.

### GetRewardIdOk

`func (o *UnlockRewardEffectProps) GetRewardIdOk() (*int64, bool)`

GetRewardIdOk returns a tuple with the RewardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardId

`func (o *UnlockRewardEffectProps) SetRewardId(v int64)`

SetRewardId sets RewardId field to given value.


### GetApplicationId

`func (o *UnlockRewardEffectProps) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *UnlockRewardEffectProps) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *UnlockRewardEffectProps) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileIntegrationId

`func (o *UnlockRewardEffectProps) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *UnlockRewardEffectProps) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *UnlockRewardEffectProps) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.


### GetUnlockedAt

`func (o *UnlockRewardEffectProps) GetUnlockedAt() time.Time`

GetUnlockedAt returns the UnlockedAt field if non-nil, zero value otherwise.

### GetUnlockedAtOk

`func (o *UnlockRewardEffectProps) GetUnlockedAtOk() (*time.Time, bool)`

GetUnlockedAtOk returns a tuple with the UnlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedAt

`func (o *UnlockRewardEffectProps) SetUnlockedAt(v time.Time)`

SetUnlockedAt sets UnlockedAt field to given value.


### GetCardIdentifier

`func (o *UnlockRewardEffectProps) GetCardIdentifier() string`

GetCardIdentifier returns the CardIdentifier field if non-nil, zero value otherwise.

### GetCardIdentifierOk

`func (o *UnlockRewardEffectProps) GetCardIdentifierOk() (*string, bool)`

GetCardIdentifierOk returns a tuple with the CardIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardIdentifier

`func (o *UnlockRewardEffectProps) SetCardIdentifier(v string)`

SetCardIdentifier sets CardIdentifier field to given value.

### HasCardIdentifier

`func (o *UnlockRewardEffectProps) HasCardIdentifier() bool`

HasCardIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


