# MessageLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the response was received. | [optional] 
**Response** | Pointer to **string** | Raw response data. | [optional] 
**Status** | Pointer to **int32** | HTTP status code of the response. | [optional] 

## Methods

### GetCreatedAt

`func (o *MessageLogResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageLogResponse) GetCreatedAtOk() (time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasCreatedAt

`func (o *MessageLogResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAt

`func (o *MessageLogResponse) SetCreatedAt(v time.Time)`

SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.

### GetResponse

`func (o *MessageLogResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MessageLogResponse) GetResponseOk() (string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponse

`func (o *MessageLogResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### SetResponse

`func (o *MessageLogResponse) SetResponse(v string)`

SetResponse gets a reference to the given string and assigns it to the Response field.

### GetStatus

`func (o *MessageLogResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageLogResponse) GetStatusOk() (int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasStatus

`func (o *MessageLogResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatus

`func (o *MessageLogResponse) SetStatus(v int32)`

SetStatus gets a reference to the given int32 and assigns it to the Status field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


