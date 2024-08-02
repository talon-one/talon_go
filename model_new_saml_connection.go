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

// NewSamlConnection A new SAML 2.0 connection.
type NewSamlConnection struct {
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used. 
	AudienceURI *string `json:"audienceURI,omitempty"`
	// Determines if this SAML connection active.
	Enabled bool `json:"enabled"`
	// Identity Provider Entity ID.
	Issuer string `json:"issuer"`
	// Metadata URL.
	MetadataURL *string `json:"metadataURL,omitempty"`
	// ID of the SAML service.
	Name string `json:"name"`
	// Single Sign-On URL.
	SignOnURL string `json:"signOnURL"`
	// Single Sign-Out URL.
	SignOutURL *string `json:"signOutURL,omitempty"`
	// X.509 Certificate.
	X509certificate string `json:"x509certificate"`
}

// GetAccountId returns the AccountId field value
func (o *NewSamlConnection) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *NewSamlConnection) SetAccountId(v int32) {
	o.AccountId = v
}

// GetAudienceURI returns the AudienceURI field value if set, zero value otherwise.
func (o *NewSamlConnection) GetAudienceURI() string {
	if o == nil || o.AudienceURI == nil {
		var ret string
		return ret
	}
	return *o.AudienceURI
}

// GetAudienceURIOk returns a tuple with the AudienceURI field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewSamlConnection) GetAudienceURIOk() (string, bool) {
	if o == nil || o.AudienceURI == nil {
		var ret string
		return ret, false
	}
	return *o.AudienceURI, true
}

// HasAudienceURI returns a boolean if a field has been set.
func (o *NewSamlConnection) HasAudienceURI() bool {
	if o != nil && o.AudienceURI != nil {
		return true
	}

	return false
}

// SetAudienceURI gets a reference to the given string and assigns it to the AudienceURI field.
func (o *NewSamlConnection) SetAudienceURI(v string) {
	o.AudienceURI = &v
}

// GetEnabled returns the Enabled field value
func (o *NewSamlConnection) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// SetEnabled sets field value
func (o *NewSamlConnection) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIssuer returns the Issuer field value
func (o *NewSamlConnection) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// SetIssuer sets field value
func (o *NewSamlConnection) SetIssuer(v string) {
	o.Issuer = v
}

// GetMetadataURL returns the MetadataURL field value if set, zero value otherwise.
func (o *NewSamlConnection) GetMetadataURL() string {
	if o == nil || o.MetadataURL == nil {
		var ret string
		return ret
	}
	return *o.MetadataURL
}

// GetMetadataURLOk returns a tuple with the MetadataURL field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewSamlConnection) GetMetadataURLOk() (string, bool) {
	if o == nil || o.MetadataURL == nil {
		var ret string
		return ret, false
	}
	return *o.MetadataURL, true
}

// HasMetadataURL returns a boolean if a field has been set.
func (o *NewSamlConnection) HasMetadataURL() bool {
	if o != nil && o.MetadataURL != nil {
		return true
	}

	return false
}

// SetMetadataURL gets a reference to the given string and assigns it to the MetadataURL field.
func (o *NewSamlConnection) SetMetadataURL(v string) {
	o.MetadataURL = &v
}

// GetName returns the Name field value
func (o *NewSamlConnection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *NewSamlConnection) SetName(v string) {
	o.Name = v
}

// GetSignOnURL returns the SignOnURL field value
func (o *NewSamlConnection) GetSignOnURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignOnURL
}

// SetSignOnURL sets field value
func (o *NewSamlConnection) SetSignOnURL(v string) {
	o.SignOnURL = v
}

// GetSignOutURL returns the SignOutURL field value if set, zero value otherwise.
func (o *NewSamlConnection) GetSignOutURL() string {
	if o == nil || o.SignOutURL == nil {
		var ret string
		return ret
	}
	return *o.SignOutURL
}

// GetSignOutURLOk returns a tuple with the SignOutURL field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *NewSamlConnection) GetSignOutURLOk() (string, bool) {
	if o == nil || o.SignOutURL == nil {
		var ret string
		return ret, false
	}
	return *o.SignOutURL, true
}

// HasSignOutURL returns a boolean if a field has been set.
func (o *NewSamlConnection) HasSignOutURL() bool {
	if o != nil && o.SignOutURL != nil {
		return true
	}

	return false
}

// SetSignOutURL gets a reference to the given string and assigns it to the SignOutURL field.
func (o *NewSamlConnection) SetSignOutURL(v string) {
	o.SignOutURL = &v
}

// GetX509certificate returns the X509certificate field value
func (o *NewSamlConnection) GetX509certificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509certificate
}

// SetX509certificate sets field value
func (o *NewSamlConnection) SetX509certificate(v string) {
	o.X509certificate = v
}

type NullableNewSamlConnection struct {
	Value NewSamlConnection
	ExplicitNull bool
}

func (v NullableNewSamlConnection) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableNewSamlConnection) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
