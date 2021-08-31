# NewApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this application. | 
**Description** | Pointer to **string** | A longer description of the application. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**Currency** | Pointer to **string** | A string describing a default currency for new customer sessions. | 
**CaseSensitivity** | Pointer to **string** | A string indicating how should campaigns in this application deal with case sensitivity on coupon codes. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Default limits for campaigns created in this application | [optional] 
**CampaignPriority** | Pointer to **string** | Default priority for campaigns created in this application, can be one of (universal, stackable, exclusive). If no value is provided, this is set to \&quot;universal\&quot; | [optional] 
**ExclusiveCampaignsStrategy** | Pointer to **string** | The strategy used when choosing exclusive campaigns for evaluation, can be one of (listOrder, lowestDiscount, highestDiscount). If no value is provided, this is set to \&quot;listOrder\&quot; | [optional] 
**DefaultDiscountScope** | Pointer to **string** | The default scope to apply \&quot;setDiscount\&quot; effects on if no scope was provided with the effect. | [optional] 
**EnableCascadingDiscounts** | Pointer to **bool** | Indicates if discounts should cascade for this application | [optional] 
**EnableFlattenedCartItems** | Pointer to **bool** | Indicates if cart items of quantity larger than one should be separated into different items of quantity one | [optional] 
**AttributesSettings** | Pointer to [**AttributesSettings**](AttributesSettings.md) |  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox application | [optional] 
**Key** | Pointer to **string** | Hex key for HMAC-signing API calls as coming from this application (16 hex digits) | [optional] 

## Methods

### GetName

`func (o *NewApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewApplication) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *NewApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *NewApplication) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *NewApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewApplication) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *NewApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *NewApplication) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetTimezone

`func (o *NewApplication) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NewApplication) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *NewApplication) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *NewApplication) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetCurrency

`func (o *NewApplication) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NewApplication) GetCurrencyOk() (string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrency

`func (o *NewApplication) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrency

`func (o *NewApplication) SetCurrency(v string)`

SetCurrency gets a reference to the given string and assigns it to the Currency field.

### GetCaseSensitivity

`func (o *NewApplication) GetCaseSensitivity() string`

GetCaseSensitivity returns the CaseSensitivity field if non-nil, zero value otherwise.

### GetCaseSensitivityOk

`func (o *NewApplication) GetCaseSensitivityOk() (string, bool)`

GetCaseSensitivityOk returns a tuple with the CaseSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCaseSensitivity

`func (o *NewApplication) HasCaseSensitivity() bool`

HasCaseSensitivity returns a boolean if a field has been set.

### SetCaseSensitivity

`func (o *NewApplication) SetCaseSensitivity(v string)`

SetCaseSensitivity gets a reference to the given string and assigns it to the CaseSensitivity field.

### GetAttributes

`func (o *NewApplication) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewApplication) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *NewApplication) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *NewApplication) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetLimits

`func (o *NewApplication) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *NewApplication) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *NewApplication) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *NewApplication) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetCampaignPriority

`func (o *NewApplication) GetCampaignPriority() string`

GetCampaignPriority returns the CampaignPriority field if non-nil, zero value otherwise.

### GetCampaignPriorityOk

`func (o *NewApplication) GetCampaignPriorityOk() (string, bool)`

GetCampaignPriorityOk returns a tuple with the CampaignPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignPriority

`func (o *NewApplication) HasCampaignPriority() bool`

HasCampaignPriority returns a boolean if a field has been set.

### SetCampaignPriority

`func (o *NewApplication) SetCampaignPriority(v string)`

SetCampaignPriority gets a reference to the given string and assigns it to the CampaignPriority field.

### GetExclusiveCampaignsStrategy

`func (o *NewApplication) GetExclusiveCampaignsStrategy() string`

GetExclusiveCampaignsStrategy returns the ExclusiveCampaignsStrategy field if non-nil, zero value otherwise.

### GetExclusiveCampaignsStrategyOk

`func (o *NewApplication) GetExclusiveCampaignsStrategyOk() (string, bool)`

