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

// OutgoingIntegrationIterablePolicy struct for OutgoingIntegrationIterablePolicy
type OutgoingIntegrationIterablePolicy struct {
	// The base URL that is based on the region key of your Iterable account.
	BaseUrl string `json:"baseUrl"`
	// The API key generated from your Iterable account. See [Iterable API Key Guide](https://support.iterable.com/hc/en-us/articles/360043464871-API-Keys-)
	ApiKey string `json:"apiKey"`
}

// GetBaseUrl returns the BaseUrl field value
func (o *OutgoingIntegrationIterablePolicy) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// SetBaseUrl sets field value
func (o *OutgoingIntegrationIterablePolicy) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetApiKey returns the ApiKey field value
func (o *OutgoingIntegrationIterablePolicy) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// SetApiKey sets field value
func (o *OutgoingIntegrationIterablePolicy) SetApiKey(v string) {
	o.ApiKey = v
}

type NullableOutgoingIntegrationIterablePolicy struct {
	Value        OutgoingIntegrationIterablePolicy
	ExplicitNull bool
}

func (v NullableOutgoingIntegrationIterablePolicy) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOutgoingIntegrationIterablePolicy) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
