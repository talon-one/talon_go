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

// Referral struct for Referral
type Referral struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// Timestamp at which point the referral code becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the referral code. Referral never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of times a referral code can be used. `0` means no limit but any campaign usage limits will still apply. 
	UsageLimit int32 `json:"usageLimit"`
	// ID of the campaign from which the referral received the referral code.
	CampaignId int32 `json:"campaignId"`
	// The Integration ID of the Advocate's Profile.
	AdvocateProfileIntegrationId string `json:"advocateProfileIntegrationId"`
	// An optional Integration ID of the Friend's Profile.
	FriendProfileIntegrationId *string `json:"friendProfileIntegrationId,omitempty"`
	// Arbitrary properties associated with this item.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	// The ID of the Import which created this referral.
	ImportId *int32 `json:"importId,omitempty"`
	// The referral code.
	Code string `json:"code"`
	// The number of times this referral code has been successfully used.
	UsageCounter int32 `json:"usageCounter"`
	// The ID of the batch the referrals belong to.
	BatchId *string `json:"batchId,omitempty"`
}

// GetId returns the Id field value
func (o *Referral) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Referral) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Referral) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Referral) SetCreated(v time.Time) {
	o.Created = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Referral) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Referral) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Referral) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Referral) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Referral) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *Referral) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetUsageLimit returns the UsageLimit field value
func (o *Referral) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *Referral) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetCampaignId returns the CampaignId field value
func (o *Referral) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *Referral) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetAdvocateProfileIntegrationId returns the AdvocateProfileIntegrationId field value
func (o *Referral) GetAdvocateProfileIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvocateProfileIntegrationId
}

// SetAdvocateProfileIntegrationId sets field value
func (o *Referral) SetAdvocateProfileIntegrationId(v string) {
	o.AdvocateProfileIntegrationId = v
}

// GetFriendProfileIntegrationId returns the FriendProfileIntegrationId field value if set, zero value otherwise.
func (o *Referral) GetFriendProfileIntegrationId() string {
	if o == nil || o.FriendProfileIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.FriendProfileIntegrationId
}

// GetFriendProfileIntegrationIdOk returns a tuple with the FriendProfileIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetFriendProfileIntegrationIdOk() (string, bool) {
	if o == nil || o.FriendProfileIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.FriendProfileIntegrationId, true
}

// HasFriendProfileIntegrationId returns a boolean if a field has been set.
func (o *Referral) HasFriendProfileIntegrationId() bool {
	if o != nil && o.FriendProfileIntegrationId != nil {
		return true
	}

	return false
}

// SetFriendProfileIntegrationId gets a reference to the given string and assigns it to the FriendProfileIntegrationId field.
func (o *Referral) SetFriendProfileIntegrationId(v string) {
	o.FriendProfileIntegrationId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Referral) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Referral) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *Referral) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *Referral) GetImportId() int32 {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetImportIdOk() (int32, bool) {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret, false
	}
	return *o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *Referral) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given int32 and assigns it to the ImportId field.
func (o *Referral) SetImportId(v int32) {
	o.ImportId = &v
}

// GetCode returns the Code field value
func (o *Referral) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// SetCode sets field value
func (o *Referral) SetCode(v string) {
	o.Code = v
}

// GetUsageCounter returns the UsageCounter field value
func (o *Referral) GetUsageCounter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageCounter
}

// SetUsageCounter sets field value
func (o *Referral) SetUsageCounter(v int32) {
	o.UsageCounter = v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *Referral) GetBatchId() string {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Referral) GetBatchIdOk() (string, bool) {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret, false
	}
	return *o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *Referral) HasBatchId() bool {
	if o != nil && o.BatchId != nil {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *Referral) SetBatchId(v string) {
	o.BatchId = &v
}

type NullableReferral struct {
	Value Referral
	ExplicitNull bool
}

func (v NullableReferral) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableReferral) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
