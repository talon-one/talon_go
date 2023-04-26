# UpdatePicklist

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of allowed values in the picklist. If type time is chosen, it must be an RFC3339 timestamp string. | 
**Values** | Pointer to **[]string** | The list of allowed values provided by this picklist. | 

## Methods

### GetType

`func (o *UpdatePicklist) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdatePicklist) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *UpdatePicklist) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *UpdatePicklist) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetValues

`func (o *UpdatePicklist) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *UpdatePicklist) GetValuesOk() ([]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValues

`func (o *UpdatePicklist) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValues

`func (o *UpdatePicklist) SetValues(v []string)`

SetValues gets a reference to the given []string and assigns it to the Values field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


