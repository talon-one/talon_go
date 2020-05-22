# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The exact moment this entity was last modified. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Name** | Pointer to **string** | The name of this application. | 
**Description** | Pointer to **string** | A longer description of the application. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**Currency** | Pointer to **string** | A string describing a default currency for new customer sessions. | 
**CaseSensitivity** | Pointer to **string** | A string indicating how should campaigns in this application deal with case sensitivity on coupon codes. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Default limits for campaigns created in this application | [optional] 
**CampaignPriority** | Pointer to **string** | Default priority for campaigns created in this application, can be one of (universal, stackable, exclusive) | [optional] 
**AttributesSettings** | Pointer to [**AttributesSettings**](AttributesSettings.md) |  | [optional] 
**LoyaltyPrograms** | Pointer to [**[]LoyaltyProgram**](LoyaltyProgram.md) | An array containing all the loyalty programs to which this application is subscribed | 

## Methods

### GetId

`func (o *Application) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *Application) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *Application) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *Application) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Application) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *Application) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *Application) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetModified

`func (o *Application) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Application) GetModifiedOk() (time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasModified

`func (o *Application) HasModified() bool`

HasModified returns a boolean if a field has been set.

### SetModified

`func (o *Application) SetModified(v time.Time)`

SetModified gets a reference to the given time.Time and assigns it to the Modified field.

### GetAccountId

`func (o *Application) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Application) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *Application) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *Application) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *Application) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetTimezone

`func (o *Application) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Application) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *Application) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *Application) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetCurrency

`func (o *Application) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Application) GetCurrencyOk() (string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrency

`func (o *Application) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrency

`func (o *Application) SetCurrency(v string)`

SetCurrency gets a reference to the given string and assigns it to the Currency field.

### GetCaseSensitivity

`func (o *Application) GetCaseSensitivity() string`

GetCaseSensitivity returns the CaseSensitivity field if non-nil, zero value otherwise.

### GetCaseSensitivityOk

`func (o *Application) GetCaseSensitivityOk() (string, bool)`

GetCaseSensitivityOk returns a tuple with the CaseSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCaseSensitivity

`func (o *Application) HasCaseSensitivity() bool`

HasCaseSensitivity returns a boolean if a field has been set.

### SetCaseSensitivity

`func (o *Application) SetCaseSensitivity(v string)`

SetCaseSensitivity gets a reference to the given string and assigns it to the CaseSensitivity field.

### GetAttributes

`func (o *Application) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Application) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *Application) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *Application) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetLimits

`func (o *Application) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Application) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *Application) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *Application) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetCampaignPriority

`func (o *Application) GetCampaignPriority() string`

GetCampaignPriority returns the CampaignPriority field if non-nil, zero value otherwise.

### GetCampaignPriorityOk

`func (o *Application) GetCampaignPriorityOk() (string, bool)`

GetCampaignPriorityOk returns a tuple with the CampaignPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignPriority

`func (o *Application) HasCampaignPriority() bool`

HasCampaignPriority returns a boolean if a field has been set.

### SetCampaignPriority

`func (o *Application) SetCampaignPriority(v string)`

SetCampaignPriority gets a reference to the given string and assigns it to the CampaignPriority field.

### GetAttributesSettings

`func (o *Application) GetAttributesSettings() AttributesSettings`

GetAttributesSettings returns the AttributesSettings field if non-nil, zero value otherwise.

### GetAttributesSettingsOk

`func (o *Application) GetAttributesSettingsOk() (AttributesSettings, bool)`

GetAttributesSettingsOk returns a tuple with the AttributesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributesSettings

`func (o *Application) HasAttributesSettings() bool`

HasAttributesSettings returns a boolean if a field has been set.

### SetAttributesSettings

`func (o *Application) SetAttributesSettings(v AttributesSettings)`

SetAttributesSettings gets a reference to the given AttributesSettings and assigns it to the AttributesSettings field.

### GetLoyaltyPrograms

`func (o *Application) GetLoyaltyPrograms() []LoyaltyProgram`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *Application) GetLoyaltyProgramsOk() ([]LoyaltyProgram, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLoyaltyPrograms

`func (o *Application) HasLoyaltyPrograms() bool`

HasLoyaltyPrograms returns a boolean if a field has been set.

### SetLoyaltyPrograms

`func (o *Application) SetLoyaltyPrograms(v []LoyaltyProgram)`

SetLoyaltyPrograms gets a reference to the given []LoyaltyProgram and assigns it to the LoyaltyPrograms field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


