# Loyalty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cards** | Pointer to [**[]LoyaltyCard**](LoyaltyCard.md) | Displays information about the balances of the loyalty cards. | [optional] 
**Programs** | Pointer to [**map[string]LoyaltyProgramLedgers**](LoyaltyProgramLedgers.md) | Displays information about point balances in profile-based programs. | 

## Methods

### NewLoyalty

`func NewLoyalty(programs map[string]LoyaltyProgramLedgers, ) *Loyalty`

NewLoyalty instantiates a new Loyalty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoyaltyWithDefaults

`func NewLoyaltyWithDefaults() *Loyalty`

NewLoyaltyWithDefaults instantiates a new Loyalty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCards

`func (o *Loyalty) GetCards() []LoyaltyCard`

GetCards returns the Cards field if non-nil, zero value otherwise.

### GetCardsOk

`func (o *Loyalty) GetCardsOk() (*[]LoyaltyCard, bool)`

GetCardsOk returns a tuple with the Cards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCards

`func (o *Loyalty) SetCards(v []LoyaltyCard)`

SetCards sets Cards field to given value.

### HasCards

`func (o *Loyalty) HasCards() bool`

HasCards returns a boolean if a field has been set.

### GetPrograms

`func (o *Loyalty) GetPrograms() map[string]LoyaltyProgramLedgers`

GetPrograms returns the Programs field if non-nil, zero value otherwise.

### GetProgramsOk

`func (o *Loyalty) GetProgramsOk() (*map[string]LoyaltyProgramLedgers, bool)`

GetProgramsOk returns a tuple with the Programs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrograms

`func (o *Loyalty) SetPrograms(v map[string]LoyaltyProgramLedgers)`

SetPrograms sets Programs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


