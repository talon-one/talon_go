# NewLoyaltyProgramAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The internal name for the Loyalty Program. This is an immutable value. | 
**Tiers** | Pointer to [**[]NewLoyaltyTier**](NewLoyaltyTier.md) | The tiers in this loyalty program. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**CardBased** | Pointer to **bool** | Defines the type of loyalty program: - &#x60;true&#x60;: the program is a card-based. - &#x60;false&#x60;: the program is profile-based.  | [default to false]

## Methods

### GetName

`func (o *NewLoyaltyProgramAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewLoyaltyProgramAllOf) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewLoyaltyProgramAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewLoyaltyProgramAllOf) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTiers

`func (o *NewLoyaltyProgramAllOf) GetTiers() []NewLoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *NewLoyaltyProgramAllOf) GetTiersOk() ([]NewLoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiers

`func (o *NewLoyaltyProgramAllOf) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### SetTiers

`func (o *NewLoyaltyProgramAllOf) SetTiers(v []NewLoyaltyTier)`

SetTiers gets a reference to the given []NewLoyaltyTier and assigns it to the Tiers field.

### GetTimezone

`func (o *NewLoyaltyProgramAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NewLoyaltyProgramAllOf) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *NewLoyaltyProgramAllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *NewLoyaltyProgramAllOf) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetCardBased

`func (o *NewLoyaltyProgramAllOf) GetCardBased() bool`

GetCardBased returns the CardBased field if non-nil, zero value otherwise.

### GetCardBasedOk

`func (o *NewLoyaltyProgramAllOf) GetCardBasedOk() (bool, bool)`

GetCardBasedOk returns a tuple with the CardBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardBased

`func (o *NewLoyaltyProgramAllOf) HasCardBased() bool`

HasCardBased returns a boolean if a field has been set.

### SetCardBased

`func (o *NewLoyaltyProgramAllOf) SetCardBased(v bool)`

SetCardBased gets a reference to the given bool and assigns it to the CardBased field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


