# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**Modified** | Pointer to [**time.Time**](time.Time.md) | The time this entity was last modified. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
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
**LoyaltyPrograms** | Pointer to [**[]LoyaltyProgram**](LoyaltyProgram.md) | An array containing all the loyalty programs to which this application is subscribed. | 

## Methods

### NewApplication

`func NewApplication(id int64, created time.Time, modified time.Time, accountId int64, name string, timezone string, currency string, loyaltyPrograms []LoyaltyProgram, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *Application) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Application) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Application) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetModified

`func (o *Application) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Application) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Application) SetModified(v time.Time)`

SetModified sets Modified field to given value.


### GetAccountId

`func (o *Application) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Application) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Application) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimezone

`func (o *Application) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Application) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Application) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetCurrency

`func (o *Application) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Application) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Application) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCaseSensitivity

`func (o *Application) GetCaseSensitivity() string`

GetCaseSensitivity returns the CaseSensitivity field if non-nil, zero value otherwise.

### GetCaseSensitivityOk

`func (o *Application) GetCaseSensitivityOk() (*string, bool)`

GetCaseSensitivityOk returns a tuple with the CaseSensitivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitivity

`func (o *Application) SetCaseSensitivity(v string)`

SetCaseSensitivity sets CaseSensitivity field to given value.

### HasCaseSensitivity

`func (o *Application) HasCaseSensitivity() bool`

HasCaseSensitivity returns a boolean if a field has been set.

### GetAttributes

`func (o *Application) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Application) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Application) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Application) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLimits

`func (o *Application) GetLimits() []LimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *Application) GetLimitsOk() (*[]LimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *Application) SetLimits(v []LimitConfig)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *Application) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetDefaultDiscountScope

`func (o *Application) GetDefaultDiscountScope() string`

GetDefaultDiscountScope returns the DefaultDiscountScope field if non-nil, zero value otherwise.

### GetDefaultDiscountScopeOk

`func (o *Application) GetDefaultDiscountScopeOk() (*string, bool)`

GetDefaultDiscountScopeOk returns a tuple with the DefaultDiscountScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountScope

`func (o *Application) SetDefaultDiscountScope(v string)`

SetDefaultDiscountScope sets DefaultDiscountScope field to given value.

### HasDefaultDiscountScope

`func (o *Application) HasDefaultDiscountScope() bool`

HasDefaultDiscountScope returns a boolean if a field has been set.

### GetEnableCascadingDiscounts

`func (o *Application) GetEnableCascadingDiscounts() bool`

GetEnableCascadingDiscounts returns the EnableCascadingDiscounts field if non-nil, zero value otherwise.

### GetEnableCascadingDiscountsOk

`func (o *Application) GetEnableCascadingDiscountsOk() (*bool, bool)`

GetEnableCascadingDiscountsOk returns a tuple with the EnableCascadingDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCascadingDiscounts

`func (o *Application) SetEnableCascadingDiscounts(v bool)`

SetEnableCascadingDiscounts sets EnableCascadingDiscounts field to given value.

### HasEnableCascadingDiscounts

`func (o *Application) HasEnableCascadingDiscounts() bool`

HasEnableCascadingDiscounts returns a boolean if a field has been set.

### GetEnableFlattenedCartItems

`func (o *Application) GetEnableFlattenedCartItems() bool`

GetEnableFlattenedCartItems returns the EnableFlattenedCartItems field if non-nil, zero value otherwise.

### GetEnableFlattenedCartItemsOk

`func (o *Application) GetEnableFlattenedCartItemsOk() (*bool, bool)`

GetEnableFlattenedCartItemsOk returns a tuple with the EnableFlattenedCartItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFlattenedCartItems

`func (o *Application) SetEnableFlattenedCartItems(v bool)`

SetEnableFlattenedCartItems sets EnableFlattenedCartItems field to given value.

### HasEnableFlattenedCartItems

`func (o *Application) HasEnableFlattenedCartItems() bool`

HasEnableFlattenedCartItems returns a boolean if a field has been set.

### GetAttributesSettings

`func (o *Application) GetAttributesSettings() AttributesSettings`

GetAttributesSettings returns the AttributesSettings field if non-nil, zero value otherwise.

### GetAttributesSettingsOk

`func (o *Application) GetAttributesSettingsOk() (*AttributesSettings, bool)`

GetAttributesSettingsOk returns a tuple with the AttributesSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributesSettings

`func (o *Application) SetAttributesSettings(v AttributesSettings)`

SetAttributesSettings sets AttributesSettings field to given value.

### HasAttributesSettings

`func (o *Application) HasAttributesSettings() bool`

HasAttributesSettings returns a boolean if a field has been set.

### GetSandbox

`func (o *Application) GetSandbox() bool`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *Application) GetSandboxOk() (*bool, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *Application) SetSandbox(v bool)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *Application) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.

### GetEnablePartialDiscounts

