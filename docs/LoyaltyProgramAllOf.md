# LoyaltyProgramAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of loyalty program. | 
**AccountID** | Pointer to **int32** | The ID of the Talon.One account that owns this program. | 
**Name** | Pointer to **string** | The internal name for the Loyalty Program. This is an immutable value. | 
**Tiers** | Pointer to [**[]LoyaltyTier**](LoyaltyTier.md) | The tiers in this loyalty program. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**CardBased** | Pointer to **bool** | Defines the type of loyalty program: - &#x60;true&#x60;: the program is a card-based. - &#x60;false&#x60;: the program is profile-based.  | [default to false]
**CanUpdateTiers** | Pointer to **bool** | &#x60;True&#x60; if the tier definitions can be updated.  | [optional] [default to false]
**CanUpdateJoinPolicy** | Pointer to **bool** | &#x60;True&#x60; if the program join policy can be updated.  | [optional] 
**CanUpdateTierExpirationPolicy** | Pointer to **bool** | &#x60;True&#x60; if the tier expiration policy can be updated.  | [optional] 
**CanUpgradeToAdvancedTiers** | Pointer to **bool** | &#x60;True&#x60; if the program can be upgraded to use the &#x60;tiersExpireIn&#x60; and &#x60;tiersDowngradePolicy&#x60; properties.  | [optional] [default to false]

## Methods

### GetId

`func (o *LoyaltyProgramAllOf) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgramAllOf) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LoyaltyProgramAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LoyaltyProgramAllOf) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetAccountID

`func (o *LoyaltyProgramAllOf) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *LoyaltyProgramAllOf) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *LoyaltyProgramAllOf) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *LoyaltyProgramAllOf) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetName

`func (o *LoyaltyProgramAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgramAllOf) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyProgramAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyProgramAllOf) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTiers

`func (o *LoyaltyProgramAllOf) GetTiers() []LoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *LoyaltyProgramAllOf) GetTiersOk() ([]LoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiers

`func (o *LoyaltyProgramAllOf) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### SetTiers

`func (o *LoyaltyProgramAllOf) SetTiers(v []LoyaltyTier)`

SetTiers gets a reference to the given []LoyaltyTier and assigns it to the Tiers field.

### GetTimezone

`func (o *LoyaltyProgramAllOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *LoyaltyProgramAllOf) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *LoyaltyProgramAllOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *LoyaltyProgramAllOf) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetCardBased

`func (o *LoyaltyProgramAllOf) GetCardBased() bool`

GetCardBased returns the CardBased field if non-nil, zero value otherwise.

### GetCardBasedOk

`func (o *LoyaltyProgramAllOf) GetCardBasedOk() (bool, bool)`

GetCardBasedOk returns a tuple with the CardBased field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCardBased

`func (o *LoyaltyProgramAllOf) HasCardBased() bool`

HasCardBased returns a boolean if a field has been set.

### SetCardBased

`func (o *LoyaltyProgramAllOf) SetCardBased(v bool)`

SetCardBased gets a reference to the given bool and assigns it to the CardBased field.

### GetCanUpdateTiers

`func (o *LoyaltyProgramAllOf) GetCanUpdateTiers() bool`

GetCanUpdateTiers returns the CanUpdateTiers field if non-nil, zero value otherwise.

### GetCanUpdateTiersOk

`func (o *LoyaltyProgramAllOf) GetCanUpdateTiersOk() (bool, bool)`

GetCanUpdateTiersOk returns a tuple with the CanUpdateTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpdateTiers

`func (o *LoyaltyProgramAllOf) HasCanUpdateTiers() bool`

HasCanUpdateTiers returns a boolean if a field has been set.

### SetCanUpdateTiers

`func (o *LoyaltyProgramAllOf) SetCanUpdateTiers(v bool)`

SetCanUpdateTiers gets a reference to the given bool and assigns it to the CanUpdateTiers field.

### GetCanUpdateJoinPolicy

`func (o *LoyaltyProgramAllOf) GetCanUpdateJoinPolicy() bool`

GetCanUpdateJoinPolicy returns the CanUpdateJoinPolicy field if non-nil, zero value otherwise.

### GetCanUpdateJoinPolicyOk

`func (o *LoyaltyProgramAllOf) GetCanUpdateJoinPolicyOk() (bool, bool)`

GetCanUpdateJoinPolicyOk returns a tuple with the CanUpdateJoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpdateJoinPolicy

`func (o *LoyaltyProgramAllOf) HasCanUpdateJoinPolicy() bool`

HasCanUpdateJoinPolicy returns a boolean if a field has been set.

### SetCanUpdateJoinPolicy

`func (o *LoyaltyProgramAllOf) SetCanUpdateJoinPolicy(v bool)`

SetCanUpdateJoinPolicy gets a reference to the given bool and assigns it to the CanUpdateJoinPolicy field.

### GetCanUpdateTierExpirationPolicy

`func (o *LoyaltyProgramAllOf) GetCanUpdateTierExpirationPolicy() bool`

GetCanUpdateTierExpirationPolicy returns the CanUpdateTierExpirationPolicy field if non-nil, zero value otherwise.

### GetCanUpdateTierExpirationPolicyOk

`func (o *LoyaltyProgramAllOf) GetCanUpdateTierExpirationPolicyOk() (bool, bool)`

GetCanUpdateTierExpirationPolicyOk returns a tuple with the CanUpdateTierExpirationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpdateTierExpirationPolicy

`func (o *LoyaltyProgramAllOf) HasCanUpdateTierExpirationPolicy() bool`

HasCanUpdateTierExpirationPolicy returns a boolean if a field has been set.

### SetCanUpdateTierExpirationPolicy

`func (o *LoyaltyProgramAllOf) SetCanUpdateTierExpirationPolicy(v bool)`

SetCanUpdateTierExpirationPolicy gets a reference to the given bool and assigns it to the CanUpdateTierExpirationPolicy field.

### GetCanUpgradeToAdvancedTiers

`func (o *LoyaltyProgramAllOf) GetCanUpgradeToAdvancedTiers() bool`

GetCanUpgradeToAdvancedTiers returns the CanUpgradeToAdvancedTiers field if non-nil, zero value otherwise.

### GetCanUpgradeToAdvancedTiersOk

`func (o *LoyaltyProgramAllOf) GetCanUpgradeToAdvancedTiersOk() (bool, bool)`

GetCanUpgradeToAdvancedTiersOk returns a tuple with the CanUpgradeToAdvancedTiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCanUpgradeToAdvancedTiers

`func (o *LoyaltyProgramAllOf) HasCanUpgradeToAdvancedTiers() bool`

HasCanUpgradeToAdvancedTiers returns a boolean if a field has been set.

### SetCanUpgradeToAdvancedTiers

`func (o *LoyaltyProgramAllOf) SetCanUpgradeToAdvancedTiers(v bool)`

SetCanUpgradeToAdvancedTiers gets a reference to the given bool and assigns it to the CanUpgradeToAdvancedTiers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


