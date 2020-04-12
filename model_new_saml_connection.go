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

// NewSamlConnection
type NewSamlConnection struct {
	// X.509 Certificate.
	X509certificate string `json:"x509certificate"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// ID of the SAML service.
	Name string `json:"name"`
	// Determines if this SAML connection active.
	Enabled bool `json:"enabled"`
	// Identity Provider Entity ID.
	Issuer string `json:"issuer"`
	// Single Sign-On URL.
	SignOnURL string `json:"signOnURL"`
	// Single Sign-Out URL.
	SignOutURL *string `json:"signOutURL,omitempty"`
	// Metadata URL.
	MetadataURL *string `json:"metadataURL,omitempty"`
	// The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used.
	AudienceURI *string `json:"audienceURI,omitempty"`
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

type NullableNewSamlConnection struct {
	Value        NewSamlConnection
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
