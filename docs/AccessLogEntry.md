# AccessLogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | UUID reference of request. | 
**Status** | Pointer to **int32** | HTTP status code of response. | 
**Method** | Pointer to **string** | HTTP method of request. | 
**RequestUri** | Pointer to **string** | target URI of request | 
**Time** | Pointer to [**time.Time**](time.Time.md) | timestamp of request | 
**RequestPayload** | Pointer to **string** | payload of request | 
**ResponsePayload** | Pointer to **string** | payload of response | 

## Methods

### GetUuid

`func (o *AccessLogEntry) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AccessLogEntry) GetUuidOk() (string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasUuid

`func (o *AccessLogEntry) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### SetUuid

`func (o *AccessLogEntry) SetUuid(v string)`

SetUuid gets a reference to the given string and assigns it to the Uuid field.

### GetStatus

`func (o *AccessLogEntry) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessLogEntry) GetStatusOk() (int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *AccessLogEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *AccessLogEntry) SetStatus(v int32)`

SetStatus gets a reference to the given int32 and assigns it to the Status field.

### GetMethod

`func (o *AccessLogEntry) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccessLogEntry) GetMethodOk() (string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasMethod

`func (o *AccessLogEntry) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### SetMethod

`func (o *AccessLogEntry) SetMethod(v string)`

SetMethod gets a reference to the given string and assigns it to the Method field.

### GetRequestUri

`func (o *AccessLogEntry) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *AccessLogEntry) GetRequestUriOk() (string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestUri

`func (o *AccessLogEntry) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### SetRequestUri

`func (o *AccessLogEntry) SetRequestUri(v string)`

SetRequestUri gets a reference to the given string and assigns it to the RequestUri field.

### GetTime

`func (o *AccessLogEntry) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AccessLogEntry) GetTimeOk() (time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTime

`func (o *AccessLogEntry) HasTime() bool`

HasTime returns a boolean if a field has been set.

### SetTime

`func (o *AccessLogEntry) SetTime(v time.Time)`

SetTime gets a reference to the given time.Time and assigns it to the Time field.

### GetRequestPayload

`func (o *AccessLogEntry) GetRequestPayload() string`

GetRequestPayload returns the RequestPayload field if non-nil, zero value otherwise.

### GetRequestPayloadOk

`func (o *AccessLogEntry) GetRequestPayloadOk() (string, bool)`

GetRequestPayloadOk returns a tuple with the RequestPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasRequestPayload

`func (o *AccessLogEntry) HasRequestPayload() bool`

HasRequestPayload returns a boolean if a field has been set.

### SetRequestPayload

`func (o *AccessLogEntry) SetRequestPayload(v string)`

SetRequestPayload gets a reference to the given string and assigns it to the RequestPayload field.

### GetResponsePayload

`func (o *AccessLogEntry) GetResponsePayload() string`

GetResponsePayload returns the ResponsePayload field if non-nil, zero value otherwise.

### GetResponsePayloadOk

`func (o *AccessLogEntry) GetResponsePayloadOk() (string, bool)`

GetResponsePayloadOk returns a tuple with the ResponsePayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponsePayload

`func (o *AccessLogEntry) HasResponsePayload() bool`

HasResponsePayload returns a boolean if a field has been set.

### SetResponsePayload

`func (o *AccessLogEntry) SetResponsePayload(v string)`

SetResponsePayload gets a reference to the given string and assigns it to the ResponsePayload field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


