# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Short description of the problem. | 
**Details** | Pointer to **string** | Longer description of this specific instance of the problem. | [optional] 
**Source** | Pointer to [**ErrorSource**](ErrorSource.md) |  | 

## Methods

### GetTitle

`func (o *ApiError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiError) GetTitleOk() (string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasTitle

`func (o *ApiError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitle

`func (o *ApiError) SetTitle(v string)`

SetTitle gets a reference to the given string and assigns it to the Title field.

### GetDetails

`func (o *ApiError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ApiError) GetDetailsOk() (string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasDetails

`func (o *ApiError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### SetDetails

`func (o *ApiError) SetDetails(v string)`

SetDetails gets a reference to the given string and assigns it to the Details field.

### GetSource

`func (o *ApiError) GetSource() ErrorSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiError) GetSourceOk() (ErrorSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasSource

`func (o *ApiError) HasSource() bool`

HasSource returns a boolean if a field has been set.

### SetSource

`func (o *ApiError) SetSource(v ErrorSource)`

SetSource gets a reference to the given ErrorSource and assigns it to the Source field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


