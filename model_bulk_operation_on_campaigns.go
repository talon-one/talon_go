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

// BulkOperationOnCampaigns struct for BulkOperationOnCampaigns
type BulkOperationOnCampaigns struct {
	// The operation to perform on the specified campaign IDs.
	Operation string `json:"operation"`
	// The list of campaign IDs on which the operation will be performed.
	CampaignIds []int32 `json:"campaignIds"`
}

// GetOperation returns the Operation field value
func (o *BulkOperationOnCampaigns) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// SetOperation sets field value
func (o *BulkOperationOnCampaigns) SetOperation(v string) {
	o.Operation = v
}

// GetCampaignIds returns the CampaignIds field value
func (o *BulkOperationOnCampaigns) GetCampaignIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.CampaignIds
}

// SetCampaignIds sets field value
func (o *BulkOperationOnCampaigns) SetCampaignIds(v []int32) {
	o.CampaignIds = v
}

type NullableBulkOperationOnCampaigns struct {
	Value        BulkOperationOnCampaigns
	ExplicitNull bool
}

func (v NullableBulkOperationOnCampaigns) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableBulkOperationOnCampaigns) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
