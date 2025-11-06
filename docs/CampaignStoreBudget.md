# CampaignStoreBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The internal ID of this entity. | 
**Created** | Pointer to [**time.Time**](time.Time.md) | The time this entity was created. | 
**CampaignId** | Pointer to **int64** | The ID of the campaign that owns this entity. | 
**StoreId** | Pointer to **int64** | The ID of the store. | 
**Limits** | Pointer to [**[]CampaignStoreBudgetLimitConfig**](CampaignStoreBudgetLimitConfig.md) | The set of budget limits for stores linked to the campaign. | 

## Methods

### NewCampaignStoreBudget

`func NewCampaignStoreBudget(id int64, created time.Time, campaignId int64, storeId int64, limits []CampaignStoreBudgetLimitConfig, ) *CampaignStoreBudget`

NewCampaignStoreBudget instantiates a new CampaignStoreBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignStoreBudgetWithDefaults

`func NewCampaignStoreBudgetWithDefaults() *CampaignStoreBudget`

NewCampaignStoreBudgetWithDefaults instantiates a new CampaignStoreBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CampaignStoreBudget) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CampaignStoreBudget) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CampaignStoreBudget) SetId(v int64)`

SetId sets Id field to given value.


### GetCreated

`func (o *CampaignStoreBudget) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignStoreBudget) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignStoreBudget) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetCampaignId

`func (o *CampaignStoreBudget) GetCampaignId() int64`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CampaignStoreBudget) GetCampaignIdOk() (*int64, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CampaignStoreBudget) SetCampaignId(v int64)`

SetCampaignId sets CampaignId field to given value.


### GetStoreId

`func (o *CampaignStoreBudget) GetStoreId() int64`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *CampaignStoreBudget) GetStoreIdOk() (*int64, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *CampaignStoreBudget) SetStoreId(v int64)`

SetStoreId sets StoreId field to given value.


### GetLimits

`func (o *CampaignStoreBudget) GetLimits() []CampaignStoreBudgetLimitConfig`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CampaignStoreBudget) GetLimitsOk() (*[]CampaignStoreBudgetLimitConfig, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CampaignStoreBudget) SetLimits(v []CampaignStoreBudgetLimitConfig)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


