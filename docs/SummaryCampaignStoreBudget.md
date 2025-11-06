# SummaryCampaignStoreBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | 
**Period** | Pointer to **string** |  | [optional] 
**StoreCount** | Pointer to **int64** |  | 
**Imported** | Pointer to **bool** |  | 

## Methods

### NewSummaryCampaignStoreBudget

`func NewSummaryCampaignStoreBudget(action string, storeCount int64, imported bool, ) *SummaryCampaignStoreBudget`

NewSummaryCampaignStoreBudget instantiates a new SummaryCampaignStoreBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSummaryCampaignStoreBudgetWithDefaults

`func NewSummaryCampaignStoreBudgetWithDefaults() *SummaryCampaignStoreBudget`

NewSummaryCampaignStoreBudgetWithDefaults instantiates a new SummaryCampaignStoreBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SummaryCampaignStoreBudget) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SummaryCampaignStoreBudget) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SummaryCampaignStoreBudget) SetAction(v string)`

SetAction sets Action field to given value.


### GetPeriod

`func (o *SummaryCampaignStoreBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SummaryCampaignStoreBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SummaryCampaignStoreBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *SummaryCampaignStoreBudget) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetStoreCount

`func (o *SummaryCampaignStoreBudget) GetStoreCount() int64`

GetStoreCount returns the StoreCount field if non-nil, zero value otherwise.

### GetStoreCountOk

`func (o *SummaryCampaignStoreBudget) GetStoreCountOk() (*int64, bool)`

GetStoreCountOk returns a tuple with the StoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreCount

`func (o *SummaryCampaignStoreBudget) SetStoreCount(v int64)`

SetStoreCount sets StoreCount field to given value.


### GetImported

`func (o *SummaryCampaignStoreBudget) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *SummaryCampaignStoreBudget) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *SummaryCampaignStoreBudget) SetImported(v bool)`

SetImported sets Imported field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


