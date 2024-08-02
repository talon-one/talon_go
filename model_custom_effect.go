/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
	"time"
)

// CustomEffect struct for CustomEffect
type CustomEffect struct {
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The IDs of the Applications that are related to this entity.
	ApplicationIds []int32 `json:"applicationIds"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// ID of the user who created this effect.
	CreatedBy int32 `json:"createdBy"`
	// The description of this effect.
	Description *string `json:"description,omitempty"`
	// Determines if this effect is active.
	Enabled bool `json:"enabled"`
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// Indicates if this effect is per item or not.
	IsPerItem *bool `json:"isPerItem,omitempty"`
	// The time this entity was last modified.
	Modified time.Time `json:"modified"`
	// ID of the user who last updated this effect if available.
	ModifiedBy *int32 `json:"modifiedBy,omitempty"`
	// The name of this effect.
	Name string `json:"name"`
	// Array of template argument definitions.
	Params *[]TemplateArgDef `json:"params,omitempty"`
	// The JSON payload of this effect.
	Payload string `json:"payload"`
	// The title of this effect.
	Title string `json:"title"`
}

// GetAccountId returns the AccountId field value
func (o *CustomEffect) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *CustomEffect) SetAccountId(v int32) {
	o.AccountId = v
}

// GetApplicationIds returns the ApplicationIds field value
func (o *CustomEffect) GetApplicationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationIds
}

// SetApplicationIds sets field value
func (o *CustomEffect) SetApplicationIds(v []int32) {
	o.ApplicationIds = v
}

// GetCreated returns the Created field value
func (o *CustomEffect) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CustomEffect) SetCreated(v time.Time) {
	o.Created = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CustomEffect) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// SetCreatedBy sets field value
func (o *CustomEffect) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomEffect) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomEffect) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomEffect) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *CustomEffect) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// SetEnabled sets field value
func (o *CustomEffect) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value
func (o *CustomEffect) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *CustomEffect) SetId(v int32) {
	o.Id = v
}

// GetIsPerItem returns the IsPerItem field value if set, zero value otherwise.
func (o *CustomEffect) GetIsPerItem() bool {
	if o == nil || o.IsPerItem == nil {
		var ret bool
		return ret
	}
	return *o.IsPerItem
}

// GetIsPerItemOk returns a tuple with the IsPerItem field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetIsPerItemOk() (bool, bool) {
	if o == nil || o.IsPerItem == nil {
		var ret bool
		return ret, false
	}
	return *o.IsPerItem, true
}

// HasIsPerItem returns a boolean if a field has been set.
func (o *CustomEffect) HasIsPerItem() bool {
	if o != nil && o.IsPerItem != nil {
		return true
	}

	return false
}

// SetIsPerItem gets a reference to the given bool and assigns it to the IsPerItem field.
func (o *CustomEffect) SetIsPerItem(v bool) {
	o.IsPerItem = &v
}

// GetModified returns the Modified field value
func (o *CustomEffect) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *CustomEffect) SetModified(v time.Time) {
	o.Modified = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *CustomEffect) GetModifiedBy() int32 {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetModifiedByOk() (int32, bool) {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret, false
	}
	return *o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *CustomEffect) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.
func (o *CustomEffect) SetModifiedBy(v int32) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value
func (o *CustomEffect) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CustomEffect) SetName(v string) {
	o.Name = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *CustomEffect) GetParams() []TemplateArgDef {
	if o == nil || o.Params == nil {
		var ret []TemplateArgDef
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CustomEffect) GetParamsOk() ([]TemplateArgDef, bool) {
	if o == nil || o.Params == nil {
		var ret []TemplateArgDef
		return ret, false
	}
	return *o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *CustomEffect) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given []TemplateArgDef and assigns it to the Params field.
func (o *CustomEffect) SetParams(v []TemplateArgDef) {
	o.Params = &v
}

// GetPayload returns the Payload field value
func (o *CustomEffect) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// SetPayload sets field value
func (o *CustomEffect) SetPayload(v string) {
	o.Payload = v
}

// GetTitle returns the Title field value
func (o *CustomEffect) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *CustomEffect) SetTitle(v string) {
	o.Title = v
}

type NullableCustomEffect struct {
	Value CustomEffect
	ExplicitNull bool
}

func (v NullableCustomEffect) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomEffect) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
