# UpdateApplicationApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOffset** | Pointer to **int32** | A time offset in nanoseconds associated with the API key. When making a request using the API key, rule evaluation is based on a date that is calculated by adding the offset to the current date.  | 

## Methods

### GetTimeOffset

`func (o *UpdateApplicationApiKey) GetTimeOffset() int32`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *UpdateApplicationApiKey) GetTimeOffsetOk() (int32, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTimeOffset

`func (o *UpdateApplicationApiKey) HasTimeOffset() bool`

HasTimeOffset returns a boolean if a field has been set.

### SetTimeOffset

`func (o *UpdateApplicationApiKey) SetTimeOffset(v int32)`

SetTimeOffset gets a reference to the given int32 and assigns it to the TimeOffset field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


