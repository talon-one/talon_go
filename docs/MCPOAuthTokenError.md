# MCPOAuthTokenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | RFC 6749 §5.2 error code. | 
**ErrorDescription** | Pointer to **string** | Human-readable description of the error. | [optional] 

## Methods

### NewMCPOAuthTokenError

`func NewMCPOAuthTokenError(error_ string, ) *MCPOAuthTokenError`

NewMCPOAuthTokenError instantiates a new MCPOAuthTokenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMCPOAuthTokenErrorWithDefaults

`func NewMCPOAuthTokenErrorWithDefaults() *MCPOAuthTokenError`

NewMCPOAuthTokenErrorWithDefaults instantiates a new MCPOAuthTokenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *MCPOAuthTokenError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MCPOAuthTokenError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MCPOAuthTokenError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *MCPOAuthTokenError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *MCPOAuthTokenError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *MCPOAuthTokenError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *MCPOAuthTokenError) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


