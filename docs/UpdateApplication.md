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
**DefaultDiscountScope** | Pointer to **string** | The default scope to apply &#x60;setDiscount&#x60; effects on if no scope was provided with the effect.  | [optional] 
**EnableCascadingDiscounts** | Pointer to **bool** | Indicates if discounts should cascade for this Application. | [optional] 
**EnableFlattenedCartItems** | Pointer to **bool** | Indicates if cart items of quantity larger than one should be separated into different items of quantity one.  | [optional] 
**AttributesSettings** | Pointer to [**AttributesSettings**](AttributesSettings.md) |  | [optional] 
**Sandbox** | Pointer to **bool** | Indicates if this is a live or sandbox Application. | [optional] 
**EnablePartialDiscounts** | Pointer to **bool** | Indicates if this Application supports partial discounts. | [optional] 
**DefaultDiscountAdditionalCostPerItemScope** | Pointer to **string** | The default scope to apply &#x60;setDiscountPerItem&#x60; effects on if no scope was provided with the effect.  | [optional] 
**DefaultEvaluationGroupId** | Pointer to **int64** | The ID of the default campaign evaluation group to which new campaigns will be added unless a different group is selected when creating the campaign. | [optional] 
**DefaultCartItemFilterId** | Pointer to **int64** | The ID of the default Cart-Item-Filter for this application. | [optional] 
**EnableCampaignStateManagement** | Pointer to **bool** | Indicates whether the campaign staging and revisions feature is enabled for the Application.  **Important:** After this feature is enabled, it cannot be disabled.  | [optional] 

## Methods

### NewUpdateApplication

`func NewUpdateApplication(name string, timezone string, currency string, ) *UpdateApplication`

NewUpdateApplication instantiates a new UpdateApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationWithDefaults

`func NewUpdateApplicationWithDefaults() *UpdateApplication`

NewUpdateApplicationWithDefaults instantiates a new UpdateApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateApplication) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateApplication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateApplication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateApplication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateApplication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateApplication) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateApplication) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateApplication) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCurrency

`func (o *UpdateApplication) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdateApplication) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdateApplication) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCaseSensitivity

`func (o *UpdateApplication) GetCaseSensitivity() string`

GetCaseSensitivity returns the CaseSensitivity field if non-nil, zero value otherwise.

### GetCaseSensitivityOk

`func (o *UpdateApplication) GetCaseSensitivityOk() (*string, bool)`

GetCaseSensitivityOk returns a tuple with the CaseSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitivity

`func (o *UpdateApplication) SetCaseSensitivity(v string)`

SetCaseSensitivity sets CaseSensitivity field to given value.

### HasCaseSensitivity

`func (o *UpdateApplication) HasCaseSensitivity() bool`

HasCaseSensitivity returns a boolean if a field has been set.

### GetAttributes

`func (o *UpdateApplication) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *UpdateApplication) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *UpdateApplication) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *UpdateApplication) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLimits

`func (o *UpdateApplication) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UpdateApplication) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *UpdateApplication) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *UpdateApplication) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetDefaultDiscountScope

`func (o *UpdateApplication) GetDefaultDiscountScope() string`

GetDefaultDiscountScope returns the DefaultDiscountScope field if non-nil, zero value otherwise.

### GetDefaultDiscountScopeOk

`func (o *UpdateApplication) GetDefaultDiscountScopeOk() (*string, bool)`

GetDefaultDiscountScopeOk returns a tuple with the DefaultDiscountScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountScope

`func (o *UpdateApplication) SetDefaultDiscountScope(v string)`

SetDefaultDiscountScope sets DefaultDiscountScope field to given value.

### HasDefaultDiscountScope

`func (o *UpdateApplication) HasDefaultDiscountScope() bool`

HasDefaultDiscountScope returns a boolean if a field has been set.

### GetEnableCascadingDiscounts

`func (o *UpdateApplication) GetEnableCascadingDiscounts() bool`

GetEnableCascadingDiscounts returns the EnableCascadingDiscounts field if non-nil, zero value otherwise.

### GetEnableCascadingDiscountsOk

`func (o *UpdateApplication) GetEnableCascadingDiscountsOk() (*bool, bool)`

GetEnableCascadingDiscountsOk returns a tuple with the EnableCascadingDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCascadingDiscounts

`func (o *UpdateApplication) SetEnableCascadingDiscounts(v bool)`

SetEnableCascadingDiscounts sets EnableCascadingDiscounts field to given value.

### HasEnableCascadingDiscounts

`func (o *UpdateApplication) HasEnableCascadingDiscounts() bool`

HasEnableCascadingDiscounts returns a boolean if a field has been set.

### GetEnableFlattenedCartItems

`func (o *UpdateApplication) GetEnableFlattenedCartItems() bool`

GetEnableFlattenedCartItems returns the EnableFlattenedCartItems field if non-nil, zero value otherwise.

### GetEnableFlattenedCartItemsOk

`func (o *UpdateApplication) GetEnableFlattenedCartItemsOk() (*bool, bool)`

GetEnableFlattenedCartItemsOk returns a tuple with the EnableFlattenedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFlattenedCartItems

`func (o *UpdateApplication) SetEnableFlattenedCartItems(v bool)`

SetEnableFlattenedCartItems sets EnableFlattenedCartItems field to given value.

### HasEnableFlattenedCartItems

`func (o *UpdateApplication) HasEnableFlattenedCartItems() bool`

HasEnableFlattenedCartItems returns a boolean if a field has been set.

### GetAttributesSettings

`func (o *UpdateApplication) GetAttributesSettings() AttributesSettings`

GetAttributesSettings returns the AttributesSettings field if non-nil, zero value otherwise.

### GetAttributesSettingsOk

`func (o *UpdateApplication) GetAttributesSettingsOk() (*AttributesSettings, bool)`

GetAttributesSettingsOk returns a tuple with the AttributesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesSettings

`func (o *UpdateApplication) SetAttributesSettings(v AttributesSettings)`

SetAttributesSettings sets AttributesSettings field to given value.

### HasAttributesSettings

`func (o *UpdateApplication) HasAttributesSettings() bool`

HasAttributesSettings returns a boolean if a field has been set.

### GetSandbox

`func (o *UpdateApplication) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *UpdateApplication) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *UpdateApplication) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *UpdateApplication) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetEnablePartialDiscounts

