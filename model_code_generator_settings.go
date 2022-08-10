/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerSession](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2) endpoint is `https://mycompany.talon.one/v2/customer_sessions/{Id}`
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// CodeGeneratorSettings struct for CodeGeneratorSettings
type CodeGeneratorSettings struct {
	// List of characters used to generate the random parts of a code. Defaults to a list equivalent to the `[A-Z, 0-9]` regexp.
	ValidCharacters []string `json:"validCharacters"`
	// The pattern used to generate coupon codes. The character `#` is a placeholder and is replaced by a random character from the `validCharacters` set.
	CouponPattern string `json:"couponPattern"`
}

// GetValidCharacters returns the ValidCharacters field value
func (o *CodeGeneratorSettings) GetValidCharacters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ValidCharacters
}

// SetValidCharacters sets field value
func (o *CodeGeneratorSettings) SetValidCharacters(v []string) {
	o.ValidCharacters = v
}

// GetCouponPattern returns the CouponPattern field value
func (o *CodeGeneratorSettings) GetCouponPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CouponPattern
}

// SetCouponPattern sets field value
func (o *CodeGeneratorSettings) SetCouponPattern(v string) {
	o.CouponPattern = v
}

type NullableCodeGeneratorSettings struct {
	Value        CodeGeneratorSettings
	ExplicitNull bool
}

func (v NullableCodeGeneratorSettings) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCodeGeneratorSettings) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
