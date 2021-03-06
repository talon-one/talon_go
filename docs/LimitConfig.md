# LimitConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The limitable action to which this limit will be applied | 
**Limit** | Pointer to **float32** | The value to set for the limit | 
**Entities** | Pointer to **[]string** | The entities that make the address of this limit | 

## Methods

### GetAction

`func (o *LimitConfig) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LimitConfig) GetActionOk() (string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasAction

`func (o *LimitConfig) HasAction() bool`

HasAction returns a boolean if a field has been set.

### SetAction

`func (o *LimitConfig) SetAction(v string)`

SetAction gets a reference to the given string and assigns it to the Action field.

### GetLimit

`func (o *LimitConfig) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *LimitConfig) GetLimitOk() (float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLimit

`func (o *LimitConfig) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### SetLimit

`func (o *LimitConfig) SetLimit(v float32)`

SetLimit gets a reference to the given float32 and assigns it to the Limit field.

### GetEntities

`func (o *LimitConfig) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *LimitConfig) GetEntitiesOk() ([]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntities

`func (o *LimitConfig) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### SetEntities

`func (o *LimitConfig) SetEntities(v []string)`

SetEntities gets a reference to the given []string and assigns it to the Entities field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


