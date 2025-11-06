# CouponDeletionJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int64** | The ID of the Application that owns this entity. | 
**AccountId** | Pointer to **int64** | The ID of the account that owns this entity. | 
**Filters** | Pointer to [**CouponDeletionFilters**](CouponDeletionFilters.md) |  | 
**Status** | Pointer to **string** | The current status of this request. Possible values: - &#x60;not_ready&#x60; - &#x60;pending&#x60; - &#x60;completed&#x60; - &#x60;failed&#x60;  | 
**DeletedAmount** | Pointer to **int64** | The number of coupon codes that were already deleted for this request. | [optional] 
**FailCount** | Pointer to **int64** | The number of times this job failed. | 
**Errors** | Pointer to **[]string** | An array of individual problems encountered during the request. | 
**CreatedBy** | Pointer to **int64** | ID of the user who created this effect. | 
**Communicated** | Pointer to **bool** | Indicates whether the user that created this job was notified of its final state. | 
**CampaignIDs** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewCouponDeletionJob

`func NewCouponDeletionJob(id int64, created time.Time, applicationId int64, accountId int64, filters CouponDeletionFilters, status string, failCount int64, errors []string, createdBy int64, communicated bool, ) *CouponDeletionJob`

NewCouponDeletionJob instantiates a new CouponDeletionJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCouponDeletionJobWithDefaults

`func NewCouponDeletionJobWithDefaults() *CouponDeletionJob`

NewCouponDeletionJobWithDefaults instantiates a new CouponDeletionJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CouponDeletionJob) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponDeletionJob) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CouponDeletionJob) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *CouponDeletionJob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CouponDeletionJob) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CouponDeletionJob) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetApplicationId

`func (o *CouponDeletionJob) GetApplicationId() int64`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CouponDeletionJob) GetApplicationIdOk() (*int64, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CouponDeletionJob) SetApplicationId(v int64)`

SetApplicationId sets ApplicationId field to given value.


### GetAccountId

`func (o *CouponDeletionJob) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CouponDeletionJob) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CouponDeletionJob) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetFilters

`func (o *CouponDeletionJob) GetFilters() CouponDeletionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CouponDeletionJob) GetFiltersOk() (*CouponDeletionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *CouponDeletionJob) SetFilters(v CouponDeletionFilters)`

SetFilters sets Filters field to given value.


### GetStatus

`func (o *CouponDeletionJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponDeletionJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CouponDeletionJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetDeletedAmount

`func (o *CouponDeletionJob) GetDeletedAmount() int64`

GetDeletedAmount returns the DeletedAmount field if non-nil, zero value otherwise.

### GetDeletedAmountOk

`func (o *CouponDeletionJob) GetDeletedAmountOk() (*int64, bool)`

GetDeletedAmountOk returns a tuple with the DeletedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAmount

`func (o *CouponDeletionJob) SetDeletedAmount(v int64)`

SetDeletedAmount sets DeletedAmount field to given value.

### HasDeletedAmount

`func (o *CouponDeletionJob) HasDeletedAmount() bool`

HasDeletedAmount returns a boolean if a field has been set.

### GetFailCount

`func (o *CouponDeletionJob) GetFailCount() int64`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *CouponDeletionJob) GetFailCountOk() (*int64, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailCount

`func (o *CouponDeletionJob) SetFailCount(v int64)`

SetFailCount sets FailCount field to given value.


### GetErrors

`func (o *CouponDeletionJob) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CouponDeletionJob) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CouponDeletionJob) SetErrors(v []string)`

SetErrors sets Errors field to given value.


### GetCreatedBy

`func (o *CouponDeletionJob) GetCreatedBy() int64`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CouponDeletionJob) GetCreatedByOk() (*int64, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *CouponDeletionJob) SetCreatedBy(v int64)`

SetCreatedBy sets CreatedBy field to given value.


### GetCommunicated

`func (o *CouponDeletionJob) GetCommunicated() bool`

GetCommunicated returns the Communicated field if non-nil, zero value otherwise.

### GetCommunicatedOk

`func (o *CouponDeletionJob) GetCommunicatedOk() (*bool, bool)`

GetCommunicatedOk returns a tuple with the Communicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunicated

`func (o *CouponDeletionJob) SetCommunicated(v bool)`

SetCommunicated sets Communicated field to given value.


### GetCampaignIDs

`func (o *CouponDeletionJob) GetCampaignIDs() []int64`

GetCampaignIDs returns the CampaignIDs field if non-nil, zero value otherwise.

### GetCampaignIDsOk

`func (o *CouponDeletionJob) GetCampaignIDsOk() (*[]int64, bool)`

GetCampaignIDsOk returns a tuple with the CampaignIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIDs

`func (o *CouponDeletionJob) SetCampaignIDs(v []int64)`

SetCampaignIDs sets CampaignIDs field to given value.

### HasCampaignIDs

`func (o *CouponDeletionJob) HasCampaignIDs() bool`

HasCampaignIDs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


