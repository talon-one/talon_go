# CouponDeletionJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**ApplicationId** | Pointer to **int32** | The ID of the application that owns this entity. | 
**AccountId** | Pointer to **int32** | The ID of the account that owns this entity. | 
**Filters** | Pointer to [**CouponDeletionFilters**](CouponDeletionFilters.md) |  | 
**Status** | Pointer to **string** | The current status of this request. Possible values: - &#x60;not_ready&#x60; - &#x60;pending&#x60; - &#x60;completed&#x60; - &#x60;failed&#x60;  | 
**DeletedAmount** | Pointer to **int32** | The number of coupon codes that were already deleted for this request. | [optional] 
**FailCount** | Pointer to **int32** | The number of times this job failed. | 
**Errors** | Pointer to **[]string** | An array of individual problems encountered during the request. | 
**CreatedBy** | Pointer to **int32** | ID of the user who created this effect. | 
**Communicated** | Pointer to **bool** | Indicates whether the user that created this job was notified of its final state. | 
**CampaignIDs** | Pointer to **[]int32** |  | [optional] 

## Methods

### GetId

`func (o *CouponDeletionJob) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CouponDeletionJob) GetIdOk() (int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasId

`func (o *CouponDeletionJob) HasId() bool`

HasId returns a boolean if a field has been set.

### SetId

`func (o *CouponDeletionJob) SetId(v int32)`

SetId gets a reference to the given int32 and assigns it to the Id field.

### GetCreated

`func (o *CouponDeletionJob) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CouponDeletionJob) GetCreatedOk() (time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreated

`func (o *CouponDeletionJob) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### SetCreated

`func (o *CouponDeletionJob) SetCreated(v time.Time)`

SetCreated gets a reference to the given time.Time and assigns it to the Created field.

### GetApplicationId

`func (o *CouponDeletionJob) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CouponDeletionJob) GetApplicationIdOk() (int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasApplicationId

`func (o *CouponDeletionJob) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationId

`func (o *CouponDeletionJob) SetApplicationId(v int32)`

SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.

### GetAccountId

`func (o *CouponDeletionJob) GetAccountId() int32`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CouponDeletionJob) GetAccountIdOk() (int32, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAccountId

`func (o *CouponDeletionJob) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountId

`func (o *CouponDeletionJob) SetAccountId(v int32)`

SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.

### GetFilters

`func (o *CouponDeletionJob) GetFilters() CouponDeletionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *CouponDeletionJob) GetFiltersOk() (CouponDeletionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilters

`func (o *CouponDeletionJob) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFilters

`func (o *CouponDeletionJob) SetFilters(v CouponDeletionFilters)`

SetFilters gets a reference to the given CouponDeletionFilters and assigns it to the Filters field.

### GetStatus

`func (o *CouponDeletionJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CouponDeletionJob) GetStatusOk() (string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *CouponDeletionJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *CouponDeletionJob) SetStatus(v string)`

SetStatus gets a reference to the given string and assigns it to the Status field.

### GetDeletedAmount

`func (o *CouponDeletionJob) GetDeletedAmount() int32`

GetDeletedAmount returns the DeletedAmount field if non-nil, zero value otherwise.

### GetDeletedAmountOk

`func (o *CouponDeletionJob) GetDeletedAmountOk() (int32, bool)`

GetDeletedAmountOk returns a tuple with the DeletedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDeletedAmount

`func (o *CouponDeletionJob) HasDeletedAmount() bool`

HasDeletedAmount returns a boolean if a field has been set.

### SetDeletedAmount

`func (o *CouponDeletionJob) SetDeletedAmount(v int32)`

SetDeletedAmount gets a reference to the given int32 and assigns it to the DeletedAmount field.

### GetFailCount

`func (o *CouponDeletionJob) GetFailCount() int32`

GetFailCount returns the FailCount field if non-nil, zero value otherwise.

### GetFailCountOk

`func (o *CouponDeletionJob) GetFailCountOk() (int32, bool)`

GetFailCountOk returns a tuple with the FailCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFailCount

`func (o *CouponDeletionJob) HasFailCount() bool`

HasFailCount returns a boolean if a field has been set.

### SetFailCount

`func (o *CouponDeletionJob) SetFailCount(v int32)`

SetFailCount gets a reference to the given int32 and assigns it to the FailCount field.

### GetErrors

`func (o *CouponDeletionJob) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CouponDeletionJob) GetErrorsOk() ([]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *CouponDeletionJob) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *CouponDeletionJob) SetErrors(v []string)`

SetErrors gets a reference to the given []string and assigns it to the Errors field.

### GetCreatedBy

`func (o *CouponDeletionJob) GetCreatedBy() int32`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *CouponDeletionJob) GetCreatedByOk() (int32, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedBy

`func (o *CouponDeletionJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedBy

`func (o *CouponDeletionJob) SetCreatedBy(v int32)`

SetCreatedBy gets a reference to the given int32 and assigns it to the CreatedBy field.

### GetCommunicated

`func (o *CouponDeletionJob) GetCommunicated() bool`

GetCommunicated returns the Communicated field if non-nil, zero value otherwise.

### GetCommunicatedOk

`func (o *CouponDeletionJob) GetCommunicatedOk() (bool, bool)`

GetCommunicatedOk returns a tuple with the Communicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCommunicated

`func (o *CouponDeletionJob) HasCommunicated() bool`

HasCommunicated returns a boolean if a field has been set.

### SetCommunicated

`func (o *CouponDeletionJob) SetCommunicated(v bool)`

SetCommunicated gets a reference to the given bool and assigns it to the Communicated field.

### GetCampaignIDs

`func (o *CouponDeletionJob) GetCampaignIDs() []int32`

GetCampaignIDs returns the CampaignIDs field if non-nil, zero value otherwise.

### GetCampaignIDsOk

`func (o *CouponDeletionJob) GetCampaignIDsOk() ([]int32, bool)`

GetCampaignIDsOk returns a tuple with the CampaignIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCampaignIDs

`func (o *CouponDeletionJob) HasCampaignIDs() bool`

HasCampaignIDs returns a boolean if a field has been set.

### SetCampaignIDs

`func (o *CouponDeletionJob) SetCampaignIDs(v []int32)`

SetCampaignIDs gets a reference to the given []int32 and assigns it to the CampaignIDs field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


