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

// ErrorEffectProps Whenever an error occurred during evaluation, we return an error effect. This should never happen for rules created in the rule builder.
type ErrorEffectProps struct {
	// The error message.
	Message string `json:"message"`
}

// GetMessage returns the Message field value
func (o *ErrorEffectProps) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// SetMessage sets field value
func (o *ErrorEffectProps) SetMessage(v string) {
	o.Message = v
}

type NullableErrorEffectProps struct {
	Value        ErrorEffectProps
	ExplicitNull bool
}

func (v NullableErrorEffectProps) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableErrorEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
