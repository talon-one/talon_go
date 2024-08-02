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
)

// TalangAttribute struct for TalangAttribute
type TalangAttribute struct {
	// The number of campaigns that refer to the attribute.
	CampaignsCount int32 `json:"campaignsCount"`
	// A description of the attribute.
	Description *string `json:"description,omitempty"`
	// The name of the entity of the attribute.
	Entity *string `json:"entity,omitempty"`
	// Examples of values that can be assigned to the attribute.
	ExampleValue *[]string `json:"exampleValue,omitempty"`
	// Indicate the kind of the attribute.
	Kind string `json:"kind"`
	// The attribute name that will be used in API requests and Talang. E.g. if `name == \"region\"` then you would set the region attribute by including an `attributes.region` property in your request payload. 
	Name string `json:"name"`
	// The name of the attribute that is displayed to the user in the Campaign Manager.
	Title *string `json:"title,omitempty"`
	// The talang type of the attribute.
	Type string `json:"type"`
	// Indicates whether the attribute is visible in the UI or not.
	Visible bool `json:"visible"`
}

// GetCampaignsCount returns the CampaignsCount field value
func (o *TalangAttribute) GetCampaignsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignsCount
}

// SetCampaignsCount sets field value
func (o *TalangAttribute) SetCampaignsCount(v int32) {
	o.CampaignsCount = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *TalangAttribute) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TalangAttribute) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *TalangAttribute) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *TalangAttribute) SetDescription(v string) {
	o.Description = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *TalangAttribute) GetEntity() string {
	if o == nil || o.Entity == nil {
		var ret string
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TalangAttribute) GetEntityOk() (string, bool) {
	if o == nil || o.Entity == nil {
		var ret string
		return ret, false
	}
	return *o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *TalangAttribute) HasEntity() bool {
	if o != nil && o.Entity != nil {
		return true
	}

	return false
}

// SetEntity gets a reference to the given string and assigns it to the Entity field.
func (o *TalangAttribute) SetEntity(v string) {
	o.Entity = &v
}

// GetExampleValue returns the ExampleValue field value if set, zero value otherwise.
func (o *TalangAttribute) GetExampleValue() []string {
	if o == nil || o.ExampleValue == nil {
		var ret []string
		return ret
	}
	return *o.ExampleValue
}

// GetExampleValueOk returns a tuple with the ExampleValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TalangAttribute) GetExampleValueOk() ([]string, bool) {
	if o == nil || o.ExampleValue == nil {
		var ret []string
		return ret, false
	}
	return *o.ExampleValue, true
}

// HasExampleValue returns a boolean if a field has been set.
func (o *TalangAttribute) HasExampleValue() bool {
	if o != nil && o.ExampleValue != nil {
		return true
	}

	return false
}

// SetExampleValue gets a reference to the given []string and assigns it to the ExampleValue field.
func (o *TalangAttribute) SetExampleValue(v []string) {
	o.ExampleValue = &v
}

// GetKind returns the Kind field value
func (o *TalangAttribute) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// SetKind sets field value
func (o *TalangAttribute) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *TalangAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *TalangAttribute) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TalangAttribute) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *TalangAttribute) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TalangAttribute) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TalangAttribute) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value
func (o *TalangAttribute) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *TalangAttribute) SetType(v string) {
	o.Type = v
}

// GetVisible returns the Visible field value
func (o *TalangAttribute) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// SetVisible sets field value
func (o *TalangAttribute) SetVisible(v bool) {
	o.Visible = v
}

type NullableTalangAttribute struct {
	Value TalangAttribute
	ExplicitNull bool
}

func (v NullableTalangAttribute) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableTalangAttribute) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
