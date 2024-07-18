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

// CouponCreationJob 
type CouponCreationJob struct {
	// Internal ID of this entity.
	Id int32 `json:"id"`
	// The time this entity was created.
	Created time.Time `json:"created"`
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
	// The ID of the application that owns this entity.
	ApplicationId int32 `json:"applicationId"`
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// The number of times the coupon code can be redeemed. `0` means unlimited redemptions but any campaign usage limits will still apply. 
	UsageLimit int32 `json:"usageLimit"`
	// The total discount value that the code can give. Typically used to represent a gift card value. 
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// The number of reservations that can be made with this coupon code. 
	ReservationLimit *int32 `json:"reservationLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiration date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of new coupon codes to generate for the campaign.
	NumberOfCoupons int32 `json:"numberOfCoupons"`
	CouponSettings *CodeGeneratorSettings `json:"couponSettings,omitempty"`
	// Arbitrary properties associated with coupons.
	Attributes map[string]interface{} `json:"attributes"`
	// The batch ID coupons created by this job will bear.
	BatchId string `json:"batchId"`
	// The current status of this request. Possible values: - `pending verification` - `pending` - `completed` - `failed` - `coupon pattern full` 
	Status string `json:"status"`
	// The number of coupon codes that were already created for this request.
	CreatedAmount int32 `json:"createdAmount"`
	// The number of times this job failed.
	FailCount int32 `json:"failCount"`
	// An array of individual problems encountered during the request.
	Errors []string `json:"errors"`
	// ID of the user who created this effect.
	CreatedBy int32 `json:"createdBy"`
	// Whether or not the user that created this job was notified of its final state.
	Communicated bool `json:"communicated"`
	// The number of times an attempt to create a chunk of coupons was made during the processing of the job.
	ChunkExecutionCount int32 `json:"chunkExecutionCount"`
	// The number of coupons that will be created in a single transactions. Coupons will be created in chunks until arriving at the requested amount.
	ChunkSize *int32 `json:"chunkSize,omitempty"`
}

// GetId returns the Id field value
func (o *CouponCreationJob) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// SetId sets field value
func (o *CouponCreationJob) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *CouponCreationJob) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// SetCreated sets field value
func (o *CouponCreationJob) SetCreated(v time.Time) {
	o.Created = v
}

// GetCampaignId returns the CampaignId field value
func (o *CouponCreationJob) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// SetCampaignId sets field value
func (o *CouponCreationJob) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CouponCreationJob) GetApplicationId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ApplicationId
}

// SetApplicationId sets field value
func (o *CouponCreationJob) SetApplicationId(v int32) {
	o.ApplicationId = v
}

// GetAccountId returns the AccountId field value
func (o *CouponCreationJob) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// SetAccountId sets field value
func (o *CouponCreationJob) SetAccountId(v int32) {
	o.AccountId = v
}

// GetUsageLimit returns the UsageLimit field value
func (o *CouponCreationJob) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// SetUsageLimit sets field value
func (o *CouponCreationJob) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *CouponCreationJob) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreationJob) GetDiscountLimitOk() (float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret, false
	}
	return *o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *CouponCreationJob) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *CouponCreationJob) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetReservationLimit returns the ReservationLimit field value if set, zero value otherwise.
func (o *CouponCreationJob) GetReservationLimit() int32 {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret
	}
	return *o.ReservationLimit
}

// GetReservationLimitOk returns a tuple with the ReservationLimit field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreationJob) GetReservationLimitOk() (int32, bool) {
	if o == nil || o.ReservationLimit == nil {
		var ret int32
		return ret, false
	}
	return *o.ReservationLimit, true
}

// HasReservationLimit returns a boolean if a field has been set.
func (o *CouponCreationJob) HasReservationLimit() bool {
	if o != nil && o.ReservationLimit != nil {
		return true
	}

	return false
}

// SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.
func (o *CouponCreationJob) SetReservationLimit(v int32) {
	o.ReservationLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CouponCreationJob) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreationJob) GetStartDateOk() (time.Time, bool) {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CouponCreationJob) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *CouponCreationJob) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *CouponCreationJob) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreationJob) GetExpiryDateOk() (time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret, false
	}
	return *o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *CouponCreationJob) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *CouponCreationJob) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetNumberOfCoupons returns the NumberOfCoupons field value
func (o *CouponCreationJob) GetNumberOfCoupons() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfCoupons
}

// SetNumberOfCoupons sets field value
func (o *CouponCreationJob) SetNumberOfCoupons(v int32) {
	o.NumberOfCoupons = v
}

// GetCouponSettings returns the CouponSettings field value if set, zero value otherwise.
func (o *CouponCreationJob) GetCouponSettings() CodeGeneratorSettings {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret
	}
	return *o.CouponSettings
}

// GetCouponSettingsOk returns a tuple with the CouponSettings field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreationJob) GetCouponSettingsOk() (CodeGeneratorSettings, bool) {
	if o == nil || o.CouponSettings == nil {
		var ret CodeGeneratorSettings
		return ret, false
	}
	return *o.CouponSettings, true
}

// HasCouponSettings returns a boolean if a field has been set.
func (o *CouponCreationJob) HasCouponSettings() bool {
	if o != nil && o.CouponSettings != nil {
		return true
	}

	return false
}

// SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.
func (o *CouponCreationJob) SetCouponSettings(v CodeGeneratorSettings) {
	o.CouponSettings = &v
}

// GetAttributes returns the Attributes field value
func (o *CouponCreationJob) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// SetAttributes sets field value
func (o *CouponCreationJob) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetBatchId returns the BatchId field value
func (o *CouponCreationJob) GetBatchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchId
}

// SetBatchId sets field value
func (o *CouponCreationJob) SetBatchId(v string) {
	o.BatchId = v
}

// GetStatus returns the Status field value
func (o *CouponCreationJob) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// SetStatus sets field value
func (o *CouponCreationJob) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAmount returns the CreatedAmount field value
func (o *CouponCreationJob) GetCreatedAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAmount
}

// SetCreatedAmount sets field value
func (o *CouponCreationJob) SetCreatedAmount(v int32) {
	o.CreatedAmount = v
}

// GetFailCount returns the FailCount field value
func (o *CouponCreationJob) GetFailCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailCount
}

// SetFailCount sets field value
func (o *CouponCreationJob) SetFailCount(v int32) {
	o.FailCount = v
}

// GetErrors returns the Errors field value
func (o *CouponCreationJob) GetErrors() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Errors
}

// SetErrors sets field value
func (o *CouponCreationJob) SetErrors(v []string) {
	o.Errors = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *CouponCreationJob) GetCreatedBy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedBy
}

// SetCreatedBy sets field value
func (o *CouponCreationJob) SetCreatedBy(v int32) {
	o.CreatedBy = v
}

// GetCommunicated returns the Communicated field value
func (o *CouponCreationJob) GetCommunicated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Communicated
}

// SetCommunicated sets field value
func (o *CouponCreationJob) SetCommunicated(v bool) {
	o.Communicated = v
}

// GetChunkExecutionCount returns the ChunkExecutionCount field value
func (o *CouponCreationJob) GetChunkExecutionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChunkExecutionCount
}

// SetChunkExecutionCount sets field value
func (o *CouponCreationJob) SetChunkExecutionCount(v int32) {
	o.ChunkExecutionCount = v
}

// GetChunkSize returns the ChunkSize field value if set, zero value otherwise.
func (o *CouponCreationJob) GetChunkSize() int32 {
	if o == nil || o.ChunkSize == nil {
		var ret int32
		return ret
	}
	return *o.ChunkSize
}

// GetChunkSizeOk returns a tuple with the ChunkSize field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *CouponCreationJob) GetChunkSizeOk() (int32, bool) {
	if o == nil || o.ChunkSize == nil {
		var ret int32
		return ret, false
	}
	return *o.ChunkSize, true
}

// HasChunkSize returns a boolean if a field has been set.
func (o *CouponCreationJob) HasChunkSize() bool {
	if o != nil && o.ChunkSize != nil {
		return true
	}

	return false
}

// SetChunkSize gets a reference to the given int32 and assigns it to the ChunkSize field.
func (o *CouponCreationJob) SetChunkSize(v int32) {
	o.ChunkSize = &v
}

type NullableCouponCreationJob struct {
	Value CouponCreationJob
	ExplicitNull bool
}

func (v NullableCouponCreationJob) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableCouponCreationJob) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
