# ErrorSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pointer** | Pointer to **string** | Pointer to the path in the payload that caused this error. | [optional] 
**Parameter** | Pointer to **string** | Query parameter that caused this error. | [optional] 
**Line** | Pointer to **string** | Line number in uploaded multipart file that caused this error. &#39;N/A&#39; if unknown. | [optional] 
**Resource** | Pointer to **string** | Pointer to the resource that caused this error. | [optional] 

## Methods

### GetPointer

`func (o *ErrorSource) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorSource) GetPointerOk() (string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPointer

`func (o *ErrorSource) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### SetPointer

`func (o *ErrorSource) SetPointer(v string)`

SetPointer gets a reference to the given string and assigns it to the Pointer field.

### GetParameter

`func (o *ErrorSource) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *ErrorSource) GetParameterOk() (string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasParameter

`func (o *ErrorSource) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### SetParameter

`func (o *ErrorSource) SetParameter(v string)`

SetParameter gets a reference to the given string and assigns it to the Parameter field.

### GetLine

`func (o *ErrorSource) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *ErrorSource) GetLineOk() (string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasLine

`func (o *ErrorSource) HasLine() bool`

HasLine returns a boolean if a field has been set.

### SetLine

`func (o *ErrorSource) SetLine(v string)`

SetLine gets a reference to the given string and assigns it to the Line field.

### GetResource

`func (o *ErrorSource) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ErrorSource) GetResourceOk() (string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResource

`func (o *ErrorSource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResource

`func (o *ErrorSource) SetResource(v string)`

SetResource gets a reference to the given string and assigns it to the Resource field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


