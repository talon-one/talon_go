# AwardGiveawayEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolId** | Pointer to **int64** | The ID of the giveaways pool the code was taken from. | 
**PoolName** | Pointer to **string** | The name of the giveaways pool the code was taken from. | 
**RecipientIntegrationId** | Pointer to **string** | The integration ID of the profile that was awarded the giveaway. | 
**GiveawayId** | Pointer to **int64** | The internal ID for the giveaway that was awarded. | 
**Code** | Pointer to **string** | The giveaway code that was awarded. | 

## Methods

### NewAwardGiveawayEffectProps

`func NewAwardGiveawayEffectProps(poolId int64, poolName string, recipientIntegrationId string, giveawayId int64, code string, ) *AwardGiveawayEffectProps`

NewAwardGiveawayEffectProps instantiates a new AwardGiveawayEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwardGiveawayEffectPropsWithDefaults

`func NewAwardGiveawayEffectPropsWithDefaults() *AwardGiveawayEffectProps`

NewAwardGiveawayEffectPropsWithDefaults instantiates a new AwardGiveawayEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolId

`func (o *AwardGiveawayEffectProps) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AwardGiveawayEffectProps) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AwardGiveawayEffectProps) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.


### GetPoolName

`func (o *AwardGiveawayEffectProps) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AwardGiveawayEffectProps) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AwardGiveawayEffectProps) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.


### GetRecipientIntegrationId

`func (o *AwardGiveawayEffectProps) GetRecipientIntegrationId() string`

GetRecipientIntegrationId returns the RecipientIntegrationId field if non-nil, zero value otherwise.

### GetRecipientIntegrationIdOk

`func (o *AwardGiveawayEffectProps) GetRecipientIntegrationIdOk() (*string, bool)`

GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIntegrationId

`func (o *AwardGiveawayEffectProps) SetRecipientIntegrationId(v string)`

SetRecipientIntegrationId sets RecipientIntegrationId field to given value.


### GetGiveawayId

`func (o *AwardGiveawayEffectProps) GetGiveawayId() int64`

GetGiveawayId returns the GiveawayId field if non-nil, zero value otherwise.

### GetGiveawayIdOk

`func (o *AwardGiveawayEffectProps) GetGiveawayIdOk() (*int64, bool)`

GetGiveawayIdOk returns a tuple with the GiveawayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGiveawayId

`func (o *AwardGiveawayEffectProps) SetGiveawayId(v int64)`

SetGiveawayId sets GiveawayId field to given value.


### GetCode

`func (o *AwardGiveawayEffectProps) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AwardGiveawayEffectProps) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AwardGiveawayEffectProps) SetCode(v string)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


