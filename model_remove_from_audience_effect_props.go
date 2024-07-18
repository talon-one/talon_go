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

// RemoveFromAudienceEffectProps The properties specific to the \"removeFromAudience\" effect. This gets triggered whenever a validated rule contains a \"removeFromAudience\" effect.
type RemoveFromAudienceEffectProps struct {
	// The internal ID of the audience.
	AudienceId *int32 `json:"audienceId,omitempty"`
	// The name of the audience.
	AudienceName *string `json:"audienceName,omitempty"`
	// The ID of the customer profile in the third-party integration platform.
	ProfileIntegrationId *string `json:"profileIntegrationId,omitempty"`
	// The internal ID of the customer profile.
	ProfileId *int32 `json:"profileId,omitempty"`
}

// GetAudienceId returns the AudienceId field value if set, zero value otherwise.
func (o *RemoveFromAudienceEffectProps) GetAudienceId() int32 {
	if o == nil || o.AudienceId == nil {
		var ret int32
		return ret
	}
	return *o.AudienceId
}

// GetAudienceIdOk returns a tuple with the AudienceId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFromAudienceEffectProps) GetAudienceIdOk() (int32, bool) {
	if o == nil || o.AudienceId == nil {
		var ret int32
		return ret, false
	}
	return *o.AudienceId, true
}

// HasAudienceId returns a boolean if a field has been set.
func (o *RemoveFromAudienceEffectProps) HasAudienceId() bool {
	if o != nil && o.AudienceId != nil {
		return true
	}

	return false
}

// SetAudienceId gets a reference to the given int32 and assigns it to the AudienceId field.
func (o *RemoveFromAudienceEffectProps) SetAudienceId(v int32) {
	o.AudienceId = &v
}

// GetAudienceName returns the AudienceName field value if set, zero value otherwise.
func (o *RemoveFromAudienceEffectProps) GetAudienceName() string {
	if o == nil || o.AudienceName == nil {
		var ret string
		return ret
	}
	return *o.AudienceName
}

// GetAudienceNameOk returns a tuple with the AudienceName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFromAudienceEffectProps) GetAudienceNameOk() (string, bool) {
	if o == nil || o.AudienceName == nil {
		var ret string
		return ret, false
	}
	return *o.AudienceName, true
}

// HasAudienceName returns a boolean if a field has been set.
func (o *RemoveFromAudienceEffectProps) HasAudienceName() bool {
	if o != nil && o.AudienceName != nil {
		return true
	}

	return false
}

// SetAudienceName gets a reference to the given string and assigns it to the AudienceName field.
func (o *RemoveFromAudienceEffectProps) SetAudienceName(v string) {
	o.AudienceName = &v
}

// GetProfileIntegrationId returns the ProfileIntegrationId field value if set, zero value otherwise.
func (o *RemoveFromAudienceEffectProps) GetProfileIntegrationId() string {
	if o == nil || o.ProfileIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.ProfileIntegrationId
}

// GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFromAudienceEffectProps) GetProfileIntegrationIdOk() (string, bool) {
	if o == nil || o.ProfileIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileIntegrationId, true
}

// HasProfileIntegrationId returns a boolean if a field has been set.
func (o *RemoveFromAudienceEffectProps) HasProfileIntegrationId() bool {
	if o != nil && o.ProfileIntegrationId != nil {
		return true
	}

	return false
}

// SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.
func (o *RemoveFromAudienceEffectProps) SetProfileIntegrationId(v string) {
	o.ProfileIntegrationId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *RemoveFromAudienceEffectProps) GetProfileId() int32 {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RemoveFromAudienceEffectProps) GetProfileIdOk() (int32, bool) {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *RemoveFromAudienceEffectProps) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *RemoveFromAudienceEffectProps) SetProfileId(v int32) {
	o.ProfileId = &v
}

type NullableRemoveFromAudienceEffectProps struct {
	Value RemoveFromAudienceEffectProps
	ExplicitNull bool
}

func (v NullableRemoveFromAudienceEffectProps) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRemoveFromAudienceEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
