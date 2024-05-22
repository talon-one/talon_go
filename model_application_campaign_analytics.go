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

// ApplicationCampaignAnalytics struct for ApplicationCampaignAnalytics
type ApplicationCampaignAnalytics struct {
	// The start of the aggregation time frame in UTC.
	StartTime *time.Time `json:"startTime,omitempty"`
	// The end of the aggregation time frame in UTC.
	EndTime *time.Time `json:"endTime,omitempty"`
	// The ID of the campaign.
	CampaignId *int32 `json:"campaignId,omitempty"`
	// The name of the campaign.
	CampaignName *string `json:"campaignName,omitempty"`
	// A list of tags for the campaign.
	CampaignTags *[]string `json:"campaignTags,omitempty"`
	// The state of the campaign.  **Note:** A disabled or archived campaign is not evaluated for rules or coupons.
	CampaignState *string `json:"campaignState,omitempty"`
	// The [ID of the ruleset](https://docs.talon.one/management-api#operation/getRulesets) this campaign applies on customer session evaluation.
	CampaignActiveRulesetId *int32 `json:"campaignActiveRulesetId,omitempty"`
	// Date and time when the campaign becomes active.
	CampaignStartTime *time.Time `json:"campaignStartTime,omitempty"`
	// Date and time when the campaign becomes inactive.
	CampaignEndTime    *time.Time                                      `json:"campaignEndTime,omitempty"`
	TotalRevenue       *ApplicationCampaignAnalyticsTotalRevenue       `json:"totalRevenue,omitempty"`
	SessionsCount      *ApplicationCampaignAnalyticsSessionsCount      `json:"sessionsCount,omitempty"`
	AvgItemsPerSession *ApplicationCampaignAnalyticsAvgItemsPerSession `json:"avgItemsPerSession,omitempty"`
	AvgSessionValue    *ApplicationCampaignAnalyticsAvgSessionValue    `json:"avgSessionValue,omitempty"`
	TotalDiscounts     *ApplicationCampaignAnalyticsTotalDiscounts     `json:"totalDiscounts,omitempty"`
	CouponsCount       *ApplicationCampaignAnalyticsCouponsCount       `json:"couponsCount,omitempty"`
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetStartTimeOk() (time.Time, bool) {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApplicationCampaignAnalytics) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetEndTime() time.Time {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetEndTimeOk() (time.Time, bool) {
	if o == nil || o.EndTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasEndTime() bool {
	if o != nil && o.EndTime != nil {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *ApplicationCampaignAnalytics) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCampaignId() int32 {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCampaignIdOk() (int32, bool) {
	if o == nil || o.CampaignId == nil {
		var ret int32
		return ret, false
	}
	return *o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCampaignId() bool {
	if o != nil && o.CampaignId != nil {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.
func (o *ApplicationCampaignAnalytics) SetCampaignId(v int32) {
	o.CampaignId = &v
}

// GetCampaignName returns the CampaignName field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCampaignName() string {
	if o == nil || o.CampaignName == nil {
		var ret string
		return ret
	}
	return *o.CampaignName
}

// GetCampaignNameOk returns a tuple with the CampaignName field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCampaignNameOk() (string, bool) {
	if o == nil || o.CampaignName == nil {
		var ret string
		return ret, false
	}
	return *o.CampaignName, true
}

// HasCampaignName returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCampaignName() bool {
	if o != nil && o.CampaignName != nil {
		return true
	}

	return false
}

// SetCampaignName gets a reference to the given string and assigns it to the CampaignName field.
func (o *ApplicationCampaignAnalytics) SetCampaignName(v string) {
	o.CampaignName = &v
}

// GetCampaignTags returns the CampaignTags field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCampaignTags() []string {
	if o == nil || o.CampaignTags == nil {
		var ret []string
		return ret
	}
	return *o.CampaignTags
}

// GetCampaignTagsOk returns a tuple with the CampaignTags field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCampaignTagsOk() ([]string, bool) {
	if o == nil || o.CampaignTags == nil {
		var ret []string
		return ret, false
	}
	return *o.CampaignTags, true
}

// HasCampaignTags returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCampaignTags() bool {
	if o != nil && o.CampaignTags != nil {
		return true
	}

	return false
}

// SetCampaignTags gets a reference to the given []string and assigns it to the CampaignTags field.
func (o *ApplicationCampaignAnalytics) SetCampaignTags(v []string) {
	o.CampaignTags = &v
}

// GetCampaignState returns the CampaignState field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCampaignState() string {
	if o == nil || o.CampaignState == nil {
		var ret string
		return ret
	}
	return *o.CampaignState
}

// GetCampaignStateOk returns a tuple with the CampaignState field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCampaignStateOk() (string, bool) {
	if o == nil || o.CampaignState == nil {
		var ret string
		return ret, false
	}
	return *o.CampaignState, true
}

// HasCampaignState returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCampaignState() bool {
	if o != nil && o.CampaignState != nil {
		return true
	}

	return false
}

