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

// LibraryAttribute struct for LibraryAttribute
type LibraryAttribute struct {
	// The name of the entity that can have this attribute. When creating or updating the entities of a given type, you can include an `attributes` object with keys corresponding to the `name` of the custom attributes for that type.
	Entity string `json:"entity"`
	// The attribute name that will be used in API requests and Talang. E.g. if `name == \"region\"` then you would set the region attribute by including an `attributes.region` property in your request payload.
	Name string `json:"name"`
	// The human-readable name for the attribute that will be shown in the Campaign Manager. Like `name`, the combination of entity and title must also be unique.
	Title string `json:"title"`
	// The data type of the attribute, a `time` attribute must be sent as a string that conforms to the [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) timestamp format.
	Type string `json:"type"`
	// A description of the attribute.
	Description string `json:"description"`
	// The presets that indicate to which industry the attribute applies to.
	Presets []string `json:"presets"`
	// Short suggestions that are used to group attributes.
	Suggestions []string `json:"suggestions"`
}

// GetEntity returns the Entity field value
func (o *LibraryAttribute) GetEntity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Entity
}

// SetEntity sets field value
func (o *LibraryAttribute) SetEntity(v string) {
	o.Entity = v
}

// GetName returns the Name field value
func (o *LibraryAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *LibraryAttribute) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value
func (o *LibraryAttribute) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *LibraryAttribute) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *LibraryAttribute) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *LibraryAttribute) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *LibraryAttribute) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *LibraryAttribute) SetDescription(v string) {
	o.Description = v
}

// GetPresets returns the Presets field value
func (o *LibraryAttribute) GetPresets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Presets
}

// SetPresets sets field value
func (o *LibraryAttribute) SetPresets(v []string) {
	o.Presets = v
}

// GetSuggestions returns the Suggestions field value
func (o *LibraryAttribute) GetSuggestions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Suggestions
}

// SetSuggestions sets field value
func (o *LibraryAttribute) SetSuggestions(v []string) {
	o.Suggestions = v
}

type NullableLibraryAttribute struct {
	Value        LibraryAttribute
	ExplicitNull bool
}

func (v NullableLibraryAttribute) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLibraryAttribute) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
