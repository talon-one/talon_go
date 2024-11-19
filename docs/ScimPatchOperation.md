# ScimPatchOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | The method that should be used in the operation. | 
**Path** | Pointer to **string** | The path specifying the attribute that should be updated. | [optional] 
**Value** | Pointer to **string** | The value that should be updated. Required if &#x60;op&#x60; is &#x60;add&#x60; or &#x60;replace&#x60;. | [optional] 

## Methods

### GetOp

`func (o *ScimPatchOperation) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ScimPatchOperation) GetOpOk() (string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasOp

`func (o *ScimPatchOperation) HasOp() bool`

HasOp returns a boolean if a field has been set.

### SetOp

`func (o *ScimPatchOperation) SetOp(v string)`

SetOp gets a reference to the given string and assigns it to the Op field.

### GetPath

`func (o *ScimPatchOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ScimPatchOperation) GetPathOk() (string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPath

`func (o *ScimPatchOperation) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPath

`func (o *ScimPatchOperation) SetPath(v string)`

SetPath gets a reference to the given string and assigns it to the Path field.

### GetValue

`func (o *ScimPatchOperation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ScimPatchOperation) GetValueOk() (string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasValue

`func (o *ScimPatchOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValue

`func (o *ScimPatchOperation) SetValue(v string)`

SetValue gets a reference to the given string and assigns it to the Value field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


