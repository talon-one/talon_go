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

// ApplicationApiKey
type ApplicationApiKey struct {
	// Title for API Key.
	Title string `json:"title"`
	// The date the API key expired.
	Expires time.Time `json:"expires"`
	// The third-party platform the API key is valid for. Use `none` for a generic API key to be used from your own integration layer.
	Platform *string `json:"platform,omitempty"`
	// ID of the API Key.
	Id int32 `json:"id"`
	// ID of user who created.
	CreatedBy int32 `json:"createdBy"`
	// ID of account the key is used for.
	AccountID int32 `json:"accountID"`
	// ID of application the key is used for.
	ApplicationID int32 `json:"applicationID"`
	// The date the API key was created.
	Created time.Time `json:"created"`
}

// GetTitle returns the Title field value
func (o *ApplicationApiKey) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// SetTitle sets field value
func (o *ApplicationApiKey) SetTitle(v string) {
	o.Title = v
}

// GetExpires returns the Expires field value
func (o *ApplicationApiKey) GetExpires() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expires
}

// SetExpires sets field value
func (o *ApplicationApiKey) SetExpires(v time.Time) {
	o.Expires = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ApplicationApiKey) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationApiKey) GetPlatformOk() (string, bool) {
	if o == nil || o.Platform == nil {
		var ret string
		return ret, false
	}
	return *o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ApplicationApiKey) HasPlatform() bool {
	if o != nil && o.Platform != nil {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *ApplicationApiKey) SetPlatform(v string) {
	o.Platform = &v
}

// GetId returns the Id field value
func (o *ApplicationApiKey) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *ApplicationApiKey) SetId(v int32) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ApplicationApiKey) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// SetCreatedBy sets field value
func (o *ApplicationApiKey) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetAccountID returns the AccountID field value
func (o *ApplicationApiKey) GetAccountID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountID
}

// SetAccountID sets field value
func (o *ApplicationApiKey) SetAccountID(v int32) {
	o.AccountID = v
}

// GetApplicationID returns the ApplicationID field value
func (o *ApplicationApiKey) GetApplicationID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationID
}

// SetApplicationID sets field value
func (o *ApplicationApiKey) SetApplicationID(v int32) {
	o.ApplicationID = v
}

// GetCreated returns the Created field value
func (o *ApplicationApiKey) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *ApplicationApiKey) SetCreated(v time.Time) {
	o.Created = v
}

type NullableApplicationApiKey struct {
	Value        ApplicationApiKey
	ExplicitNull bool
}

func (v NullableApplicationApiKey) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationApiKey) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
