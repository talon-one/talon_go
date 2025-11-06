# NewAppWideCouponDeletionJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**CouponDeletionFilters**](CouponDeletionFilters.md) |  | 
**Campaignids** | Pointer to **[]int64** |  | 

## Methods

### NewNewAppWideCouponDeletionJob

`func NewNewAppWideCouponDeletionJob(filters CouponDeletionFilters, campaignids []int64, ) *NewAppWideCouponDeletionJob`

NewNewAppWideCouponDeletionJob instantiates a new NewAppWideCouponDeletionJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewAppWideCouponDeletionJobWithDefaults

`func NewNewAppWideCouponDeletionJobWithDefaults() *NewAppWideCouponDeletionJob`

NewNewAppWideCouponDeletionJobWithDefaults instantiates a new NewAppWideCouponDeletionJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *NewAppWideCouponDeletionJob) GetFilters() CouponDeletionFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *NewAppWideCouponDeletionJob) GetFiltersOk() (*CouponDeletionFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *NewAppWideCouponDeletionJob) SetFilters(v CouponDeletionFilters)`

SetFilters sets Filters field to given value.


### GetCampaignids

`func (o *NewAppWideCouponDeletionJob) GetCampaignids() []int64`

GetCampaignids returns the Campaignids field if non-nil, zero value otherwise.

### GetCampaignidsOk

`func (o *NewAppWideCouponDeletionJob) GetCampaignidsOk() (*[]int64, bool)`

GetCampaignidsOk returns a tuple with the Campaignids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignids

`func (o *NewAppWideCouponDeletionJob) SetCampaignids(v []int64)`

SetCampaignids sets Campaignids field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


