# MessageTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpResponse** | Pointer to **string** | The returned http response. | 
**HttpStatus** | Pointer to **int64** | The returned http status code. | 

## Methods

### NewMessageTest

`func NewMessageTest(httpResponse string, httpStatus int64, ) *MessageTest`

NewMessageTest instantiates a new MessageTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageTestWithDefaults

`func NewMessageTestWithDefaults() *MessageTest`

NewMessageTestWithDefaults instantiates a new MessageTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpResponse

`func (o *MessageTest) GetHttpResponse() string`

GetHttpResponse returns the HttpResponse field if non-nil, zero value otherwise.

### GetHttpResponseOk

`func (o *MessageTest) GetHttpResponseOk() (*string, bool)`

GetHttpResponseOk returns a tuple with the HttpResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponse

`func (o *MessageTest) SetHttpResponse(v string)`

SetHttpResponse sets HttpResponse field to given value.


### GetHttpStatus

`func (o *MessageTest) GetHttpStatus() int64`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *MessageTest) GetHttpStatusOk() (*int64, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *MessageTest) SetHttpStatus(v int64)`

SetHttpStatus sets HttpStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


