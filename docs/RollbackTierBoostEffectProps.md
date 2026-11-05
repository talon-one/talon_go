# RollbackTierBoostEffectProps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProgramId** | Pointer to **int64** | The ID of the loyalty program. | 
**SubLedgerId** | Pointer to **string** | The ID of the subledger within the loyalty program. | 
**TierName** | Pointer to **string** | The name of the boosted tier that was rolled back. | 
**BoostUuid** | Pointer to **string** | The unique identifier of the tier boost that was rolled back. Matches the &#x60;boostUuid&#x60; of the original &#x60;boostLoyaltyTier&#x60; effect. | 

## Methods

### NewRollbackTierBoostEffectProps

`func NewRollbackTierBoostEffectProps(programId int64, subLedgerId string, tierName string, boostUuid string, ) *RollbackTierBoostEffectProps`

NewRollbackTierBoostEffectProps instantiates a new RollbackTierBoostEffectProps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollbackTierBoostEffectPropsWithDefaults

`func NewRollbackTierBoostEffectPropsWithDefaults() *RollbackTierBoostEffectProps`

NewRollbackTierBoostEffectPropsWithDefaults instantiates a new RollbackTierBoostEffectProps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgramId

`func (o *RollbackTierBoostEffectProps) GetProgramId() int64`

GetProgramId returns the ProgramId field if non-nil, zero value otherwise.

### GetProgramIdOk

`func (o *RollbackTierBoostEffectProps) GetProgramIdOk() (*int64, bool)`

GetProgramIdOk returns a tuple with the ProgramId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgramId

`func (o *RollbackTierBoostEffectProps) SetProgramId(v int64)`

SetProgramId sets ProgramId field to given value.


### GetSubLedgerId

`func (o *RollbackTierBoostEffectProps) GetSubLedgerId() string`

GetSubLedgerId returns the SubLedgerId field if non-nil, zero value otherwise.

### GetSubLedgerIdOk

`func (o *RollbackTierBoostEffectProps) GetSubLedgerIdOk() (*string, bool)`

GetSubLedgerIdOk returns a tuple with the SubLedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubLedgerId

`func (o *RollbackTierBoostEffectProps) SetSubLedgerId(v string)`

SetSubLedgerId sets SubLedgerId field to given value.


### GetTierName

`func (o *RollbackTierBoostEffectProps) GetTierName() string`

GetTierName returns the TierName field if non-nil, zero value otherwise.

### GetTierNameOk

`func (o *RollbackTierBoostEffectProps) GetTierNameOk() (*string, bool)`

GetTierNameOk returns a tuple with the TierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierName

`func (o *RollbackTierBoostEffectProps) SetTierName(v string)`

SetTierName sets TierName field to given value.


### GetBoostUuid

`func (o *RollbackTierBoostEffectProps) GetBoostUuid() string`

GetBoostUuid returns the BoostUuid field if non-nil, zero value otherwise.

### GetBoostUuidOk

`func (o *RollbackTierBoostEffectProps) GetBoostUuidOk() (*string, bool)`

GetBoostUuidOk returns a tuple with the BoostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostUuid

`func (o *RollbackTierBoostEffectProps) SetBoostUuid(v string)`

SetBoostUuid sets BoostUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


