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
)

// NewTemplateDef struct for NewTemplateDef
type NewTemplateDef struct {
	// Campaigner-friendly name for the template that will be shown in the rule editor.
	Title string `json:"title"`
	// A short description of the template that will be shown in the rule editor.
	Description *string `json:"description,omitempty"`
	// Extended help text for the template.
	Help *string `json:"help,omitempty"`
	// Used for grouping templates in the rule editor sidebar.
	Category string `json:"category"`
	// A Talang expression that contains variable bindings referring to args.
	Expr []map[string]interface{} `json:"expr"`
	// An array of argument definitions.
	Args []TemplateArgDef `json:"args"`
	// A flag to control exposure in Rule Builder.
	Expose *bool `json:"expose,omitempty"`
}

// GetTitle returns the Title field value
func (o *NewTemplateDef) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *NewTemplateDef) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NewTemplateDef) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewTemplateDef) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NewTemplateDef) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NewTemplateDef) SetDescription(v string) {
	o.Description = &v
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *NewTemplateDef) GetHelp() string {
	if o == nil || o.Help == nil {
		var ret string
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewTemplateDef) GetHelpOk() (string, bool) {
	if o == nil || o.Help == nil {
		var ret string
		return ret, false
	}
	return *o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *NewTemplateDef) HasHelp() bool {
	if o != nil && o.Help != nil {
		return true
	}

	return false
}

// SetHelp gets a reference to the given string and assigns it to the Help field.
func (o *NewTemplateDef) SetHelp(v string) {
	o.Help = &v
}

// GetCategory returns the Category field value
func (o *NewTemplateDef) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// SetCategory sets field value
func (o *NewTemplateDef) SetCategory(v string) {
	o.Category = v
}

// GetExpr returns the Expr field value
func (o *NewTemplateDef) GetExpr() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}

	return o.Expr
}

// SetExpr sets field value
func (o *NewTemplateDef) SetExpr(v []map[string]interface{}) {
	o.Expr = v
}

// GetArgs returns the Args field value
func (o *NewTemplateDef) GetArgs() []TemplateArgDef {
	if o == nil {
		var ret []TemplateArgDef
		return ret
	}

	return o.Args
}

// SetArgs sets field value
func (o *NewTemplateDef) SetArgs(v []TemplateArgDef) {
	o.Args = v
}

// GetExpose returns the Expose field value if set, zero value otherwise.
func (o *NewTemplateDef) GetExpose() bool {
	if o == nil || o.Expose == nil {
		var ret bool
		return ret
	}
	return *o.Expose
}

// GetExposeOk returns a tuple with the Expose field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewTemplateDef) GetExposeOk() (bool, bool) {
	if o == nil || o.Expose == nil {
		var ret bool
		return ret, false
	}
	return *o.Expose, true
}

// HasExpose returns a boolean if a field has been set.
func (o *NewTemplateDef) HasExpose() bool {
	if o != nil && o.Expose != nil {
		return true
	}

	return false
}

// SetExpose gets a reference to the given bool and assigns it to the Expose field.
func (o *NewTemplateDef) SetExpose(v bool) {
	o.Expose = &v
}

type NullableNewTemplateDef struct {
	Value        NewTemplateDef
	ExplicitNull bool
}

func (v NullableNewTemplateDef) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewTemplateDef) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
