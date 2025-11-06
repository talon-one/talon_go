# ListCampaignStoreBudgets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Store** | Pointer to [**ListCampaignStoreBudgetsStore**](ListCampaignStoreBudgetsStore.md) |  | 
**Limit** | Pointer to **int64** |  | 
**Action** | Pointer to **string** |  | 
**Period** | Pointer to **string** |  | [optional] 

## Methods

### NewListCampaignStoreBudgets

`func NewListCampaignStoreBudgets(store ListCampaignStoreBudgetsStore, limit int64, action string, ) *ListCampaignStoreBudgets`

NewListCampaignStoreBudgets instantiates a new ListCampaignStoreBudgets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCampaignStoreBudgetsWithDefaults

`func NewListCampaignStoreBudgetsWithDefaults() *ListCampaignStoreBudgets`

NewListCampaignStoreBudgetsWithDefaults instantiates a new ListCampaignStoreBudgets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStore

`func (o *ListCampaignStoreBudgets) GetStore() ListCampaignStoreBudgetsStore`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *ListCampaignStoreBudgets) GetStoreOk() (*ListCampaignStoreBudgetsStore, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *ListCampaignStoreBudgets) SetStore(v ListCampaignStoreBudgetsStore)`

SetStore sets Store field to given value.


### GetLimit

`func (o *ListCampaignStoreBudgets) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ListCampaignStoreBudgets) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ListCampaignStoreBudgets) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetAction

`func (o *ListCampaignStoreBudgets) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListCampaignStoreBudgets) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListCampaignStoreBudgets) SetAction(v string)`

SetAction sets Action field to given value.


### GetPeriod

`func (o *ListCampaignStoreBudgets) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ListCampaignStoreBudgets) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ListCampaignStoreBudgets) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ListCampaignStoreBudgets) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


