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
	StartTime time.Time `json:"startTime"`
	// The end of the aggregation time frame in UTC.
	EndTime time.Time `json:"endTime"`
	// The ID of the campaign.
	CampaignId int32 `json:"campaignId"`
	// The name of the campaign.
	CampaignName string `json:"campaignName"`
	// A list of tags for the campaign.
	CampaignTags []string `json:"campaignTags"`
	// The state of the campaign.  **Note:** A disabled or archived campaign is not evaluated for rules or coupons. 
	CampaignState string `json:"campaignState"`
	TotalRevenue *AnalyticsDataPointWithTrendAndInfluencedRate `json:"totalRevenue,omitempty"`
	SessionsCount *AnalyticsDataPointWithTrendAndInfluencedRate `json:"sessionsCount,omitempty"`
	AvgItemsPerSession *AnalyticsDataPointWithTrendAndUplift `json:"avgItemsPerSession,omitempty"`
	AvgSessionValue *AnalyticsDataPointWithTrendAndUplift `json:"avgSessionValue,omitempty"`
	TotalDiscounts *AnalyticsDataPointWithTrend `json:"totalDiscounts,omitempty"`
	CouponsCount *AnalyticsDataPointWithTrend `json:"couponsCount,omitempty"`
}

// GetStartTime returns the StartTime field value
func (o *ApplicationCampaignAnalytics) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// SetStartTime sets field value
func (o *ApplicationCampaignAnalytics) SetStartTime(v time.Time) {
	o.StartTime = v
}

// GetEndTime returns the EndTime field value
func (o *ApplicationCampaignAnalytics) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// SetEndTime sets field value
func (o *ApplicationCampaignAnalytics) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetCampaignId returns the CampaignId field value
func (o *ApplicationCampaignAnalytics) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *ApplicationCampaignAnalytics) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetCampaignName returns the CampaignName field value
func (o *ApplicationCampaignAnalytics) GetCampaignName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignName
}

// SetCampaignName sets field value
func (o *ApplicationCampaignAnalytics) SetCampaignName(v string) {
	o.CampaignName = v
}

// GetCampaignTags returns the CampaignTags field value
func (o *ApplicationCampaignAnalytics) GetCampaignTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CampaignTags
}

// SetCampaignTags sets field value
func (o *ApplicationCampaignAnalytics) SetCampaignTags(v []string) {
	o.CampaignTags = v
}

// GetCampaignState returns the CampaignState field value
func (o *ApplicationCampaignAnalytics) GetCampaignState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignState
}

// SetCampaignState sets field value
func (o *ApplicationCampaignAnalytics) SetCampaignState(v string) {
	o.CampaignState = v
}

// GetTotalRevenue returns the TotalRevenue field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetTotalRevenue() AnalyticsDataPointWithTrendAndInfluencedRate {
	if o == nil || o.TotalRevenue == nil {
		var ret AnalyticsDataPointWithTrendAndInfluencedRate
		return ret
	}
	return *o.TotalRevenue
}

// GetTotalRevenueOk returns a tuple with the TotalRevenue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetTotalRevenueOk() (AnalyticsDataPointWithTrendAndInfluencedRate, bool) {
	if o == nil || o.TotalRevenue == nil {
		var ret AnalyticsDataPointWithTrendAndInfluencedRate
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

// SetTotalRevenue gets a reference to the given AnalyticsDataPointWithTrendAndInfluencedRate and assigns it to the TotalRevenue field.
func (o *ApplicationCampaignAnalytics) SetTotalRevenue(v AnalyticsDataPointWithTrendAndInfluencedRate) {
	o.TotalRevenue = &v
}

// GetSessionsCount returns the SessionsCount field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetSessionsCount() AnalyticsDataPointWithTrendAndInfluencedRate {
	if o == nil || o.SessionsCount == nil {
		var ret AnalyticsDataPointWithTrendAndInfluencedRate
		return ret
	}
	return *o.SessionsCount
}

