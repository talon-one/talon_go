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

// FeaturesFeed struct for FeaturesFeed
type FeaturesFeed struct {
	Title   *string `json:"title,omitempty"`
	PubDate *string `json:"pubDate,omitempty"`
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *FeaturesFeed) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesFeed) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *FeaturesFeed) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *FeaturesFeed) SetTitle(v string) {
	o.Title = &v
}

// GetPubDate returns the PubDate field value if set, zero value otherwise.
func (o *FeaturesFeed) GetPubDate() string {
	if o == nil || o.PubDate == nil {
		var ret string
		return ret
	}
	return *o.PubDate
}

// GetPubDateOk returns a tuple with the PubDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *FeaturesFeed) GetPubDateOk() (string, bool) {
	if o == nil || o.PubDate == nil {
		var ret string
		return ret, false
	}
	return *o.PubDate, true
}

// HasPubDate returns a boolean if a field has been set.
func (o *FeaturesFeed) HasPubDate() bool {
	if o != nil && o.PubDate != nil {
		return true
	}

	return false
}

// SetPubDate gets a reference to the given string and assigns it to the PubDate field.
func (o *FeaturesFeed) SetPubDate(v string) {
	o.PubDate = &v
}

type NullableFeaturesFeed struct {
	Value        FeaturesFeed
	ExplicitNull bool
}

func (v NullableFeaturesFeed) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableFeaturesFeed) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
