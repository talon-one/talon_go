# NewCampaignStoreBudgetStoreLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | Pointer to **int64** | The ID of the store. You can get this ID with the [List stores](#tag/Stores/operation/listStores) endpoint.  | 
**Limit** | Pointer to **float32** | The value to set for the limit. | 

## Methods

### NewNewCampaignStoreBudgetStoreLimit

`func NewNewCampaignStoreBudgetStoreLimit(storeId int64, limit float32, ) *NewCampaignStoreBudgetStoreLimit`

NewNewCampaignStoreBudgetStoreLimit instantiates a new NewCampaignStoreBudgetStoreLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignStoreBudgetStoreLimitWithDefaults

`func NewNewCampaignStoreBudgetStoreLimitWithDefaults() *NewCampaignStoreBudgetStoreLimit`

NewNewCampaignStoreBudgetStoreLimitWithDefaults instantiates a new NewCampaignStoreBudgetStoreLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *NewCampaignStoreBudgetStoreLimit) GetStoreId() int64`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *NewCampaignStoreBudgetStoreLimit) GetStoreIdOk() (*int64, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *NewCampaignStoreBudgetStoreLimit) SetStoreId(v int64)`

SetStoreId sets StoreId field to given value.


### GetLimit

`func (o *NewCampaignStoreBudgetStoreLimit) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *NewCampaignStoreBudgetStoreLimit) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *NewCampaignStoreBudgetStoreLimit) SetLimit(v float32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


