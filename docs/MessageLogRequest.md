# MessageLogRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to [**time.Time**](time.Time.md) | Timestamp when the request was made. | 
**Request** | Pointer to **string** | Raw request data. | 

## Methods

### NewMessageLogRequest

`func NewMessageLogRequest(createdAt time.Time, request string, ) *MessageLogRequest`

NewMessageLogRequest instantiates a new MessageLogRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageLogRequestWithDefaults

`func NewMessageLogRequestWithDefaults() *MessageLogRequest`

NewMessageLogRequestWithDefaults instantiates a new MessageLogRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *MessageLogRequest) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MessageLogRequest) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MessageLogRequest) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetRequest

`func (o *MessageLogRequest) GetRequest() string`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *MessageLogRequest) GetRequestOk() (*string, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *MessageLogRequest) SetRequest(v string)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


