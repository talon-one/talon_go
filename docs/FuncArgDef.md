# FuncArgDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of value this argument expects. | 
**Description** | Pointer to **string** | A campaigner-friendly description of the argument, this will also be shown in the rule editor. | [optional] 

## Methods

### NewFuncArgDef

`func NewFuncArgDef(type_ string, ) *FuncArgDef`

NewFuncArgDef instantiates a new FuncArgDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuncArgDefWithDefaults

`func NewFuncArgDefWithDefaults() *FuncArgDef`

NewFuncArgDefWithDefaults instantiates a new FuncArgDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FuncArgDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FuncArgDef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FuncArgDef) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *FuncArgDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FuncArgDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FuncArgDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FuncArgDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


