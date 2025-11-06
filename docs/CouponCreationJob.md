# CouponCreationJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CampaignId** | Pointer to **int64** | The ID of the campaign that owns this entity. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**UsageLimit** | Pointer to **int64** | The number of times the coupon code can be redeemed. &#x60;0&#x60; means unlimited redemptions but any campaign usage limits will still apply.  | 
**DiscountLimit** | Pointer to **float32** | The total discount value that the code can give. Typically used to represent a gift card value.  | [optional] 
**ReservationLimit** | Pointer to **int64** | The number of reservations that can be made with this coupon code.  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) | Timestamp at which point the coupon becomes valid. | [optional] 
**ExpiryDate** | Pointer to [**time.Time**](time.Time.md) | Expiration date of the coupon. Coupon never expires if this is omitted. | [optional] 
**NumberOfCoupons** | Pointer to **int64** | The number of new coupon codes to generate for the campaign. | 
**CouponSettings** | Pointer to [**CodeGeneratorSettings**](CodeGeneratorSettings.md) |  | [optional] 
**Attributes** | Pointer to [**map[string]interface{}**](.md) | Arbitrary properties associated with coupons. | 
**BatchId** | Pointer to **string** | The batch ID coupons created by this job will bear. | 
**Status** | Pointer to **string** | The current status of this request. Possible values: - &#x60;pending verification&#x60; - &#x60;pending&#x60; - &#x60;completed&#x60; - &#x60;failed&#x60; - &#x60;coupon pattern full&#x60;  | 
**CreatedAmount** | Pointer to **int64** | The number of coupon codes that were already created for this request. | 
**FailCount** | Pointer to **int64** | The number of times this job failed. | 
**Errors** | Pointer to **[]string** | An array of individual problems encountered during the request. | 
**CreatedBy** | Pointer to **int64** | ID of the user who created this effect. | 
**Communicated** | Pointer to **bool** | Whether or not the user that created this job was notified of its final state. | 
**ChunkExecutionCount** | Pointer to **int64** | The number of times an attempt to create a chunk of coupons was made during the processing of the job. | 
**ChunkSize** | Pointer to **int64** | The number of coupons that will be created in a single transactions. Coupons will be created in chunks until arriving at the requested amount. | [optional] 

## Methods

### NewCouponCreationJob

`func NewCouponCreationJob(id int64, created time.Time, campaignId int64, applicationId int64, accountId int64, usageLimit int64, numberOfCoupons int64, attributes map[string]interface{}, batchId string, status string, createdAmount int64, failCount int64, errors []string, createdBy int64, communicated bool, chunkExecutionCount int64, ) *CouponCreationJob`

NewCouponCreationJob instantiates a new CouponCreationJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponCreationJobWithDefaults

`func NewCouponCreationJobWithDefaults() *CouponCreationJob`

NewCouponCreationJobWithDefaults instantiates a new CouponCreationJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CouponCreationJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponCreationJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponCreationJob) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *CouponCreationJob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CouponCreationJob) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CouponCreationJob) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCampaignId

`func (o *CouponCreationJob) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CouponCreationJob) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CouponCreationJob) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetApplicationId

`func (o *CouponCreationJob) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CouponCreationJob) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CouponCreationJob) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetAccountId

`func (o *CouponCreationJob) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CouponCreationJob) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CouponCreationJob) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetUsageLimit

`func (o *CouponCreationJob) GetUsageLimit() int64`

GetUsageLimit returns the UsageLimit field if non-nil, zero value otherwise.

### GetUsageLimitOk

`func (o *CouponCreationJob) GetUsageLimitOk() (*int64, bool)`

GetUsageLimitOk returns a tuple with the UsageLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageLimit

`func (o *CouponCreationJob) SetUsageLimit(v int64)`

SetUsageLimit sets UsageLimit field to given value.


### GetDiscountLimit

`func (o *CouponCreationJob) GetDiscountLimit() float32`

GetDiscountLimit returns the DiscountLimit field if non-nil, zero value otherwise.

### GetDiscountLimitOk

`func (o *CouponCreationJob) GetDiscountLimitOk() (*float32, bool)`

GetDiscountLimitOk returns a tuple with the DiscountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountLimit

`func (o *CouponCreationJob) SetDiscountLimit(v float32)`

SetDiscountLimit sets DiscountLimit field to given value.

### HasDiscountLimit

`func (o *CouponCreationJob) HasDiscountLimit() bool`

HasDiscountLimit returns a boolean if a field has been set.

### GetReservationLimit

`func (o *CouponCreationJob) GetReservationLimit() int64`

GetReservationLimit returns the ReservationLimit field if non-nil, zero value otherwise.

### GetReservationLimitOk

`func (o *CouponCreationJob) GetReservationLimitOk() (*int64, bool)`

GetReservationLimitOk returns a tuple with the ReservationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservationLimit

`func (o *CouponCreationJob) SetReservationLimit(v int64)`

SetReservationLimit sets ReservationLimit field to given value.

### HasReservationLimit

`func (o *CouponCreationJob) HasReservationLimit() bool`

HasReservationLimit returns a boolean if a field has been set.

### GetStartDate

`func (o *CouponCreationJob) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CouponCreationJob) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CouponCreationJob) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CouponCreationJob) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpiryDate

