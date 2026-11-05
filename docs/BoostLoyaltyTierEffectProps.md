# BoostLoyaltyTierEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int64** | The ID of the loyalty program. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program. | 
**TierName** | Pointer to **string** | The name of the tier to which the customer is temporarily boosted. | 
**Reason** | Pointer to **string** | A reason for the tier boost. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | The date when the tier boost expires. | 
**BoostUuid** | Pointer to **string** | The unique identifier of the tier boost. Used to match the boost to its rollback effect when a session is cancelled. | 

## Methods

### NewBoostLoyaltyTierEffectProps

`func NewBoostLoyaltyTierEffectProps(programId int64, subLedgerId string, tierName string, expiryDate time.Time, boostUuid string, ) *BoostLoyaltyTierEffectProps`

NewBoostLoyaltyTierEffectProps instantiates a new BoostLoyaltyTierEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoostLoyaltyTierEffectPropsWithDefaults

`func NewBoostLoyaltyTierEffectPropsWithDefaults() *BoostLoyaltyTierEffectProps`

NewBoostLoyaltyTierEffectPropsWithDefaults instantiates a new BoostLoyaltyTierEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *BoostLoyaltyTierEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *BoostLoyaltyTierEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *BoostLoyaltyTierEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *BoostLoyaltyTierEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *BoostLoyaltyTierEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *BoostLoyaltyTierEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetTierName

`func (o *BoostLoyaltyTierEffectProps) GetTierName() string`

GetTierName returns the TierName field if non-nil, zero value otherwise.

### GetTierNameOk

`func (o *BoostLoyaltyTierEffectProps) GetTierNameOk() (*string, bool)`

GetTierNameOk returns a tuple with the TierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierName

`func (o *BoostLoyaltyTierEffectProps) SetTierName(v string)`

SetTierName sets TierName field to given value.


### GetReason

`func (o *BoostLoyaltyTierEffectProps) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BoostLoyaltyTierEffectProps) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BoostLoyaltyTierEffectProps) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BoostLoyaltyTierEffectProps) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetExpiryDate

`func (o *BoostLoyaltyTierEffectProps) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *BoostLoyaltyTierEffectProps) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *BoostLoyaltyTierEffectProps) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.


### GetBoostUuid

`func (o *BoostLoyaltyTierEffectProps) GetBoostUuid() string`

GetBoostUuid returns the BoostUuid field if non-nil, zero value otherwise.

### GetBoostUuidOk

`func (o *BoostLoyaltyTierEffectProps) GetBoostUuidOk() (*string, bool)`

GetBoostUuidOk returns a tuple with the BoostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostUuid

`func (o *BoostLoyaltyTierEffectProps) SetBoostUuid(v string)`

SetBoostUuid sets BoostUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


