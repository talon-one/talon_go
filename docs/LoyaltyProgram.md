# LoyaltyProgram

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of loyalty program. | 
**AccountID** | Pointer to **int32** | The ID of the Talon.One account that owns this program. | 
**Name** | Pointer to **string** | The internal name for the Loyalty Program. | 
**Title** | Pointer to **string** | The display title for the Loyalty Program. | 
**Description** | Pointer to **string** | Description of our Loyalty Program. | 
**SubscribedApplications** | Pointer to **[]int32** | A list containing the IDs of all applications that are subscribed to this Loyalty Program. | 
**DefaultValidity** | Pointer to **string** | Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like &#39;1h&#39; or &#39;40m&#39;. | 
**DefaultPending** | Pointer to **string** | Indicates the default duration for the pending time, after which points will be valid. The format is a number followed by a duration unit, like &#39;1h&#39; or &#39;40m&#39;. | 
**AllowSubledger** | Pointer to **bool** | Indicates if this program supports subledgers inside the program | 
**Tiers** | Pointer to [**[]LoyaltyTier**](LoyaltyTier.md) | The tiers in this loyalty program | [optional] 

## Methods

### GetId

`func (o *LoyaltyProgram) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoyaltyProgram) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *LoyaltyProgram) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *LoyaltyProgram) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetAccountID

`func (o *LoyaltyProgram) GetAccountID() int32`

GetAccountID returns the AccountID field if non-nil, zero value otherwise.

### GetAccountIDOk

`func (o *LoyaltyProgram) GetAccountIDOk() (int32, bool)`

GetAccountIDOk returns a tuple with the AccountID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountID

`func (o *LoyaltyProgram) HasAccountID() bool`

HasAccountID returns a boolean if a field has been set.

### SetAccountID

`func (o *LoyaltyProgram) SetAccountID(v int32)`

SetAccountID gets a reference to the given int32 and assigns it to the AccountID field.

### GetName

`func (o *LoyaltyProgram) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoyaltyProgram) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *LoyaltyProgram) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *LoyaltyProgram) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetTitle

`func (o *LoyaltyProgram) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LoyaltyProgram) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *LoyaltyProgram) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *LoyaltyProgram) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDescription

`func (o *LoyaltyProgram) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoyaltyProgram) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *LoyaltyProgram) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *LoyaltyProgram) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetSubscribedApplications

`func (o *LoyaltyProgram) GetSubscribedApplications() []int32`

GetSubscribedApplications returns the SubscribedApplications field if non-nil, zero value otherwise.

### GetSubscribedApplicationsOk

`func (o *LoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool)`

GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSubscribedApplications

`func (o *LoyaltyProgram) HasSubscribedApplications() bool`

HasSubscribedApplications returns a boolean if a field has been set.

### SetSubscribedApplications

`func (o *LoyaltyProgram) SetSubscribedApplications(v []int32)`

SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.

### GetDefaultValidity

`func (o *LoyaltyProgram) GetDefaultValidity() string`

GetDefaultValidity returns the DefaultValidity field if non-nil, zero value otherwise.

### GetDefaultValidityOk

`func (o *LoyaltyProgram) GetDefaultValidityOk() (string, bool)`

GetDefaultValidityOk returns a tuple with the DefaultValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultValidity

`func (o *LoyaltyProgram) HasDefaultValidity() bool`

HasDefaultValidity returns a boolean if a field has been set.

### SetDefaultValidity

`func (o *LoyaltyProgram) SetDefaultValidity(v string)`

SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.

### GetDefaultPending

`func (o *LoyaltyProgram) GetDefaultPending() string`

GetDefaultPending returns the DefaultPending field if non-nil, zero value otherwise.

### GetDefaultPendingOk

`func (o *LoyaltyProgram) GetDefaultPendingOk() (string, bool)`

GetDefaultPendingOk returns a tuple with the DefaultPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultPending

`func (o *LoyaltyProgram) HasDefaultPending() bool`

HasDefaultPending returns a boolean if a field has been set.

### SetDefaultPending

`func (o *LoyaltyProgram) SetDefaultPending(v string)`

SetDefaultPending gets a reference to the given string and assigns it to the DefaultPending field.

### GetAllowSubledger

`func (o *LoyaltyProgram) GetAllowSubledger() bool`

GetAllowSubledger returns the AllowSubledger field if non-nil, zero value otherwise.

### GetAllowSubledgerOk

`func (o *LoyaltyProgram) GetAllowSubledgerOk() (bool, bool)`

GetAllowSubledgerOk returns a tuple with the AllowSubledger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAllowSubledger

`func (o *LoyaltyProgram) HasAllowSubledger() bool`

HasAllowSubledger returns a boolean if a field has been set.

### SetAllowSubledger

`func (o *LoyaltyProgram) SetAllowSubledger(v bool)`

SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.

### GetTiers

`func (o *LoyaltyProgram) GetTiers() []LoyaltyTier`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *LoyaltyProgram) GetTiersOk() ([]LoyaltyTier, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTiers

`func (o *LoyaltyProgram) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### SetTiers

`func (o *LoyaltyProgram) SetTiers(v []LoyaltyTier)`

SetTiers gets a reference to the given []LoyaltyTier and assigns it to the Tiers field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


