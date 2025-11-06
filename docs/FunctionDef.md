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

### NewFunctionDef

`func NewFunctionDef(name string, type_ string, args []FuncArgDef, ) *FunctionDef`

NewFunctionDef instantiates a new FunctionDef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFunctionDefWithDefaults

`func NewFunctionDefWithDefaults() *FunctionDef`

NewFunctionDefWithDefaults instantiates a new FunctionDef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FunctionDef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FunctionDef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FunctionDef) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *FunctionDef) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FunctionDef) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FunctionDef) SetType(v string)`

SetType sets Type field to given value.


### GetDescription

`func (o *FunctionDef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FunctionDef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FunctionDef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FunctionDef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHelp

`func (o *FunctionDef) GetHelp() string`

GetHelp returns the Help field if non-nil, zero value otherwise.

### GetHelpOk

`func (o *FunctionDef) GetHelpOk() (*string, bool)`

GetHelpOk returns a tuple with the Help field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelp

`func (o *FunctionDef) SetHelp(v string)`

SetHelp sets Help field to given value.

### HasHelp

`func (o *FunctionDef) HasHelp() bool`

HasHelp returns a boolean if a field has been set.

### GetArgs

`func (o *FunctionDef) GetArgs() []FuncArgDef`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *FunctionDef) GetArgsOk() (*[]FuncArgDef, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *FunctionDef) SetArgs(v []FuncArgDef)`

SetArgs sets Args field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


