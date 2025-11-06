# ErrorResponseWithStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]APIError**](APIError.md) | An array of individual problems encountered during the request. | [optional] 
**StatusCode** | Pointer to **int64** | The error code | [optional] 

## Methods

### NewErrorResponseWithStatus

`func NewErrorResponseWithStatus() *ErrorResponseWithStatus`

NewErrorResponseWithStatus instantiates a new ErrorResponseWithStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithStatusWithDefaults

`func NewErrorResponseWithStatusWithDefaults() *ErrorResponseWithStatus`

NewErrorResponseWithStatusWithDefaults instantiates a new ErrorResponseWithStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorResponseWithStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseWithStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseWithStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorResponseWithStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorResponseWithStatus) GetErrors() []APIError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorResponseWithStatus) GetErrorsOk() (*[]APIError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorResponseWithStatus) SetErrors(v []APIError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorResponseWithStatus) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatusCode

`func (o *ErrorResponseWithStatus) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ErrorResponseWithStatus) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ErrorResponseWithStatus) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ErrorResponseWithStatus) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