// SetCampaignState gets a reference to the given string and assigns it to the CampaignState field.
func (o *ApplicationCampaignAnalytics) SetCampaignState(v string) {
	o.CampaignState = &v
}

// GetCampaignActiveRulesetId returns the CampaignActiveRulesetId field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCampaignActiveRulesetId() int32 {
	if o == nil || o.CampaignActiveRulesetId == nil {
		var ret int32
		return ret
	}
	return *o.CampaignActiveRulesetId
}

// GetCampaignActiveRulesetIdOk returns a tuple with the CampaignActiveRulesetId field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCampaignActiveRulesetIdOk() (int32, bool) {
	if o == nil || o.CampaignActiveRulesetId == nil {
		var ret int32
		return ret, false
	}
	return *o.CampaignActiveRulesetId, true
}

// HasCampaignActiveRulesetId returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCampaignActiveRulesetId() bool {
	if o != nil && o.CampaignActiveRulesetId != nil {
		return true
	}

	return false
}

// SetCampaignActiveRulesetId gets a reference to the given int32 and assigns it to the CampaignActiveRulesetId field.
func (o *ApplicationCampaignAnalytics) SetCampaignActiveRulesetId(v int32) {
	o.CampaignActiveRulesetId = &v
}

// GetCampaignStartTime returns the CampaignStartTime field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCampaignStartTime() time.Time {
	if o == nil || o.CampaignStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CampaignStartTime
}

// GetCampaignStartTimeOk returns a tuple with the CampaignStartTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCampaignStartTimeOk() (time.Time, bool) {
	if o == nil || o.CampaignStartTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CampaignStartTime, true
}

// HasCampaignStartTime returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCampaignStartTime() bool {
	if o != nil && o.CampaignStartTime != nil {
		return true
	}

	return false
}

// SetCampaignStartTime gets a reference to the given time.Time and assigns it to the CampaignStartTime field.
func (o *ApplicationCampaignAnalytics) SetCampaignStartTime(v time.Time) {
	o.CampaignStartTime = &v
}

// GetCampaignEndTime returns the CampaignEndTime field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCampaignEndTime() time.Time {
	if o == nil || o.CampaignEndTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CampaignEndTime
}

// GetCampaignEndTimeOk returns a tuple with the CampaignEndTime field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCampaignEndTimeOk() (time.Time, bool) {
	if o == nil || o.CampaignEndTime == nil {
		var ret time.Time
		return ret, false
	}
	return *o.CampaignEndTime, true
}

// HasCampaignEndTime returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCampaignEndTime() bool {
	if o != nil && o.CampaignEndTime != nil {
		return true
	}

	return false
}

// SetCampaignEndTime gets a reference to the given time.Time and assigns it to the CampaignEndTime field.
func (o *ApplicationCampaignAnalytics) SetCampaignEndTime(v time.Time) {
	o.CampaignEndTime = &v
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetTotalRevenue() ApplicationCampaignAnalyticsTotalRevenue {
	if o == nil || o.TotalRevenue == nil {
		var ret ApplicationCampaignAnalyticsTotalRevenue
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetTotalRevenueOk() (ApplicationCampaignAnalyticsTotalRevenue, bool) {
	if o == nil || o.TotalRevenue == nil {
		var ret ApplicationCampaignAnalyticsTotalRevenue
		return ret, false
	}
	return *o.TotalRevenue, true
}

// HasTotalRevenue returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasTotalRevenue() bool {
	if o != nil && o.TotalRevenue != nil {
		return true
	}

	return false
}

// SetTotalRevenue gets a reference to the given ApplicationCampaignAnalyticsTotalRevenue and assigns it to the TotalRevenue field.
func (o *ApplicationCampaignAnalytics) SetTotalRevenue(v ApplicationCampaignAnalyticsTotalRevenue) {
	o.TotalRevenue = &v
}

// GetSessionsCount returns the SessionsCount field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetSessionsCount() ApplicationCampaignAnalyticsSessionsCount {
	if o == nil || o.SessionsCount == nil {
		var ret ApplicationCampaignAnalyticsSessionsCount
		return ret
	}
	return *o.SessionsCount
}

// GetSessionsCountOk returns a tuple with the SessionsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetSessionsCountOk() (ApplicationCampaignAnalyticsSessionsCount, bool) {
	if o == nil || o.SessionsCount == nil {
		var ret ApplicationCampaignAnalyticsSessionsCount
		return ret, false
	}
	return *o.SessionsCount, true
}

// HasSessionsCount returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasSessionsCount() bool {
	if o != nil && o.SessionsCount != nil {
		return true
	}

	return false
}

// SetSessionsCount gets a reference to the given ApplicationCampaignAnalyticsSessionsCount and assigns it to the SessionsCount field.
func (o *ApplicationCampaignAnalytics) SetSessionsCount(v ApplicationCampaignAnalyticsSessionsCount) {
	o.SessionsCount = &v
}

// GetAvgItemsPerSession returns the AvgItemsPerSession field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetAvgItemsPerSession() ApplicationCampaignAnalyticsAvgItemsPerSession {
	if o == nil || o.AvgItemsPerSession == nil {
		var ret ApplicationCampaignAnalyticsAvgItemsPerSession
		return ret
	}
	return *o.AvgItemsPerSession
}

