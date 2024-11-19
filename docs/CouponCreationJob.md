# CouponCreationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CampaignId** | Pointer to **int32** | The ID of the campaign that owns this entity. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**UsageLimit** | Pointer to **int32** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int32** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**NumberOfCoupons** | Pointer to **int32** | The number of new coupon codes to generate for the campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with coupons. | 
**BatchId** | Pointer to **string** | The batch ID coupons created by this job will bear. | 
**Status** | Pointer to **string** | The current status of this request. Possible values: - &#x60;pending verification&#x60; - &#x60;pending&#x60; - &#x60;completed&#x60; - &#x60;failed&#x60; - &#x60;coupon pattern full&#x60;  | 
**CreatedAmount** | Pointer to **int32** | The number of coupon codes that were already created for this request. | 
**FailCount** | Pointer to **int32** | The number of times this job failed. | 
**Errors** | Pointer to **[]string** | An array of individual problems encountered during the request. | 
**CreatedBy** | Pointer to **int32** | ID of the user who created this effect. | 
**Communicated** | Pointer to **bool** | Whether or not the user that created this job was notified of its final state. | 
**ChunkExecutionCount** | Pointer to **int32** | The number of times an attempt to create a chunk of coupons was made during the processing of the job. | 
**ChunkSize** | Pointer to **int32** | The number of coupons that will be created in a single transactions. Coupons will be created in chunks until arriving at the requested amount. | [optional] 

## Methods

### GetId

`func (o *CouponCreationJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponCreationJob) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CouponCreationJob) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CouponCreationJob) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CouponCreationJob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CouponCreationJob) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CouponCreationJob) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CouponCreationJob) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetCampaignId

`func (o *CouponCreationJob) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CouponCreationJob) GetCampaignIdOk() (int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignId

`func (o *CouponCreationJob) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignId

`func (o *CouponCreationJob) SetCampaignId(v int32)`

SetCampaignId gets a reference to the given int32 and assigns it to the CampaignId field.

### GetApplicationId

`func (o *CouponCreationJob) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CouponCreationJob) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CouponCreationJob) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CouponCreationJob) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetAccountId

`func (o *CouponCreationJob) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CouponCreationJob) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CouponCreationJob) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CouponCreationJob) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetUsageLimit

`func (o *CouponCreationJob) GetUsageLimit() int32`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *CouponCreationJob) GetUsageLimitOk() (int32, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUsageLimit

`func (o *CouponCreationJob) HasUsageLimit() bool`

HasUsageLimit returns a boolean if a field has been set.

### SetUsageLimit

`func (o *CouponCreationJob) SetUsageLimit(v int32)`

SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.

### GetDiscountLimit

`func (o *CouponCreationJob) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *CouponCreationJob) GetDiscountLimitOk() (float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDiscountLimit

`func (o *CouponCreationJob) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### SetDiscountLimit

`func (o *CouponCreationJob) SetDiscountLimit(v float32)`

SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.

### GetReservationLimit

`func (o *CouponCreationJob) GetReservationLimit() int32`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *CouponCreationJob) GetReservationLimitOk() (int32, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReservationLimit

`func (o *CouponCreationJob) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### SetReservationLimit

`func (o *CouponCreationJob) SetReservationLimit(v int32)`

SetReservationLimit gets a reference to the given int32 and assigns it to the ReservationLimit field.

### GetStartDate

`func (o *CouponCreationJob) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CouponCreationJob) GetStartDateOk() (time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStartDate

`func (o *CouponCreationJob) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### SetStartDate

`func (o *CouponCreationJob) SetStartDate(v time.Time)`

SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.

### GetExpiryDate

`func (o *CouponCreationJob) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CouponCreationJob) GetExpiryDateOk() (time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasExpiryDate

`func (o *CouponCreationJob) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### SetExpiryDate

`func (o *CouponCreationJob) SetExpiryDate(v time.Time)`

SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.

### GetNumberOfCoupons

`func (o *CouponCreationJob) GetNumberOfCoupons() int32`

GetNumberOfCoupons returns the NumberOfCoupons field if non-nil, zero value otherwise.

### GetNumberOfCouponsOk

`func (o *CouponCreationJob) GetNumberOfCouponsOk() (int32, bool)`

GetNumberOfCouponsOk returns a tuple with the NumberOfCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasNumberOfCoupons

`func (o *CouponCreationJob) HasNumberOfCoupons() bool`

HasNumberOfCoupons returns a boolean if a field has been set.

### SetNumberOfCoupons

`func (o *CouponCreationJob) SetNumberOfCoupons(v int32)`

SetNumberOfCoupons gets a reference to the given int32 and assigns it to the NumberOfCoupons field.

### GetCouponSettings

`func (o *CouponCreationJob) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *CouponCreationJob) GetCouponSettingsOk() (CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCouponSettings

`func (o *CouponCreationJob) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### SetCouponSettings

`func (o *CouponCreationJob) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings gets a reference to the given CodeGeneratorSettings and assigns it to the CouponSettings field.

### GetAttributes

`func (o *CouponCreationJob) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCreationJob) GetAttributesOk() (map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAttributes

`func (o *CouponCreationJob) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### SetAttributes

`func (o *CouponCreationJob) SetAttributes(v map[string]interface{})`

SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.

### GetBatchId

`func (o *CouponCreationJob) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CouponCreationJob) GetBatchIdOk() (string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasBatchId

`func (o *CouponCreationJob) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchId

`func (o *CouponCreationJob) SetBatchId(v string)`

SetBatchId gets a reference to the given string and assigns it to the BatchId field.

### GetStatus

`func (o *CouponCreationJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponCreationJob) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *CouponCreationJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *CouponCreationJob) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetCreatedAmount

