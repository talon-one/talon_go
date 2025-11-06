# CampaignStoreBudgetLimitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The limitable action to which this limit applies. For example: - &#x60;setDiscount&#x60; - &#x60;setDiscountEffect&#x60; - &#x60;redeemCoupon&#x60; - &#x60;createCoupon&#x60;  | 
**Limit** | Pointer to **float32** | The value to set for the limit. | 
**Period** | Pointer to **string** | The period on which the budget limit recurs. | [optional] 
**Entities** | Pointer to **[]string** | The entity that this limit applies to. | 
**Imported** | Pointer to **bool** | Indicates whether this limit configuration is managed via a CSV file. | 

## Methods

### NewCampaignStoreBudgetLimitConfig

`func NewCampaignStoreBudgetLimitConfig(action string, limit float32, entities []string, imported bool, ) *CampaignStoreBudgetLimitConfig`

NewCampaignStoreBudgetLimitConfig instantiates a new CampaignStoreBudgetLimitConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignStoreBudgetLimitConfigWithDefaults

`func NewCampaignStoreBudgetLimitConfigWithDefaults() *CampaignStoreBudgetLimitConfig`

NewCampaignStoreBudgetLimitConfigWithDefaults instantiates a new CampaignStoreBudgetLimitConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CampaignStoreBudgetLimitConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CampaignStoreBudgetLimitConfig) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CampaignStoreBudgetLimitConfig) SetAction(v string)`

SetAction sets Action field to given value.


### GetLimit

`func (o *CampaignStoreBudgetLimitConfig) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CampaignStoreBudgetLimitConfig) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CampaignStoreBudgetLimitConfig) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetPeriod

`func (o *CampaignStoreBudgetLimitConfig) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CampaignStoreBudgetLimitConfig) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CampaignStoreBudgetLimitConfig) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CampaignStoreBudgetLimitConfig) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetEntities

`func (o *CampaignStoreBudgetLimitConfig) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *CampaignStoreBudgetLimitConfig) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *CampaignStoreBudgetLimitConfig) SetEntities(v []string)`

SetEntities sets Entities field to given value.


### GetImported

`func (o *CampaignStoreBudgetLimitConfig) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *CampaignStoreBudgetLimitConfig) GetImportedOk() (*bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *CampaignStoreBudgetLimitConfig) SetImported(v bool)`

SetImported sets Imported field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


