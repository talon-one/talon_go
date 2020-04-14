# FunctionDef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The function name used in Talang. | 
**Type** | Pointer to **string** | The type of this function argument. | 
**Description** | Pointer to **string** | A short description of the function. | [optional] 
**Help** | Pointer to **string** | Extended help text for the function. | [optional] 
**Args** | Pointer to [**[]FuncArgDef**](FuncArgDef.md) | An array of argument definitions. | 

## Methods

### GetName

`func (o *FunctionDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionDef) GetNameOk() (string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasName

`func (o *FunctionDef) HasName() bool`

HasName returns a boolean if a field has been set.

### SetName

`func (o *FunctionDef) SetName(v string)`

SetName gets a reference to the given string and assigns it to the Name field.

### GetType

`func (o *FunctionDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionDef) GetTypeOk() (string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasType

`func (o *FunctionDef) HasType() bool`

HasType returns a boolean if a field has been set.

### SetType

`func (o *FunctionDef) SetType(v string)`

SetType gets a reference to the given string and assigns it to the Type field.

### GetDescription

`func (o *FunctionDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionDef) GetDescriptionOk() (string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDescription

`func (o *FunctionDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescription

`func (o *FunctionDef) SetDescription(v string)`

SetDescription gets a reference to the given string and assigns it to the Description field.

### GetHelp

`func (o *FunctionDef) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *FunctionDef) GetHelpOk() (string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasHelp

`func (o *FunctionDef) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### SetHelp

`func (o *FunctionDef) SetHelp(v string)`

SetHelp gets a reference to the given string and assigns it to the Help field.

### GetArgs

`func (o *FunctionDef) GetArgs() []FuncArgDef`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *FunctionDef) GetArgsOk() ([]FuncArgDef, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasArgs

`func (o *FunctionDef) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgs

`func (o *FunctionDef) SetArgs(v []FuncArgDef)`

SetArgs gets a reference to the given []FuncArgDef and assigns it to the Args field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


