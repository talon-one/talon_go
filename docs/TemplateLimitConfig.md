# TemplateLimitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The limitable action to which this limit will be applied | 
**Limit** | Pointer to **float32** | The value to set for the limit | 
**Period** | Pointer to **string** | The period on which the budget limit recurs | [optional] 
**Entities** | Pointer to **[]string** | The entities that make the address of this limit | 
**Description** | Pointer to **string** | The description of this budget configuration | 

## Methods

### GetAction

`func (o *TemplateLimitConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TemplateLimitConfig) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *TemplateLimitConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *TemplateLimitConfig) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetLimit

`func (o *TemplateLimitConfig) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TemplateLimitConfig) GetLimitOk() (float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *TemplateLimitConfig) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *TemplateLimitConfig) SetLimit(v float32)`

SetLimit gets a reference to the given float32 and assigns it to the Limit field.

### GetPeriod

`func (o *TemplateLimitConfig) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TemplateLimitConfig) GetPeriodOk() (string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPeriod

`func (o *TemplateLimitConfig) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriod

`func (o *TemplateLimitConfig) SetPeriod(v string)`

SetPeriod gets a reference to the given string and assigns it to the Period field.

### GetEntities

`func (o *TemplateLimitConfig) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *TemplateLimitConfig) GetEntitiesOk() ([]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntities

`func (o *TemplateLimitConfig) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### SetEntities

`func (o *TemplateLimitConfig) SetEntities(v []string)`

SetEntities gets a reference to the given []string and assigns it to the Entities field.

### GetDescription

`func (o *TemplateLimitConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateLimitConfig) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *TemplateLimitConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *TemplateLimitConfig) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


