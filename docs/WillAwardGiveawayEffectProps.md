# WillAwardGiveawayEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | Pointer to **int64** | The ID of the giveaways pool the code will be taken from. | 
**PoolName** | Pointer to **string** | The name of the giveaways pool the code will be taken from. | 
**RecipientIntegrationId** | Pointer to **string** | The integration ID of the profile that will be awarded the giveaway. | 

## Methods

### NewWillAwardGiveawayEffectProps

`func NewWillAwardGiveawayEffectProps(poolId int64, poolName string, recipientIntegrationId string, ) *WillAwardGiveawayEffectProps`

NewWillAwardGiveawayEffectProps instantiates a new WillAwardGiveawayEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWillAwardGiveawayEffectPropsWithDefaults

`func NewWillAwardGiveawayEffectPropsWithDefaults() *WillAwardGiveawayEffectProps`

NewWillAwardGiveawayEffectPropsWithDefaults instantiates a new WillAwardGiveawayEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *WillAwardGiveawayEffectProps) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *WillAwardGiveawayEffectProps) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *WillAwardGiveawayEffectProps) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetPoolName

`func (o *WillAwardGiveawayEffectProps) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *WillAwardGiveawayEffectProps) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *WillAwardGiveawayEffectProps) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetRecipientIntegrationId

`func (o *WillAwardGiveawayEffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *WillAwardGiveawayEffectProps) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *WillAwardGiveawayEffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


