# Loyalty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Programs** | Pointer to [**map[string]LoyaltyProgramLedgers**](LoyaltyProgramLedgers.md) | A map holding information about the loyalty programs balance | 

## Methods

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


