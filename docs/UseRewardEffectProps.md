# UseRewardEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationId** | Pointer to **string** | The integration ID of the customer reward that was used. | 
**RewardId** | Pointer to **int64** | The ID of the reward that was used. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application the reward belongs to. | 

## Methods

### NewUseRewardEffectProps

`func NewUseRewardEffectProps(integrationId string, rewardId int64, applicationId int64, ) *UseRewardEffectProps`

NewUseRewardEffectProps instantiates a new UseRewardEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUseRewardEffectPropsWithDefaults

`func NewUseRewardEffectPropsWithDefaults() *UseRewardEffectProps`

NewUseRewardEffectPropsWithDefaults instantiates a new UseRewardEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationId

`func (o *UseRewardEffectProps) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *UseRewardEffectProps) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *UseRewardEffectProps) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetRewardId

`func (o *UseRewardEffectProps) GetRewardId() int64`

GetRewardId returns the RewardId field if non-nil, zero value otherwise.

### GetRewardIdOk

`func (o *UseRewardEffectProps) GetRewardIdOk() (*int64, bool)`

GetRewardIdOk returns a tuple with the RewardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRewardId

`func (o *UseRewardEffectProps) SetRewardId(v int64)`

SetRewardId sets RewardId field to given value.


### GetApplicationId

`func (o *UseRewardEffectProps) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *UseRewardEffectProps) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *UseRewardEffectProps) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


