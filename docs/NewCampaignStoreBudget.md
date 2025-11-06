# NewCampaignStoreBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | 
**StoreLimits** | Pointer to [**[]NewCampaignStoreBudgetStoreLimit**](NewCampaignStoreBudgetStoreLimit.md) | The set of budget limits for stores linked to the campaign. | 
**Period** | Pointer to **string** |  | [optional] 

## Methods

### NewNewCampaignStoreBudget

`func NewNewCampaignStoreBudget(action string, storeLimits []NewCampaignStoreBudgetStoreLimit, ) *NewCampaignStoreBudget`

NewNewCampaignStoreBudget instantiates a new NewCampaignStoreBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewCampaignStoreBudgetWithDefaults

`func NewNewCampaignStoreBudgetWithDefaults() *NewCampaignStoreBudget`

NewNewCampaignStoreBudgetWithDefaults instantiates a new NewCampaignStoreBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *NewCampaignStoreBudget) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *NewCampaignStoreBudget) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *NewCampaignStoreBudget) SetAction(v string)`

SetAction sets Action field to given value.


### GetStoreLimits

`func (o *NewCampaignStoreBudget) GetStoreLimits() []NewCampaignStoreBudgetStoreLimit`

GetStoreLimits returns the StoreLimits field if non-nil, zero value otherwise.

### GetStoreLimitsOk

`func (o *NewCampaignStoreBudget) GetStoreLimitsOk() (*[]NewCampaignStoreBudgetStoreLimit, bool)`

GetStoreLimitsOk returns a tuple with the StoreLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreLimits

`func (o *NewCampaignStoreBudget) SetStoreLimits(v []NewCampaignStoreBudgetStoreLimit)`

SetStoreLimits sets StoreLimits field to given value.


### GetPeriod

`func (o *NewCampaignStoreBudget) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *NewCampaignStoreBudget) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *NewCampaignStoreBudget) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *NewCampaignStoreBudget) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


