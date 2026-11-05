# OutboundMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | Pointer to **int64** | HTTP status code returned by the receiver. | 
**RawBody** | Pointer to **string** | Raw HTTP response. | 
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the log entry was created. | 
**ProcessingTimeMs** | Pointer to **int64** | Processing time of the outbound request in milliseconds. | 

## Methods

### NewOutboundMessageResponse

`func NewOutboundMessageResponse(statusCode int64, rawBody string, createdAt time.Time, processingTimeMs int64, ) *OutboundMessageResponse`

NewOutboundMessageResponse instantiates a new OutboundMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundMessageResponseWithDefaults

`func NewOutboundMessageResponseWithDefaults() *OutboundMessageResponse`

NewOutboundMessageResponseWithDefaults instantiates a new OutboundMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusCode

`func (o *OutboundMessageResponse) GetStatusCode() int64`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *OutboundMessageResponse) GetStatusCodeOk() (*int64, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *OutboundMessageResponse) SetStatusCode(v int64)`

SetStatusCode sets StatusCode field to given value.


### GetRawBody

`func (o *OutboundMessageResponse) GetRawBody() string`

GetRawBody returns the RawBody field if non-nil, zero value otherwise.

### GetRawBodyOk

`func (o *OutboundMessageResponse) GetRawBodyOk() (*string, bool)`

GetRawBodyOk returns a tuple with the RawBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawBody

`func (o *OutboundMessageResponse) SetRawBody(v string)`

SetRawBody sets RawBody field to given value.


### GetCreatedAt

`func (o *OutboundMessageResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OutboundMessageResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OutboundMessageResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetProcessingTimeMs

`func (o *OutboundMessageResponse) GetProcessingTimeMs() int64`

GetProcessingTimeMs returns the ProcessingTimeMs field if non-nil, zero value otherwise.

### GetProcessingTimeMsOk

`func (o *OutboundMessageResponse) GetProcessingTimeMsOk() (*int64, bool)`

GetProcessingTimeMsOk returns a tuple with the ProcessingTimeMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingTimeMs

`func (o *OutboundMessageResponse) SetProcessingTimeMs(v int64)`

SetProcessingTimeMs sets ProcessingTimeMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


