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

// CampaignTemplateParams struct for CampaignTemplateParams
type CampaignTemplateParams struct {
	// Name of the campaign template parameter.
	Name string `json:"name"`
	// Defines the type of parameter value.
	Type string `json:"type"`
	// Explains the meaning of this template parameter and the placeholder value that will define it. It is used on campaign creation from this template.
	Description string `json:"description"`
	// ID of the corresponding attribute.
	AttributeId *int32 `json:"attributeId,omitempty"`
}

// GetName returns the Name field value
func (o *CampaignTemplateParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CampaignTemplateParams) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CampaignTemplateParams) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// SetType sets field value
func (o *CampaignTemplateParams) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *CampaignTemplateParams) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// SetDescription sets field value
func (o *CampaignTemplateParams) SetDescription(v string) {
	o.Description = v
}

// GetAttributeId returns the AttributeId field value if set, zero value otherwise.
func (o *CampaignTemplateParams) GetAttributeId() int32 {
	if o == nil || o.AttributeId == nil {
		var ret int32
		return ret
	}
	return *o.AttributeId
}

// GetAttributeIdOk returns a tuple with the AttributeId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplateParams) GetAttributeIdOk() (int32, bool) {
	if o == nil || o.AttributeId == nil {
		var ret int32
		return ret, false
	}
	return *o.AttributeId, true
}

// HasAttributeId returns a boolean if a field has been set.
func (o *CampaignTemplateParams) HasAttributeId() bool {
	if o != nil && o.AttributeId != nil {
		return true
	}

	return false
}

// SetAttributeId gets a reference to the given int32 and assigns it to the AttributeId field.
func (o *CampaignTemplateParams) SetAttributeId(v int32) {
	o.AttributeId = &v
}

type NullableCampaignTemplateParams struct {
	Value        CampaignTemplateParams
	ExplicitNull bool
}

func (v NullableCampaignTemplateParams) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignTemplateParams) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