`func (o *UpdateApplication) GetEnablePartialDiscounts() bool`

GetEnablePartialDiscounts returns the EnablePartialDiscounts field if non-nil, zero value otherwise.

### GetEnablePartialDiscountsOk

`func (o *UpdateApplication) GetEnablePartialDiscountsOk() (*bool, bool)`

GetEnablePartialDiscountsOk returns a tuple with the EnablePartialDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartialDiscounts

`func (o *UpdateApplication) SetEnablePartialDiscounts(v bool)`

SetEnablePartialDiscounts sets EnablePartialDiscounts field to given value.

### HasEnablePartialDiscounts

`func (o *UpdateApplication) HasEnablePartialDiscounts() bool`

HasEnablePartialDiscounts returns a boolean if a field has been set.

### GetDefaultDiscountAdditionalCostPerItemScope

`func (o *UpdateApplication) GetDefaultDiscountAdditionalCostPerItemScope() string`

GetDefaultDiscountAdditionalCostPerItemScope returns the DefaultDiscountAdditionalCostPerItemScope field if non-nil, zero value otherwise.

### GetDefaultDiscountAdditionalCostPerItemScopeOk

`func (o *UpdateApplication) GetDefaultDiscountAdditionalCostPerItemScopeOk() (*string, bool)`

GetDefaultDiscountAdditionalCostPerItemScopeOk returns a tuple with the DefaultDiscountAdditionalCostPerItemScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountAdditionalCostPerItemScope

`func (o *UpdateApplication) SetDefaultDiscountAdditionalCostPerItemScope(v string)`

SetDefaultDiscountAdditionalCostPerItemScope sets DefaultDiscountAdditionalCostPerItemScope field to given value.

### HasDefaultDiscountAdditionalCostPerItemScope

`func (o *UpdateApplication) HasDefaultDiscountAdditionalCostPerItemScope() bool`

HasDefaultDiscountAdditionalCostPerItemScope returns a boolean if a field has been set.

### GetDefaultEvaluationGroupId

`func (o *UpdateApplication) GetDefaultEvaluationGroupId() int64`

GetDefaultEvaluationGroupId returns the DefaultEvaluationGroupId field if non-nil, zero value otherwise.

### GetDefaultEvaluationGroupIdOk

`func (o *UpdateApplication) GetDefaultEvaluationGroupIdOk() (*int64, bool)`

GetDefaultEvaluationGroupIdOk returns a tuple with the DefaultEvaluationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEvaluationGroupId

`func (o *UpdateApplication) SetDefaultEvaluationGroupId(v int64)`

SetDefaultEvaluationGroupId sets DefaultEvaluationGroupId field to given value.

### HasDefaultEvaluationGroupId

`func (o *UpdateApplication) HasDefaultEvaluationGroupId() bool`

HasDefaultEvaluationGroupId returns a boolean if a field has been set.

### GetDefaultCartItemFilterId

`func (o *UpdateApplication) GetDefaultCartItemFilterId() int64`

GetDefaultCartItemFilterId returns the DefaultCartItemFilterId field if non-nil, zero value otherwise.

### GetDefaultCartItemFilterIdOk

`func (o *UpdateApplication) GetDefaultCartItemFilterIdOk() (*int64, bool)`

GetDefaultCartItemFilterIdOk returns a tuple with the DefaultCartItemFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCartItemFilterId

`func (o *UpdateApplication) SetDefaultCartItemFilterId(v int64)`

SetDefaultCartItemFilterId sets DefaultCartItemFilterId field to given value.

### HasDefaultCartItemFilterId

`func (o *UpdateApplication) HasDefaultCartItemFilterId() bool`

HasDefaultCartItemFilterId returns a boolean if a field has been set.

### GetEnableCampaignStateManagement

`func (o *UpdateApplication) GetEnableCampaignStateManagement() bool`

GetEnableCampaignStateManagement returns the EnableCampaignStateManagement field if non-nil, zero value otherwise.

### GetEnableCampaignStateManagementOk

`func (o *UpdateApplication) GetEnableCampaignStateManagementOk() (*bool, bool)`

GetEnableCampaignStateManagementOk returns a tuple with the EnableCampaignStateManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCampaignStateManagement

`func (o *UpdateApplication) SetEnableCampaignStateManagement(v bool)`

SetEnableCampaignStateManagement sets EnableCampaignStateManagement field to given value.

### HasEnableCampaignStateManagement

`func (o *UpdateApplication) HasEnableCampaignStateManagement() bool`

HasEnableCampaignStateManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


