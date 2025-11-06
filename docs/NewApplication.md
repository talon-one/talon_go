# NewApplication

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
**DefaultDiscountScope** | Pointer to **string** | The default scope to apply &#x60;setDiscount&#x60; effects on if no scope was provided with the effect.  | [optional] 
**EnableCascadingDiscounts** | Pointer to **bool** | Indicates if discounts should cascade for this Application. | [optional] 
**EnableFlattenedCartItems** | Pointer to **bool** | Indicates if cart items of quantity larger than one should be separated into different items of quantity one.  | [optional] 
**AttributesSettings** | Pointer to [**AttributesSettings**](AttributesSettings.md) |  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox Application. | [optional] 
**EnablePartialDiscounts** | Pointer to **bool** | Indicates if this Application supports partial discounts. | [optional] 
**DefaultDiscountAdditionalCostPerItemScope** | Pointer to **string** | The default scope to apply &#x60;setDiscountPerItem&#x60; effects on if no scope was provided with the effect.  | [optional] 
**Key** | Pointer to **string** | Hex key for HMAC-signing API calls as coming from this application (16 hex digits). | [optional] 
**EnableCampaignStateManagement** | Pointer to **bool** | Indicates whether the campaign staging and revisions feature is enabled for the Application.  **Important:** After this feature is enabled, it cannot be disabled.  | [optional] 

## Methods

### NewNewApplication

`func NewNewApplication(name string, timezone string, currency string, ) *NewApplication`

NewNewApplication instantiates a new NewApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewApplicationWithDefaults

`func NewNewApplicationWithDefaults() *NewApplication`

NewNewApplicationWithDefaults instantiates a new NewApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NewApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NewApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NewApplication) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NewApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NewApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NewApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NewApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimezone

`func (o *NewApplication) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *NewApplication) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *NewApplication) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCurrency

`func (o *NewApplication) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NewApplication) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *NewApplication) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCaseSensitivity

`func (o *NewApplication) GetCaseSensitivity() string`

GetCaseSensitivity returns the CaseSensitivity field if non-nil, zero value otherwise.

### GetCaseSensitivityOk

`func (o *NewApplication) GetCaseSensitivityOk() (*string, bool)`

GetCaseSensitivityOk returns a tuple with the CaseSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitivity

`func (o *NewApplication) SetCaseSensitivity(v string)`

SetCaseSensitivity sets CaseSensitivity field to given value.

### HasCaseSensitivity

`func (o *NewApplication) HasCaseSensitivity() bool`

HasCaseSensitivity returns a boolean if a field has been set.

### GetAttributes

`func (o *NewApplication) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *NewApplication) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *NewApplication) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *NewApplication) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLimits

`func (o *NewApplication) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *NewApplication) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *NewApplication) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *NewApplication) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetDefaultDiscountScope

`func (o *NewApplication) GetDefaultDiscountScope() string`

GetDefaultDiscountScope returns the DefaultDiscountScope field if non-nil, zero value otherwise.

### GetDefaultDiscountScopeOk

`func (o *NewApplication) GetDefaultDiscountScopeOk() (*string, bool)`

GetDefaultDiscountScopeOk returns a tuple with the DefaultDiscountScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountScope

`func (o *NewApplication) SetDefaultDiscountScope(v string)`

SetDefaultDiscountScope sets DefaultDiscountScope field to given value.

### HasDefaultDiscountScope

`func (o *NewApplication) HasDefaultDiscountScope() bool`

HasDefaultDiscountScope returns a boolean if a field has been set.

### GetEnableCascadingDiscounts

`func (o *NewApplication) GetEnableCascadingDiscounts() bool`

GetEnableCascadingDiscounts returns the EnableCascadingDiscounts field if non-nil, zero value otherwise.

### GetEnableCascadingDiscountsOk

`func (o *NewApplication) GetEnableCascadingDiscountsOk() (*bool, bool)`

GetEnableCascadingDiscountsOk returns a tuple with the EnableCascadingDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCascadingDiscounts

`func (o *NewApplication) SetEnableCascadingDiscounts(v bool)`

SetEnableCascadingDiscounts sets EnableCascadingDiscounts field to given value.

### HasEnableCascadingDiscounts

`func (o *NewApplication) HasEnableCascadingDiscounts() bool`

HasEnableCascadingDiscounts returns a boolean if a field has been set.

### GetEnableFlattenedCartItems

