# NewLoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The display title for the Loyalty Program. | 
**Description** | Pointer to **string** | Description of our Loyalty Program. | [optional] 
**SubscribedApplications** | Pointer to **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | [optional] 
**DefaultValidity** | Pointer to **string** | Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like &#39;1h&#39; or &#39;40m&#39;. | 
**DefaultPending** | Pointer to **string** | Indicates the default duration for the pending time, after which points will be valid. The format is a number followed by a duration unit, like &#39;1h&#39; or &#39;40m&#39;. | 
**AllowSubledger** | Pointer to **bool** | Indicates if this program supports subledgers inside the program | 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | [optional] 
**Name** | Pointer to **string** | The internal name for the Loyalty Program. This is an immutable value. | 
**Tiers** | Pointer to [**[]NewLoyaltyTier**](NewLoyaltyTier.md) | The tiers in this loyalty program | [optional] 

## Methods

### GetTitle

`func (o *NewLoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NewLoyaltyProgram) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *NewLoyaltyProgram) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *NewLoyaltyProgram) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *NewLoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewLoyaltyProgram) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewLoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewLoyaltyProgram) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplications

`func (o *NewLoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *NewLoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplications

`func (o *NewLoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### SetSubscribedApplications

`func (o *NewLoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.

### GetDefaultValidity

`func (o *NewLoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *NewLoyaltyProgram) GetDefaultValidityOk() (string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultValidity

`func (o *NewLoyaltyProgram) HasDefaultValidity() bool`

HasDefaultValidity returns a boolean if a field has been set.

### SetDefaultValidity

`func (o *NewLoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.

### GetDefaultPending

`func (o *NewLoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *NewLoyaltyProgram) GetDefaultPendingOk() (string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultPending

`func (o *NewLoyaltyProgram) HasDefaultPending() bool`

HasDefaultPending returns a boolean if a field has been set.

### SetDefaultPending

`func (o *NewLoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending gets a reference to the given string and assigns it to the DefaultPending field.

### GetAllowSubledger

`func (o *NewLoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *NewLoyaltyProgram) GetAllowSubledgerOk() (bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowSubledger

`func (o *NewLoyaltyProgram) HasAllowSubledger() bool`

HasAllowSubledger returns a boolean if a field has been set.

### SetAllowSubledger

`func (o *NewLoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.

### GetTimezone

`func (o *NewLoyaltyProgram) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NewLoyaltyProgram) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *NewLoyaltyProgram) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *NewLoyaltyProgram) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetName

`func (o *NewLoyaltyProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewLoyaltyProgram) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewLoyaltyProgram) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewLoyaltyProgram) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTiers

`func (o *NewLoyaltyProgram) GetTiers() []NewLoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *NewLoyaltyProgram) GetTiersOk() ([]NewLoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiers

`func (o *NewLoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### SetTiers

`func (o *NewLoyaltyProgram) SetTiers(v []NewLoyaltyTier)`

SetTiers gets a reference to the given []NewLoyaltyTier and assigns it to the Tiers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


