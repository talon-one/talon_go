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

// ScimServiceProviderConfigResponse Service provider configuration details.
type ScimServiceProviderConfigResponse struct {
	Bulk           *ScimServiceProviderConfigResponseBulk           `json:"bulk,omitempty"`
	ChangePassword *ScimServiceProviderConfigResponseChangePassword `json:"changePassword,omitempty"`
	// The URI that points to the SCIM service provider's documentation, providing further details about the service's capabilities and usage.
	DocumentationUri *string                                  `json:"documentationUri,omitempty"`
	Filter           *ScimServiceProviderConfigResponseFilter `json:"filter,omitempty"`
	Patch            *ScimServiceProviderConfigResponsePatch  `json:"patch,omitempty"`
	// A list of SCIM schemas that define the structure and data types supported by the service provider.
	Schemas *[]string                              `json:"schemas,omitempty"`
	Sort    *ScimServiceProviderConfigResponseSort `json:"sort,omitempty"`
}

// GetBulk returns the Bulk field value if set, zero value otherwise.
func (o *ScimServiceProviderConfigResponse) GetBulk() ScimServiceProviderConfigResponseBulk {
	if o == nil || o.Bulk == nil {
		var ret ScimServiceProviderConfigResponseBulk
		return ret
	}
	return *o.Bulk
}

// GetBulkOk returns a tuple with the Bulk field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigResponse) GetBulkOk() (ScimServiceProviderConfigResponseBulk, bool) {
	if o == nil || o.Bulk == nil {
		var ret ScimServiceProviderConfigResponseBulk
		return ret, false
	}
	return *o.Bulk, true
}

// HasBulk returns a boolean if a field has been set.
func (o *ScimServiceProviderConfigResponse) HasBulk() bool {
	if o != nil && o.Bulk != nil {
		return true
	}

	return false
}

// SetBulk gets a reference to the given ScimServiceProviderConfigResponseBulk and assigns it to the Bulk field.
func (o *ScimServiceProviderConfigResponse) SetBulk(v ScimServiceProviderConfigResponseBulk) {
	o.Bulk = &v
}

// GetChangePassword returns the ChangePassword field value if set, zero value otherwise.
func (o *ScimServiceProviderConfigResponse) GetChangePassword() ScimServiceProviderConfigResponseChangePassword {
	if o == nil || o.ChangePassword == nil {
		var ret ScimServiceProviderConfigResponseChangePassword
		return ret
	}
	return *o.ChangePassword
}

// GetChangePasswordOk returns a tuple with the ChangePassword field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigResponse) GetChangePasswordOk() (ScimServiceProviderConfigResponseChangePassword, bool) {
	if o == nil || o.ChangePassword == nil {
		var ret ScimServiceProviderConfigResponseChangePassword
		return ret, false
	}
	return *o.ChangePassword, true
}

// HasChangePassword returns a boolean if a field has been set.
func (o *ScimServiceProviderConfigResponse) HasChangePassword() bool {
	if o != nil && o.ChangePassword != nil {
		return true
	}

	return false
}

// SetChangePassword gets a reference to the given ScimServiceProviderConfigResponseChangePassword and assigns it to the ChangePassword field.
func (o *ScimServiceProviderConfigResponse) SetChangePassword(v ScimServiceProviderConfigResponseChangePassword) {
	o.ChangePassword = &v
}

// GetDocumentationUri returns the DocumentationUri field value if set, zero value otherwise.
func (o *ScimServiceProviderConfigResponse) GetDocumentationUri() string {
	if o == nil || o.DocumentationUri == nil {
		var ret string
		return ret
	}
	return *o.DocumentationUri
}

// GetDocumentationUriOk returns a tuple with the DocumentationUri field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigResponse) GetDocumentationUriOk() (string, bool) {
	if o == nil || o.DocumentationUri == nil {
		var ret string
		return ret, false
	}
	return *o.DocumentationUri, true
}

// HasDocumentationUri returns a boolean if a field has been set.
func (o *ScimServiceProviderConfigResponse) HasDocumentationUri() bool {
	if o != nil && o.DocumentationUri != nil {
		return true
	}

	return false
}

// SetDocumentationUri gets a reference to the given string and assigns it to the DocumentationUri field.
func (o *ScimServiceProviderConfigResponse) SetDocumentationUri(v string) {
	o.DocumentationUri = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ScimServiceProviderConfigResponse) GetFilter() ScimServiceProviderConfigResponseFilter {
	if o == nil || o.Filter == nil {
		var ret ScimServiceProviderConfigResponseFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigResponse) GetFilterOk() (ScimServiceProviderConfigResponseFilter, bool) {
	if o == nil || o.Filter == nil {
		var ret ScimServiceProviderConfigResponseFilter
		return ret, false
	}
	return *o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ScimServiceProviderConfigResponse) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ScimServiceProviderConfigResponseFilter and assigns it to the Filter field.
func (o *ScimServiceProviderConfigResponse) SetFilter(v ScimServiceProviderConfigResponseFilter) {
	o.Filter = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *ScimServiceProviderConfigResponse) GetPatch() ScimServiceProviderConfigResponsePatch {
	if o == nil || o.Patch == nil {
		var ret ScimServiceProviderConfigResponsePatch
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigResponse) GetPatchOk() (ScimServiceProviderConfigResponsePatch, bool) {
	if o == nil || o.Patch == nil {
		var ret ScimServiceProviderConfigResponsePatch
		return ret, false
	}
	return *o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *ScimServiceProviderConfigResponse) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given ScimServiceProviderConfigResponsePatch and assigns it to the Patch field.
func (o *ScimServiceProviderConfigResponse) SetPatch(v ScimServiceProviderConfigResponsePatch) {
	o.Patch = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *ScimServiceProviderConfigResponse) GetSchemas() []string {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret
	}
	return *o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigResponse) GetSchemasOk() ([]string, bool) {
	if o == nil || o.Schemas == nil {
		var ret []string
		return ret, false
	}
	return *o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *ScimServiceProviderConfigResponse) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given []string and assigns it to the Schemas field.
func (o *ScimServiceProviderConfigResponse) SetSchemas(v []string) {
	o.Schemas = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ScimServiceProviderConfigResponse) GetSort() ScimServiceProviderConfigResponseSort {
	if o == nil || o.Sort == nil {
		var ret ScimServiceProviderConfigResponseSort
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ScimServiceProviderConfigResponse) GetSortOk() (ScimServiceProviderConfigResponseSort, bool) {
	if o == nil || o.Sort == nil {
		var ret ScimServiceProviderConfigResponseSort
		return ret, false
	}
	return *o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ScimServiceProviderConfigResponse) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given ScimServiceProviderConfigResponseSort and assigns it to the Sort field.
func (o *ScimServiceProviderConfigResponse) SetSort(v ScimServiceProviderConfigResponseSort) {
	o.Sort = &v
}

type NullableScimServiceProviderConfigResponse struct {
	Value        ScimServiceProviderConfigResponse
	ExplicitNull bool
}

func (v NullableScimServiceProviderConfigResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableScimServiceProviderConfigResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
