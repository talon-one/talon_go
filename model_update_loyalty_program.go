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

// UpdateLoyaltyProgram
type UpdateLoyaltyProgram struct {
	// The display title for the Loyalty Program.
	Title *string `json:"title,omitempty"`
	// Description of our Loyalty Program.
	Description *string `json:"description,omitempty"`
	// A list containing the IDs of all applications that are subscribed to this Loyalty Program.
	SubscribedApplications *[]int32 `json:"subscribedApplications,omitempty"`
	// The default duration after which new loyalty points should expire. Can be 'unlimited' or a specific time. The time format is a number followed by one letter indicating the time unit, like '30s', '40m', '1h', '5D', '7W', or 10M'. These rounding suffixes are also supported: - '_D' for rounding down. Can be used as a suffix after 'D', and signifies the start of the day. - '_U' for rounding up. Can be used as a suffix after 'D', 'W', and 'M', and signifies the end of the day, week, and month.
	DefaultValidity *string `json:"defaultValidity,omitempty"`
	// The default duration of the pending time after which points should be valid. Can be 'immediate' or a specific time. The time format is a number followed by one letter indicating the time unit, like '30s', '40m', '1h', '5D', '7W', or 10M'. These rounding suffixes are also supported: - '_D' for rounding down. Can be used as a suffix after 'D', and signifies the start of the day. - '_U' for rounding up. Can be used as a suffix after 'D', 'W', and 'M', and signifies the end of the day, week, and month.
	DefaultPending *string `json:"defaultPending,omitempty"`
	// Indicates if this program supports subledgers inside the program.
	AllowSubledger *bool `json:"allowSubledger,omitempty"`
	// The max amount of user profiles with whom a card can be shared. This can be set to 0 for no limit. This property is only used when `cardBased` is `true`.
	UsersPerCardLimit *int32 `json:"usersPerCardLimit,omitempty"`
	// Indicates if this program is a live or sandbox program. Programs of a given type can only be connected to Applications of the same type.
	Sandbox *bool `json:"sandbox,omitempty"`
	// The policy that defines which date is used to calculate the expiration date of a customer's current tier.  - `tier_start_date`: The tier expiration date is calculated based on when the customer joined the current tier.  - `program_join_date`: The tier expiration date is calculated based on when the customer joined the loyalty program.
	TiersExpirationPolicy *string `json:"tiersExpirationPolicy,omitempty"`
	// The amount of time after which the tier expires.  The time format is an **integer** followed by one letter indicating the time unit. Examples: `30s`, `40m`, `1h`, `5D`, `7W`, `10M`, `15Y`.  Available units:  - `s`: seconds - `m`: minutes - `h`: hours - `D`: days - `W`: weeks - `M`: months - `Y`: years  You can round certain units up or down: - `_D` for rounding down days only. Signifies the start of the day. - `_U` for rounding up days, weeks, months and years. Signifies the end of the day, week, month or year.
	TiersExpireIn *string `json:"tiersExpireIn,omitempty"`
	// Customers's tier downgrade policy.  - `one_down`: Once the tier expires and if the user doesn't have enough points to stay in the tier, the user is downgraded one tier down.  - `balance_based`: Once the tier expires, the user's tier is evaluated based on the amount of active points the user has at this instant.
	TiersDowngradePolicy *string `json:"tiersDowngradePolicy,omitempty"`
	// The policy that defines when the customer joins the loyalty program.   - `not_join`: The customer does not join the loyalty program but can still earn and spend loyalty points.       **Note**: The customer does not have a program join date.   - `points_activated`: The customer joins the loyalty program only when their earned loyalty points become active for the first time.   - `points_earned`: The customer joins the loyalty program when they earn loyalty points for the first time.
	ProgramJoinPolicy *string `json:"programJoinPolicy,omitempty"`
	// The tiers in this loyalty program.
	Tiers *[]NewLoyaltyTier `json:"tiers,omitempty"`
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetTitleOk() (string, bool) {
	if o == nil || o.Title == nil {
		var ret string
		return ret, false
	}
	return *o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateLoyaltyProgram) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetDescriptionOk() (string, bool) {
	if o == nil || o.Description == nil {
		var ret string
		return ret, false
	}
	return *o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateLoyaltyProgram) SetDescription(v string) {
	o.Description = &v
}

// GetSubscribedApplications returns the SubscribedApplications field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetSubscribedApplications() []int32 {
	if o == nil || o.SubscribedApplications == nil {
		var ret []int32
		return ret
	}
	return *o.SubscribedApplications
}

// GetSubscribedApplicationsOk returns a tuple with the SubscribedApplications field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetSubscribedApplicationsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplications == nil {
		var ret []int32
		return ret, false
	}
	return *o.SubscribedApplications, true
}

