# TemplateArgDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of value this argument expects. | 
**Description** | Pointer to **string** | A campaigner-friendly description of the argument, this will also be shown in the rule editor. | [optional] 
**Title** | Pointer to **string** | A campaigner friendly name for the argument, this will be shown in the rule editor. | 
**Ui** | Pointer to [**map[string]interface{}**](.md) | Arbitrary metadata that may be used to render an input for this argument. | 
**Key** | Pointer to **string** | The identifier for the associated value within the JSON object. | [optional] 
**PicklistID** | Pointer to **int64** | ID of the picklist linked to a template. | [optional] 
**RestrictedByPicklist** | Pointer to **bool** | Whether or not this attribute&#39;s value is restricted by picklist (&#x60;picklist&#x60; property) | [optional] 

## Methods

### NewTemplateArgDef

`func NewTemplateArgDef(type_ string, title string, ui map[string]interface{}, ) *TemplateArgDef`

NewTemplateArgDef instantiates a new TemplateArgDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateArgDefWithDefaults

`func NewTemplateArgDefWithDefaults() *TemplateArgDef`

NewTemplateArgDefWithDefaults instantiates a new TemplateArgDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateArgDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateArgDef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateArgDef) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *TemplateArgDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplateArgDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplateArgDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplateArgDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *TemplateArgDef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplateArgDef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplateArgDef) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUi

`func (o *TemplateArgDef) GetUi() map[string]interface{}`

GetUi returns the Ui field if non-nil, zero value otherwise.

### GetUiOk

`func (o *TemplateArgDef) GetUiOk() (*map[string]interface{}, bool)`

GetUiOk returns a tuple with the Ui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUi

`func (o *TemplateArgDef) SetUi(v map[string]interface{})`

SetUi sets Ui field to given value.


### GetKey

`func (o *TemplateArgDef) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TemplateArgDef) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TemplateArgDef) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TemplateArgDef) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetPicklistID

`func (o *TemplateArgDef) GetPicklistID() int64`

GetPicklistID returns the PicklistID field if non-nil, zero value otherwise.

### GetPicklistIDOk

`func (o *TemplateArgDef) GetPicklistIDOk() (*int64, bool)`

GetPicklistIDOk returns a tuple with the PicklistID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicklistID

`func (o *TemplateArgDef) SetPicklistID(v int64)`

SetPicklistID sets PicklistID field to given value.

### HasPicklistID

`func (o *TemplateArgDef) HasPicklistID() bool`

HasPicklistID returns a boolean if a field has been set.

### GetRestrictedByPicklist

`func (o *TemplateArgDef) GetRestrictedByPicklist() bool`

GetRestrictedByPicklist returns the RestrictedByPicklist field if non-nil, zero value otherwise.

### GetRestrictedByPicklistOk

`func (o *TemplateArgDef) GetRestrictedByPicklistOk() (*bool, bool)`

GetRestrictedByPicklistOk returns a tuple with the RestrictedByPicklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedByPicklist

`func (o *TemplateArgDef) SetRestrictedByPicklist(v bool)`

SetRestrictedByPicklist sets RestrictedByPicklist field to given value.

### HasRestrictedByPicklist

`func (o *TemplateArgDef) HasRestrictedByPicklist() bool`

HasRestrictedByPicklist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


