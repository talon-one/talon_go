# ExportAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | The name of the entity that was exported. | 
**Filter** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Map of keys and values that were used to filter the exported rows. | 

## Methods

### GetEntity

`func (o *ExportAllOf) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ExportAllOf) GetEntityOk() (string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasEntity

`func (o *ExportAllOf) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntity

`func (o *ExportAllOf) SetEntity(v string)`

SetEntity gets a reference to the given string and assigns it to the Entity field.

### GetFilter

`func (o *ExportAllOf) GetFilter() map[string]map[string]interface{}`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ExportAllOf) GetFilterOk() (map[string]map[string]interface{}, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasFilter

`func (o *ExportAllOf) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### SetFilter

`func (o *ExportAllOf) SetFilter(v map[string]map[string]interface{})`

SetFilter gets a reference to the given map[string]map[string]interface{} and assigns it to the Filter field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