// HasSubscribedApplications returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasSubscribedApplications() bool {
	if o != nil && o.SubscribedApplications != nil {
		return true
	}

	return false
}

// SetSubscribedApplications gets a reference to the given []int32 and assigns it to the SubscribedApplications field.
func (o *UpdateLoyaltyProgram) SetSubscribedApplications(v []int32) {
	o.SubscribedApplications = &v
}

// GetDefaultValidity returns the DefaultValidity field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetDefaultValidity() string {
	if o == nil || o.DefaultValidity == nil {
		var ret string
		return ret
	}
	return *o.DefaultValidity
}

// GetDefaultValidityOk returns a tuple with the DefaultValidity field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetDefaultValidityOk() (string, bool) {
	if o == nil || o.DefaultValidity == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultValidity, true
}

// HasDefaultValidity returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasDefaultValidity() bool {
	if o != nil && o.DefaultValidity != nil {
		return true
	}

	return false
}

// SetDefaultValidity gets a reference to the given string and assigns it to the DefaultValidity field.
func (o *UpdateLoyaltyProgram) SetDefaultValidity(v string) {
	o.DefaultValidity = &v
}

// GetDefaultPending returns the DefaultPending field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetDefaultPending() string {
	if o == nil || o.DefaultPending == nil {
		var ret string
		return ret
	}
	return *o.DefaultPending
}

// GetDefaultPendingOk returns a tuple with the DefaultPending field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetDefaultPendingOk() (string, bool) {
	if o == nil || o.DefaultPending == nil {
		var ret string
		return ret, false
	}
	return *o.DefaultPending, true
}

// HasDefaultPending returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasDefaultPending() bool {
	if o != nil && o.DefaultPending != nil {
		return true
	}

	return false
}

// SetDefaultPending gets a reference to the given string and assigns it to the DefaultPending field.
func (o *UpdateLoyaltyProgram) SetDefaultPending(v string) {
	o.DefaultPending = &v
}

// GetAllowSubledger returns the AllowSubledger field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetAllowSubledger() bool {
	if o == nil || o.AllowSubledger == nil {
		var ret bool
		return ret
	}
	return *o.AllowSubledger
}

// GetAllowSubledgerOk returns a tuple with the AllowSubledger field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetAllowSubledgerOk() (bool, bool) {
	if o == nil || o.AllowSubledger == nil {
		var ret bool
		return ret, false
	}
	return *o.AllowSubledger, true
}

// HasAllowSubledger returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasAllowSubledger() bool {
	if o != nil && o.AllowSubledger != nil {
		return true
	}

	return false
}

// SetAllowSubledger gets a reference to the given bool and assigns it to the AllowSubledger field.
func (o *UpdateLoyaltyProgram) SetAllowSubledger(v bool) {
	o.AllowSubledger = &v
}

// GetUsersPerCardLimit returns the UsersPerCardLimit field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetUsersPerCardLimit() int32 {
	if o == nil || o.UsersPerCardLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsersPerCardLimit
}

// GetUsersPerCardLimitOk returns a tuple with the UsersPerCardLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetUsersPerCardLimitOk() (int32, bool) {
	if o == nil || o.UsersPerCardLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.UsersPerCardLimit, true
}

// HasUsersPerCardLimit returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasUsersPerCardLimit() bool {
	if o != nil && o.UsersPerCardLimit != nil {
		return true
	}

	return false
}

// SetUsersPerCardLimit gets a reference to the given int32 and assigns it to the UsersPerCardLimit field.
func (o *UpdateLoyaltyProgram) SetUsersPerCardLimit(v int32) {
	o.UsersPerCardLimit = &v
}

// GetSandbox returns the Sandbox field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetSandbox() bool {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret
	}
	return *o.Sandbox
}

// GetSandboxOk returns a tuple with the Sandbox field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetSandboxOk() (bool, bool) {
	if o == nil || o.Sandbox == nil {
		var ret bool
		return ret, false
	}
	return *o.Sandbox, true
}

// HasSandbox returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasSandbox() bool {
	if o != nil && o.Sandbox != nil {
		return true
	}

	return false
}

// SetSandbox gets a reference to the given bool and assigns it to the Sandbox field.
func (o *UpdateLoyaltyProgram) SetSandbox(v bool) {
	o.Sandbox = &v
}

// GetTiersExpirationPolicy returns the TiersExpirationPolicy field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetTiersExpirationPolicy() string {
	if o == nil || o.TiersExpirationPolicy == nil {
		var ret string
		return ret
	}
	return *o.TiersExpirationPolicy
}

