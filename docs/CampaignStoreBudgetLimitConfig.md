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

### GetAction

`func (o *CampaignStoreBudgetLimitConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CampaignStoreBudgetLimitConfig) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *CampaignStoreBudgetLimitConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *CampaignStoreBudgetLimitConfig) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetLimit

`func (o *CampaignStoreBudgetLimitConfig) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CampaignStoreBudgetLimitConfig) GetLimitOk() (float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *CampaignStoreBudgetLimitConfig) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *CampaignStoreBudgetLimitConfig) SetLimit(v float32)`

SetLimit gets a reference to the given float32 and assigns it to the Limit field.

### GetPeriod

`func (o *CampaignStoreBudgetLimitConfig) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CampaignStoreBudgetLimitConfig) GetPeriodOk() (string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *CampaignStoreBudgetLimitConfig) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *CampaignStoreBudgetLimitConfig) SetPeriod(v string)`

SetPeriod gets a reference to the given string and assigns it to the Period field.

### GetEntities

`func (o *CampaignStoreBudgetLimitConfig) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *CampaignStoreBudgetLimitConfig) GetEntitiesOk() ([]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntities

`func (o *CampaignStoreBudgetLimitConfig) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### SetEntities

`func (o *CampaignStoreBudgetLimitConfig) SetEntities(v []string)`

SetEntities gets a reference to the given []string and assigns it to the Entities field.

### GetImported

`func (o *CampaignStoreBudgetLimitConfig) GetImported() bool`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *CampaignStoreBudgetLimitConfig) GetImportedOk() (bool, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasImported

`func (o *CampaignStoreBudgetLimitConfig) HasImported() bool`

HasImported returns a boolean if a field has been set.

### SetImported

`func (o *CampaignStoreBudgetLimitConfig) SetImported(v bool)`

SetImported gets a reference to the given bool and assigns it to the Imported field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


