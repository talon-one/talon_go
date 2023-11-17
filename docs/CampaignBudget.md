# CampaignBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The limitable action to which this limit applies. For example: - &#x60;setDiscount&#x60; - &#x60;setDiscountEffect&#x60; - &#x60;redeemCoupon&#x60; - &#x60;createCoupon&#x60;  | 
**Limit** | Pointer to **float32** | The value to set for the limit. | 
**Counter** | Pointer to **float32** | The number of occurrences of the limited action in the context of the campaign. | 

## Methods

### GetAction

`func (o *CampaignBudget) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CampaignBudget) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *CampaignBudget) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *CampaignBudget) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetLimit

`func (o *CampaignBudget) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CampaignBudget) GetLimitOk() (float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *CampaignBudget) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *CampaignBudget) SetLimit(v float32)`

SetLimit gets a reference to the given float32 and assigns it to the Limit field.

### GetCounter

`func (o *CampaignBudget) GetCounter() float32`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *CampaignBudget) GetCounterOk() (float32, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCounter

`func (o *CampaignBudget) HasCounter() bool`

HasCounter returns a boolean if a field has been set.

### SetCounter

`func (o *CampaignBudget) SetCounter(v float32)`

SetCounter gets a reference to the given float32 and assigns it to the Counter field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


