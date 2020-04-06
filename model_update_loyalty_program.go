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
)

// UpdateLoyaltyProgram Update Loyalty Program
type UpdateLoyaltyProgram struct {
	// The display title for the Loyalty Program.
	Title *string `json:"title,omitempty"`
	// Description of our Loyalty Program.
	Description *string `json:"description,omitempty"`
	// A list containing the IDs of all applications that are subscribed to this Loyalty Program.
	SubscribedApplications *[]int32 `json:"subscribedApplications,omitempty"`
	// Indicates the default duration after which new loyalty points should expire. The format is a number, followed by one letter indicating the unit; like '1h' or '40m' or '30d'.
	DefaultValidity *string `json:"defaultValidity,omitempty"`
	// Indicates if this program supports subledgers inside the program
	AllowSubledger *bool `json:"allowSubledger,omitempty"`
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateLoyaltyProgram) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateLoyaltyProgram) SetDescription(v string) {
	o.Description = &v
}

// GetSubscribedApplications returns the SubscribedApplications field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetSubscribedApplications() []int32 {
	if o == nil || o.SubscribedApplications == nil {
		var ret []int32
		return ret
	}
	return *o.SubscribedApplications
}

// GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplications == nil {
		var ret []int32
		return ret, false
	}
	return *o.SubscribedApplications, true
}

// HasSubscribedApplications returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasSubscribedApplications() bool {
	if o != nil && o.SubscribedApplications != nil {
		return true
	}

	return false
}

// SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.
func (o *UpdateLoyaltyProgram) SetSubscribedApplications(v []int32) {
	o.SubscribedApplications = &v
}

// GetDefaultValidity returns the DefaultValidity field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetDefaultValidity() string {
	if o == nil || o.DefaultValidity == nil {
		var ret string
		return ret
	}
	return *o.DefaultValidity
}

// GetDefaultValidityOk returns a tuple with the DefaultValidity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetDefaultValidityOk() (string, bool) {
	if o == nil || o.DefaultValidity == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultValidity, true
}

// HasDefaultValidity returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasDefaultValidity() bool {
	if o != nil && o.DefaultValidity != nil {
		return true
	}

	return false
}

// SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.
func (o *UpdateLoyaltyProgram) SetDefaultValidity(v string) {
	o.DefaultValidity = &v
}

// GetAllowSubledger returns the AllowSubledger field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetAllowSubledger() bool {
	if o == nil || o.AllowSubledger == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubledger
}

// GetAllowSubledgerOk returns a tuple with the AllowSubledger field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetAllowSubledgerOk() (bool, bool) {
	if o == nil || o.AllowSubledger == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowSubledger, true
}

// HasAllowSubledger returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasAllowSubledger() bool {
	if o != nil && o.AllowSubledger != nil {
		return true
	}

	return false
}

// SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.
func (o *UpdateLoyaltyProgram) SetAllowSubledger(v bool) {
	o.AllowSubledger = &v
}

type NullableUpdateLoyaltyProgram struct {
	Value UpdateLoyaltyProgram
	ExplicitNull bool
}

func (v NullableUpdateLoyaltyProgram) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateLoyaltyProgram) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
