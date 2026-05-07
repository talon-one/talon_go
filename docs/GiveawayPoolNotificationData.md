# GiveawayPoolNotificationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** | The event type of the notification. | 
**PoolId** | Pointer to **int64** | The ID of the giveaway pool. | 
**PoolName** | Pointer to **string** | The name of the giveaway pool. | 
**PoolDescription** | Pointer to **string** | The description of the giveaway pool. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns the giveaway pool. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application connected to the giveaway pool. | 
**TotalCodes** | Pointer to **int64** | The total number of codes in the giveaway pool. | 
**UsedCodes** | Pointer to **int64** | The number of codes that have been used. | 
**RemainingCodes** | Pointer to **int64** | The number of codes remaining in the giveaway pool. | 
**ThresholdPercent** | Pointer to **int64** | The percentage threshold for the notification. The notification is triggered when the number of codes drops below this threshold. | 

## Methods

### NewGiveawayPoolNotificationData

`func NewGiveawayPoolNotificationData(event string, poolId int64, poolName string, poolDescription string, accountId int64, applicationId int64, totalCodes int64, usedCodes int64, remainingCodes int64, thresholdPercent int64, ) *GiveawayPoolNotificationData`

NewGiveawayPoolNotificationData instantiates a new GiveawayPoolNotificationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGiveawayPoolNotificationDataWithDefaults

`func NewGiveawayPoolNotificationDataWithDefaults() *GiveawayPoolNotificationData`

NewGiveawayPoolNotificationDataWithDefaults instantiates a new GiveawayPoolNotificationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *GiveawayPoolNotificationData) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GiveawayPoolNotificationData) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GiveawayPoolNotificationData) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetPoolId

`func (o *GiveawayPoolNotificationData) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GiveawayPoolNotificationData) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GiveawayPoolNotificationData) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetPoolName

`func (o *GiveawayPoolNotificationData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *GiveawayPoolNotificationData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *GiveawayPoolNotificationData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetPoolDescription

`func (o *GiveawayPoolNotificationData) GetPoolDescription() string`

GetPoolDescription returns the PoolDescription field if non-nil, zero value otherwise.

### GetPoolDescriptionOk

`func (o *GiveawayPoolNotificationData) GetPoolDescriptionOk() (*string, bool)`

GetPoolDescriptionOk returns a tuple with the PoolDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolDescription

`func (o *GiveawayPoolNotificationData) SetPoolDescription(v string)`

SetPoolDescription sets PoolDescription field to given value.


### GetAccountId

`func (o *GiveawayPoolNotificationData) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GiveawayPoolNotificationData) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GiveawayPoolNotificationData) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetApplicationId

`func (o *GiveawayPoolNotificationData) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *GiveawayPoolNotificationData) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *GiveawayPoolNotificationData) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetTotalCodes

`func (o *GiveawayPoolNotificationData) GetTotalCodes() int64`

GetTotalCodes returns the TotalCodes field if non-nil, zero value otherwise.

### GetTotalCodesOk

`func (o *GiveawayPoolNotificationData) GetTotalCodesOk() (*int64, bool)`

GetTotalCodesOk returns a tuple with the TotalCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCodes

`func (o *GiveawayPoolNotificationData) SetTotalCodes(v int64)`

SetTotalCodes sets TotalCodes field to given value.


### GetUsedCodes

`func (o *GiveawayPoolNotificationData) GetUsedCodes() int64`

GetUsedCodes returns the UsedCodes field if non-nil, zero value otherwise.

### GetUsedCodesOk

`func (o *GiveawayPoolNotificationData) GetUsedCodesOk() (*int64, bool)`

GetUsedCodesOk returns a tuple with the UsedCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCodes

`func (o *GiveawayPoolNotificationData) SetUsedCodes(v int64)`

SetUsedCodes sets UsedCodes field to given value.


### GetRemainingCodes

`func (o *GiveawayPoolNotificationData) GetRemainingCodes() int64`

GetRemainingCodes returns the RemainingCodes field if non-nil, zero value otherwise.

### GetRemainingCodesOk

`func (o *GiveawayPoolNotificationData) GetRemainingCodesOk() (*int64, bool)`

GetRemainingCodesOk returns a tuple with the RemainingCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCodes

`func (o *GiveawayPoolNotificationData) SetRemainingCodes(v int64)`

SetRemainingCodes sets RemainingCodes field to given value.


### GetThresholdPercent

`func (o *GiveawayPoolNotificationData) GetThresholdPercent() int64`

GetThresholdPercent returns the ThresholdPercent field if non-nil, zero value otherwise.

### GetThresholdPercentOk

`func (o *GiveawayPoolNotificationData) GetThresholdPercentOk() (*int64, bool)`

GetThresholdPercentOk returns a tuple with the ThresholdPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercent

`func (o *GiveawayPoolNotificationData) SetThresholdPercent(v int64)`

SetThresholdPercent sets ThresholdPercent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


