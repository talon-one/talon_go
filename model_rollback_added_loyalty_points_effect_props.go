/*
 * Talon.One API
 *
 * Use the Talon.One API to integrate with your application and to manage applications and campaigns:  - Use the operations in the [Integration API section](#integration-api) are used to integrate with our platform - Use the operation in the [Management API section](#management-api) to manage applications and campaigns.  ## Determining the base URL of the endpoints  The API is available at the same hostname as your Campaign Manager deployment. For example, if you access the Campaign Manager at `https://yourbaseurl.talon.one/`, the URL for the [updateCustomerSessionV2](https://docs.talon.one/integration-api#operation/updateCustomerSessionV2) endpoint is `https://yourbaseurl.talon.one/v2/customer_sessions/{Id}` 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package talon

import (
	"bytes"
	"encoding/json"
)

// RollbackAddedLoyaltyPointsEffectProps The properties specific to the \"rollbackAddedLoyaltyPoints\" effect. This gets triggered whenever previously a closed session with an addLoyaltyPoints effect is cancelled.
type RollbackAddedLoyaltyPointsEffectProps struct {
	// The ID of the loyalty program where the points were originally added.
	ProgramId int32 `json:"programId"`
	// The ID of the subledger within the loyalty program where these points were originally added.
	SubLedgerId string `json:"subLedgerId"`
	// The amount of points that were rolled back.
	Value float32 `json:"value"`
	// The user for whom these points were originally added.
	RecipientIntegrationId string `json:"recipientIntegrationId"`
	// The identifier of 'deduction' entry added to the ledger as the `addLoyaltyPoints` effect is rolled back.
	TransactionUUID string `json:"transactionUUID"`
	// The index of the item in the cart items for which the loyalty points were rolled back.
	CartItemPosition *float32 `json:"cartItemPosition,omitempty"`
	// For cart items with `quantity` > 1, the sub-position indicates to which item the loyalty points were rolled back. 
	CartItemSubPosition *float32 `json:"cartItemSubPosition,omitempty"`
	// The alphanumeric identifier of the loyalty card. 
	CardIdentifier *string `json:"cardIdentifier,omitempty"`
}

// GetProgramId returns the ProgramId field value
func (o *RollbackAddedLoyaltyPointsEffectProps) GetProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramId
}

// SetProgramId sets field value
func (o *RollbackAddedLoyaltyPointsEffectProps) SetProgramId(v int32) {
	o.ProgramId = v
}

// GetSubLedgerId returns the SubLedgerId field value
func (o *RollbackAddedLoyaltyPointsEffectProps) GetSubLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubLedgerId
}

// SetSubLedgerId sets field value
func (o *RollbackAddedLoyaltyPointsEffectProps) SetSubLedgerId(v string) {
	o.SubLedgerId = v
}

// GetValue returns the Value field value
func (o *RollbackAddedLoyaltyPointsEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// SetValue sets field value
func (o *RollbackAddedLoyaltyPointsEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value
func (o *RollbackAddedLoyaltyPointsEffectProps) GetRecipientIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientIntegrationId
}

// SetRecipientIntegrationId sets field value
func (o *RollbackAddedLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = v
}

// GetTransactionUUID returns the TransactionUUID field value
func (o *RollbackAddedLoyaltyPointsEffectProps) GetTransactionUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionUUID
}

// SetTransactionUUID sets field value
func (o *RollbackAddedLoyaltyPointsEffectProps) SetTransactionUUID(v string) {
	o.TransactionUUID = v
}

// GetCartItemPosition returns the CartItemPosition field value if set, zero value otherwise.
func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemPosition() float32 {
	if o == nil || o.CartItemPosition == nil {
		var ret float32
		return ret
	}
	return *o.CartItemPosition
}

// GetCartItemPositionOk returns a tuple with the CartItemPosition field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemPositionOk() (float32, bool) {
	if o == nil || o.CartItemPosition == nil {
		var ret float32
		return ret, false
	}
	return *o.CartItemPosition, true
}

// HasCartItemPosition returns a boolean if a field has been set.
func (o *RollbackAddedLoyaltyPointsEffectProps) HasCartItemPosition() bool {
	if o != nil && o.CartItemPosition != nil {
		return true
	}

	return false
}

// SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.
func (o *RollbackAddedLoyaltyPointsEffectProps) SetCartItemPosition(v float32) {
	o.CartItemPosition = &v
}

// GetCartItemSubPosition returns the CartItemSubPosition field value if set, zero value otherwise.
func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemSubPosition() float32 {
	if o == nil || o.CartItemSubPosition == nil {
		var ret float32
		return ret
	}
	return *o.CartItemSubPosition
}

// GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RollbackAddedLoyaltyPointsEffectProps) GetCartItemSubPositionOk() (float32, bool) {
	if o == nil || o.CartItemSubPosition == nil {
		var ret float32
		return ret, false
	}
	return *o.CartItemSubPosition, true
}

// HasCartItemSubPosition returns a boolean if a field has been set.
func (o *RollbackAddedLoyaltyPointsEffectProps) HasCartItemSubPosition() bool {
	if o != nil && o.CartItemSubPosition != nil {
		return true
	}

	return false
}

// SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.
func (o *RollbackAddedLoyaltyPointsEffectProps) SetCartItemSubPosition(v float32) {
	o.CartItemSubPosition = &v
}

// GetCardIdentifier returns the CardIdentifier field value if set, zero value otherwise.
func (o *RollbackAddedLoyaltyPointsEffectProps) GetCardIdentifier() string {
	if o == nil || o.CardIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CardIdentifier
}

// GetCardIdentifierOk returns a tuple with the CardIdentifier field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *RollbackAddedLoyaltyPointsEffectProps) GetCardIdentifierOk() (string, bool) {
	if o == nil || o.CardIdentifier == nil {
		var ret string
		return ret, false
	}
	return *o.CardIdentifier, true
}

// HasCardIdentifier returns a boolean if a field has been set.
func (o *RollbackAddedLoyaltyPointsEffectProps) HasCardIdentifier() bool {
	if o != nil && o.CardIdentifier != nil {
		return true
	}

	return false
}

// SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.
func (o *RollbackAddedLoyaltyPointsEffectProps) SetCardIdentifier(v string) {
	o.CardIdentifier = &v
}

type NullableRollbackAddedLoyaltyPointsEffectProps struct {
	Value RollbackAddedLoyaltyPointsEffectProps
	ExplicitNull bool
}

func (v NullableRollbackAddedLoyaltyPointsEffectProps) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableRollbackAddedLoyaltyPointsEffectProps) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
