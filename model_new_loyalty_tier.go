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

// NewLoyaltyTier A tier in a loyalty program.
type NewLoyaltyTier struct {
	// The minimum amount of points required to be eligible for the tier.
	MinPoints float32 `json:"minPoints"`
	// The name of the tier
	Name string `json:"name"`
}

// GetMinPoints returns the MinPoints field value
func (o *NewLoyaltyTier) GetMinPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinPoints
}

// SetMinPoints sets field value
func (o *NewLoyaltyTier) SetMinPoints(v float32) {
	o.MinPoints = v
}

// GetName returns the Name field value
func (o *NewLoyaltyTier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewLoyaltyTier) SetName(v string) {
	o.Name = v
}

type NullableNewLoyaltyTier struct {
	Value NewLoyaltyTier
	ExplicitNull bool
}

func (v NullableNewLoyaltyTier) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewLoyaltyTier) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
