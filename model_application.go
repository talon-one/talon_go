/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
	"time"
)

// Application
type Application struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The name of this application.
	Name string `json:"name"`
	// A longer description of the application.
	Description *string `json:"description,omitempty"`
	// A string containing an IANA timezone descriptor.
	Timezone string `json:"timezone"`
	// A string describing a default currency for new customer sessions.
	Currency string `json:"currency"`
	// A string indicating how should campaigns in this application deal with case sensitivity on coupon codes.
	CaseSensitivity *string `json:"caseSensitivity,omitempty"`
	// Arbitrary properties associated with this campaign
	Attributes *map[string]interface{} `json:"attributes,omitempty"`
	// Default limits for campaigns created in this application
	Limits *[]LimitConfig `json:"limits,omitempty"`
	// Default priority for campaigns created in this application, can be one of (universal, stackable, exclusive). If no value is provided, this is set to \"universal\"
	CampaignPriority *string `json:"campaignPriority,omitempty"`
	// The strategy used when choosing exclusive campaigns for evaluation, can be one of (listOrder, lowestDiscount, highestDiscount). If no value is provided, this is set to \"listOrder\"
	ExclusiveCampaignsStrategy *string `json:"exclusiveCampaignsStrategy,omitempty"`
	// The default scope to apply \"setDiscount\" effects on if no scope was provided with the effect.
	DefaultDiscountScope *string `json:"defaultDiscountScope,omitempty"`
	// Flag indicating if discounts should cascade for this application
	EnableCascadingDiscounts *bool `json:"enableCascadingDiscounts,omitempty"`
	// Flag indicating if cart items of quantity larger than one should be separated into different items of quantity one
	EnableFlattenedCartItems *bool               `json:"enableFlattenedCartItems,omitempty"`
	AttributesSettings       *AttributesSettings `json:"attributesSettings,omitempty"`
	// Flag indicating if this is a live or sandbox application
	Sandbox *bool `json:"sandbox,omitempty"`
	// An array containing all the loyalty programs to which this application is subscribed
	LoyaltyPrograms []LoyaltyProgram `json:"loyaltyPrograms"`
}

// GetId returns the Id field value
func (o *Application) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Application) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Application) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Application) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *Application) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *Application) SetModified(v time.Time) {
	o.Modified = v
}

// GetAccountId returns the AccountId field value
func (o *Application) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *Application) SetAccountId(v int32) {
	o.AccountId = v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Application) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Application) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Application) SetDescription(v string) {
	o.Description = &v
}

// GetTimezone returns the Timezone field value
func (o *Application) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// SetTimezone sets field value
func (o *Application) SetTimezone(v string) {
	o.Timezone = v
}

// GetCurrency returns the Currency field value
func (o *Application) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// SetCurrency sets field value
func (o *Application) SetCurrency(v string) {
	o.Currency = v
}

// GetCaseSensitivity returns the CaseSensitivity field value if set, zero value otherwise.
func (o *Application) GetCaseSensitivity() string {
	if o == nil || o.CaseSensitivity == nil {
		var ret string
		return ret
	}
	return *o.CaseSensitivity
}

// GetCaseSensitivityOk returns a tuple with the CaseSensitivity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCaseSensitivityOk() (string, bool) {
	if o == nil || o.CaseSensitivity == nil {
		var ret string
		return ret, false
	}
	return *o.CaseSensitivity, true
}

// HasCaseSensitivity returns a boolean if a field has been set.
func (o *Application) HasCaseSensitivity() bool {
	if o != nil && o.CaseSensitivity != nil {
		return true
	}

	return false
}

// SetCaseSensitivity gets a reference to the given string and assigns it to the CaseSensitivity field.
func (o *Application) SetCaseSensitivity(v string) {
	o.CaseSensitivity = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Application) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Application) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Application) SetAttributes(v map[string]interface{}) {
	o.Attributes = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *Application) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret, false
	}
	return *o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *Application) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *Application) SetLimits(v []LimitConfig) {
	o.Limits = &v
}

// GetCampaignPriority returns the CampaignPriority field value if set, zero value otherwise.
func (o *Application) GetCampaignPriority() string {
	if o == nil || o.CampaignPriority == nil {
		var ret string
		return ret
	}
	return *o.CampaignPriority
}

// GetCampaignPriorityOk returns a tuple with the CampaignPriority field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCampaignPriorityOk() (string, bool) {
	if o == nil || o.CampaignPriority == nil {
		var ret string
		return ret, false
	}
	return *o.CampaignPriority, true
}

// HasCampaignPriority returns a boolean if a field has been set.
func (o *Application) HasCampaignPriority() bool {
	if o != nil && o.CampaignPriority != nil {
		return true
	}

	return false
}

// SetCampaignPriority gets a reference to the given string and assigns it to the CampaignPriority field.
func (o *Application) SetCampaignPriority(v string) {
	o.CampaignPriority = &v
}

// GetExclusiveCampaignsStrategy returns the ExclusiveCampaignsStrategy field value if set, zero value otherwise.
func (o *Application) GetExclusiveCampaignsStrategy() string {
	if o == nil || o.ExclusiveCampaignsStrategy == nil {
		var ret string
		return ret
	}
	return *o.ExclusiveCampaignsStrategy
}

// GetExclusiveCampaignsStrategyOk returns a tuple with the ExclusiveCampaignsStrategy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetExclusiveCampaignsStrategyOk() (string, bool) {
	if o == nil || o.ExclusiveCampaignsStrategy == nil {
		var ret string
		return ret, false
	}
	return *o.ExclusiveCampaignsStrategy, true
}

