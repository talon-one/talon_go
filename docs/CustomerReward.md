# CustomerReward

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int64** | The ID of the Application in which the reward was unlocked. | 
**ProfileIntegrationId** | Pointer to **string** | The integration ID of the customer profile that unlocked this reward. | 
**IntegrationId** | Pointer to **string** | The integration ID assigned to this reward unlock. | 
**UnlockedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the reward was unlocked. | 
**UsedAt** | Pointer to [**time.Time**](time.Time.md) | The date and time when the reward was used. | [optional] 

## Methods

### NewCustomerReward

`func NewCustomerReward(applicationId int64, profileIntegrationId string, integrationId string, unlockedAt time.Time, ) *CustomerReward`

NewCustomerReward instantiates a new CustomerReward object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRewardWithDefaults

`func NewCustomerRewardWithDefaults() *CustomerReward`

NewCustomerRewardWithDefaults instantiates a new CustomerReward object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CustomerReward) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CustomerReward) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CustomerReward) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetProfileIntegrationId

`func (o *CustomerReward) GetProfileIntegrationId() string`

GetProfileIntegrationId returns the ProfileIntegrationId field if non-nil, zero value otherwise.

### GetProfileIntegrationIdOk

`func (o *CustomerReward) GetProfileIntegrationIdOk() (*string, bool)`

GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIntegrationId

`func (o *CustomerReward) SetProfileIntegrationId(v string)`

SetProfileIntegrationId sets ProfileIntegrationId field to given value.


### GetIntegrationId

`func (o *CustomerReward) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *CustomerReward) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *CustomerReward) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.


### GetUnlockedAt

`func (o *CustomerReward) GetUnlockedAt() time.Time`

GetUnlockedAt returns the UnlockedAt field if non-nil, zero value otherwise.

### GetUnlockedAtOk

`func (o *CustomerReward) GetUnlockedAtOk() (*time.Time, bool)`

GetUnlockedAtOk returns a tuple with the UnlockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlockedAt

`func (o *CustomerReward) SetUnlockedAt(v time.Time)`

SetUnlockedAt sets UnlockedAt field to given value.


### GetUsedAt

`func (o *CustomerReward) GetUsedAt() time.Time`

GetUsedAt returns the UsedAt field if non-nil, zero value otherwise.

### GetUsedAtOk

`func (o *CustomerReward) GetUsedAtOk() (*time.Time, bool)`

GetUsedAtOk returns a tuple with the UsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedAt

`func (o *CustomerReward) SetUsedAt(v time.Time)`

SetUsedAt sets UsedAt field to given value.

### HasUsedAt

`func (o *CustomerReward) HasUsedAt() bool`

HasUsedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


