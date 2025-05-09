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

// LoyaltyTier A tier in a loyalty program.
type LoyaltyTier struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The ID of the loyalty program that owns this entity.
	ProgramID int32 `json:"programID"`
	// The integration name of the loyalty program that owns this entity.
	ProgramName *string `json:"programName,omitempty"`
	// The Campaign Manager-displayed name of the loyalty program that owns this entity.
	ProgramTitle *string `json:"programTitle,omitempty"`
	// The name of the tier.
	Name string `json:"name"`
	// The minimum amount of points required to enter the tier.
	MinPoints float32 `json:"minPoints"`
}

// GetId returns the Id field value
func (o *LoyaltyTier) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *LoyaltyTier) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *LoyaltyTier) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *LoyaltyTier) SetCreated(v time.Time) {
	o.Created = v
}

// GetProgramID returns the ProgramID field value
func (o *LoyaltyTier) GetProgramID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramID
}

// SetProgramID sets field value
func (o *LoyaltyTier) SetProgramID(v int32) {
	o.ProgramID = v
}

// GetProgramName returns the ProgramName field value if set, zero value otherwise.
func (o *LoyaltyTier) GetProgramName() string {
	if o == nil || o.ProgramName == nil {
		var ret string
		return ret
	}
	return *o.ProgramName
}

// GetProgramNameOk returns a tuple with the ProgramName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyTier) GetProgramNameOk() (string, bool) {
	if o == nil || o.ProgramName == nil {
		var ret string
		return ret, false
	}
	return *o.ProgramName, true
}

// HasProgramName returns a boolean if a field has been set.
func (o *LoyaltyTier) HasProgramName() bool {
	if o != nil && o.ProgramName != nil {
		return true
	}

	return false
}

// SetProgramName gets a reference to the given string and assigns it to the ProgramName field.
func (o *LoyaltyTier) SetProgramName(v string) {
	o.ProgramName = &v
}

// GetProgramTitle returns the ProgramTitle field value if set, zero value otherwise.
func (o *LoyaltyTier) GetProgramTitle() string {
	if o == nil || o.ProgramTitle == nil {
		var ret string
		return ret
	}
	return *o.ProgramTitle
}

// GetProgramTitleOk returns a tuple with the ProgramTitle field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *LoyaltyTier) GetProgramTitleOk() (string, bool) {
	if o == nil || o.ProgramTitle == nil {
		var ret string
		return ret, false
	}
	return *o.ProgramTitle, true
}

// HasProgramTitle returns a boolean if a field has been set.
func (o *LoyaltyTier) HasProgramTitle() bool {
	if o != nil && o.ProgramTitle != nil {
		return true
	}

	return false
}

// SetProgramTitle gets a reference to the given string and assigns it to the ProgramTitle field.
func (o *LoyaltyTier) SetProgramTitle(v string) {
	o.ProgramTitle = &v
}

// GetName returns the Name field value
func (o *LoyaltyTier) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *LoyaltyTier) SetName(v string) {
	o.Name = v
}

// GetMinPoints returns the MinPoints field value
func (o *LoyaltyTier) GetMinPoints() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MinPoints
}

// SetMinPoints sets field value
func (o *LoyaltyTier) SetMinPoints(v float32) {
	o.MinPoints = v
}

type NullableLoyaltyTier struct {
	Value        LoyaltyTier
	ExplicitNull bool
}

func (v NullableLoyaltyTier) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableLoyaltyTier) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