// HasExclusiveCampaignsStrategy returns a boolean if a field has been set.
func (o *Application) HasExclusiveCampaignsStrategy() bool {
	if o != nil && o.ExclusiveCampaignsStrategy != nil {
		return true
	}

	return false
}

// SetExclusiveCampaignsStrategy gets a reference to the given string and assigns it to the ExclusiveCampaignsStrategy field.
func (o *Application) SetExclusiveCampaignsStrategy(v string) {
	o.ExclusiveCampaignsStrategy = &v
}

// GetDefaultDiscountScope returns the DefaultDiscountScope field value if set, zero value otherwise.
func (o *Application) GetDefaultDiscountScope() string {
	if o == nil || o.DefaultDiscountScope == nil {
		var ret string
		return ret
	}
	return *o.DefaultDiscountScope
}

// GetDefaultDiscountScopeOk returns a tuple with the DefaultDiscountScope field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDefaultDiscountScopeOk() (string, bool) {
	if o == nil || o.DefaultDiscountScope == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultDiscountScope, true
}

// HasDefaultDiscountScope returns a boolean if a field has been set.
func (o *Application) HasDefaultDiscountScope() bool {
	if o != nil && o.DefaultDiscountScope != nil {
		return true
	}

	return false
}

// SetDefaultDiscountScope gets a reference to the given string and assigns it to the DefaultDiscountScope field.
func (o *Application) SetDefaultDiscountScope(v string) {
	o.DefaultDiscountScope = &v
}

// GetEnableCascadingDiscounts returns the EnableCascadingDiscounts field value if set, zero value otherwise.
func (o *Application) GetEnableCascadingDiscounts() bool {
	if o == nil || o.EnableCascadingDiscounts == nil {
		var ret bool
		return ret
	}
	return *o.EnableCascadingDiscounts
}

// GetEnableCascadingDiscountsOk returns a tuple with the EnableCascadingDiscounts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetEnableCascadingDiscountsOk() (bool, bool) {
	if o == nil || o.EnableCascadingDiscounts == nil {
		var ret bool
		return ret, false
	}
	return *o.EnableCascadingDiscounts, true
}

// HasEnableCascadingDiscounts returns a boolean if a field has been set.
func (o *Application) HasEnableCascadingDiscounts() bool {
	if o != nil && o.EnableCascadingDiscounts != nil {
		return true
	}

	return false
}

// SetEnableCascadingDiscounts gets a reference to the given bool and assigns it to the EnableCascadingDiscounts field.
func (o *Application) SetEnableCascadingDiscounts(v bool) {
	o.EnableCascadingDiscounts = &v
}

// GetEnableFlattenedCartItems returns the EnableFlattenedCartItems field value if set, zero value otherwise.
func (o *Application) GetEnableFlattenedCartItems() bool {
	if o == nil || o.EnableFlattenedCartItems == nil {
		var ret bool
		return ret
	}
	return *o.EnableFlattenedCartItems
}

// GetEnableFlattenedCartItemsOk returns a tuple with the EnableFlattenedCartItems field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetEnableFlattenedCartItemsOk() (bool, bool) {
	if o == nil || o.EnableFlattenedCartItems == nil {
		var ret bool
		return ret, false
	}
	return *o.EnableFlattenedCartItems, true
}

// HasEnableFlattenedCartItems returns a boolean if a field has been set.
func (o *Application) HasEnableFlattenedCartItems() bool {
	if o != nil && o.EnableFlattenedCartItems != nil {
		return true
	}

	return false
}

// SetEnableFlattenedCartItems gets a reference to the given bool and assigns it to the EnableFlattenedCartItems field.
func (o *Application) SetEnableFlattenedCartItems(v bool) {
	o.EnableFlattenedCartItems = &v
}

// GetAttributesSettings returns the AttributesSettings field value if set, zero value otherwise.
func (o *Application) GetAttributesSettings() AttributesSettings {
	if o == nil || o.AttributesSettings == nil {
		var ret AttributesSettings
		return ret
	}
	return *o.AttributesSettings
}

// GetAttributesSettingsOk returns a tuple with the AttributesSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAttributesSettingsOk() (AttributesSettings, bool) {
	if o == nil || o.AttributesSettings == nil {
		var ret AttributesSettings
		return ret, false
	}
	return *o.AttributesSettings, true
}

// HasAttributesSettings returns a boolean if a field has been set.
func (o *Application) HasAttributesSettings() bool {
	if o != nil && o.AttributesSettings != nil {
		return true
	}

	return false
}

// SetAttributesSettings gets a reference to the given AttributesSettings and assigns it to the AttributesSettings field.
func (o *Application) SetAttributesSettings(v AttributesSettings) {
	o.AttributesSettings = &v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *Application) GetSandbox() bool {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetSandboxOk() (bool, bool) {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret, false
	}
	return *o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *Application) HasSandbox() bool {
	if o != nil && o.Sandbox != nil {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *Application) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetLoyaltyPrograms returns the LoyaltyPrograms field value
func (o *Application) GetLoyaltyPrograms() []LoyaltyProgram {
	if o == nil {
		var ret []LoyaltyProgram
		return ret
	}

	return o.LoyaltyPrograms
}

// SetLoyaltyPrograms sets field value
func (o *Application) SetLoyaltyPrograms(v []LoyaltyProgram) {
	o.LoyaltyPrograms = v
}

type NullableApplication struct {
	Value        Application
	ExplicitNull bool
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
