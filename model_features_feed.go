/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you are reading this page at `https://mycompany.talon.one/docs/api/`, the URL for the [updateCustomerSession](https://docs.talon.one/integration-api/#operation/updateCustomerSessionV2) endpoint is `https://mycompany.talon.one/v2/customer_sessions/{Id}`
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// FeaturesFeed
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
