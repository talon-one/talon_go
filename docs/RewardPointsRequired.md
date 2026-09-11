# RewardPointsRequired

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the &#x60;pointsRequired&#x60; entry. When updating a reward, include this property to update an existing entry. Omit it to create a new one.  | [optional] 
**Amount** | Pointer to **float32** | The number of loyalty points required to activate the reward. | 
**LoyaltyProgramId** | Pointer to **int64** | The ID of the associated loyalty program. | 
**SubledgerId** | Pointer to **string** | The ID of the subledger within the loyalty program from which points are deducted when activating the reward.  To specify the main ledger, provide an empty string (\&quot;\&quot;).  | 

## Methods

### NewRewardPointsRequired

`func NewRewardPointsRequired(amount float32, loyaltyProgramId int64, subledgerId string, ) *RewardPointsRequired`

NewRewardPointsRequired instantiates a new RewardPointsRequired object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardPointsRequiredWithDefaults

`func NewRewardPointsRequiredWithDefaults() *RewardPointsRequired`

NewRewardPointsRequiredWithDefaults instantiates a new RewardPointsRequired object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RewardPointsRequired) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RewardPointsRequired) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RewardPointsRequired) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RewardPointsRequired) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *RewardPointsRequired) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RewardPointsRequired) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RewardPointsRequired) SetAmount(v float32)`

SetAmount sets Amount field to given value.


### GetLoyaltyProgramId

`func (o *RewardPointsRequired) GetLoyaltyProgramId() int64`

GetLoyaltyProgramId returns the LoyaltyProgramId field if non-nil, zero value otherwise.

### GetLoyaltyProgramIdOk

`func (o *RewardPointsRequired) GetLoyaltyProgramIdOk() (*int64, bool)`

GetLoyaltyProgramIdOk returns a tuple with the LoyaltyProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyProgramId

`func (o *RewardPointsRequired) SetLoyaltyProgramId(v int64)`

SetLoyaltyProgramId sets LoyaltyProgramId field to given value.


### GetSubledgerId

`func (o *RewardPointsRequired) GetSubledgerId() string`

GetSubledgerId returns the SubledgerId field if non-nil, zero value otherwise.

### GetSubledgerIdOk

`func (o *RewardPointsRequired) GetSubledgerIdOk() (*string, bool)`

GetSubledgerIdOk returns a tuple with the SubledgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubledgerId

`func (o *RewardPointsRequired) SetSubledgerId(v string)`

SetSubledgerId sets SubledgerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


