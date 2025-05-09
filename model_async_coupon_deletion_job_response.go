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

// AsyncCouponDeletionJobResponse struct for AsyncCouponDeletionJobResponse
type AsyncCouponDeletionJobResponse struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
}

// GetId returns the Id field value
func (o *AsyncCouponDeletionJobResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *AsyncCouponDeletionJobResponse) SetId(v int32) {
	o.Id = v
}

type NullableAsyncCouponDeletionJobResponse struct {
	Value        AsyncCouponDeletionJobResponse
	ExplicitNull bool
}

func (v NullableAsyncCouponDeletionJobResponse) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableAsyncCouponDeletionJobResponse) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
