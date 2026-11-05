# IntegrationUnlockRewardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerProfile** | Pointer to [**CustomerProfile**](CustomerProfile.md) |  | [optional] 
**Loyalty** | Pointer to [**Loyalty**](Loyalty.md) |  | [optional] 
**Effects** | Pointer to **[]map[string]interface{}** | The effects generated when evaluating this reward unlock, after the reward&#39;s eligibility conditions are met. See [API effects](https://docs.talon.one/docs/dev/integration-api/api-effects). | 

## Methods

### NewIntegrationUnlockRewardResponse

`func NewIntegrationUnlockRewardResponse(effects []map[string]interface{}, ) *IntegrationUnlockRewardResponse`

NewIntegrationUnlockRewardResponse instantiates a new IntegrationUnlockRewardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationUnlockRewardResponseWithDefaults

`func NewIntegrationUnlockRewardResponseWithDefaults() *IntegrationUnlockRewardResponse`

NewIntegrationUnlockRewardResponseWithDefaults instantiates a new IntegrationUnlockRewardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerProfile

`func (o *IntegrationUnlockRewardResponse) GetCustomerProfile() CustomerProfile`

GetCustomerProfile returns the CustomerProfile field if non-nil, zero value otherwise.

### GetCustomerProfileOk

`func (o *IntegrationUnlockRewardResponse) GetCustomerProfileOk() (*CustomerProfile, bool)`

GetCustomerProfileOk returns a tuple with the CustomerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerProfile

`func (o *IntegrationUnlockRewardResponse) SetCustomerProfile(v CustomerProfile)`

SetCustomerProfile sets CustomerProfile field to given value.

### HasCustomerProfile

`func (o *IntegrationUnlockRewardResponse) HasCustomerProfile() bool`

HasCustomerProfile returns a boolean if a field has been set.

### GetLoyalty

`func (o *IntegrationUnlockRewardResponse) GetLoyalty() Loyalty`

GetLoyalty returns the Loyalty field if non-nil, zero value otherwise.

### GetLoyaltyOk

`func (o *IntegrationUnlockRewardResponse) GetLoyaltyOk() (*Loyalty, bool)`

GetLoyaltyOk returns a tuple with the Loyalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyalty

`func (o *IntegrationUnlockRewardResponse) SetLoyalty(v Loyalty)`

SetLoyalty sets Loyalty field to given value.

### HasLoyalty

`func (o *IntegrationUnlockRewardResponse) HasLoyalty() bool`

HasLoyalty returns a boolean if a field has been set.

### GetEffects

`func (o *IntegrationUnlockRewardResponse) GetEffects() []map[string]interface{}`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *IntegrationUnlockRewardResponse) GetEffectsOk() (*[]map[string]interface{}, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *IntegrationUnlockRewardResponse) SetEffects(v []map[string]interface{})`

SetEffects sets Effects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


