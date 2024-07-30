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

// CampaignCollection 
type CampaignCollection struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The time this entity was last modified.
	Modified time.Time `json:"modified"`
	// A short description of the purpose of this collection.
	Description *string `json:"description,omitempty"`
	// The name of this collection.
	Name string `json:"name"`
	// ID of the user who last updated this effect if available.
	ModifiedBy *int32 `json:"modifiedBy,omitempty"`
	// ID of the user who created this effect.
	CreatedBy int32 `json:"createdBy"`
	// The ID of the Application that owns this entity.
	ApplicationId *int32 `json:"applicationId,omitempty"`
	// The ID of the campaign that owns this entity.
	CampaignId *int32 `json:"campaignId,omitempty"`
	// The content of the collection.
	Payload *[]string `json:"payload,omitempty"`
}

// GetId returns the Id field value
func (o *CampaignCollection) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *CampaignCollection) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CampaignCollection) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CampaignCollection) SetCreated(v time.Time) {
	o.Created = v
}

// GetAccountId returns the AccountId field value
func (o *CampaignCollection) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *CampaignCollection) SetAccountId(v int32) {
	o.AccountId = v
}

// GetModified returns the Modified field value
func (o *CampaignCollection) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *CampaignCollection) SetModified(v time.Time) {
	o.Modified = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignCollection) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignCollection) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignCollection) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *CampaignCollection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CampaignCollection) SetName(v string) {
	o.Name = v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *CampaignCollection) GetModifiedBy() int32 {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetModifiedByOk() (int32, bool) {
	if o == nil || o.ModifiedBy == nil {
		var ret int32
		return ret, false
	}
	return *o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *CampaignCollection) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given int32 and assigns it to the ModifiedBy field.
func (o *CampaignCollection) SetModifiedBy(v int32) {
	o.ModifiedBy = &v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CampaignCollection) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// SetCreatedBy sets field value
func (o *CampaignCollection) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *CampaignCollection) GetApplicationId() int32 {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetApplicationIdOk() (int32, bool) {
	if o == nil || o.ApplicationId == nil {
		var ret int32
		return ret, false
	}
	return *o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *CampaignCollection) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *CampaignCollection) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *CampaignCollection) GetCampaignId() int32 {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetCampaignIdOk() (int32, bool) {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret, false
	}
	return *o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *CampaignCollection) HasCampaignId() bool {
	if o != nil && o.CampaignId != nil {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.
func (o *CampaignCollection) SetCampaignId(v int32) {
	o.CampaignId = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CampaignCollection) GetPayload() []string {
	if o == nil || o.Payload == nil {
		var ret []string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignCollection) GetPayloadOk() ([]string, bool) {
	if o == nil || o.Payload == nil {
		var ret []string
		return ret, false
	}
	return *o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CampaignCollection) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given []string and assigns it to the Payload field.
func (o *CampaignCollection) SetPayload(v []string) {
	o.Payload = &v
}

type NullableCampaignCollection struct {
	Value CampaignCollection
	ExplicitNull bool
}

func (v NullableCampaignCollection) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignCollection) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