GetExclusiveCampaignsStrategyOk returns a tuple with the ExclusiveCampaignsStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExclusiveCampaignsStrategy

`func (o *NewApplication) HasExclusiveCampaignsStrategy() bool`

HasExclusiveCampaignsStrategy returns a boolean if a field has been set.

### SetExclusiveCampaignsStrategy

`func (o *NewApplication) SetExclusiveCampaignsStrategy(v string)`

SetExclusiveCampaignsStrategy gets a reference to the given string and assigns it to the ExclusiveCampaignsStrategy field.

### GetDefaultDiscountScope

`func (o *NewApplication) GetDefaultDiscountScope() string`

GetDefaultDiscountScope returns the DefaultDiscountScope field if non-nil, zero value otherwise.

### GetDefaultDiscountScopeOk

`func (o *NewApplication) GetDefaultDiscountScopeOk() (string, bool)`

GetDefaultDiscountScopeOk returns a tuple with the DefaultDiscountScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultDiscountScope

`func (o *NewApplication) HasDefaultDiscountScope() bool`

HasDefaultDiscountScope returns a boolean if a field has been set.

### SetDefaultDiscountScope

`func (o *NewApplication) SetDefaultDiscountScope(v string)`

SetDefaultDiscountScope gets a reference to the given string and assigns it to the DefaultDiscountScope field.

### GetEnableCascadingDiscounts

`func (o *NewApplication) GetEnableCascadingDiscounts() bool`

GetEnableCascadingDiscounts returns the EnableCascadingDiscounts field if non-nil, zero value otherwise.

### GetEnableCascadingDiscountsOk

`func (o *NewApplication) GetEnableCascadingDiscountsOk() (bool, bool)`

GetEnableCascadingDiscountsOk returns a tuple with the EnableCascadingDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnableCascadingDiscounts

`func (o *NewApplication) HasEnableCascadingDiscounts() bool`

HasEnableCascadingDiscounts returns a boolean if a field has been set.

### SetEnableCascadingDiscounts

`func (o *NewApplication) SetEnableCascadingDiscounts(v bool)`

SetEnableCascadingDiscounts gets a reference to the given bool and assigns it to the EnableCascadingDiscounts field.

### GetEnableFlattenedCartItems

`func (o *NewApplication) GetEnableFlattenedCartItems() bool`

GetEnableFlattenedCartItems returns the EnableFlattenedCartItems field if non-nil, zero value otherwise.

### GetEnableFlattenedCartItemsOk

`func (o *NewApplication) GetEnableFlattenedCartItemsOk() (bool, bool)`

GetEnableFlattenedCartItemsOk returns a tuple with the EnableFlattenedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnableFlattenedCartItems

`func (o *NewApplication) HasEnableFlattenedCartItems() bool`

HasEnableFlattenedCartItems returns a boolean if a field has been set.

### SetEnableFlattenedCartItems

`func (o *NewApplication) SetEnableFlattenedCartItems(v bool)`

SetEnableFlattenedCartItems gets a reference to the given bool and assigns it to the EnableFlattenedCartItems field.

### GetAttributesSettings

`func (o *NewApplication) GetAttributesSettings() AttributesSettings`

GetAttributesSettings returns the AttributesSettings field if non-nil, zero value otherwise.

### GetAttributesSettingsOk

`func (o *NewApplication) GetAttributesSettingsOk() (AttributesSettings, bool)`

GetAttributesSettingsOk returns a tuple with the AttributesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributesSettings

`func (o *NewApplication) HasAttributesSettings() bool`

HasAttributesSettings returns a boolean if a field has been set.

### SetAttributesSettings

`func (o *NewApplication) SetAttributesSettings(v AttributesSettings)`

SetAttributesSettings gets a reference to the given AttributesSettings and assigns it to the AttributesSettings field.

### GetSandbox

`func (o *NewApplication) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewApplication) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *NewApplication) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *NewApplication) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.

### GetKey

`func (o *NewApplication) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NewApplication) GetKeyOk() (string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasKey

`func (o *NewApplication) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKey

`func (o *NewApplication) SetKey(v string)`

SetKey gets a reference to the given string and assigns it to the Key field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


