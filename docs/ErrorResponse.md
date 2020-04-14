# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | A message describing the error. | 
**Errors** | Pointer to [**[]ApiError**](APIError.md) | An array of individual problems encountered during the request. | [optional] 

## Methods

### GetMessage

`func (o *ErrorResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponse) GetMessageOk() (string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMessage

`func (o *ErrorResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessage

`func (o *ErrorResponse) SetMessage(v string)`

SetMessage gets a reference to the given string and assigns it to the Message field.

### GetErrors

`func (o *ErrorResponse) GetErrors() []ApiError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorResponse) GetErrorsOk() ([]ApiError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasErrors

`func (o *ErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrors

`func (o *ErrorResponse) SetErrors(v []ApiError)`

SetErrors gets a reference to the given []ApiError and assigns it to the Errors field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


