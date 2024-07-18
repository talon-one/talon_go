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

// Giveaway struct for Giveaway
type Giveaway struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The code value of this giveaway.
	Code string `json:"code"`
	// The ID of the pool to return giveaway codes from.
	PoolId int32 `json:"poolId"`
	// Timestamp at which point the giveaway becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Timestamp at which point the giveaway becomes invalid.
	EndDate *time.Time `json:"endDate,omitempty"`
	// Arbitrary properties associated with this giveaway.
	Attributes *map[string]map[string]interface{} `json:"attributes,omitempty"`
	// Indicates whether this giveaway code was given before.
	Used *bool `json:"used,omitempty"`
	// The ID of the Import which created this giveaway.
	ImportId *int32 `json:"importId,omitempty"`
	// The third-party integration ID of the customer profile that was awarded the giveaway, if the giveaway was awarded.
	ProfileIntegrationId *string `json:"profileIntegrationId,omitempty"`
	// The internal ID of the customer profile that was awarded the giveaway, if the giveaway was awarded and an internal ID exists.
	ProfileId *int32 `json:"profileId,omitempty"`
}

// GetId returns the Id field value
func (o *Giveaway) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *Giveaway) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Giveaway) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *Giveaway) SetCreated(v time.Time) {
	o.Created = v
}

// GetCode returns the Code field value
func (o *Giveaway) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// SetCode sets field value
func (o *Giveaway) SetCode(v string) {
	o.Code = v
}

// GetPoolId returns the PoolId field value
func (o *Giveaway) GetPoolId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PoolId
}

// SetPoolId sets field value
func (o *Giveaway) SetPoolId(v int32) {
	o.PoolId = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Giveaway) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Giveaway) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Giveaway) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Giveaway) GetEndDate() time.Time {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetEndDateOk() (time.Time, bool) {
	if o == nil || o.EndDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Giveaway) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *Giveaway) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Giveaway) GetAttributes() map[string]map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetAttributesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		var ret map[string]map[string]interface{}
		return ret, false
	}
	return *o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Giveaway) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]map[string]interface{} and assigns it to the Attributes field.
func (o *Giveaway) SetAttributes(v map[string]map[string]interface{}) {
	o.Attributes = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *Giveaway) GetUsed() bool {
	if o == nil || o.Used == nil {
		var ret bool
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetUsedOk() (bool, bool) {
	if o == nil || o.Used == nil {
		var ret bool
		return ret, false
	}
	return *o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *Giveaway) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given bool and assigns it to the Used field.
func (o *Giveaway) SetUsed(v bool) {
	o.Used = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *Giveaway) GetImportId() int32 {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetImportIdOk() (int32, bool) {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret, false
	}
	return *o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *Giveaway) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given int32 and assigns it to the ImportId field.
func (o *Giveaway) SetImportId(v int32) {
	o.ImportId = &v
}

// GetProfileIntegrationId returns the ProfileIntegrationId field value if set, zero value otherwise.
func (o *Giveaway) GetProfileIntegrationId() string {
	if o == nil || o.ProfileIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.ProfileIntegrationId
}

// GetProfileIntegrationIdOk returns a tuple with the ProfileIntegrationId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetProfileIntegrationIdOk() (string, bool) {
	if o == nil || o.ProfileIntegrationId == nil {
		var ret string
		return ret, false
	}
	return *o.ProfileIntegrationId, true
}

// HasProfileIntegrationId returns a boolean if a field has been set.
func (o *Giveaway) HasProfileIntegrationId() bool {
	if o != nil && o.ProfileIntegrationId != nil {
		return true
	}

	return false
}

// SetProfileIntegrationId gets a reference to the given string and assigns it to the ProfileIntegrationId field.
func (o *Giveaway) SetProfileIntegrationId(v string) {
	o.ProfileIntegrationId = &v
}

// GetProfileId returns the ProfileId field value if set, zero value otherwise.
func (o *Giveaway) GetProfileId() int32 {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret
	}
	return *o.ProfileId
}

// GetProfileIdOk returns a tuple with the ProfileId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *Giveaway) GetProfileIdOk() (int32, bool) {
	if o == nil || o.ProfileId == nil {
		var ret int32
		return ret, false
	}
	return *o.ProfileId, true
}

// HasProfileId returns a boolean if a field has been set.
func (o *Giveaway) HasProfileId() bool {
	if o != nil && o.ProfileId != nil {
		return true
	}

	return false
}

// SetProfileId gets a reference to the given int32 and assigns it to the ProfileId field.
func (o *Giveaway) SetProfileId(v int32) {
	o.ProfileId = &v
}

type NullableGiveaway struct {
	Value Giveaway
	ExplicitNull bool
}

func (v NullableGiveaway) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableGiveaway) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