// GetTiersExpirationPolicyOk returns a tuple with the TiersExpirationPolicy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetTiersExpirationPolicyOk() (string, bool) {
	if o == nil || o.TiersExpirationPolicy == nil {
		var ret string
		return ret, false
	}
	return *o.TiersExpirationPolicy, true
}

// HasTiersExpirationPolicy returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasTiersExpirationPolicy() bool {
	if o != nil && o.TiersExpirationPolicy != nil {
		return true
	}

	return false
}

// SetTiersExpirationPolicy gets a reference to the given string and assigns it to the TiersExpirationPolicy field.
func (o *UpdateLoyaltyProgram) SetTiersExpirationPolicy(v string) {
	o.TiersExpirationPolicy = &v
}

// GetTiersExpireIn returns the TiersExpireIn field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetTiersExpireIn() string {
	if o == nil || o.TiersExpireIn == nil {
		var ret string
		return ret
	}
	return *o.TiersExpireIn
}

// GetTiersExpireInOk returns a tuple with the TiersExpireIn field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetTiersExpireInOk() (string, bool) {
	if o == nil || o.TiersExpireIn == nil {
		var ret string
		return ret, false
	}
	return *o.TiersExpireIn, true
}

// HasTiersExpireIn returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasTiersExpireIn() bool {
	if o != nil && o.TiersExpireIn != nil {
		return true
	}

	return false
}

// SetTiersExpireIn gets a reference to the given string and assigns it to the TiersExpireIn field.
func (o *UpdateLoyaltyProgram) SetTiersExpireIn(v string) {
	o.TiersExpireIn = &v
}

// GetTiersDowngradePolicy returns the TiersDowngradePolicy field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetTiersDowngradePolicy() string {
	if o == nil || o.TiersDowngradePolicy == nil {
		var ret string
		return ret
	}
	return *o.TiersDowngradePolicy
}

// GetTiersDowngradePolicyOk returns a tuple with the TiersDowngradePolicy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetTiersDowngradePolicyOk() (string, bool) {
	if o == nil || o.TiersDowngradePolicy == nil {
		var ret string
		return ret, false
	}
	return *o.TiersDowngradePolicy, true
}

// HasTiersDowngradePolicy returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasTiersDowngradePolicy() bool {
	if o != nil && o.TiersDowngradePolicy != nil {
		return true
	}

	return false
}

// SetTiersDowngradePolicy gets a reference to the given string and assigns it to the TiersDowngradePolicy field.
func (o *UpdateLoyaltyProgram) SetTiersDowngradePolicy(v string) {
	o.TiersDowngradePolicy = &v
}

// GetProgramJoinPolicy returns the ProgramJoinPolicy field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetProgramJoinPolicy() string {
	if o == nil || o.ProgramJoinPolicy == nil {
		var ret string
		return ret
	}
	return *o.ProgramJoinPolicy
}

// GetProgramJoinPolicyOk returns a tuple with the ProgramJoinPolicy field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetProgramJoinPolicyOk() (string, bool) {
	if o == nil || o.ProgramJoinPolicy == nil {
		var ret string
		return ret, false
	}
	return *o.ProgramJoinPolicy, true
}

// HasProgramJoinPolicy returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasProgramJoinPolicy() bool {
	if o != nil && o.ProgramJoinPolicy != nil {
		return true
	}

	return false
}

// SetProgramJoinPolicy gets a reference to the given string and assigns it to the ProgramJoinPolicy field.
func (o *UpdateLoyaltyProgram) SetProgramJoinPolicy(v string) {
	o.ProgramJoinPolicy = &v
}

// GetTiers returns the Tiers field value if set, zero value otherwise.
func (o *UpdateLoyaltyProgram) GetTiers() []NewLoyaltyTier {
	if o == nil || o.Tiers == nil {
		var ret []NewLoyaltyTier
		return ret
	}
	return *o.Tiers
}

// GetTiersOk returns a tuple with the Tiers field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UpdateLoyaltyProgram) GetTiersOk() ([]NewLoyaltyTier, bool) {
	if o == nil || o.Tiers == nil {
		var ret []NewLoyaltyTier
		return ret, false
	}
	return *o.Tiers, true
}

// HasTiers returns a boolean if a field has been set.
func (o *UpdateLoyaltyProgram) HasTiers() bool {
	if o != nil && o.Tiers != nil {
		return true
	}

	return false
}

// SetTiers gets a reference to the given []NewLoyaltyTier and assigns it to the Tiers field.
func (o *UpdateLoyaltyProgram) SetTiers(v []NewLoyaltyTier) {
	o.Tiers = &v
}

type NullableUpdateLoyaltyProgram struct {
	Value        UpdateLoyaltyProgram
	ExplicitNull bool
}

func (v NullableUpdateLoyaltyProgram) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableUpdateLoyaltyProgram) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
