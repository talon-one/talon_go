# ResponseContentObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResponseContent** | Pointer to **[]string** | Extends the response with the chosen data entities. Use this property to get as much data back as needed from one request instead of sending extra requests to other endpoints.  | [optional] 

## Methods

### NewResponseContentObject

`func NewResponseContentObject() *ResponseContentObject`

NewResponseContentObject instantiates a new ResponseContentObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseContentObjectWithDefaults

`func NewResponseContentObjectWithDefaults() *ResponseContentObject`

NewResponseContentObjectWithDefaults instantiates a new ResponseContentObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponseContent

`func (o *ResponseContentObject) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *ResponseContentObject) GetResponseContentOk() (*[]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseContent

`func (o *ResponseContentObject) SetResponseContent(v []string)`

SetResponseContent sets ResponseContent field to given value.

### HasResponseContent

`func (o *ResponseContentObject) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


