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

// ErrorSource The source of the current error, exactly one of `pointer`, `parameter` or `line` will be defined. 
type ErrorSource struct {
	// Pointer to the path in the payload that caused this error.
	Pointer *string `json:"pointer,omitempty"`
	// Query parameter that caused this error.
	Parameter *string `json:"parameter,omitempty"`
	// Line number in uploaded multipart file that caused this error. 'N/A' if unknown.
	Line *string `json:"line,omitempty"`
	// Pointer to the resource that caused this error
	Resource *string `json:"resource,omitempty"`
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *ErrorSource) GetPointer() string {
	if o == nil || o.Pointer == nil {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetPointerOk() (string, bool) {
	if o == nil || o.Pointer == nil {
		var ret string
		return ret, false
	}
	return *o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *ErrorSource) HasPointer() bool {
	if o != nil && o.Pointer != nil {
		return true
	}

	return false
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *ErrorSource) SetPointer(v string) {
	o.Pointer = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ErrorSource) GetParameter() string {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetParameterOk() (string, bool) {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret, false
	}
	return *o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ErrorSource) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *ErrorSource) SetParameter(v string) {
	o.Parameter = &v
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *ErrorSource) GetLine() string {
	if o == nil || o.Line == nil {
		var ret string
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetLineOk() (string, bool) {
	if o == nil || o.Line == nil {
		var ret string
		return ret, false
	}
	return *o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *ErrorSource) HasLine() bool {
	if o != nil && o.Line != nil {
		return true
	}

	return false
}

// SetLine gets a reference to the given string and assigns it to the Line field.
func (o *ErrorSource) SetLine(v string) {
	o.Line = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ErrorSource) GetResource() string {
	if o == nil || o.Resource == nil {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetResourceOk() (string, bool) {
	if o == nil || o.Resource == nil {
		var ret string
		return ret, false
	}
	return *o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ErrorSource) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *ErrorSource) SetResource(v string) {
	o.Resource = &v
}

type NullableErrorSource struct {
	Value ErrorSource
	ExplicitNull bool
}

func (v NullableErrorSource) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableErrorSource) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
