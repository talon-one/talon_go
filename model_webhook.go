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
	"time"
)

// Webhook struct for Webhook
type Webhook struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The time this entity was last modified.
	Modified time.Time `json:"modified"`
	// The IDs of the Applications in which this webhook is available. An empty array means the webhook is available in `All Applications`.
	ApplicationIds []int32 `json:"applicationIds"`
	// Name or title for this webhook.
	Title string `json:"title"`
	// A description of the webhook.
	Description *string `json:"description,omitempty"`
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
}

// GetId returns the Id field value
func (o *Webhook) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Webhook) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Webhook) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Webhook) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *Webhook) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *Webhook) SetModified(v time.Time) {
	o.Modified = v
}

// GetApplicationIds returns the ApplicationIds field value
func (o *Webhook) GetApplicationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.ApplicationIds
}

// SetApplicationIds sets field value
func (o *Webhook) SetApplicationIds(v []int32) {
	o.ApplicationIds = v
}

// GetTitle returns the Title field value
func (o *Webhook) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *Webhook) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Webhook) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Webhook) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Webhook) SetDescription(v string) {
	o.Description = &v
}

// GetVerb returns the Verb field value
func (o *Webhook) GetVerb() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Verb
}

// SetVerb sets field value
func (o *Webhook) SetVerb(v string) {
	o.Verb = v
}

// GetUrl returns the Url field value
func (o *Webhook) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// SetUrl sets field value
func (o *Webhook) SetUrl(v string) {
	o.Url = v
}

// GetHeaders returns the Headers field value
func (o *Webhook) GetHeaders() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Headers
}

// SetHeaders sets field value
func (o *Webhook) SetHeaders(v []string) {
	o.Headers = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *Webhook) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Webhook) GetPayloadOk() (string, bool) {
	if o == nil || o.Payload == nil {
		var ret string
		return ret, false
	}
	return *o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *Webhook) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *Webhook) SetPayload(v string) {
	o.Payload = &v
}

// GetParams returns the Params field value
func (o *Webhook) GetParams() []TemplateArgDef {
	if o == nil {
		var ret []TemplateArgDef
		return ret
	}

	return o.Params
}

// SetParams sets field value
func (o *Webhook) SetParams(v []TemplateArgDef) {
	o.Params = v
}

// GetEnabled returns the Enabled field value
func (o *Webhook) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// SetEnabled sets field value
func (o *Webhook) SetEnabled(v bool) {
	o.Enabled = v
}

type NullableWebhook struct {
	Value        Webhook
	ExplicitNull bool
}

func (v NullableWebhook) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableWebhook) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
