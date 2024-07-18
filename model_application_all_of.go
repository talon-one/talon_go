/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// ApplicationAllOf struct for ApplicationAllOf
type ApplicationAllOf struct {
	// An array containing all the loyalty programs to which this application is subscribed.
	LoyaltyPrograms []LoyaltyProgram `json:"loyaltyPrograms"`
}

// GetLoyaltyPrograms returns the LoyaltyPrograms field value
func (o *ApplicationAllOf) GetLoyaltyPrograms() []LoyaltyProgram {
	if o == nil {
		var ret []LoyaltyProgram
		return ret
	}

	return o.LoyaltyPrograms
}

// SetLoyaltyPrograms sets field value
func (o *ApplicationAllOf) SetLoyaltyPrograms(v []LoyaltyProgram) {
	o.LoyaltyPrograms = v
}

type NullableApplicationAllOf struct {
	Value ApplicationAllOf
	ExplicitNull bool
}

func (v NullableApplicationAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}