# OutboundLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int64** | HTTP status code returned by the receiver. | 
**RawBody** | Pointer to **string** | Raw HTTP response. | 

## Methods

### NewOutboundLogResponse

`func NewOutboundLogResponse(statusCode int64, rawBody string, ) *OutboundLogResponse`

NewOutboundLogResponse instantiates a new OutboundLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundLogResponseWithDefaults

`func NewOutboundLogResponseWithDefaults() *OutboundLogResponse`

NewOutboundLogResponseWithDefaults instantiates a new OutboundLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *OutboundLogResponse) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *OutboundLogResponse) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *OutboundLogResponse) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.


### GetRawBody

`func (o *OutboundLogResponse) GetRawBody() string`

GetRawBody returns the RawBody field if non-nil, zero value otherwise.

### GetRawBodyOk

`func (o *OutboundLogResponse) GetRawBodyOk() (*string, bool)`

GetRawBodyOk returns a tuple with the RawBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawBody

`func (o *OutboundLogResponse) SetRawBody(v string)`

SetRawBody sets RawBody field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


