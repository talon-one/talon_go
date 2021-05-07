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

// CampaignGroup struct for CampaignGroup
type CampaignGroup struct {
	// Unique ID for this entity.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The exact moment this entity was last modified.
	Modified time.Time `json:"modified"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The name of this campaign group.
	Name string `json:"name"`
	// A longer description of the campaign group.
	Description *string `json:"description,omitempty"`
	// A list of the IDs of the applications that this campaign group is enabled for
	SubscribedApplicationsIds *[]int32 `json:"subscribedApplicationsIds,omitempty"`
	// A list of the IDs of the campaigns that this campaign group owns
	CampaignIds *[]int32 `json:"campaignIds,omitempty"`
}

// GetId returns the Id field value
func (o *CampaignGroup) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *CampaignGroup) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CampaignGroup) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CampaignGroup) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *CampaignGroup) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// SetModified sets field value
func (o *CampaignGroup) SetModified(v time.Time) {
	o.Modified = v
}

// GetAccountId returns the AccountId field value
func (o *CampaignGroup) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *CampaignGroup) SetAccountId(v int32) {
	o.AccountId = v
}

// GetName returns the Name field value
func (o *CampaignGroup) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// SetName sets field value
func (o *CampaignGroup) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CampaignGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignGroup) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CampaignGroup) SetDescription(v string) {
	o.Description = &v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *CampaignGroup) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return *o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignGroup) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *CampaignGroup) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *CampaignGroup) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = &v
}

// GetCampaignIds returns the CampaignIds field value if set, zero value otherwise.
func (o *CampaignGroup) GetCampaignIds() []int32 {
	if o == nil || o.CampaignIds == nil {
		var ret []int32
		return ret
	}
	return *o.CampaignIds
}

// GetCampaignIdsOk returns a tuple with the CampaignIds field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CampaignGroup) GetCampaignIdsOk() ([]int32, bool) {
	if o == nil || o.CampaignIds == nil {
		var ret []int32
		return ret, false
	}
	return *o.CampaignIds, true
}

// HasCampaignIds returns a boolean if a field has been set.
func (o *CampaignGroup) HasCampaignIds() bool {
	if o != nil && o.CampaignIds != nil {
		return true
	}

	return false
}

// SetCampaignIds gets a reference to the given []int32 and assigns it to the CampaignIds field.
func (o *CampaignGroup) SetCampaignIds(v []int32) {
	o.CampaignIds = &v
}

type NullableCampaignGroup struct {
	Value        CampaignGroup
	ExplicitNull bool
}

func (v NullableCampaignGroup) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCampaignGroup) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
