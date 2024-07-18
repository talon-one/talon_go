# TemplateArgDefAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A campaigner friendly name for the argument, this will be shown in the rule editor. | 
**Ui** | Pointer to [**map[string]map[string]interface{}**](map[string]interface{}.md) | Arbitrary metadata that may be used to render an input for this argument. | 
**PicklistID** | Pointer to **int32** | ID of the picklist linked to a template. | [optional] 
**RestrictedByPicklist** | Pointer to **bool** | Whether or not this attribute&#39;s value is restricted by picklist (&#x60;picklist&#x60; property) | [optional] 

## Methods

### GetTitle

`func (o *TemplateArgDefAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateArgDefAllOf) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *TemplateArgDefAllOf) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *TemplateArgDefAllOf) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetUi

`func (o *TemplateArgDefAllOf) GetUi() map[string]map[string]interface{}`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *TemplateArgDefAllOf) GetUiOk() (map[string]map[string]interface{}, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUi

`func (o *TemplateArgDefAllOf) HasUi() bool`

HasUi returns a boolean if a field has been set.

### SetUi

`func (o *TemplateArgDefAllOf) SetUi(v map[string]map[string]interface{})`

SetUi gets a reference to the given map[string]map[string]interface{} and assigns it to the Ui field.

### GetPicklistID

`func (o *TemplateArgDefAllOf) GetPicklistID() int32`

GetPicklistID returns the PicklistID field if non-nil, zero value otherwise.

### GetPicklistIDOk

`func (o *TemplateArgDefAllOf) GetPicklistIDOk() (int32, bool)`

GetPicklistIDOk returns a tuple with the PicklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPicklistID

`func (o *TemplateArgDefAllOf) HasPicklistID() bool`

HasPicklistID returns a boolean if a field has been set.

### SetPicklistID

`func (o *TemplateArgDefAllOf) SetPicklistID(v int32)`

SetPicklistID gets a reference to the given int32 and assigns it to the PicklistID field.

### GetRestrictedByPicklist

`func (o *TemplateArgDefAllOf) GetRestrictedByPicklist() bool`

GetRestrictedByPicklist returns the RestrictedByPicklist field if non-nil, zero value otherwise.

### GetRestrictedByPicklistOk

`func (o *TemplateArgDefAllOf) GetRestrictedByPicklistOk() (bool, bool)`

GetRestrictedByPicklistOk returns a tuple with the RestrictedByPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRestrictedByPicklist

`func (o *TemplateArgDefAllOf) HasRestrictedByPicklist() bool`

HasRestrictedByPicklist returns a boolean if a field has been set.

### SetRestrictedByPicklist

`func (o *TemplateArgDefAllOf) SetRestrictedByPicklist(v bool)`

SetRestrictedByPicklist gets a reference to the given bool and assigns it to the RestrictedByPicklist field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