`func (o *CouponCreationJob) GetCreatedAmount() int32`

GetCreatedAmount returns the CreatedAmount field if non-nil, zero value otherwise.

### GetCreatedAmountOk

`func (o *CouponCreationJob) GetCreatedAmountOk() (int32, bool)`

GetCreatedAmountOk returns a tuple with the CreatedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAmount

`func (o *CouponCreationJob) HasCreatedAmount() bool`

HasCreatedAmount returns a boolean if a field has been set.

### SetCreatedAmount

`func (o *CouponCreationJob) SetCreatedAmount(v int32)`

SetCreatedAmount gets a reference to the given int32 and assigns it to the CreatedAmount field.

### GetFailCount

`func (o *CouponCreationJob) GetFailCount() int32`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *CouponCreationJob) GetFailCountOk() (int32, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailCount

`func (o *CouponCreationJob) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### SetFailCount

`func (o *CouponCreationJob) SetFailCount(v int32)`

SetFailCount gets a reference to the given int32 and assigns it to the FailCount field.

### GetErrors

`func (o *CouponCreationJob) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CouponCreationJob) GetErrorsOk() ([]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *CouponCreationJob) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *CouponCreationJob) SetErrors(v []string)`

SetErrors gets a reference to the given []string and assigns it to the Errors field.

### GetCreatedBy

`func (o *CouponCreationJob) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CouponCreationJob) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *CouponCreationJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *CouponCreationJob) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetCommunicated

`func (o *CouponCreationJob) GetCommunicated() bool`

GetCommunicated returns the Communicated field if non-nil, zero value otherwise.

### GetCommunicatedOk

`func (o *CouponCreationJob) GetCommunicatedOk() (bool, bool)`

GetCommunicatedOk returns a tuple with the Communicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommunicated

`func (o *CouponCreationJob) HasCommunicated() bool`

HasCommunicated returns a boolean if a field has been set.

### SetCommunicated

`func (o *CouponCreationJob) SetCommunicated(v bool)`

SetCommunicated gets a reference to the given bool and assigns it to the Communicated field.

### GetChunkExecutionCount

`func (o *CouponCreationJob) GetChunkExecutionCount() int32`

GetChunkExecutionCount returns the ChunkExecutionCount field if non-nil, zero value otherwise.

### GetChunkExecutionCountOk

`func (o *CouponCreationJob) GetChunkExecutionCountOk() (int32, bool)`

GetChunkExecutionCountOk returns a tuple with the ChunkExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChunkExecutionCount

`func (o *CouponCreationJob) HasChunkExecutionCount() bool`

HasChunkExecutionCount returns a boolean if a field has been set.

### SetChunkExecutionCount

`func (o *CouponCreationJob) SetChunkExecutionCount(v int32)`

SetChunkExecutionCount gets a reference to the given int32 and assigns it to the ChunkExecutionCount field.

### GetChunkSize

`func (o *CouponCreationJob) GetChunkSize() int32`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *CouponCreationJob) GetChunkSizeOk() (int32, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasChunkSize

`func (o *CouponCreationJob) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.

### SetChunkSize

`func (o *CouponCreationJob) SetChunkSize(v int32)`

SetChunkSize gets a reference to the given int32 and assigns it to the ChunkSize field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


