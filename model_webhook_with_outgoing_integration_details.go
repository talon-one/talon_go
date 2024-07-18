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
	"time"
)

// WebhookWithOutgoingIntegrationDetails struct for WebhookWithOutgoingIntegrationDetails
type WebhookWithOutgoingIntegrationDetails struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The time this entity was last modified.
	Modified time.Time `json:"modified"`
	// The IDs of the Applications that are related to this entity.
	ApplicationIds []int32 `json:"applicationIds"`
	// Name or title for this webhook.
	Title string `json:"title"`
	// API method for this webhook.
	Verb string `json:"verb"`
	// API URL (supports templating using parameters) for this webhook.
	Url string `json:"url"`
	// List of API HTTP headers for this webhook.
	Headers []string `json:"headers"`
	// API payload (supports templating using parameters) for this webhook.
	Payload *string `json:"payload,omitempty"`
	// Array of template argument definitions.
	Params []TemplateArgDef `json:"params"`
	// Enables or disables webhook from showing in the Rule Builder.
	Enabled bool `json:"enabled"`
	// Identifier of the outgoing integration template.
	OutgoingIntegrationTemplateId *int32 `json:"outgoingIntegrationTemplateId,omitempty"`
	// Identifier of the outgoing integration type.
	OutgoingIntegrationTypeId *int32 `json:"outgoingIntegrationTypeId,omitempty"`
	// Name of the outgoing integration.
	OutgoingIntegrationTypeName *string `json:"outgoingIntegrationTypeName,omitempty"`
}

// GetId returns the Id field value
func (o *WebhookWithOutgoingIntegrationDetails) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *WebhookWithOutgoingIntegrationDetails) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *WebhookWithOutgoingIntegrationDetails) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetModified(v time.Time) {
	o.Modified = v
}

// GetApplicationIds returns the ApplicationIds field value
func (o *WebhookWithOutgoingIntegrationDetails) GetApplicationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationIds
}

// SetApplicationIds sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetApplicationIds(v []int32) {
	o.ApplicationIds = v
}

// GetTitle returns the Title field value
func (o *WebhookWithOutgoingIntegrationDetails) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetTitle(v string) {
	o.Title = v
}

// GetVerb returns the Verb field value
func (o *WebhookWithOutgoingIntegrationDetails) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// SetVerb sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetVerb(v string) {
	o.Verb = v
}

// GetUrl returns the Url field value
func (o *WebhookWithOutgoingIntegrationDetails) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// SetUrl sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value
func (o *WebhookWithOutgoingIntegrationDetails) GetHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Headers
}

// SetHeaders sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetHeaders(v []string) {
	o.Headers = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *WebhookWithOutgoingIntegrationDetails) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWithOutgoingIntegrationDetails) GetPayloadOk() (string, bool) {
	if o == nil || o.Payload == nil {
		var ret string
		return ret, false
	}
	return *o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *WebhookWithOutgoingIntegrationDetails) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *WebhookWithOutgoingIntegrationDetails) SetPayload(v string) {
	o.Payload = &v
}

// GetParams returns the Params field value
func (o *WebhookWithOutgoingIntegrationDetails) GetParams() []TemplateArgDef {
	if o == nil {
		var ret []TemplateArgDef
		return ret
	}

	return o.Params
}

// SetParams sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetParams(v []TemplateArgDef) {
	o.Params = v
}

// GetEnabled returns the Enabled field value
func (o *WebhookWithOutgoingIntegrationDetails) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// SetEnabled sets field value
func (o *WebhookWithOutgoingIntegrationDetails) SetEnabled(v bool) {
	o.Enabled = v
}

// GetOutgoingIntegrationTemplateId returns the OutgoingIntegrationTemplateId field value if set, zero value otherwise.
func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTemplateId() int32 {
	if o == nil || o.OutgoingIntegrationTemplateId == nil {
		var ret int32
		return ret
	}
	return *o.OutgoingIntegrationTemplateId
}

// GetOutgoingIntegrationTemplateIdOk returns a tuple with the OutgoingIntegrationTemplateId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTemplateIdOk() (int32, bool) {
	if o == nil || o.OutgoingIntegrationTemplateId == nil {
		var ret int32
		return ret, false
	}
	return *o.OutgoingIntegrationTemplateId, true
}

// HasOutgoingIntegrationTemplateId returns a boolean if a field has been set.
func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTemplateId() bool {
	if o != nil && o.OutgoingIntegrationTemplateId != nil {
		return true
	}

	return false
}

// SetOutgoingIntegrationTemplateId gets a reference to the given int32 and assigns it to the OutgoingIntegrationTemplateId field.
func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTemplateId(v int32) {
	o.OutgoingIntegrationTemplateId = &v
}

// GetOutgoingIntegrationTypeId returns the OutgoingIntegrationTypeId field value if set, zero value otherwise.
func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeId() int32 {
	if o == nil || o.OutgoingIntegrationTypeId == nil {
		var ret int32
		return ret
	}
	return *o.OutgoingIntegrationTypeId
}

// GetOutgoingIntegrationTypeIdOk returns a tuple with the OutgoingIntegrationTypeId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeIdOk() (int32, bool) {
	if o == nil || o.OutgoingIntegrationTypeId == nil {
		var ret int32
		return ret, false
	}
	return *o.OutgoingIntegrationTypeId, true
}

// HasOutgoingIntegrationTypeId returns a boolean if a field has been set.
func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTypeId() bool {
	if o != nil && o.OutgoingIntegrationTypeId != nil {
		return true
	}

	return false
}

// SetOutgoingIntegrationTypeId gets a reference to the given int32 and assigns it to the OutgoingIntegrationTypeId field.
func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTypeId(v int32) {
	o.OutgoingIntegrationTypeId = &v
}

// GetOutgoingIntegrationTypeName returns the OutgoingIntegrationTypeName field value if set, zero value otherwise.
func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeName() string {
	if o == nil || o.OutgoingIntegrationTypeName == nil {
		var ret string
		return ret
	}
	return *o.OutgoingIntegrationTypeName
}

// GetOutgoingIntegrationTypeNameOk returns a tuple with the OutgoingIntegrationTypeName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *WebhookWithOutgoingIntegrationDetails) GetOutgoingIntegrationTypeNameOk() (string, bool) {
	if o == nil || o.OutgoingIntegrationTypeName == nil {
		var ret string
		return ret, false
	}
	return *o.OutgoingIntegrationTypeName, true
}

// HasOutgoingIntegrationTypeName returns a boolean if a field has been set.
func (o *WebhookWithOutgoingIntegrationDetails) HasOutgoingIntegrationTypeName() bool {
	if o != nil && o.OutgoingIntegrationTypeName != nil {
		return true
	}

	return false
}

// SetOutgoingIntegrationTypeName gets a reference to the given string and assigns it to the OutgoingIntegrationTypeName field.
func (o *WebhookWithOutgoingIntegrationDetails) SetOutgoingIntegrationTypeName(v string) {
	o.OutgoingIntegrationTypeName = &v
}

type NullableWebhookWithOutgoingIntegrationDetails struct {
	Value WebhookWithOutgoingIntegrationDetails
	ExplicitNull bool
}

func (v NullableWebhookWithOutgoingIntegrationDetails) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableWebhookWithOutgoingIntegrationDetails) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
