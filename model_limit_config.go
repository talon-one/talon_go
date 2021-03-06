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

// LimitConfig struct for LimitConfig
type LimitConfig struct {
	// The limitable action to which this limit will be applied
	Action string `json:"action"`
	// The value to set for the limit
	Limit float32 `json:"limit"`
	// The entities that make the address of this limit
	Entities []string `json:"entities"`
}

// GetAction returns the Action field value
func (o *LimitConfig) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// SetAction sets field value
func (o *LimitConfig) SetAction(v string) {
	o.Action = v
}

// GetLimit returns the Limit field value
func (o *LimitConfig) GetLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Limit
}

// SetLimit sets field value
func (o *LimitConfig) SetLimit(v float32) {
	o.Limit = v
}

// GetEntities returns the Entities field value
func (o *LimitConfig) GetEntities() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Entities
}

// SetEntities sets field value
func (o *LimitConfig) SetEntities(v []string) {
	o.Entities = v
}

type NullableLimitConfig struct {
	Value        LimitConfig
	ExplicitNull bool
}

func (v NullableLimitConfig) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLimitConfig) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
