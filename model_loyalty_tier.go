/*
 * Talon.One API
 *
 * The Talon.One API is used to manage applications and campaigns, as well as to integrate with your application. The operations in the _Integration API_ section are used to integrate with our platform, while the other operations are used to manage applications and campaigns.  ### Where is the API?  The API is available at the same hostname as these docs. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerProfile][] operation is `https://mycompany.talon.one/v1/customer_profiles/id`  [updateCustomerProfile]: #operation--v1-customer_profiles--integrationId--put
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

// LoyaltyTier
type LoyaltyTier struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the loyalty program that owns this entity.
	ProgramID int32 `json:"programID"`
	// The name of the tier
	Name string `json:"name"`
	// The minimum amount of points required to be eligible for the tier
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