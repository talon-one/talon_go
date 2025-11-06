# AccessLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | UUID reference of request. | 
**Status** | Pointer to **int64** | HTTP status code of response. | 
**Method** | Pointer to **string** | HTTP method of request. | 
**RequestUri** | Pointer to **string** | target URI of request | 
**Time** | Pointer to [**time.Time**](time.Time.md) | timestamp of request | 
**RequestPayload** | Pointer to **string** | payload of request | 
**ResponsePayload** | Pointer to **string** | payload of response | 

## Methods

### NewAccessLogEntry

`func NewAccessLogEntry(uuid string, status int64, method string, requestUri string, time time.Time, requestPayload string, responsePayload string, ) *AccessLogEntry`

NewAccessLogEntry instantiates a new AccessLogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessLogEntryWithDefaults

`func NewAccessLogEntryWithDefaults() *AccessLogEntry`

NewAccessLogEntryWithDefaults instantiates a new AccessLogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AccessLogEntry) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AccessLogEntry) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AccessLogEntry) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetStatus

`func (o *AccessLogEntry) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessLogEntry) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessLogEntry) SetStatus(v int64)`

SetStatus sets Status field to given value.


### GetMethod

`func (o *AccessLogEntry) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccessLogEntry) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccessLogEntry) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRequestUri

`func (o *AccessLogEntry) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *AccessLogEntry) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *AccessLogEntry) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.


### GetTime

`func (o *AccessLogEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AccessLogEntry) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AccessLogEntry) SetTime(v time.Time)`

SetTime sets Time field to given value.


### GetRequestPayload

`func (o *AccessLogEntry) GetRequestPayload() string`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *AccessLogEntry) GetRequestPayloadOk() (*string, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPayload

`func (o *AccessLogEntry) SetRequestPayload(v string)`

SetRequestPayload sets RequestPayload field to given value.


### GetResponsePayload

`func (o *AccessLogEntry) GetResponsePayload() string`

GetResponsePayload returns the ResponsePayload field if non-nil, zero value otherwise.

### GetResponsePayloadOk

`func (o *AccessLogEntry) GetResponsePayloadOk() (*string, bool)`

GetResponsePayloadOk returns a tuple with the ResponsePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsePayload

`func (o *AccessLogEntry) SetResponsePayload(v string)`

SetResponsePayload sets ResponsePayload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


