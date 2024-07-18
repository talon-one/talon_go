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

// OutgoingIntegrationTemplateWithConfigurationDetails struct for OutgoingIntegrationTemplateWithConfigurationDetails
type OutgoingIntegrationTemplateWithConfigurationDetails struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// Unique ID of outgoing integration type.
	IntegrationType int32 `json:"integrationType"`
	// The title of the integration template.
	Title string `json:"title"`
	// The description of the specific outgoing integration template.
	Description string `json:"description"`
	// The API payload (supports templating using parameters) for this integration template.
	Payload string `json:"payload"`
	// API method for this webhook.
	Method string `json:"method"`
	// The relative URL corresponding to each integration template.
	RelativeUrl string `json:"relativeUrl"`
	// The list of HTTP headers for this integration template.
	Headers []string `json:"headers"`
	Policy OutgoingIntegrationConfigurationPolicy `json:"policy"`
}

// GetId returns the Id field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetId(v int32) {
	o.Id = v
}

// GetIntegrationType returns the IntegrationType field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetIntegrationType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IntegrationType
}

// SetIntegrationType sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetIntegrationType(v int32) {
	o.IntegrationType = v
}

// GetTitle returns the Title field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetDescription(v string) {
	o.Description = v
}

// GetPayload returns the Payload field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// SetPayload sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetPayload(v string) {
	o.Payload = v
}

// GetMethod returns the Method field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// SetMethod sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetMethod(v string) {
	o.Method = v
}

// GetRelativeUrl returns the RelativeUrl field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetRelativeUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelativeUrl
}

// SetRelativeUrl sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetRelativeUrl(v string) {
	o.RelativeUrl = v
}

// GetHeaders returns the Headers field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Headers
}

// SetHeaders sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetHeaders(v []string) {
	o.Headers = v
}

// GetPolicy returns the Policy field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) GetPolicy() OutgoingIntegrationConfigurationPolicy {
	if o == nil {
		var ret OutgoingIntegrationConfigurationPolicy
		return ret
	}

	return o.Policy
}

// SetPolicy sets field value
func (o *OutgoingIntegrationTemplateWithConfigurationDetails) SetPolicy(v OutgoingIntegrationConfigurationPolicy) {
	o.Policy = v
}

type NullableOutgoingIntegrationTemplateWithConfigurationDetails struct {
	Value OutgoingIntegrationTemplateWithConfigurationDetails
	ExplicitNull bool
}

func (v NullableOutgoingIntegrationTemplateWithConfigurationDetails) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableOutgoingIntegrationTemplateWithConfigurationDetails) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
