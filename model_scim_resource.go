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

// ScimResource Resource definition for the SCIM provisioning protocol.
type ScimResource struct {
	// ID of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// Human-readable description of the resource.
	Description *string `json:"description,omitempty"`
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ScimResource) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimResource) GetIdOk() (string, bool) {
	if o == nil || o.Id == nil {
		var ret string
		return ret, false
	}
	return *o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ScimResource) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ScimResource) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScimResource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimResource) GetNameOk() (string, bool) {
	if o == nil || o.Name == nil {
		var ret string
		return ret, false
	}
	return *o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScimResource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScimResource) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ScimResource) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimResource) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ScimResource) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ScimResource) SetDescription(v string) {
	o.Description = &v
}

type NullableScimResource struct {
	Value        ScimResource
	ExplicitNull bool
}

func (v NullableScimResource) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableScimResource) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
