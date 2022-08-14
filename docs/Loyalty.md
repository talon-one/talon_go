# Loyalty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**[]LoyaltyCard**](LoyaltyCard.md) | Displays information about the balances of the loyalty cards. | [optional] 
**Programs** | Pointer to [**map[string]LoyaltyProgramLedgers**](LoyaltyProgramLedgers.md) | Displays information about point balances in profile-based programs. | 

## Methods

### GetCards

`func (o *Loyalty) GetCards() []LoyaltyCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *Loyalty) GetCardsOk() ([]LoyaltyCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCards

`func (o *Loyalty) HasCards() bool`

HasCards returns a boolean if a field has been set.

### SetCards

`func (o *Loyalty) SetCards(v []LoyaltyCard)`

SetCards gets a reference to the given []LoyaltyCard and assigns it to the Cards field.

### GetPrograms

`func (o *Loyalty) GetPrograms() map[string]LoyaltyProgramLedgers`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *Loyalty) GetProgramsOk() (map[string]LoyaltyProgramLedgers, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPrograms

`func (o *Loyalty) HasPrograms() bool`

HasPrograms returns a boolean if a field has been set.

### SetPrograms

`func (o *Loyalty) SetPrograms(v map[string]LoyaltyProgramLedgers)`

SetPrograms gets a reference to the given map[string]LoyaltyProgramLedgers and assigns it to the Programs field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


