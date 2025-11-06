# UpdateApplicationAPIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeOffset** | Pointer to **int64** | A time offset in nanoseconds associated with the API key. When making a request using the API key, rule evaluation is based on a date that is calculated by adding the offset to the current date.  | 

## Methods

### NewUpdateApplicationAPIKey

`func NewUpdateApplicationAPIKey(timeOffset int64, ) *UpdateApplicationAPIKey`

NewUpdateApplicationAPIKey instantiates a new UpdateApplicationAPIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApplicationAPIKeyWithDefaults

`func NewUpdateApplicationAPIKeyWithDefaults() *UpdateApplicationAPIKey`

NewUpdateApplicationAPIKeyWithDefaults instantiates a new UpdateApplicationAPIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeOffset

`func (o *UpdateApplicationAPIKey) GetTimeOffset() int64`

GetTimeOffset returns the TimeOffset field if non-nil, zero value otherwise.

### GetTimeOffsetOk

`func (o *UpdateApplicationAPIKey) GetTimeOffsetOk() (*int64, bool)`

GetTimeOffsetOk returns a tuple with the TimeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOffset

`func (o *UpdateApplicationAPIKey) SetTimeOffset(v int64)`

SetTimeOffset sets TimeOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


