# MessageLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the response was received. | [optional] 
**Response** | Pointer to **string** | Raw response data. | [optional] 
**Status** | Pointer to **int64** | HTTP status code of the response. | [optional] 

## Methods

### NewMessageLogResponse

`func NewMessageLogResponse() *MessageLogResponse`

NewMessageLogResponse instantiates a new MessageLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageLogResponseWithDefaults

`func NewMessageLogResponseWithDefaults() *MessageLogResponse`

NewMessageLogResponseWithDefaults instantiates a new MessageLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MessageLogResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageLogResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageLogResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MessageLogResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetResponse

`func (o *MessageLogResponse) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *MessageLogResponse) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *MessageLogResponse) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *MessageLogResponse) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *MessageLogResponse) GetStatus() int64`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessageLogResponse) GetStatusOk() (*int64, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessageLogResponse) SetStatus(v int64)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessageLogResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


