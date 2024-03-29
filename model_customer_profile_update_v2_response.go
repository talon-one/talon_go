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

// CustomerProfileUpdateV2Response Contains the updated customer profiles entities. This is the response type returned by the V2 PUT customer_profiles endpoint
type CustomerProfileUpdateV2Response struct {
	CustomerProfile CustomerProfile `json:"customerProfile"`
}

// GetCustomerProfile returns the CustomerProfile field value
func (o *CustomerProfileUpdateV2Response) GetCustomerProfile() CustomerProfile {
	if o == nil {
		var ret CustomerProfile
		return ret
	}

	return o.CustomerProfile
}

// SetCustomerProfile sets field value
func (o *CustomerProfileUpdateV2Response) SetCustomerProfile(v CustomerProfile) {
	o.CustomerProfile = v
}

type NullableCustomerProfileUpdateV2Response struct {
	Value        CustomerProfileUpdateV2Response
	ExplicitNull bool
}

func (v NullableCustomerProfileUpdateV2Response) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCustomerProfileUpdateV2Response) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
