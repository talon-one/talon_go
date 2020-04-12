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

// InlineResponse20030 struct for InlineResponse20030
type InlineResponse20030 struct {
	TotalResultSize int32  `json:"totalResultSize"`
	Data            []Role `json:"data"`
}

// GetTotalResultSize returns the TotalResultSize field value
func (o *InlineResponse20030) GetTotalResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResultSize
}

// SetTotalResultSize sets field value
func (o *InlineResponse20030) SetTotalResultSize(v int32) {
	o.TotalResultSize = v
}

// GetData returns the Data field value
func (o *InlineResponse20030) GetData() []Role {
	if o == nil {
		var ret []Role
		return ret
	}

	return o.Data
}

// SetData sets field value
func (o *InlineResponse20030) SetData(v []Role) {
	o.Data = v
}

type NullableInlineResponse20030 struct {
	Value        InlineResponse20030
	ExplicitNull bool
}

func (v NullableInlineResponse20030) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableInlineResponse20030) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
