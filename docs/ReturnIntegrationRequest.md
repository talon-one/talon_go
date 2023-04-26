# ReturnIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Return** | Pointer to [**NewReturn**](NewReturn.md) |  | 
**ResponseContent** | Pointer to **[]string** | Extends the response with the chosen data entities. Use this property to get as much data as you need in one _Update customer session_ request instead of sending extra requests to other endpoints.  | [optional] 

## Methods

### GetReturn

`func (o *ReturnIntegrationRequest) GetReturn() NewReturn`

GetReturn returns the Return field if non-nil, zero value otherwise.

### GetReturnOk

`func (o *ReturnIntegrationRequest) GetReturnOk() (NewReturn, bool)`

GetReturnOk returns a tuple with the Return field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasReturn

`func (o *ReturnIntegrationRequest) HasReturn() bool`

HasReturn returns a boolean if a field has been set.

### SetReturn

`func (o *ReturnIntegrationRequest) SetReturn(v NewReturn)`

SetReturn gets a reference to the given NewReturn and assigns it to the Return field.

### GetResponseContent

`func (o *ReturnIntegrationRequest) GetResponseContent() []string`

GetResponseContent returns the ResponseContent field if non-nil, zero value otherwise.

### GetResponseContentOk

`func (o *ReturnIntegrationRequest) GetResponseContentOk() ([]string, bool)`

GetResponseContentOk returns a tuple with the ResponseContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasResponseContent

`func (o *ReturnIntegrationRequest) HasResponseContent() bool`

HasResponseContent returns a boolean if a field has been set.

### SetResponseContent

`func (o *ReturnIntegrationRequest) SetResponseContent(v []string)`

SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