`func (o *CouponCreationJob) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CouponCreationJob) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CouponCreationJob) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CouponCreationJob) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetNumberOfCoupons

`func (o *CouponCreationJob) GetNumberOfCoupons() int64`

GetNumberOfCoupons returns the NumberOfCoupons field if non-nil, zero value otherwise.

### GetNumberOfCouponsOk

`func (o *CouponCreationJob) GetNumberOfCouponsOk() (*int64, bool)`

GetNumberOfCouponsOk returns a tuple with the NumberOfCoupons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCoupons

`func (o *CouponCreationJob) SetNumberOfCoupons(v int64)`

SetNumberOfCoupons sets NumberOfCoupons field to given value.


### GetCouponSettings

`func (o *CouponCreationJob) GetCouponSettings() CodeGeneratorSettings`

GetCouponSettings returns the CouponSettings field if non-nil, zero value otherwise.

### GetCouponSettingsOk

`func (o *CouponCreationJob) GetCouponSettingsOk() (*CodeGeneratorSettings, bool)`

GetCouponSettingsOk returns a tuple with the CouponSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCouponSettings

`func (o *CouponCreationJob) SetCouponSettings(v CodeGeneratorSettings)`

SetCouponSettings sets CouponSettings field to given value.

### HasCouponSettings

`func (o *CouponCreationJob) HasCouponSettings() bool`

HasCouponSettings returns a boolean if a field has been set.

### GetAttributes

`func (o *CouponCreationJob) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CouponCreationJob) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CouponCreationJob) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.


### GetBatchId

`func (o *CouponCreationJob) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *CouponCreationJob) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *CouponCreationJob) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### GetStatus

`func (o *CouponCreationJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponCreationJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponCreationJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAmount

`func (o *CouponCreationJob) GetCreatedAmount() int64`

GetCreatedAmount returns the CreatedAmount field if non-nil, zero value otherwise.

### GetCreatedAmountOk

`func (o *CouponCreationJob) GetCreatedAmountOk() (*int64, bool)`

GetCreatedAmountOk returns a tuple with the CreatedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAmount

`func (o *CouponCreationJob) SetCreatedAmount(v int64)`

SetCreatedAmount sets CreatedAmount field to given value.


### GetFailCount

`func (o *CouponCreationJob) GetFailCount() int64`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *CouponCreationJob) GetFailCountOk() (*int64, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *CouponCreationJob) SetFailCount(v int64)`

SetFailCount sets FailCount field to given value.


### GetErrors

`func (o *CouponCreationJob) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CouponCreationJob) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CouponCreationJob) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetCreatedBy

`func (o *CouponCreationJob) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CouponCreationJob) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CouponCreationJob) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetCommunicated

`func (o *CouponCreationJob) GetCommunicated() bool`

GetCommunicated returns the Communicated field if non-nil, zero value otherwise.

### GetCommunicatedOk

`func (o *CouponCreationJob) GetCommunicatedOk() (*bool, bool)`

GetCommunicatedOk returns a tuple with the Communicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicated

`func (o *CouponCreationJob) SetCommunicated(v bool)`

SetCommunicated sets Communicated field to given value.


### GetChunkExecutionCount

`func (o *CouponCreationJob) GetChunkExecutionCount() int64`

GetChunkExecutionCount returns the ChunkExecutionCount field if non-nil, zero value otherwise.

### GetChunkExecutionCountOk

`func (o *CouponCreationJob) GetChunkExecutionCountOk() (*int64, bool)`

GetChunkExecutionCountOk returns a tuple with the ChunkExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkExecutionCount

`func (o *CouponCreationJob) SetChunkExecutionCount(v int64)`

SetChunkExecutionCount sets ChunkExecutionCount field to given value.


### GetChunkSize

`func (o *CouponCreationJob) GetChunkSize() int64`

GetChunkSize returns the ChunkSize field if non-nil, zero value otherwise.

### GetChunkSizeOk

`func (o *CouponCreationJob) GetChunkSizeOk() (*int64, bool)`

GetChunkSizeOk returns a tuple with the ChunkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChunkSize

`func (o *CouponCreationJob) SetChunkSize(v int64)`

SetChunkSize sets ChunkSize field to given value.

### HasChunkSize

`func (o *CouponCreationJob) HasChunkSize() bool`

HasChunkSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