// GetAvgItemsPerSessionOk returns a tuple with the AvgItemsPerSession field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetAvgItemsPerSessionOk() (ApplicationCampaignAnalyticsAvgItemsPerSession, bool) {
	if o == nil || o.AvgItemsPerSession == nil {
		var ret ApplicationCampaignAnalyticsAvgItemsPerSession
		return ret, false
	}
	return *o.AvgItemsPerSession, true
}

// HasAvgItemsPerSession returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasAvgItemsPerSession() bool {
	if o != nil && o.AvgItemsPerSession != nil {
		return true
	}

	return false
}

// SetAvgItemsPerSession gets a reference to the given ApplicationCampaignAnalyticsAvgItemsPerSession and assigns it to the AvgItemsPerSession field.
func (o *ApplicationCampaignAnalytics) SetAvgItemsPerSession(v ApplicationCampaignAnalyticsAvgItemsPerSession) {
	o.AvgItemsPerSession = &v
}

// GetAvgSessionValue returns the AvgSessionValue field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetAvgSessionValue() ApplicationCampaignAnalyticsAvgSessionValue {
	if o == nil || o.AvgSessionValue == nil {
		var ret ApplicationCampaignAnalyticsAvgSessionValue
		return ret
	}
	return *o.AvgSessionValue
}

// GetAvgSessionValueOk returns a tuple with the AvgSessionValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetAvgSessionValueOk() (ApplicationCampaignAnalyticsAvgSessionValue, bool) {
	if o == nil || o.AvgSessionValue == nil {
		var ret ApplicationCampaignAnalyticsAvgSessionValue
		return ret, false
	}
	return *o.AvgSessionValue, true
}

// HasAvgSessionValue returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasAvgSessionValue() bool {
	if o != nil && o.AvgSessionValue != nil {
		return true
	}

	return false
}

// SetAvgSessionValue gets a reference to the given ApplicationCampaignAnalyticsAvgSessionValue and assigns it to the AvgSessionValue field.
func (o *ApplicationCampaignAnalytics) SetAvgSessionValue(v ApplicationCampaignAnalyticsAvgSessionValue) {
	o.AvgSessionValue = &v
}

// GetTotalDiscounts returns the TotalDiscounts field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetTotalDiscounts() ApplicationCampaignAnalyticsTotalDiscounts {
	if o == nil || o.TotalDiscounts == nil {
		var ret ApplicationCampaignAnalyticsTotalDiscounts
		return ret
	}
	return *o.TotalDiscounts
}

// GetTotalDiscountsOk returns a tuple with the TotalDiscounts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetTotalDiscountsOk() (ApplicationCampaignAnalyticsTotalDiscounts, bool) {
	if o == nil || o.TotalDiscounts == nil {
		var ret ApplicationCampaignAnalyticsTotalDiscounts
		return ret, false
	}
	return *o.TotalDiscounts, true
}

// HasTotalDiscounts returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasTotalDiscounts() bool {
	if o != nil && o.TotalDiscounts != nil {
		return true
	}

	return false
}

// SetTotalDiscounts gets a reference to the given ApplicationCampaignAnalyticsTotalDiscounts and assigns it to the TotalDiscounts field.
func (o *ApplicationCampaignAnalytics) SetTotalDiscounts(v ApplicationCampaignAnalyticsTotalDiscounts) {
	o.TotalDiscounts = &v
}

// GetCouponsCount returns the CouponsCount field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCouponsCount() ApplicationCampaignAnalyticsCouponsCount {
	if o == nil || o.CouponsCount == nil {
		var ret ApplicationCampaignAnalyticsCouponsCount
		return ret
	}
	return *o.CouponsCount
}

// GetCouponsCountOk returns a tuple with the CouponsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCouponsCountOk() (ApplicationCampaignAnalyticsCouponsCount, bool) {
	if o == nil || o.CouponsCount == nil {
		var ret ApplicationCampaignAnalyticsCouponsCount
		return ret, false
	}
	return *o.CouponsCount, true
}

// HasCouponsCount returns a boolean if a field has been set.
func (o *ApplicationCampaignAnalytics) HasCouponsCount() bool {
	if o != nil && o.CouponsCount != nil {
		return true
	}

	return false
}

// SetCouponsCount gets a reference to the given ApplicationCampaignAnalyticsCouponsCount and assigns it to the CouponsCount field.
func (o *ApplicationCampaignAnalytics) SetCouponsCount(v ApplicationCampaignAnalyticsCouponsCount) {
	o.CouponsCount = &v
}

type NullableApplicationCampaignAnalytics struct {
	Value        ApplicationCampaignAnalytics
	ExplicitNull bool
}

func (v NullableApplicationCampaignAnalytics) MarshalJSON() ([]byte, error) {
	switch {
	case v.ExplicitNull:
		return []byte("null"), nil
	default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableApplicationCampaignAnalytics) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
