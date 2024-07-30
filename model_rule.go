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

// Rule struct for Rule
type Rule struct {
	// A unique identifier for the rule.
	Id *string `json:"id,omitempty"`
	// The ID of the rule that was copied to create this rule.
	ParentId *string `json:"parentId,omitempty"`
	// A short description of the rule.
	Title string `json:"title"`
	// A longer, more detailed description of the rule.
	Description *string `json:"description,omitempty"`
	// An array that provides objects with variable names (name) and talang expressions to whose result they are bound (expression) during rule evaluation. The order of the evaluation is decided by the position in the array.
	Bindings *[]Binding `json:"bindings,omitempty"`
	// A Talang expression that will be evaluated in the context of the given event.
	Condition []map[string]interface{} `json:"condition"`
	// An array of effectful Talang expressions in arrays that will be evaluated when a rule matches.
	Effects []map[string]interface{} `json:"effects"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Rule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Rule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Rule) SetId(v string) {
	o.Id = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *Rule) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetParentIdOk() (string, bool) {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret, false
	}
	return *o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *Rule) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *Rule) SetParentId(v string) {
	o.ParentId = &v
}

// GetTitle returns the Title field value
func (o *Rule) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *Rule) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Rule) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Rule) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Rule) SetDescription(v string) {
	o.Description = &v
}

// GetBindings returns the Bindings field value if set, zero value otherwise.
func (o *Rule) GetBindings() []Binding {
	if o == nil || o.Bindings == nil {
		var ret []Binding
		return ret
	}
	return *o.Bindings
}

// GetBindingsOk returns a tuple with the Bindings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Rule) GetBindingsOk() ([]Binding, bool) {
	if o == nil || o.Bindings == nil {
		var ret []Binding
		return ret, false
	}
	return *o.Bindings, true
}

// HasBindings returns a boolean if a field has been set.
func (o *Rule) HasBindings() bool {
	if o != nil && o.Bindings != nil {
		return true
	}

	return false
}

// SetBindings gets a reference to the given []Binding and assigns it to the Bindings field.
func (o *Rule) SetBindings(v []Binding) {
	o.Bindings = &v
}

// GetCondition returns the Condition field value
func (o *Rule) GetCondition() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Condition
}

// SetCondition sets field value
func (o *Rule) SetCondition(v []map[string]interface{}) {
	o.Condition = v
}

// GetEffects returns the Effects field value
func (o *Rule) GetEffects() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Effects
}

// SetEffects sets field value
func (o *Rule) SetEffects(v []map[string]interface{}) {
	o.Effects = v
}

type NullableRule struct {
	Value Rule
	ExplicitNull bool
}

func (v NullableRule) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRule) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