`func (o *NewApplication) GetEnableFlattenedCartItems() bool`

GetEnableFlattenedCartItems returns the EnableFlattenedCartItems field if non-nil, zero value otherwise.

### GetEnableFlattenedCartItemsOk

`func (o *NewApplication) GetEnableFlattenedCartItemsOk() (*bool, bool)`

GetEnableFlattenedCartItemsOk returns a tuple with the EnableFlattenedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFlattenedCartItems

`func (o *NewApplication) SetEnableFlattenedCartItems(v bool)`

SetEnableFlattenedCartItems sets EnableFlattenedCartItems field to given value.

### HasEnableFlattenedCartItems

`func (o *NewApplication) HasEnableFlattenedCartItems() bool`

HasEnableFlattenedCartItems returns a boolean if a field has been set.

### GetAttributesSettings

`func (o *NewApplication) GetAttributesSettings() AttributesSettings`

GetAttributesSettings returns the AttributesSettings field if non-nil, zero value otherwise.

### GetAttributesSettingsOk

`func (o *NewApplication) GetAttributesSettingsOk() (*AttributesSettings, bool)`

GetAttributesSettingsOk returns a tuple with the AttributesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesSettings

`func (o *NewApplication) SetAttributesSettings(v AttributesSettings)`

SetAttributesSettings sets AttributesSettings field to given value.

### HasAttributesSettings

`func (o *NewApplication) HasAttributesSettings() bool`

HasAttributesSettings returns a boolean if a field has been set.

### GetSandbox

`func (o *NewApplication) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *NewApplication) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *NewApplication) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *NewApplication) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetEnablePartialDiscounts

`func (o *NewApplication) GetEnablePartialDiscounts() bool`

GetEnablePartialDiscounts returns the EnablePartialDiscounts field if non-nil, zero value otherwise.

### GetEnablePartialDiscountsOk

`func (o *NewApplication) GetEnablePartialDiscountsOk() (*bool, bool)`

GetEnablePartialDiscountsOk returns a tuple with the EnablePartialDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartialDiscounts

`func (o *NewApplication) SetEnablePartialDiscounts(v bool)`

SetEnablePartialDiscounts sets EnablePartialDiscounts field to given value.

### HasEnablePartialDiscounts

`func (o *NewApplication) HasEnablePartialDiscounts() bool`

HasEnablePartialDiscounts returns a boolean if a field has been set.

### GetDefaultDiscountAdditionalCostPerItemScope

`func (o *NewApplication) GetDefaultDiscountAdditionalCostPerItemScope() string`

GetDefaultDiscountAdditionalCostPerItemScope returns the DefaultDiscountAdditionalCostPerItemScope field if non-nil, zero value otherwise.

### GetDefaultDiscountAdditionalCostPerItemScopeOk

`func (o *NewApplication) GetDefaultDiscountAdditionalCostPerItemScopeOk() (*string, bool)`

GetDefaultDiscountAdditionalCostPerItemScopeOk returns a tuple with the DefaultDiscountAdditionalCostPerItemScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountAdditionalCostPerItemScope

`func (o *NewApplication) SetDefaultDiscountAdditionalCostPerItemScope(v string)`

SetDefaultDiscountAdditionalCostPerItemScope sets DefaultDiscountAdditionalCostPerItemScope field to given value.

### HasDefaultDiscountAdditionalCostPerItemScope

`func (o *NewApplication) HasDefaultDiscountAdditionalCostPerItemScope() bool`

HasDefaultDiscountAdditionalCostPerItemScope returns a boolean if a field has been set.

### GetKey

`func (o *NewApplication) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *NewApplication) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *NewApplication) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *NewApplication) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetEnableCampaignStateManagement

`func (o *NewApplication) GetEnableCampaignStateManagement() bool`

GetEnableCampaignStateManagement returns the EnableCampaignStateManagement field if non-nil, zero value otherwise.

### GetEnableCampaignStateManagementOk

`func (o *NewApplication) GetEnableCampaignStateManagementOk() (*bool, bool)`

GetEnableCampaignStateManagementOk returns a tuple with the EnableCampaignStateManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCampaignStateManagement

`func (o *NewApplication) SetEnableCampaignStateManagement(v bool)`

SetEnableCampaignStateManagement sets EnableCampaignStateManagement field to given value.

### HasEnableCampaignStateManagement

`func (o *NewApplication) HasEnableCampaignStateManagement() bool`

HasEnableCampaignStateManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