`func (o *Application) GetEnablePartialDiscounts() bool`

GetEnablePartialDiscounts returns the EnablePartialDiscounts field if non-nil, zero value otherwise.

### GetEnablePartialDiscountsOk

`func (o *Application) GetEnablePartialDiscountsOk() (*bool, bool)`

GetEnablePartialDiscountsOk returns a tuple with the EnablePartialDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePartialDiscounts

`func (o *Application) SetEnablePartialDiscounts(v bool)`

SetEnablePartialDiscounts sets EnablePartialDiscounts field to given value.

### HasEnablePartialDiscounts

`func (o *Application) HasEnablePartialDiscounts() bool`

HasEnablePartialDiscounts returns a boolean if a field has been set.

### GetDefaultDiscountAdditionalCostPerItemScope

`func (o *Application) GetDefaultDiscountAdditionalCostPerItemScope() string`

GetDefaultDiscountAdditionalCostPerItemScope returns the DefaultDiscountAdditionalCostPerItemScope field if non-nil, zero value otherwise.

### GetDefaultDiscountAdditionalCostPerItemScopeOk

`func (o *Application) GetDefaultDiscountAdditionalCostPerItemScopeOk() (*string, bool)`

GetDefaultDiscountAdditionalCostPerItemScopeOk returns a tuple with the DefaultDiscountAdditionalCostPerItemScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDiscountAdditionalCostPerItemScope

`func (o *Application) SetDefaultDiscountAdditionalCostPerItemScope(v string)`

SetDefaultDiscountAdditionalCostPerItemScope sets DefaultDiscountAdditionalCostPerItemScope field to given value.

### HasDefaultDiscountAdditionalCostPerItemScope

`func (o *Application) HasDefaultDiscountAdditionalCostPerItemScope() bool`

HasDefaultDiscountAdditionalCostPerItemScope returns a boolean if a field has been set.

### GetDefaultEvaluationGroupId

`func (o *Application) GetDefaultEvaluationGroupId() int64`

GetDefaultEvaluationGroupId returns the DefaultEvaluationGroupId field if non-nil, zero value otherwise.

### GetDefaultEvaluationGroupIdOk

`func (o *Application) GetDefaultEvaluationGroupIdOk() (*int64, bool)`

GetDefaultEvaluationGroupIdOk returns a tuple with the DefaultEvaluationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEvaluationGroupId

`func (o *Application) SetDefaultEvaluationGroupId(v int64)`

SetDefaultEvaluationGroupId sets DefaultEvaluationGroupId field to given value.

### HasDefaultEvaluationGroupId

`func (o *Application) HasDefaultEvaluationGroupId() bool`

HasDefaultEvaluationGroupId returns a boolean if a field has been set.

### GetDefaultCartItemFilterId

`func (o *Application) GetDefaultCartItemFilterId() int64`

GetDefaultCartItemFilterId returns the DefaultCartItemFilterId field if non-nil, zero value otherwise.

### GetDefaultCartItemFilterIdOk

`func (o *Application) GetDefaultCartItemFilterIdOk() (*int64, bool)`

GetDefaultCartItemFilterIdOk returns a tuple with the DefaultCartItemFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCartItemFilterId

`func (o *Application) SetDefaultCartItemFilterId(v int64)`

SetDefaultCartItemFilterId sets DefaultCartItemFilterId field to given value.

### HasDefaultCartItemFilterId

`func (o *Application) HasDefaultCartItemFilterId() bool`

HasDefaultCartItemFilterId returns a boolean if a field has been set.

### GetEnableCampaignStateManagement

`func (o *Application) GetEnableCampaignStateManagement() bool`

GetEnableCampaignStateManagement returns the EnableCampaignStateManagement field if non-nil, zero value otherwise.

### GetEnableCampaignStateManagementOk

`func (o *Application) GetEnableCampaignStateManagementOk() (*bool, bool)`

GetEnableCampaignStateManagementOk returns a tuple with the EnableCampaignStateManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCampaignStateManagement

`func (o *Application) SetEnableCampaignStateManagement(v bool)`

SetEnableCampaignStateManagement sets EnableCampaignStateManagement field to given value.

### HasEnableCampaignStateManagement

`func (o *Application) HasEnableCampaignStateManagement() bool`

HasEnableCampaignStateManagement returns a boolean if a field has been set.

### GetLoyaltyPrograms

`func (o *Application) GetLoyaltyPrograms() []LoyaltyProgram`

GetLoyaltyPrograms returns the LoyaltyPrograms field if non-nil, zero value otherwise.

### GetLoyaltyProgramsOk

`func (o *Application) GetLoyaltyProgramsOk() (*[]LoyaltyProgram, bool)`

GetLoyaltyProgramsOk returns a tuple with the LoyaltyPrograms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoyaltyPrograms

`func (o *Application) SetLoyaltyPrograms(v []LoyaltyProgram)`

SetLoyaltyPrograms sets LoyaltyPrograms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


