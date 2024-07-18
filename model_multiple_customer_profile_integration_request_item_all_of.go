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

// MultipleCustomerProfileIntegrationRequestItemAllOf The body of a V2 integration API request (customer profile update). Next to the customer profile details, this contains an optional listing of extra properties that should be returned in the response. 
type MultipleCustomerProfileIntegrationRequestItemAllOf struct {
	// The identifier of this profile, set by your integration layer. It must be unique within the account.  To get the `integrationId` of the profile from a `sessionId`, use the [Update customer session](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2). 
	IntegrationId string `json:"integrationId"`
}

// GetIntegrationId returns the IntegrationId field value
func (o *MultipleCustomerProfileIntegrationRequestItemAllOf) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IntegrationId
}

// SetIntegrationId sets field value
func (o *MultipleCustomerProfileIntegrationRequestItemAllOf) SetIntegrationId(v string) {
	o.IntegrationId = v
}

type NullableMultipleCustomerProfileIntegrationRequestItemAllOf struct {
	Value MultipleCustomerProfileIntegrationRequestItemAllOf
	ExplicitNull bool
}

func (v NullableMultipleCustomerProfileIntegrationRequestItemAllOf) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableMultipleCustomerProfileIntegrationRequestItemAllOf) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
