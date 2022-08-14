# UpdateApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this application. | 
**Description** | Pointer to **string** | A longer description of the application. | [optional] 
**Timezone** | Pointer to **string** | A string containing an IANA timezone descriptor. | 
**Currency** | Pointer to **string** | The default currency for new customer sessions. | 
**CaseSensitivity** | Pointer to **string** | The case sensitivity behavior to check coupon codes in the campaigns of this Application. | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with this campaign. | [optional] 
**Limits** | Pointer to [**[]LimitConfig**](LimitConfig.md) | Default limits for campaigns created in this application. | [optional] 
**CampaignPriority** | Pointer to **string** | Default [priority](https://docs.talon.one/docs/product/applications/setting-up-campaign-priorities) for campaigns created in this Application.  | [optional] [default to CAMPAIGN_PRIORITY_UNIVERSAL]
**ExclusiveCampaignsStrategy** | Pointer to **string** | The strategy used when choosing exclusive campaigns for evaluation. | [optional] [default to EXCLUSIVE_CAMPAIGNS_STRATEGY_LIST_ORDER]
**DefaultDiscountScope** | Pointer to **string** | The default scope to apply &#x60;setDiscount&#x60; effects on if no scope was provided with the effect.  | [optional] 
**EnableCascadingDiscounts** | Pointer to **bool** | Indicates if discounts should cascade for this Application. | [optional] 
**EnableFlattenedCartItems** | Pointer to **bool** | Indicates if cart items of quantity larger than one should be separated into different items of quantity one. See [the docs](https://docs.talon.one/docs/product/campaigns/campaign-evaluation/#flattened-cart-items).  | [optional] 
**AttributesSettings** | Pointer to [**AttributesSettings**](AttributesSettings.md) |  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox Application. | [optional] 
**EnablePartialDiscounts** | Pointer to **bool** | Indicates if this Application supports partial discounts. | [optional] 
**DefaultDiscountAdditionalCostPerItemScope** | Pointer to **string** | The default scope to apply &#x60;setDiscountPerItem&#x60; effects on if no scope was provided with the effect.  | [optional] 

## Methods

### GetName

`func (o *UpdateApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApplication) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *UpdateApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *UpdateApplication) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetDescription

`func (o *UpdateApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApplication) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *UpdateApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *UpdateApplication) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetTimezone

`func (o *UpdateApplication) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateApplication) GetTimezoneOk() (string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimezone

`func (o *UpdateApplication) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezone

`func (o *UpdateApplication) SetTimezone(v string)`

SetTimezone gets a reference to the given string and assigns it to the Timezone field.

### GetCurrency

`func (o *UpdateApplication) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateApplication) GetCurrencyOk() (string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCurrency

`func (o *UpdateApplication) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrency

`func (o *UpdateApplication) SetCurrency(v string)`

SetCurrency gets a reference to the given string and assigns it to the Currency field.

### GetCaseSensitivity

`func (o *UpdateApplication) GetCaseSensitivity() string`

GetCaseSensitivity returns the CaseSensitivity field if non-nil, zero value otherwise.

### GetCaseSensitivityOk

`func (o *UpdateApplication) GetCaseSensitivityOk() (string, bool)`

GetCaseSensitivityOk returns a tuple with the CaseSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCaseSensitivity

`func (o *UpdateApplication) HasCaseSensitivity() bool`

HasCaseSensitivity returns a boolean if a field has been set.

### SetCaseSensitivity

`func (o *UpdateApplication) SetCaseSensitivity(v string)`

SetCaseSensitivity gets a reference to the given string and assigns it to the CaseSensitivity field.

### GetAttributes

`func (o *UpdateApplication) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateApplication) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *UpdateApplication) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *UpdateApplication) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetLimits

`func (o *UpdateApplication) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateApplication) GetLimitsOk() ([]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimits

`func (o *UpdateApplication) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### SetLimits

`func (o *UpdateApplication) SetLimits(v []LimitConfig)`

SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.

### GetCampaignPriority

`func (o *UpdateApplication) GetCampaignPriority() string`

GetCampaignPriority returns the CampaignPriority field if non-nil, zero value otherwise.

### GetCampaignPriorityOk

`func (o *UpdateApplication) GetCampaignPriorityOk() (string, bool)`

GetCampaignPriorityOk returns a tuple with the CampaignPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignPriority

`func (o *UpdateApplication) HasCampaignPriority() bool`

HasCampaignPriority returns a boolean if a field has been set.

### SetCampaignPriority

`func (o *UpdateApplication) SetCampaignPriority(v string)`

SetCampaignPriority gets a reference to the given string and assigns it to the CampaignPriority field.

### GetExclusiveCampaignsStrategy

`func (o *UpdateApplication) GetExclusiveCampaignsStrategy() string`

GetExclusiveCampaignsStrategy returns the ExclusiveCampaignsStrategy field if non-nil, zero value otherwise.

### GetExclusiveCampaignsStrategyOk

`func (o *UpdateApplication) GetExclusiveCampaignsStrategyOk() (string, bool)`

GetExclusiveCampaignsStrategyOk returns a tuple with the ExclusiveCampaignsStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExclusiveCampaignsStrategy

`func (o *UpdateApplication) HasExclusiveCampaignsStrategy() bool`

HasExclusiveCampaignsStrategy returns a boolean if a field has been set.

### SetExclusiveCampaignsStrategy

`func (o *UpdateApplication) SetExclusiveCampaignsStrategy(v string)`

SetExclusiveCampaignsStrategy gets a reference to the given string and assigns it to the ExclusiveCampaignsStrategy field.

### GetDefaultDiscountScope

`func (o *UpdateApplication) GetDefaultDiscountScope() string`

GetDefaultDiscountScope returns the DefaultDiscountScope field if non-nil, zero value otherwise.

### GetDefaultDiscountScopeOk

`func (o *UpdateApplication) GetDefaultDiscountScopeOk() (string, bool)`

GetDefaultDiscountScopeOk returns a tuple with the DefaultDiscountScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultDiscountScope

`func (o *UpdateApplication) HasDefaultDiscountScope() bool`

HasDefaultDiscountScope returns a boolean if a field has been set.

### SetDefaultDiscountScope

`func (o *UpdateApplication) SetDefaultDiscountScope(v string)`

SetDefaultDiscountScope gets a reference to the given string and assigns it to the DefaultDiscountScope field.

### GetEnableCascadingDiscounts

`func (o *UpdateApplication) GetEnableCascadingDiscounts() bool`

GetEnableCascadingDiscounts returns the EnableCascadingDiscounts field if non-nil, zero value otherwise.

### GetEnableCascadingDiscountsOk

`func (o *UpdateApplication) GetEnableCascadingDiscountsOk() (bool, bool)`

GetEnableCascadingDiscountsOk returns a tuple with the EnableCascadingDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnableCascadingDiscounts

`func (o *UpdateApplication) HasEnableCascadingDiscounts() bool`

HasEnableCascadingDiscounts returns a boolean if a field has been set.

### SetEnableCascadingDiscounts

`func (o *UpdateApplication) SetEnableCascadingDiscounts(v bool)`

SetEnableCascadingDiscounts gets a reference to the given bool and assigns it to the EnableCascadingDiscounts field.

### GetEnableFlattenedCartItems

`func (o *UpdateApplication) GetEnableFlattenedCartItems() bool`

GetEnableFlattenedCartItems returns the EnableFlattenedCartItems field if non-nil, zero value otherwise.

### GetEnableFlattenedCartItemsOk

`func (o *UpdateApplication) GetEnableFlattenedCartItemsOk() (bool, bool)`

GetEnableFlattenedCartItemsOk returns a tuple with the EnableFlattenedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnableFlattenedCartItems

`func (o *UpdateApplication) HasEnableFlattenedCartItems() bool`

HasEnableFlattenedCartItems returns a boolean if a field has been set.

### SetEnableFlattenedCartItems

`func (o *UpdateApplication) SetEnableFlattenedCartItems(v bool)`

SetEnableFlattenedCartItems gets a reference to the given bool and assigns it to the EnableFlattenedCartItems field.

### GetAttributesSettings

`func (o *UpdateApplication) GetAttributesSettings() AttributesSettings`

GetAttributesSettings returns the AttributesSettings field if non-nil, zero value otherwise.

### GetAttributesSettingsOk

`func (o *UpdateApplication) GetAttributesSettingsOk() (AttributesSettings, bool)`

GetAttributesSettingsOk returns a tuple with the AttributesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributesSettings

`func (o *UpdateApplication) HasAttributesSettings() bool`

HasAttributesSettings returns a boolean if a field has been set.

### SetAttributesSettings

`func (o *UpdateApplication) SetAttributesSettings(v AttributesSettings)`

SetAttributesSettings gets a reference to the given AttributesSettings and assigns it to the AttributesSettings field.

### GetSandbox

`func (o *UpdateApplication) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *UpdateApplication) GetSandboxOk() (bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSandbox

`func (o *UpdateApplication) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### SetSandbox

`func (o *UpdateApplication) SetSandbox(v bool)`

SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.

### GetEnablePartialDiscounts

`func (o *UpdateApplication) GetEnablePartialDiscounts() bool`

GetEnablePartialDiscounts returns the EnablePartialDiscounts field if non-nil, zero value otherwise.

### GetEnablePartialDiscountsOk

`func (o *UpdateApplication) GetEnablePartialDiscountsOk() (bool, bool)`

GetEnablePartialDiscountsOk returns a tuple with the EnablePartialDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEnablePartialDiscounts

`func (o *UpdateApplication) HasEnablePartialDiscounts() bool`

HasEnablePartialDiscounts returns a boolean if a field has been set.

### SetEnablePartialDiscounts

`func (o *UpdateApplication) SetEnablePartialDiscounts(v bool)`

SetEnablePartialDiscounts gets a reference to the given bool and assigns it to the EnablePartialDiscounts field.

### GetDefaultDiscountAdditionalCostPerItemScope

`func (o *UpdateApplication) GetDefaultDiscountAdditionalCostPerItemScope() string`

GetDefaultDiscountAdditionalCostPerItemScope returns the DefaultDiscountAdditionalCostPerItemScope field if non-nil, zero value otherwise.

### GetDefaultDiscountAdditionalCostPerItemScopeOk

`func (o *UpdateApplication) GetDefaultDiscountAdditionalCostPerItemScopeOk() (string, bool)`

GetDefaultDiscountAdditionalCostPerItemScopeOk returns a tuple with the DefaultDiscountAdditionalCostPerItemScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDefaultDiscountAdditionalCostPerItemScope

`func (o *UpdateApplication) HasDefaultDiscountAdditionalCostPerItemScope() bool`

HasDefaultDiscountAdditionalCostPerItemScope returns a boolean if a field has been set.

### SetDefaultDiscountAdditionalCostPerItemScope

`func (o *UpdateApplication) SetDefaultDiscountAdditionalCostPerItemScope(v string)`

SetDefaultDiscountAdditionalCostPerItemScope gets a reference to the given string and assigns it to the DefaultDiscountAdditionalCostPerItemScope field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