// GetSessionsCountOk returns a tuple with the SessionsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetSessionsCountOk() (AnalyticsDataPointWithTrendAndInfluencedRate, bool) {
	if o == nil || o.SessionsCount == nil {
		var ret AnalyticsDataPointWithTrendAndInfluencedRate
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

// SetSessionsCount gets a reference to the given AnalyticsDataPointWithTrendAndInfluencedRate and assigns it to the SessionsCount field.
func (o *ApplicationCampaignAnalytics) SetSessionsCount(v AnalyticsDataPointWithTrendAndInfluencedRate) {
	o.SessionsCount = &v
}

// GetAvgItemsPerSession returns the AvgItemsPerSession field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetAvgItemsPerSession() AnalyticsDataPointWithTrendAndUplift {
	if o == nil || o.AvgItemsPerSession == nil {
		var ret AnalyticsDataPointWithTrendAndUplift
		return ret
	}
	return *o.AvgItemsPerSession
}

// GetAvgItemsPerSessionOk returns a tuple with the AvgItemsPerSession field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetAvgItemsPerSessionOk() (AnalyticsDataPointWithTrendAndUplift, bool) {
	if o == nil || o.AvgItemsPerSession == nil {
		var ret AnalyticsDataPointWithTrendAndUplift
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

// SetAvgItemsPerSession gets a reference to the given AnalyticsDataPointWithTrendAndUplift and assigns it to the AvgItemsPerSession field.
func (o *ApplicationCampaignAnalytics) SetAvgItemsPerSession(v AnalyticsDataPointWithTrendAndUplift) {
	o.AvgItemsPerSession = &v
}

// GetAvgSessionValue returns the AvgSessionValue field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetAvgSessionValue() AnalyticsDataPointWithTrendAndUplift {
	if o == nil || o.AvgSessionValue == nil {
		var ret AnalyticsDataPointWithTrendAndUplift
		return ret
	}
	return *o.AvgSessionValue
}

// GetAvgSessionValueOk returns a tuple with the AvgSessionValue field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetAvgSessionValueOk() (AnalyticsDataPointWithTrendAndUplift, bool) {
	if o == nil || o.AvgSessionValue == nil {
		var ret AnalyticsDataPointWithTrendAndUplift
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

// SetAvgSessionValue gets a reference to the given AnalyticsDataPointWithTrendAndUplift and assigns it to the AvgSessionValue field.
func (o *ApplicationCampaignAnalytics) SetAvgSessionValue(v AnalyticsDataPointWithTrendAndUplift) {
	o.AvgSessionValue = &v
}

// GetTotalDiscounts returns the TotalDiscounts field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetTotalDiscounts() AnalyticsDataPointWithTrend {
	if o == nil || o.TotalDiscounts == nil {
		var ret AnalyticsDataPointWithTrend
		return ret
	}
	return *o.TotalDiscounts
}

// GetTotalDiscountsOk returns a tuple with the TotalDiscounts field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetTotalDiscountsOk() (AnalyticsDataPointWithTrend, bool) {
	if o == nil || o.TotalDiscounts == nil {
		var ret AnalyticsDataPointWithTrend
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

// SetTotalDiscounts gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the TotalDiscounts field.
func (o *ApplicationCampaignAnalytics) SetTotalDiscounts(v AnalyticsDataPointWithTrend) {
	o.TotalDiscounts = &v
}

// GetCouponsCount returns the CouponsCount field value if set, zero value otherwise.
func (o *ApplicationCampaignAnalytics) GetCouponsCount() AnalyticsDataPointWithTrend {
	if o == nil || o.CouponsCount == nil {
		var ret AnalyticsDataPointWithTrend
		return ret
	}
	return *o.CouponsCount
}

// GetCouponsCountOk returns a tuple with the CouponsCount field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCampaignAnalytics) GetCouponsCountOk() (AnalyticsDataPointWithTrend, bool) {
	if o == nil || o.CouponsCount == nil {
		var ret AnalyticsDataPointWithTrend
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

// SetCouponsCount gets a reference to the given AnalyticsDataPointWithTrend and assigns it to the CouponsCount field.
func (o *ApplicationCampaignAnalytics) SetCouponsCount(v AnalyticsDataPointWithTrend) {
	o.CouponsCount = &v
}

type NullableApplicationCampaignAnalytics struct {
	Value ApplicationCampaignAnalytics
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
