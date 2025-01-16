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

// ApplicationReferee struct for ApplicationReferee
type ApplicationReferee struct {
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// Integration ID of the session in which the customer redeemed the referral.
	SessionId string `json:"sessionId"`
	// Integration ID of the Advocate's Profile.
	AdvocateIntegrationId string `json:"advocateIntegrationId"`
	// Integration ID of the Friend's Profile.
	FriendIntegrationId string `json:"friendIntegrationId"`
	// Advocate's referral code.
	Code string `json:"code"`
	// Timestamp of the moment the customer redeemed the referral.
	Created time.Time `json:"created"`
}

// GetApplicationId returns the ApplicationId field value
func (o *ApplicationReferee) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *ApplicationReferee) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetSessionId returns the SessionId field value
func (o *ApplicationReferee) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// SetSessionId sets field value
func (o *ApplicationReferee) SetSessionId(v string) {
	o.SessionId = v
}

// GetAdvocateIntegrationId returns the AdvocateIntegrationId field value
func (o *ApplicationReferee) GetAdvocateIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdvocateIntegrationId
}

// SetAdvocateIntegrationId sets field value
func (o *ApplicationReferee) SetAdvocateIntegrationId(v string) {
	o.AdvocateIntegrationId = v
}

// GetFriendIntegrationId returns the FriendIntegrationId field value
func (o *ApplicationReferee) GetFriendIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FriendIntegrationId
}

// SetFriendIntegrationId sets field value
func (o *ApplicationReferee) SetFriendIntegrationId(v string) {
	o.FriendIntegrationId = v
}

// GetCode returns the Code field value
func (o *ApplicationReferee) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// SetCode sets field value
func (o *ApplicationReferee) SetCode(v string) {
	o.Code = v
}

// GetCreated returns the Created field value
func (o *ApplicationReferee) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *ApplicationReferee) SetCreated(v time.Time) {
	o.Created = v
}

type NullableApplicationReferee struct {
	Value        ApplicationReferee
	ExplicitNull bool
}

func (v NullableApplicationReferee) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationReferee) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
